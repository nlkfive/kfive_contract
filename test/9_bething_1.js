const RaceList = artifacts.require("RaceList")
const NLGGT = artifacts.require("NLGGT")
const LINK = artifacts.require("BEP20")
const STORAGE = artifacts.require("MarketplaceStorage")
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

contract("Bething 1", (accounts) => {
    const root = accounts[0]
    const account1 = accounts[1]
    const account2 = accounts[2]
    const pauser = accounts[3]
    const admin = accounts[4]
    const account5 = accounts[5]
    const account6 = accounts[6]
    const account7 = accounts[7]
    const account8 = accounts[8]
    const receiver = accounts[9]
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

    describe('Beth / Update Result / Claim', async () => {
        var raceId = null;
        var raceStartedAt = null;
        var raceEndedAt = null;
        var slots = null;
        var commission = null;
        var rewardRate = null;

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

        it('(Account1) Bet value slot0. Cannot because race has not started', async () => {
            const betValue = web3.utils.toHex(10 * (10 ** tokenDecimals));

            await u.assertRevert(beth.bet(0, raceId, betValue, {
                from: account1
            }));

            totalBet = await beth.totalSlotBet(raceId, 0, {
                from: account1
            });
            eq(totalBet, 0);
        });

        it('Wait until raceStartedAt', async () => {
            while (Math.floor(new Date().getTime() / 1000 < raceStartedAt)) {}
        });

        it('(Account1) Bet value slot0. Cannot because account1 has no token', async () => {
            const betValue = web3.utils.toHex(10 * (10 ** tokenDecimals));

            await u.assertRevert(beth.bet(0, raceId, betValue, {
                from: account1
            }));

            totalBet = await beth.totalSlotBet(raceId, 0, {
                from: account1
            });
            eq(totalBet, 0);
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

        it('(Root) transfer KFIVE token to account5', async () => {
            const i = {
                to: account5,
                value: web3.utils.toHex(100 * (10 ** tokenDecimals)),
            }

            await kfive.mint(i.value, {
                from: root
            })

            await kfive.transfer(i.to, i.value, {
                from: root
            });

            const o = {
                account5: 100 * (10 ** tokenDecimals),
                root_balance: 0,
            }

            let account5_balance = await kfive.balanceOf(account5, {
                from: root
            });
            eq(o.account5, account5_balance.toString());
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

            userSlotBet = await beth.userSlotBet(raceId, 0, account1, {
                from: account1
            });
            eq(userSlotBet.toString(), 10 * (10 ** tokenDecimals));
        });

        it('(Account1) Bet value slot0 successful 2', async () => {
            const value = web3.utils.toHex(10 * (10 ** tokenDecimals));

            await kfive.approve(beth.address, value, {
                from: account1
            });

            await beth.bet(0, raceId, value, {
                from: account1
            });

            const account1_balance = await kfive.balanceOf(account1, {from: root});
            eq(account1_balance.toString(), 80 * (10 ** tokenDecimals))

            totalBet = await beth.totalSlotBet(raceId, 0, {
                from: account1
            });
            eq(totalBet.toString(), 20 * (10 ** tokenDecimals));

            userSlotBet = await beth.userSlotBet(raceId, 0, account1, {
                from: account1
            });
            eq(userSlotBet.toString(), 20 * (10 ** tokenDecimals));
        });

        it('(Account2) Bet value slot1 successful', async () => {
            const value = web3.utils.toHex(10 * (10 ** tokenDecimals));

            await kfive.approve(beth.address, value, {
                from: account2
            });

            await beth.bet(1, raceId, value, {
                from: account2
            });

            const account2_balance = await kfive.balanceOf(account2, {from: root});
            eq(account2_balance.toString(), 90 * (10 ** tokenDecimals))


            totalBet = await beth.totalSlotBet(raceId, 1, {
                from: account2
            });
            eq(totalBet.toString(), 10 * (10 ** tokenDecimals));

            userSlotBet = await beth.userSlotBet(raceId, 1, account2, {
                from: account2
            });
            eq(userSlotBet.toString(), 10 * (10 ** tokenDecimals));
        });

        it('(Account5) Bet value slot1 successful', async () => {
            const value = web3.utils.toHex(5 * (10 ** tokenDecimals));

            await kfive.approve(beth.address, value, {
                from: account5
            });

            await beth.bet(1, raceId, value, {
                from: account5
            });

            const account2_balance = await kfive.balanceOf(account5, {from: root});
            eq(account2_balance.toString(), 95 * (10 ** tokenDecimals))


            totalBet = await beth.totalSlotBet(raceId, 1, {
                from: account5
            });
            eq(totalBet.toString(), 15 * (10 ** tokenDecimals));

            userSlotBet = await beth.userSlotBet(raceId, 1, account5, {
                from: account5
            });
            eq(userSlotBet.toString(), 5 * (10 ** tokenDecimals));
        });

        it('(Account5) Bet value slot2 successful', async () => {
            const value = web3.utils.toHex(5 * (10 ** tokenDecimals));

            await kfive.approve(beth.address, value, {
                from: account5
            });

            await beth.bet(2, raceId, value, {
                from: account5
            });

            const account5_balance = await kfive.balanceOf(account5, {from: root});
            eq(account5_balance.toString(), 90 * (10 ** tokenDecimals))

            totalBet = await beth.totalSlotBet(raceId, 2, {
                from: account5
            });
            eq(totalBet.toString(), 5 * (10 ** tokenDecimals));

            userSlotBet = await beth.userSlotBet(raceId, 2, account5, {
                from: account5
            });
            eq(userSlotBet.toString(), 5 * (10 ** tokenDecimals));
        });

        it('(Root) Claim commission. Cannot, race is not ended', async () => {
            await u.assertRevert(beth.claim(raceId, receiver, {
                from: root
            }));
        });

        it('Wait until raceEnded', async () => {
            while (Math.floor(new Date().getTime() / 1000 < raceEndedAt)) {}
        });

        it('Check total race bet', async () => {
            const totalBetValue = 40 * (10 ** tokenDecimals);

            const total = await beth.totalRaceBet(raceId, { from: root});
            eq(totalBetValue, total.toString());
        });

        it('(Account1) Update Result. Cannot, only admin role', async () => {
            const result = '0x0001000000000000000000000000000000000000000000000000000000000000'
            await u.assertRevert(racelist.updateResult(raceId, result, {
                from: root
            }));
        });
        
        it('(Root) Update Result', async () => {
            const result = '0x0001000000000000000000000000000000000000000000000000000000000000'
            await racelist.updateResult(raceId, result, {
                from: root
            });

            const getRaceResult = await racelist.raceResult(raceId, {from: root})
            eq(getRaceResult, result)
        });

        it('(Account1) Claim result. Cannot, not the winner', async () => {
            await u.assertRevert(beth.claim(raceId, 0, {
                from: account1
            }));
        });

        it('(Account2) Claim result', async () => {
            const account2_balance_before = await kfive.balanceOf(account2, {from: root});

            // Slot1 winner:
            // Account2 bet slot1: 10 Kfive -> position 1
            // Account5 bet slot2: 5 Kfive -> position 2

            const total1stBet = await beth.userSlotBet(raceId, 1, account2, {from: root});
            const total2ndBet = await beth.userSlotBet(raceId, 1, account5, {from: root});
            const totalRaceBet = await beth.totalRaceBet(raceId, {from: root});
            const commissionPercent = totalRaceBet - totalRaceBet * (1000 - commission) / 1000;

            await beth.claim(raceId, 1, {
                from: account2
            });

            // * 1st = total1stBet * rewardRate^2 * [totalRaceBet - %commission] 
            // * div (total1stBet*rewardRate^2 + total2ndBet*rewardRate + total3rdBet)

            reward = ((total1stBet * (rewardRate*rewardRate)) * commissionPercent) / (total1stBet * rewardRate * rewardRate + total2ndBet * rewardRate);

            const account2_balance = await kfive.balanceOf(account2, {from: root});
            eq(reward, (web3.utils.toBN(account2_balance).sub(web3.utils.toBN(account2_balance_before))).toString());
        });

        it('(Account5) Claim result', async () => {
            const account5_balance_before = await kfive.balanceOf(account5, {from: root});

            const total1stBet = await beth.userSlotBet(raceId, 1, account2, {from: root});
            const total2ndBet = await beth.userSlotBet(raceId, 1, account5, {from: root});
            const totalRaceBet = await beth.totalRaceBet(raceId, {from: root});
            const commissionPercent = totalRaceBet - totalRaceBet * (1000 - commission) / 1000;

            await beth.claim(raceId, 1, {
                from: account5
            });

            // * 2nd = total2ndBet * rewardRate * [totalRaceBet - %commission] 
            // * div (total1stBet*rewardRate^2 + total2ndBet*rewardRate + total3rdBet)

            reward = ((total2ndBet * (rewardRate)) * commissionPercent) / (total1stBet * rewardRate * rewardRate + total2ndBet * rewardRate);

            const account5_balance = await kfive.balanceOf(account5, {from: root});
            eq(reward, (web3.utils.toBN(account5_balance).sub(web3.utils.toBN(account5_balance_before))).toString());
        });

        it('(Receiver) Claim commission. Cannot, only admin', async () => {
            await u.assertRevert(beth.claimCommission(raceId, receiver, {
                from: receiver
            }));
        });

        it('(Root) Grant admin_role to admin', async () => {
            await beth.grantRole(adminRole, admin, {
                from: root
            });
        });

        it('(Admin) Claim commission successfully', async () => {
            await beth.claimCommission(raceId, receiver, {
                from: admin
            });
        });
    });
});