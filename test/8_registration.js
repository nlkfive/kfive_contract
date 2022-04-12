const RaceList = artifacts.require("RaceList")
const Registration = artifacts.require("RegistrationList")
const RaceReward = artifacts.require("RaceReward")
const NLGGT = artifacts.require("NLGGT")
const LINK = artifacts.require("BEP20")
const STORAGE = artifacts.require("MarketplaceStorage")

const eq = assert.equal
const u = require('./utils.js')

const { soliditySha3 } = require("web3-utils");
const keccak256 = require('js-sha3').keccak256;
const tokenDecimals = 10;
const defaultParticipant = "0x0000000000000000000000000000000000000000";

var racelist, registration, nlggt, storage, link, racereward;
const adminRole = "0x" + keccak256("ADMIN_ROLE");
const pauserRole = "0x" + keccak256("PAUSER_ROLE");

contract("Registration", (accounts) => {
    const root = accounts[0]
    const account1 = accounts[1]
    const account2 = accounts[2]
    const pauser = accounts[3]
    const admin = accounts[4]
    const account5 = accounts[5]
    const account6 = accounts[6]
    const account7 = accounts[7]
    const account8 = accounts[8]
    const account9 = accounts[9]
    const OFFCHAIN = web3.utils.fromAscii('1')

    const baseURL = "https://ipfs.io/ipfs/";

    const vrfCoor = "0x747973a5A2a4Ae1D3a8fDF5479f1514F65Db9C31";
    const keyHash = "0xc251acd21ec4fb7f31bb8868288bfdbaeb4fbfec2df3735ddbd4f7dc8d60103c";
    subscriptionId = 311;

    const delay = ms => new Promise(res => setTimeout(res, ms));

    before(async () => {
        storage = await STORAGE.new( {from: root} );
        nlggt = await NLGGT.new(storage.address, baseURL, {from: root});
        racereward = await RaceReward.new(storage.address, baseURL, {from: root})
        link = await LINK.new("LINK Token", "LINK", 18, {from: root})
        racelist = await RaceList.new({from: root});
        registration = await Registration.new(vrfCoor, nlggt.address, racereward.address, racelist.address, subscriptionId, keyHash, {from: root});
    });

    afterEach(function() {
        const {currentTest} = this
        if (currentTest.state === 'failed') {
          currentTest.parent.pending = true
        }
    });

    describe('Deployment stage', async () => {
        it('NLGGT information', async () => {
            let o = {
                balance: 0,
                name: "Ngoc Linh Ginseng Garden Token",
                symbol: "NLGGT"
            }

            let owner_balance = await nlggt.balanceOf(root, {
                from: root
            });

            eq(o.balance, owner_balance.toString());
            eq(o.name, await nlggt.name.call());
            eq(o.symbol, await nlggt.symbol.call());
        });

        it('LINK information', async () => {
            let o = {
                balance: 0,
                name: "LINK Token",
                symbol: "LINK"
            }

            let owner_balance = await link.balanceOf(root, {
                from: root
            });

            eq(o.balance, owner_balance.toString());
            eq(o.name, await link.name.call());
            eq(o.symbol, await link.symbol.call());
        });

        it('Race List information', async () => {
            const owner = await racelist.owner();
            eq(owner, root);
        });

        it('Registration information', async () => {
            const owner = await registration.owner();
            eq(owner, root);
        });
    });

    describe('Register slots', async () => {
        var raceId = null;
        var raceStartedAt = null;
        var raceEndedAt = null;
        var slots = null;
        var commission = null;
        var rewardRate = null;

        it('(Root) Create a race successful', async () => {
            const i = {
                slots: 2,
                betStarted: Math.floor(new Date().getTime() / 1000 + 10),
                betEnded: Math.floor(new Date().getTime() / 1000 + 60),
                commission: 20,
                rewardRate: 1000,
            }

            await racelist.createRace(i.slots, i.betStarted, i.betEnded, i.commission, i.rewardRate, {
                from: root
            });

            const lastestEvent = await racelist.getPastEvents("RaceCreated");
            raceId = lastestEvent[0].returnValues.id;

            raceEndedAt = lastestEvent[0].returnValues.betEnded;
            eq(raceEndedAt, i.betEnded);

            raceStartedAt = lastestEvent[0].returnValues.betStarted;
            eq(raceStartedAt, i.betStarted);

            slots = lastestEvent[0].returnValues.slots;
            eq(slots, i.slots);

            commission = lastestEvent[0].returnValues.commission;
            eq(commission, i.commission);

            rewardRate = lastestEvent[0].returnValues.rewardRate;
            eq(rewardRate, i.rewardRate);
        });

        it('(Account1) Register a slot. Cannot because race has not started', async () => {
            await u.assertRevert(registration.register(1, raceId, {
                from: account1
            }));
        });

        it('Wait until raceStartedAt', async () => {
            while (Math.floor(new Date().getTime() / 1000 < raceStartedAt)) {}
        });

        it('(Account1) Register a slot. Cannot because Account is not a NLGGT holder', async () => {
            await u.assertRevert(registration.register(1, raceId, {
                from: account1
            }));
        });

        it('(Owner) Mint NLGGT to account1, account2, account5, account6, account7, account8', async () => {
            await nlggt.mint(account1, 1, "1", {
                from: root
            });

            await nlggt.mint(account2, 2, "2", {
                from: root
            });

            await nlggt.mint(account5, 3, "3", {
                from: root
            });

            await nlggt.mint(account6, 4, "4", {
                from: root
            });

            await nlggt.mint(account7, 5, "5", {
                from: root
            });

            await nlggt.mint(account8, 6, "6", {
                from: root
            });
        });

        it('(Account1) Register slot0 successful', async () => {
            await registration.register(0, raceId, {
                from: account1
            });

            selectedParticipant = await registration.selectedParticipant(raceId, 0, {from: root});
            eq(selectedParticipant, defaultParticipant);
        });

        it('(Account2) Register slot0 successful', async () => {
            await registration.register(0, raceId, {
                from: account2
            });

            selectedParticipant = await registration.selectedParticipant(raceId, 0, {from: root});
            eq(selectedParticipant, defaultParticipant);
        });

        it('(Account5) Register slot1 successful', async () => {
            await registration.register(1, raceId, {
                from: account5
            });

            selectedParticipant = await registration.selectedParticipant(raceId, 1, {from: root});
            eq(selectedParticipant, defaultParticipant);
        });

        it('(Account5) Register slot1. Cannot because account5 has already registered', async () => {
            await u.assertRevert(registration.register(1, raceId, {
                from: account5
            }));
        });

        it('(Account8) Register a slot2. Cannot because of invalid slot', async () => {
            await u.assertRevert(registration.register(2, raceId, {
                from: account8
            }));
        });

        it('(Root) Select participant. Cannot because race is not ended', async () => {
            await u.assertRevert(registration.selectParticipant(raceId, {
                from: root
            }));
        });

        // selectParticipant and receiveReward  are tested manually
    });

    describe('Pause/Unpause', async () => {
        var raceId2 = null;
        var raceStartedAt = null;
        var raceEndedAt = null;
        var slots = null;
        var commission = null;
        var rewardRate = null;

        it('(Root) Create a race successful', async () => {
            const i = {
                slots: 5,
                betStarted: Math.floor(new Date().getTime() / 1000 + 10),
                betEnded: Math.floor(new Date().getTime() / 1000 + 60),
                commission: 20,
                rewardRate: 1000,
            }

            await racelist.createRace(i.slots, i.betStarted, i.betEnded, i.commission, i.rewardRate, {
                from: root
            });

            const lastestEvent = await racelist.getPastEvents("RaceCreated");
            raceId2 = lastestEvent[0].returnValues.id;

            raceEndedAt = lastestEvent[0].returnValues.betEnded;
            eq(raceEndedAt, i.betEnded);

            raceStartedAt = lastestEvent[0].returnValues.betStarted;
            eq(raceStartedAt, i.betStarted);

            slots = lastestEvent[0].returnValues.slots;
            eq(slots, i.slots);

            commission = lastestEvent[0].returnValues.commission;
            eq(commission, i.commission);

            rewardRate = lastestEvent[0].returnValues.rewardRate;
            eq(rewardRate, i.rewardRate);
        });

        it('(Owner) Mint NLGGT to account1, account2, account5, account6, account7, account8', async () => {
            await nlggt.mint(account1, 7, "7", {
                from: root
            });

            await nlggt.mint(account2, 8, "8", {
                from: root
            });

            await nlggt.mint(account5, 9, "9", {
                from: root
            });

            await nlggt.mint(account6, 10, "10", {
                from: root
            });

            await nlggt.mint(account7, 11, "11", {
                from: root
            });

            await nlggt.mint(account8, 12, "12", {
                from: root
            });
        });

        it('(Account1) Pause. Cannot, need pauser_role', async () => {
            await u.assertRevert(registration.pause({
                from: account1
            }));
        });

        it('(Account2) Register slot0 successful', async () => {
            await registration.register(0, raceId2, {
                from: account2
            });
        });

        it('(Root) Grant pauser_role to pauser', async () => {
            await registration.grantRole(pauserRole, pauser, {
                from: root
            });
        });

        it('(Pauser) Pause', async () => {
            await registration.pause({
                from: pauser
            });

            var alreadyPaused = await registration.paused({ from: pauser });
            eq(alreadyPaused, true)
        });

        it('(Account5) Register slot1. Cannot because contract is pausing', async () => {
            await u.assertRevert(registration.register(1, raceId2, {
                from: account5
            }));
        });
        
        it('(Root) UnPause', async () => {
            await registration.unpause({
                from: root
            });
        });

        it('(Account5) Register slot1 successful', async () => {
            await registration.register(1, raceId2, {
                from: account5
            });
        });
    });

    describe('Update NLGGT/Update raceAddress', async () => {
        var raceId3 = null;
        var raceStartedAt = null;
        var raceEndedAt = null;
        var slots = null;
        var commission = null;
        var rewardRate = null;
        var new_nlggt = null;
        var new_racelist = null;

        it('(Root) Create a race successful', async () => {
            const i = {
                slots: 5,
                betStarted: Math.floor(new Date().getTime() / 1000 + 10),
                betEnded: Math.floor(new Date().getTime() / 1000 + 60),
                commission: 20,
                rewardRate: 1000,
            }

            await racelist.createRace(i.slots, i.betStarted, i.betEnded, i.commission, i.rewardRate, {
                from: root
            });

            const lastestEvent = await racelist.getPastEvents("RaceCreated");
            raceId3 = lastestEvent[0].returnValues.id;

            raceEndedAt = lastestEvent[0].returnValues.betEnded;
            eq(raceEndedAt, i.betEnded);

            raceStartedAt = lastestEvent[0].returnValues.betStarted;
            eq(raceStartedAt, i.betStarted);

            slots = lastestEvent[0].returnValues.slots;
            eq(slots, i.slots);

            commission = lastestEvent[0].returnValues.commission;
            eq(commission, i.commission);

            rewardRate = lastestEvent[0].returnValues.rewardRate;
            eq(rewardRate, i.rewardRate);
        });

        it('(Owner) Mint NLGGT to account2', async () => {
            await nlggt.mint(account8, 13, "13", {
                from: root
            });

            await nlggt.mint(account5, 14, "14", {
                from: root
            });
        });

        it('Wait until raceStartedAt', async () => {
            while (Math.floor(new Date().getTime() / 1000 < raceStartedAt)) {}
        });

        it('(Account8) Register a slot raceId3 successful', async () => {
            await registration.register(3, raceId3, {
                from: account8
            });
        });

        it('(Root) Update new NLGGT address', async () => {
            new_nlggt = await NLGGT.new(storage.address, baseURL, {from: root});

            await registration.updateNlggtAddress(new_nlggt.address, {
                from: root
            });
        });

        it('(Root) Deploy new NLGGT', async () => {
            new_nlggt = await NLGGT.new(storage.address, baseURL, {from: root});
        });

        it('(Account1) Update new NLGGT address. Cannot because only admin can update', async () => {
            await u.assertRevert(registration.updateNlggtAddress(new_nlggt.address, {
                from: account1
            }));
        });

        it('(Root) Grant admin_role to admin', async () => {
            await registration.grantRole(adminRole, admin, {
                from: root
            });
        });

        it('(Admin) Update new NLGGT address', async () => {
            await registration.updateNlggtAddress(new_nlggt.address, {
                from: admin
            });

            const lastestEvent = await registration.getPastEvents("NlggtUpdated");
            new_nlggt_address = lastestEvent[0].returnValues.nlggt;
            eq(new_nlggt_address, new_nlggt.address);
        });

        it('(Account5) Register a slot. Cannot because no NLGGT (nlggt contract has been changed', async () => {
            await u.assertRevert(registration.register(1, raceId3, {
                from: account1
            }));
        });

        it('(Account5) Register a slot using new_nlggt', async () => {
            await new_nlggt.mint(account5, 1, "1", {
                from: root
            });

            await registration.register(1, raceId3, {
                from: account5
            });
        });

        it('(Root) Deploy new racelist address', async () => {
            new_racelist = await RaceList.new({from: root});
        });

        it('(Admin) Update new Race address', async () => {
            await registration.updateRaceAddress(new_racelist.address, {
                from: admin
            });

            const lastestEvent = await registration.getPastEvents("RaceListUpdated");
            new_race_address = lastestEvent[0].returnValues.race;
            eq(new_race_address, new_racelist.address);
        });

        it('(Root) Create a new race successful', async () => {
            const i = {
                slots: 2,
                betStarted: Math.floor(new Date().getTime() / 1000 + 10),
                betEnded: Math.floor(new Date().getTime() / 1000 + 60),
                commission: 20,
                rewardRate: 1000,
            }

            await new_racelist.createRace(i.slots, i.betStarted, i.betEnded, i.commission, i.rewardRate, {
                from: root
            });

            const lastestEvent = await new_racelist.getPastEvents("RaceCreated");
            raceId3 = lastestEvent[0].returnValues.id;

            raceEndedAt = lastestEvent[0].returnValues.betEnded;
            eq(raceEndedAt, i.betEnded);

            raceStartedAt = lastestEvent[0].returnValues.betStarted;
            eq(raceStartedAt, i.betStarted);

            slots = lastestEvent[0].returnValues.slots;
            eq(slots, i.slots);

            commission = lastestEvent[0].returnValues.commission;
            eq(commission, i.commission);

            rewardRate = lastestEvent[0].returnValues.rewardRate;
            eq(rewardRate, i.rewardRate);
        });

        it('Wait until raceStartedAt', async () => {
            while (Math.floor(new Date().getTime() / 1000 < raceStartedAt)) {}
        });

        it('(Account5) Register a slot after updating racelist', async () => {
            await new_nlggt.mint(account1, 2, "2", {
                from: root
            });

            await registration.register(0, raceId3, {
                from: account1
            });
        });
    });
});