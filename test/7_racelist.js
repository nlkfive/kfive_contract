const RaceList = artifacts.require("RaceList")

const eq = assert.equal
const u = require('./utils.js')

const { soliditySha3 } = require("web3-utils");
const keccak256 = require('js-sha3').keccak256;
const tokenDecimals = 10;
const secret = "0x0000000000000000000000000000000000000000000000000000000000000001";

var racelist, chainlink;
const adminRole = "0x" + keccak256("ADMIN_ROLE");
const pauserRole = "0x" + keccak256("PAUSER_ROLE");

contract("Race List", (accounts) => {
    const root = accounts[0]
    const account1 = accounts[1]
    const account2 = accounts[2]
    const pauser = accounts[3]
    const admin = accounts[4]
    const account5 = accounts[5]
    const account6 = accounts[6]
    const new_owner = accounts[7]
    const account8 = accounts[8]
    const canceller = accounts[9]
    const OFFCHAIN = web3.utils.fromAscii('1')

    const delay = ms => new Promise(res => setTimeout(res, ms));

    before(async () => {
        racelist = await RaceList.new({from: root});
    });

    afterEach(function() {
        const {currentTest} = this
        if (currentTest.state === 'failed') {
          currentTest.parent.pending = true
        }
    });

    describe('Deployment stage', async () => {
        it('Race List information', async () => {
            const owner = await racelist.owner();
            eq(owner, root);
        });
    });

    describe('Create a race', async () => {
        var raceId = null;
        var raceStartedAt = null;
        var raceEndedAt = null;
        var slots = null;
        var commission = null;
        var rewardRate = null;

        it('(Root) Create a race. Cannot because slot > 32', async () => {
            const i = {
                slots: 100,
                betStarted: Math.floor(new Date().getTime() / 1000 + 10),
                betEnded: Math.floor(new Date().getTime() / 1000 + 60),
                commission: 20,
                rewardRate: 1100,
            }

            await u.assertRevert(racelist.createRace(i.slots, i.betStarted, i.betEnded, i.commission, i.rewardRate, {
                from: root
            }));
        });

        it('(Root) Create a race. Cannot because slot == 0', async () => {
            const i = {
                slots: 0,
                betStarted: Math.floor(new Date().getTime() / 1000 + 10),
                betEnded: Math.floor(new Date().getTime() / 1000 + 60),
                commission: 20,
                rewardRate: 1100,
            }

            await u.assertRevert(racelist.createRace(i.slots, i.betStarted, i.betEnded, i.commission, i.rewardRate, {
                from: root
            }));
        });

        it('(Root) Create a race. Cannot because beStarted > betEnded', async () => {
            const i = {
                slots: 10,
                betStarted: Math.floor(new Date().getTime() / 1000 + 10),
                betEnded: Math.floor(new Date().getTime() / 1000 - 10),
                commission: 20,
                rewardRate: 1100,
            }

            await u.assertRevert(racelist.createRace(i.slots, i.betStarted, i.betEnded, i.commission, i.rewardRate, {
                from: root
            }));
        });

        it('(Root) Create a race. Cannot because commission > 1000', async () => {
            const i = {
                slots: 10,
                betStarted: Math.floor(new Date().getTime() / 1000 + 10),
                betEnded: Math.floor(new Date().getTime() / 1000 + 60),
                commission: 1001,
                rewardRate: 1100,
            }

            await u.assertRevert(racelist.createRace(i.slots, i.betStarted, i.betEnded, i.commission, i.rewardRate, {
                from: root
            }));
        });

        it('(Root) Create a race. Cannot because commission > 1000', async () => {
            const i = {
                slots: 10,
                betStarted: Math.floor(new Date().getTime() / 1000 + 10),
                betEnded: Math.floor(new Date().getTime() / 1000 + 60),
                commission: 1001,
                rewardRate: 1000,
            }

            await u.assertRevert(racelist.createRace(i.slots, i.betStarted, i.betEnded, i.commission, i.rewardRate, {
                from: root
            }));
        });

        it('(Root) Create a race. Cannot because rewardRate < 1000', async () => {
            const i = {
                slots: 10,
                betStarted: Math.floor(new Date().getTime() / 1000 + 10),
                betEnded: Math.floor(new Date().getTime() / 1000 + 60),
                commission: 20,
                rewardRate: 999,
            }

            await u.assertRevert(racelist.createRace(i.slots, i.betStarted, i.betEnded, i.commission, i.rewardRate, {
                from: root
            }));
        });

        it('(Account1) Create a race. Cannot because onlyAdmin can call', async () => {
            const i = {
                slots: 10,
                betStarted: Math.floor(new Date().getTime() / 1000 + 10),
                betEnded: Math.floor(new Date().getTime() / 1000 + 60),
                commission: 20,
                rewardRate: 1000,
            }

            await u.assertRevert(racelist.createRace(i.slots, i.betStarted, i.betEnded, i.commission, i.rewardRate, {
                from: account1
            }));
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
    });

    describe('Cancel a race', async () => {
        var raceId = null;
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

        it('(Account1) Cancel race. Cannot because account1 is not an admin', async () => {
            await u.assertRevert(racelist.cancelRace(raceId, {
                from: account1
            }));
        });

        it('(Root) Grant admin_role to admin', async () => {
            await racelist.grantRole(adminRole, admin, {
                from: root
            });
        });

        it('(Admin) Cancel race successful', async () => {
            await racelist.cancelRace(raceId, {
                from: admin
            });
        });

        it('(Admin) Create race2 successful', async () => {
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

        it('(Root) Cancel race 2. Cannot because bet time has started', async () => {
            await u.assertRevert(racelist.cancelRace(raceId, {
                from: root
            }));

            await racelist.getRace(raceId, {
                from: root
            });
        });
    });

    describe('Update commission / rewardRate', async () => {
        var raceId = null;
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

        it('(Account1) Update commission. Cannot because only admin can update', async () => {
            await u.assertRevert(racelist.updateCommission(raceId, 50, {
                from: account1
            }));
        });

        it('(Root) Grant admin_role to admin', async () => {
            await racelist.grantRole(adminRole, admin, {
                from: root
            });
        });

        it('(Admin) Update commission = 200', async () => {
            await racelist.updateCommission(raceId, 200, {
                from: admin
            });

            const lastestEvent = await racelist.getPastEvents("RaceCommissionUpdated");
            commission = lastestEvent[0].returnValues.commission;
            eq(commission, 200);
        });

        it('(Root) Update commission = 1100. Cannot because commission need to < 1000', async () => {
            await u.assertRevert(racelist.updateCommission(raceId, 1100, {
                from: root
            }));

            const getRace = await racelist.getRace(raceId, {
                from: root
            });

            const Commission = getRace[0];
            eq(Commission, 200)
        });

        it('(Account1) Update rewardRate. Cannot because only admin can update', async () => {
            await u.assertRevert(racelist.updateRewardRate(raceId, 1500, {
                from: account1
            }));
        });

        it('(Admin) Update rewardRate = 1500', async () => {
            await racelist.updateRewardRate(raceId, 1500, {
                from: admin
            });

            const lastestEvent = await racelist.getPastEvents("RaceRewardRateUpdated");
            rewardRate = lastestEvent[0].returnValues.rewardRate;
            eq(rewardRate, 1500);
        });

        it('(Root) Update rewardRate = 700. Cannot because rewardRate must >= 1000', async () => {
            await u.assertRevert(racelist.updateRewardRate(raceId, 700, {
                from: root
            }));

            const getRace = await racelist.getRace(raceId, {
                from: root
            });

            const RewardRate = getRace[1];
            eq(RewardRate, 1500)
        });

        it('Wait until raceEnded', async () => {
            while (Math.floor(new Date().getTime() / 1000 < raceEndedAt)) {}
        });

        it('(Root) Update commission = 111. Cannot because betEnded', async () => {
            await u.assertRevert(racelist.updateCommission(raceId, 111, {
                from: root
            }));

            const getRace = await racelist.getRace(raceId, {
                from: root
            });

            const Commission = getRace[0];
            eq(Commission, 200);
        });

        it('(Root) Update rewardRate = 1111. Cannot because betEnded', async () => {
            await u.assertRevert(racelist.updateRewardRate(raceId, 1111, {
                from: root
            }));

            const getRace = await racelist.getRace(raceId, {
                from: root
            });

            const rewardRate = getRace[1];
            eq(rewardRate, 1500);
        });
    });

    describe('Pause', async () => {
        var raceId = null;
        var raceStartedAt = null;
        var raceEndedAt = null;
        var slots = null;
        var commission = null;
        var rewardRate = null;

        it('(Account1) Pause. Cannot, need pauser_role', async () => {
            await u.assertRevert(racelist.pause({
                from: account1
            }));
        });

        it('(Root) Grant pauser_role to pauser', async () => {
            await racelist.grantRole(pauserRole, pauser, {
                from: root
            });
        });

        it('(Root) Grant admin_role to admin', async () => {
            await racelist.grantRole(adminRole, admin, {
                from: root
            });
        });

        it('(Pauser) Pause', async () => {
            await racelist.pause({
                from: pauser
            });
        });

        it('(Root) Create a race. Cannot because race is pausing', async () => {
            const i = {
                slots: 5,
                betStarted: Math.floor(new Date().getTime() / 1000 + 10),
                betEnded: Math.floor(new Date().getTime() / 1000 + 60),
                commission: 20,
                rewardRate: 1000,
            }

            await u.assertRevert(racelist.createRace(i.slots, i.betStarted, i.betEnded, i.commission, i.rewardRate, {
                from: root
            }));
        });

        it('(Account1) UnPause. Cannot, only pauser_role can unpause', async () => {
            await u.assertRevert(racelist.unpause({
                from: account1
            }));
        });

        it('(Root) UnPause', async () => {
            await racelist.unpause({
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

        it('(root) Pause', async () => {
            await racelist.pause({
                from: root
            });
        });

        it('(Admin) Update commission = 200. Cannot because race is pausing', async () => {
            await u.assertRevert(racelist.updateCommission(raceId, 200, {
                from: admin
            }));

            const getRace = await racelist.getRace(raceId, {
                from: root
            });

            const Commission = getRace[0];
            eq(Commission, 20);
        });

        it('(Root) Update rewardRate = 1500. Cannot because race is pausing', async () => {
            await u.assertRevert(racelist.updateRewardRate(raceId, 700, {
                from: root
            }));

            const getRace = await racelist.getRace(raceId, {
                from: root
            });

            const RewardRate = getRace[1];
            eq(RewardRate, 1000);
        });

        it('(Pauser) UnPause', async () => {
            await racelist.unpause({
                from: pauser
            });
        });

        it('(Root) Update commission = 200', async () => {
            await racelist.updateCommission(raceId, 200, {
                from: root
            });

            const lastestEvent = await racelist.getPastEvents("RaceCommissionUpdated");
            commission = lastestEvent[0].returnValues.commission;
            eq(commission, 200);
        });

        it('(Root) Update rewardRate = 1500', async () => {
            await racelist.updateRewardRate(raceId, 1500, {
                from: root
            });

            const lastestEvent = await racelist.getPastEvents("RaceRewardRateUpdated");
            rewardRate = lastestEvent[0].returnValues.rewardRate;
            eq(rewardRate, 1500);
        });
    });
});