const RaceList = artifacts.require("RaceList")
const NLGGT = artifacts.require("NLGGT")
const STORAGE = artifacts.require("MarketplaceStorage")
const LINK = artifacts.require("BEP20")
const BETH = artifacts.require("Bething")
const KFIVE = artifacts.require("KFIVE")

const eq = assert.equal
const u = require('./utils.js')

const { soliditySha3 } = require("web3-utils");
const { latestTime } = require('./utils.js')
const keccak256 = require('js-sha3').keccak256;
const tokenDecimals = 10;
const secret = "0x0000000000000000000000000000000000000000000000000000000000000001";
const superAdminRole = "0x0000000000000000000000000000000000000000000000000000000000000000";

var racelist, kfive, nlggt, storage, link, beth;
const adminRole = "0x" + keccak256("ADMIN_ROLE");
const pauserRole = "0x" + keccak256("PAUSER_ROLE");

contract("Bething 2", (accounts) => {
    const root = accounts[0]
    const account1 = accounts[1]
    const account2 = accounts[2]
    const pauser = accounts[3]
    const admin = accounts[4]
    const account5 = accounts[5]
    const account6 = accounts[6]
    const account7 = accounts[7]
    const funder1 = accounts[8]
    const funder2 = accounts[9]
    const OFFCHAIN = web3.utils.fromAscii('1')

    const baseURL = "https://ipfs.io/ipfs/";

    before(async () => {
        storage = await STORAGE.new( {from: root});
        nlggt = await NLGGT.new(storage.address, baseURL, {from: root});
        link = await LINK.new("LINK Token", "LINK", 18, {from: root});
        kfive = await KFIVE.new( {from: root});
        racelist = await RaceList.new({from: root});
        beth = await BETH.new(kfive.address, racelist.address, {from: root});
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

        it('KFIVE information', async () => {
            let o = {
                balance: 0,
                name: 'KFIVE',
                symbol: 'KFIVE',
                decimals: tokenDecimals,
            }

            let owner_balance = await kfive.balanceOf(root, {
                from: root
            });

            eq(o.balance, owner_balance.toString());
            eq(o.name, await kfive.name.call());
            eq(o.symbol, await kfive.symbol.call());
            eq(o.decimals, await kfive.decimals.call());
        });

        it('Race List information', async () => {
            const owner = await racelist.owner();
            eq(owner, root);
        });

        it('Bething information', async () => {
            const owner = await beth.owner();
            eq(owner, root);

            const token = await beth._acceptedToken();
            eq(token, kfive.address);

            const racelistAdd = await beth._raceList();
            eq(racelistAdd, racelist.address)

            const pauseStatus = await beth.paused();
            eq(pauseStatus, false);
        });
    });

    describe('Fund Race / Pause / Unpause', async () => {
        var raceId = null;
        var raceStartedAt = null;
        var raceEndedAt = null;
        var slots = null;
        var commission = null;
        var rewardRate = null;

        it('(Owner) Mint NLGGT to account1, account2, funders', async () => {
            await nlggt.mint(account1, 1, "1", {
                from: root
            });

            await nlggt.mint(account2, 2, "2", {
                from: root
            });

            await nlggt.mint(funder1, 3, "3", {
                from: root
            });

            await nlggt.mint(funder2, 4, "4", {
                from: root
            });
        });

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

        it('Wait until raceStartedAt', async () => {
            while (Math.floor(new Date().getTime() / 1000 < raceStartedAt)) {}
        });

        it('(Root) transfer KFIVE token to account1', async () => {
            const i = {
                to: account1,
                value: web3.utils.toHex(100 * (10 ** tokenDecimals)),
            }

            await kfive.mint(i.value, {
                from: root
            })

            await kfive.transfer(i.to, i.value, {
                from: root
            });

            const o = {
                account1: 100 * (10 ** tokenDecimals),
                root_balance: 0,
            }

            let account1_balance = await kfive.balanceOf(account1, {
                from: root
            });
            eq(o.account1, account1_balance.toString());
        });

        it('(Root) transfer KFIVE token to account2', async () => {
            const i = {
                to: account2,
                value: web3.utils.toHex(100 * (10 ** tokenDecimals)),
            }

            await kfive.mint(i.value, {
                from: root
            })

            await kfive.transfer(i.to, i.value, {
                from: root
            });

            const o = {
                account2: 100 * (10 ** tokenDecimals),
                root_balance: 0,
            }

            let account2_balance = await kfive.balanceOf(account2, {
                from: root
            });
            eq(o.account2, account2_balance.toString());
        });

        it('(Root) transfer KFIVE token to funders', async () => {
            const i = {
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            await kfive.mint(i.value, {
                from: root
            })

            await kfive.transfer(funder1, i.value, {
                from: root
            });

            await kfive.mint(i.value, {
                from: root
            })

            await kfive.transfer(funder2, i.value, {
                from: root
            });

            const o = {
                funder1_balance: 10 * (10 ** tokenDecimals),
                funder2_balance: 10 * (10 ** tokenDecimals),
            }

            let funder1_balance = await kfive.balanceOf(funder1, {
                from: root
            });
            eq(o.funder1_balance, funder1_balance.toString());

            let funder2_balance = await kfive.balanceOf(funder2, {
                from: root
            });
            eq(o.funder2_balance, funder2_balance.toString());
        });

        it('(Account1) Bet value slot0 successful', async () => {
            const value = web3.utils.toHex(10 * (10 ** tokenDecimals));

            await kfive.approve(beth.address, value, {
                from: account1
            });

            await beth.bet(0, raceId, value, {
                from: account1
            });

            const account1_balance = await kfive.balanceOf(account1, {from: root});
            eq(account1_balance.toString(), 90 * (10 ** tokenDecimals))

            totalBet = await beth.totalSlotBet(raceId, 0, {
                from: account1
            });
            eq(totalBet.toString(), 10 * (10 ** tokenDecimals));
        });

        it('(Account7) Pause. Cannot, only pauser role', async () => {
            await u.assertRevert(beth.grantRole(pauserRole, pauser, {
                from: account7  
            }));
        });
        
        it('(Root) Grant pauser_role to pauser', async () => {
            await beth.grantRole(pauserRole, pauser, {
                from: root
            });
        });

        it('(Pauser) Pause', async () => {
            await beth.pause({
                from: pauser
            });

            var alreadyPaused = await beth.paused({ from: pauser });
            eq(alreadyPaused, true)
        });

        it('(Account2) Bet value slot1. Cannot, pausing', async () => {
            const value = web3.utils.toHex(10 * (10 ** tokenDecimals));

            await kfive.approve(beth.address, value, {
                from: account2
            });

            await u.assertRevert(beth.bet(1, raceId, value, {
                from: account2
            }));
        });

        it('(Root) Grant adminRole to funder1', async () => {
            await beth.grantRole(adminRole, funder1, {
                from: root
            });
        });

        it('(Funder1) Fund race. Cannot, pausing', async () => {
            const value = web3.utils.toHex(5 * (10 ** tokenDecimals));

            await kfive.approve(beth.address, value, {
                from: funder1
            });

            await u.assertRevert(beth.fundRace(raceId, value, {
                from: funder1
            }));

            const funder1_balance = await kfive.balanceOf(funder1, {from: root});
            eq(funder1_balance, 10 * (10 ** tokenDecimals));
        });

        it('(Root) UnPause', async () => {
            await beth.unpause({
                from: root
            });

            var alreadyPaused = await beth.paused({ from: root });
            eq(alreadyPaused, false)
        });

        it('(Account2) Bet value slot1 successful', async () => {
            const value = web3.utils.toHex(10 * (10 ** tokenDecimals));

            await kfive.approve(beth.address, value, {
                from: account2
            });

            await beth.bet(1, raceId, value, {
                from: account2
            });

            totalBet = await beth.totalSlotBet(raceId, 1, {
                from: account2
            });
            eq(totalBet.toString(), 10 * (10 ** tokenDecimals));
        });

        it('(Funder1) Fund race. Cannot, over balance', async () => {
            const value = web3.utils.toHex(100 * (10 ** tokenDecimals));

            await kfive.approve(beth.address, value, {
                from: funder1
            });

            await u.assertRevert(beth.fundRace(raceId, value, {
                from: funder1
            }));

            const funder1_balance = await kfive.balanceOf(funder1, {from: root});
            eq(funder1_balance, 10 * (10 ** tokenDecimals));
        });

        it('(Funder1) Fund race successful', async () => {
            const value = web3.utils.toHex(5 * (10 ** tokenDecimals));

            await kfive.approve(beth.address, value, {
                from: funder1
            });

            await beth.fundRace(raceId, value, {
                from: funder1
            });

            const funder1_balance = await kfive.balanceOf(funder1, {from: root});
            eq(funder1_balance, 5 * (10 ** tokenDecimals));
        });

        it('(Funder2) Fund race successful', async () => {
            const value = web3.utils.toHex(10 * (10 ** tokenDecimals));

            await kfive.approve(beth.address, value, {
                from: funder2
            });

            await beth.fundRace(raceId, value, {
                from: funder2
            })

            const funder2_balance = await kfive.balanceOf(funder2, {from: root});
            eq(funder2_balance, 0);
        });

        it('Wait until raceEnded', async () => {
            while (Math.floor(new Date().getTime() / 1000 < raceEndedAt)) {}
        });

        it('Check total race bet', async () => {
            const totalBetValue = 35 * (10 ** tokenDecimals);

            const total = await beth.totalRaceBet(raceId, { from: root});
            eq(totalBetValue, total.toString());
        });
    });
});