const ICO = artifacts.require("BEP20KFIVECrowdsale")
const KFIVE = artifacts.require("KFIVE")
const BUSD = artifacts.require("BUSD")

const eq = assert.equal
const u = require('./utils.js')
var kfive, ico, busd;

const tokenCap = 1080000;
const tokenDecimals = 10;
const tokenDecimalsBUSD = 18;

contract("Presale ICO", (accounts) => {
    const root = accounts[9]
    const account1 = accounts[8]
    const account2 = accounts[7]
    const account3 = accounts[6]
    const account4 = accounts[5]
    const new_receiver = accounts[4]
    const account5 = accounts[3]
    const account6 = accounts[2]
    const account7 = accounts[1]
    const account8 = accounts[0]
    const OFFCHAIN = web3.utils.fromAscii('1')

    const rate = 2631578947;
    const wallet = root;
    const cap = 100 * (10 ** tokenDecimals);
    const delay = ms => new Promise(res => setTimeout(res, ms));

    before(async () => {
        busd = await BUSD.new({
            from: root
        });

        kfive = await KFIVE.new({
            from: root
        });

        const soldToken = kfive.address;
        const acceptedToken = busd.address;

        const openingTime = Math.floor(new Date().getTime() / 1000 + 60);
        const closingTime = Math.floor(openingTime + 150);

        ico = await ICO.new(rate, wallet, soldToken, acceptedToken, cap, openingTime, closingTime, {
            from: root
        });
    });

    describe('Deployment stage', async () => {
        it('KFIVE information - Sold Token', async () => {
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

        it('BUSD information - Accepted Token: Used to buy KFIVE', async () => {
            let o = {
                balance: 0,
                name: 'Binance USD',
                symbol: 'BUSD',
                decimals: tokenDecimalsBUSD,
            }

            let owner_balance = await busd.balanceOf(root, {
                from: root
            });

            eq(o.balance, owner_balance.toString());
            eq(o.name, await busd.name.call());
            eq(o.symbol, await busd.symbol.call());
            eq(o.decimals, await busd.decimals.call());
        });

        it('BEP20ICO information', async () => {
            let o = {
                soldToken: kfive.address,
                acceptedToken: busd.address,
                wallet: root,
                weiRaised: 0,
                rate: 2631578947,
            }

            let soldToken = await ico.token({
                from: root
            });
            eq(o.soldToken, soldToken);

            let receiver = await ico.wallet({
                from: root
            });
            eq(o.wallet, receiver);

            let initial_weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, initial_weiRaised);

            let rate = await ico.rate({
                from: root
            })
            eq(o.rate, rate);
        });
    });

    describe('Pre ICO stage', async () => {
        it('Issue (owner) 10 KFIVE token to Crowdsale contract', async () => {
            const i = {
                to: ico.address,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            await kfive.issue(i.to, i.value, OFFCHAIN, {
                from: root
            });

            const o = {
                ico_balance: 10 * (10 ** tokenDecimals),
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());
        });
    });

    describe('BUSD Stage', async () => {
        it('Issue (account1) BUSB token to Account1. Cannot because account1 is not the owner', async () => {
            const i = {
                to: account1,
                value: web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD)),
            }

            await u.assertRevert(busd.issue(i.to, i.value, OFFCHAIN, {
                from: account1
            }));
        });

        it('Issue (owner) BUSB token to Account1 and Account2', async () => {
            const i = {
                value: web3.utils.toHex(10 * (10 ** tokenDecimalsBUSD)),
            }

            await busd.issue(account1, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(account2, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(account3, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(account4, web3.utils.toHex(20 * (10 ** tokenDecimalsBUSD)), OFFCHAIN, {
                from: root
            });

            await busd.issue(account5, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(account6, i.value, OFFCHAIN, {
                from: root
            });

            const o = {
                account1_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
                account2_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
                account3_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
                account4_busd_balance: 20 * (10 ** tokenDecimalsBUSD),
                account5_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
                account6_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
            }

            let account1_busd_balance = await busd.balanceOf(account1, {
                from: root
            });
            eq(o.account1_busd_balance, account1_busd_balance.toString());

            let account2_busd_balance = await busd.balanceOf(account2, {
                from: root
            });
            eq(o.account2_busd_balance, account2_busd_balance.toString());

            let account3_busd_balance = await busd.balanceOf(account3, {
                from: root
            });
            eq(o.account3_busd_balance, account3_busd_balance.toString());

            let account4_busd_balance = await busd.balanceOf(account4, {
                from: root
            });
            eq(o.account4_busd_balance, account4_busd_balance.toString());

            let account5_busd_balance = await busd.balanceOf(account5, {
                from: root
            });
            eq(o.account5_busd_balance, account5_busd_balance.toString());

            let account6_busd_balance = await busd.balanceOf(account6, {
                from: root
            });
            eq(o.account6_busd_balance, account6_busd_balance.toString());
        });

        it('Approve BUSD to crowdsale', async () => {
            const i = {
                to: ico.address,
                value: web3.utils.toHex(10 * (10 ** tokenDecimalsBUSD)),
            }

            await busd.approve(i.to, i.value, {
                from: account1
            });

            await busd.approve(i.to, i.value, {
                from: account2
            });

            await busd.approve(i.to, i.value, {
                from: account3
            });

            await busd.approve(i.to, web3.utils.toHex(20 * (10 ** tokenDecimalsBUSD)), {
                from: account4
            });

            await busd.approve(i.to, i.value, {
                from: account5
            });

            await busd.approve(i.to, i.value, {
                from: account6
            });
        });
    });

    describe('Running ICO stage. 3.8 BUSD = 1 KFIVE', async () => {
        it('Account1 transfer 3.8 * 10^18 wei to get 1 KFIVE token. Cannot because crowdsale is not opening', async () => {
            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await u.assertRevert(ico.buyTokens(account1, value, { from: account1, value: value }));
        });

        it('Wait until crowdsale opens', async () => {
            await delay(60000);
        });

        it('Account1 transfer BUSD to get 1 token', async () => {
            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await ico.buyTokens(account1, value, { from: account1, value: value });

            const o = {
                ico_balance: 10 * (10 ** tokenDecimals) - (rate * 4),
                account1_balance_KFIVE: rate * 4,
                account1_balance_BUSD: 6 * (10 ** tokenDecimalsBUSD),
                root_balance_BUSD: 4 * (10 ** tokenDecimalsBUSD),
                weiRaised: 4,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account1_balance_BUSD = await busd.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance_BUSD, account1_balance_BUSD.toString());

            let root_balance_BUSD = await busd.balanceOf(root, {
                from: root
            });
            eq(o.root_balance_BUSD, root_balance_BUSD.toString());

            let account1_balance_KFIVE = await kfive.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance_KFIVE, account1_balance_KFIVE.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('Account2 transfer BUSD to get 2 KFIVE', async () => {
            const value = web3.utils.toHex(8 * (10 ** tokenDecimalsBUSD));
            await ico.buyTokens(account2, value, { from: account2, value: value });

            const o = {
                ico_balance: 10 * (10 ** tokenDecimals) - (rate * 12),
                account2_balance_KFIVE: rate * 8,
                account2_balance_BUSD: 2 * (10 ** tokenDecimalsBUSD),
                root_balance_BUSD: 12 * (10 ** tokenDecimalsBUSD),
                weiRaised: 12,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account2_balance_BUSD = await busd.balanceOf(account2, {
                from: root
            });
            eq(o.account2_balance_BUSD, account2_balance_BUSD.toString());

            let root_balance_BUSD = await busd.balanceOf(root, {
                from: root
            });
            eq(o.root_balance_BUSD, root_balance_BUSD.toString());

            let account2_balance_KFIVE = await kfive.balanceOf(account2, {
                from: root
            });
            eq(o.account2_balance_KFIVE, account2_balance_KFIVE.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('Account2 transfer BUSD to get 1 KFIVE. Cannot because account does not have enough BUSD', async () => {
            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await u.assertRevert(ico.buyTokens(account2, value, { from: account2, value: value }));

            const o = {
                account2_balance_BUSD: 2 * (10 ** tokenDecimalsBUSD),
            }

            let account2_balance_BUSD = await busd.balanceOf(account2, {
                from: root
            });
            eq(o.account2_balance_BUSD, account2_balance_BUSD.toString());
        });

        it('Root transfer BUSD to get 1 token', async () => {
            await busd.issue(root, web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD)), OFFCHAIN, {
                from: root
            });

            await busd.approve(ico.address, web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD)), {
                from: root
            });

            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await ico.buyTokens(root, value, { from: root, value: value });

            const o = {
                ico_balance: 10 * (10 ** tokenDecimals) - (rate * 16),
                root_balance_KFIVE: rate * 4,
                root_balance_BUSD: 16 * (10 ** tokenDecimalsBUSD),
                weiRaised: 16,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let root_balance_BUSD = await busd.balanceOf(root, {
                from: root
            });
            eq(o.root_balance_BUSD, root_balance_BUSD.toString());

            let root_balance_KFIVE = await kfive.balanceOf(root, {
                from: root
            });
            eq(o.root_balance_KFIVE, root_balance_KFIVE.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('Account3 transfer BUSD to get 2 token', async () => {
            const value = web3.utils.toHex(8 * (10 ** tokenDecimalsBUSD));
            await ico.buyTokens(account3, value, { from: account3, value: value });

            const o = {
                ico_balance: 10 * (10 ** tokenDecimals) - (rate * 24),
                account3_balance_KFIVE: rate * 8,
                account3_balance_BUSD: 2 * (10 ** tokenDecimalsBUSD),
                root_balance_BUSD: 24 * (10 ** tokenDecimalsBUSD),
                weiRaised: 24,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account3_balance_BUSD = await busd.balanceOf(account2, {
                from: root
            });
            eq(o.account3_balance_BUSD, account3_balance_BUSD.toString());

            let root_balance_BUSD = await busd.balanceOf(root, {
                from: root
            });
            eq(o.root_balance_BUSD, root_balance_BUSD.toString());

            let account3_balance_KFIVE = await kfive.balanceOf(account2, {
                from: root
            });
            eq(o.account3_balance_KFIVE, account3_balance_KFIVE.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('Account4 transfer BUSD to get 4 KFIVE. Cannot because crowdsale does not have enough KFIVE', async () => {
            const value = web3.utils.toHex(16 * (10 ** tokenDecimalsBUSD));
            await u.assertRevert(ico.buyTokens(account4, value, { from: account4, value: value }));

            let account4_balance_BUSD = await busd.balanceOf(account4, {
                from: root
            });
            eq(20 * (10 ** tokenDecimalsBUSD), account4_balance_BUSD.toString());
        });

        it('Change receiver (root) to new_receiver', async () => {
            const i = {
                new_receiver: new_receiver,
            }

            await ico.changeWallet(i.new_receiver, {
                from: root
            });
        });

        it('Account4 transfer BUSD to get 1 KFIVE', async () => {
            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await ico.buyTokens(account4, value, { from: account4, value: value });

            const o = {
                ico_balance: 10 * (10 ** tokenDecimals) - (rate * 28),
                account4_balance_KFIVE: rate * 4,
                account4_balance_BUSD: 16 * (10 ** tokenDecimalsBUSD),
                root_balance_BUSD: 24 * (10 ** tokenDecimalsBUSD),
                receiver_balance_BUSD: 4 * (10 ** tokenDecimalsBUSD),
                weiRaised: 28,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account4_balance_BUSD = await busd.balanceOf(account4, {
                from: root
            });
            eq(o.account4_balance_BUSD, account4_balance_BUSD.toString());

            let root_balance_BUSD = await busd.balanceOf(root, {
                from: root
            });
            eq(o.root_balance_BUSD, root_balance_BUSD.toString());

            let receiver_balance_BUSD = await busd.balanceOf(new_receiver, {
                from: root
            });
            eq(o.receiver_balance_BUSD, receiver_balance_BUSD.toString());

            let account4_balance_KFIVE = await kfive.balanceOf(account4, {
                from: root
            });
            eq(o.account4_balance_KFIVE, account4_balance_KFIVE.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });
    });

    describe('Pause ICO stage/ TransferOwnership', async () => {
        it('Account1: Pause ICO. Cannot because only can pause ICO', async () => {
            await u.assertRevert(ico.pause({
                from: account1
            }));
        });

        it('Owner: Pause ICO', async () => {
            await ico.pause({
                from: root
            });
        });

        it('(Account1) transfer ownership to New_receiver. Cannot because account1 is not the owner', async () => {
            await u.assertRevert(ico.transferOwnership(new_receiver, {
                from: account1
            }));
        });

        it('(Root) transfer ownership to New_receiver', async () => {
            await ico.transferOwnership(new_receiver, {
                from: root
            });
        });

        it('(Root) can not call OnlyOwner function', async () => {
            await u.assertRevert(ico.changeWallet(account1, {
                from: root
            }));

            await u.assertRevert(ico.unpause({
                from: root
            }));
        });

        it('Account5 transfer BUSD to get 1 token. Cannot because ICO is pausing', async () => {
            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await u.assertRevert(ico.buyTokens(account5, value, { from: account5, value: value }));
        });

        it('(New owner) unpause ICO contract', async () => {
            await ico.unpause({
                from: new_receiver
            });
        });

        it('(Root) pause BUSD contract again. Cannot because root is not owner anymore', async () => {
            await u.assertRevert(ico.unpause({
                from: root
            }));
        });
    });

    describe('Closing ICO stage', async () => {
        it('Delay 50s', async () => {
            await delay(50000);
        });

        it('Account5 transfer BUSD to get 1 token (1)', async () => {
            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await ico.buyTokens(account5, value, { from: account5, value: value });

            const o = {
                ico_balance: 10 * (10 ** tokenDecimals) - (rate * 32),
                account5_balance_KFIVE: rate * 4,
                account5_balance_BUSD: 6 * (10 ** tokenDecimalsBUSD),
                receiver_balance_BUSD: 8 * (10 ** tokenDecimalsBUSD),
                weiRaised: 32,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account5_balance_BUSD = await busd.balanceOf(account5, {
                from: root
            });
            eq(o.account5_balance_BUSD, account5_balance_BUSD.toString());

            let receiver_balance_BUSD = await busd.balanceOf(new_receiver, {
                from: root
            });
            eq(o.receiver_balance_BUSD, receiver_balance_BUSD.toString());

            let account5_balance_KFIVE = await kfive.balanceOf(account5, {
                from: root
            });
            eq(o.account5_balance_KFIVE, account5_balance_KFIVE.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('Account5 transfer wei to get 1 token (2). Cannot because crowdsale has closed', async () => {
            await delay(100000);

            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await u.assertRevert(ico.buyTokens(account5, value, { from: account5, value: value }));

            // Try to buy token again
            await u.assertRevert(ico.buyTokens(account5, value, { from: account5, value: value }));
        });
    });

    describe('Destroy contract', async () => {
        it('Account1 Destroy contract. Cannot because only owner can destroy', async () => {
            await u.assertRevert(ico.destroySmartContract(account7, {
                from: account1
            }));
        });

        it('(Root) destroy ICO contract. Cannot because root is not the owner', async () => {
            await u.assertRevert(ico.destroySmartContract(account8, {
                from: root
            }));
        });

        it('Owner Destroy contract', async () => {
            await ico.destroySmartContract(account7, {
                from: new_receiver
            });

            const o = {
                ico_KFIVE_balance: 0,
                account7_KFIVE_balance: 10 * (10 ** tokenDecimals) - (rate * 32),
            }

            let ico_KFIVE_balance = await kfive.balanceOf(ico.address, {from: root});
            eq(ico_KFIVE_balance.toString(), o.ico_KFIVE_balance);

            let account7_KFIVE_balance = await kfive.balanceOf(account7, {from: root});
            eq(account7_KFIVE_balance.toString(), o.account7_KFIVE_balance);
        });

        it('Owner Destroy contract again. Cannot because there is no contract', async () => {
            await u.assertRevert(ico.destroySmartContract(account7, {
                from: new_receiver
            }));
        });

        it('Account5 transfer BUSD to get 1 KFIVE. Cannot because contract has been destroyed', async () => {
            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await u.assertRevert(ico.buyTokens(account5, value, { from: account5, value: value }));
        });
    });

    describe('Pause BUSD / Destroy contract during ICO time', async () => {
        it('Start ICO crowdsale. Cannot because openingTime < currentTime', async () => {
            const soldToken = kfive.address;
            const acceptedToken = busd.address;
            const openingTime = Math.floor(new Date().getTime() / 1000 - 10);
            const closingTime = Math.floor(openingTime + 150);

            ico = await u.assertRevert(ICO.new(rate, wallet, soldToken, acceptedToken, cap, openingTime, closingTime, {
                from: root
            }));
        });

        it('Start ICO crowdsale. Cannot because closingTime < openingTime', async () => {
            const soldToken = kfive.address;
            const acceptedToken = busd.address;
            const openingTime = Math.floor(new Date().getTime() / 1000 + 5);
            const closingTime = Math.floor(openingTime + 1);

            ico = await u.assertRevert(ICO.new(rate, wallet, soldToken, acceptedToken, cap, openingTime, closingTime, {
                from: root
            }));
        });

        it('Start ICO crowdsale', async () => {
            const soldToken = kfive.address;
            const acceptedToken = busd.address;
            const openingTime = Math.floor(new Date().getTime() / 1000 + 1);
            const closingTime = Math.floor(openingTime + 150);

            ico = await ICO.new(rate, wallet, soldToken, acceptedToken, cap, openingTime, closingTime, {
                from: root
            });
        });

        it('Wait until crowdsale opens', async () => {
            await delay(1000);
        });

        it('Account1 transfer BUSD to get 1 KFIVE', async () => {
            const i = {
                to: ico.address,
                value: web3.utils.toHex(5 * (10 ** tokenDecimals)),
            }

            await kfive.issue(i.to, i.value, OFFCHAIN, {
                from: root
            });

            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));

            await busd.approve(i.to, value, {
                from: account1
            });

            await ico.buyTokens(account1, value, { from: account1, value: value });

            const o = {
                ico_balance: 5 * (10 ** tokenDecimals) - (rate * 4),
                account1_balance_KFIVE: rate * 8,
                account1_balance_BUSD: 2 * (10 ** tokenDecimalsBUSD),
                receiver_balance_BUSD: 28 * (10 ** tokenDecimalsBUSD),
                weiRaised: 4,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account1_balance_BUSD = await busd.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance_BUSD, account1_balance_BUSD.toString());

            let receiver_balance_BUSD = await busd.balanceOf(root, {
                from: root
            });
            eq(o.receiver_balance_BUSD, receiver_balance_BUSD.toString());

            let account1_balance_KFIVE = await kfive.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance_KFIVE, account1_balance_KFIVE.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('(New_receier) pause BUSD contract. Cannot because onlyOwner can pause', async () => {
            await u.assertRevert(busd.pause({
                from: new_receiver
            }));
        });

        it('(Root) pause BUSD contract', async () => {
            await busd.pause({
                from: root
            });
        });

        it('Account2 transfer BUSD to get 1 KFIVE. Cannot beacuse BUSD is pausing', async () => {
            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await u.assertRevert(ico.buyTokens(account2, value, { from: account2, value: value }));
        });

        it('(Root) unpause BUSD contract', async () => {
            await busd.unpause({
                from: root
            });
        });

        it('(Root) Destroy contract during ICO time', async () => {
            await ico.destroySmartContract(account8, {
                from: root
            });

            const o = {
                ico_token_balance: 0,
                account8_KFIVE_balance: 5 * (10 ** tokenDecimals) - (rate * 4),
            }

            let ico_token_balance = await kfive.balanceOf(ico.address, {from: root});
            eq(ico_token_balance.toString(), o.ico_token_balance);

            let account8_KFIVE_balance = await kfive.balanceOf(account8, {from: root});
            eq(account8_KFIVE_balance.toString(), o.account8_KFIVE_balance);
        });

        it('Account6 transfer BUSD to get 1 KFIVE. Cannot because contract has been destroyed', async () => {
            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await u.assertRevert(ico.buyTokens(account6, value, { from: account6, value: value }));

            const o = {
                account6_BUSD_balance: 10 * (10 ** tokenDecimalsBUSD),
            }

            let account6_BUSD_balance = await busd.balanceOf(account6, {from: root})
            eq(o.account6_BUSD_balance, account6_BUSD_balance.toString());
        });
    });

    describe('Blacklist', async () => {
        it('Start new ICO crowdsale and issue 5 KFIVE to ICO', async () => {
            const soldToken = kfive.address;
            const acceptedToken = busd.address;
            const openingTime = Math.floor(new Date().getTime() / 1000 + 1);
            const closingTime = Math.floor(openingTime + 150);

            ico = await ICO.new(rate, wallet, soldToken, acceptedToken, cap, openingTime, closingTime, {
                from: root
            });

            const i = {
                to: ico.address,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            await kfive.issue(i.to, i.value, OFFCHAIN, {
                from: root
            });

            const o = {
                ico_balance: 10 * (10 ** tokenDecimals),
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            await busd.approve(ico.address, web3.utils.toHex(16 * (10 ** tokenDecimalsBUSD)), {
                from: account4
            });

            await busd.approve(ico.address, web3.utils.toHex(10 * (10 ** tokenDecimalsBUSD)), {
                from: account6
            });
        });

        it('Wait until crowdsale opens', async () => {
            await delay(1000);
        });

        it('Root add Account6 to Blacklist', async () => {
            await busd.addBlackList(account6, {
                from: root
            });
        });

        it('Account6 transfer BUSD to get 1 KFIVE. Cannot because account6 has been added to Blacklist', async () => {
            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await u.assertRevert(ico.buyTokens(account6, value, { from: account6, value: value }));
        });

        it('Account6 approve BUSD to root. Root cannot exchange token for account6', async () => {
            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));

            await busd.approve(root, value, {
                from: account6
            });

            await u.assertRevert(ico.buyTokens(account6, value, { from: root, value: value }));
        });

        it('Root remove Account6 to Blacklist', async () => {
            await busd.removeBlackList(account6, {
                from: root
            });
        });

        it('Account6 transfer BUSD to get 1 KFIVE successfully', async () => {
            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await ico.buyTokens(account6, value, { from: account6, value: value });

            const o = {
                ico_balance: 10 * (10 ** tokenDecimals) - (rate * 4),
                account6_balance_KFIVE: rate * 4,
                account6_balance_BUSD: 6 * (10 ** tokenDecimalsBUSD),
                receiver_balance_BUSD: 32 * (10 ** tokenDecimalsBUSD),
                weiRaised: 4,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account6_balance_BUSD = await busd.balanceOf(account6, {
                from: root
            });
            eq(o.account6_balance_BUSD, account6_balance_BUSD.toString());

            let receiver_balance_BUSD = await busd.balanceOf(root, {
                from: root
            });
            eq(o.receiver_balance_BUSD, receiver_balance_BUSD.toString());

            let account6_balance_KFIVE = await kfive.balanceOf(account6, {
                from: root
            });
            eq(o.account6_balance_KFIVE, account6_balance_KFIVE.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('Root add Account6 to Blacklist. For future transaction', async () => {
            await busd.addBlackList(account6, {
                from: root
            });
        });

        it('Root change receiver to Account6 (has been added to Blacklist). Cannot because Account6 is in Blacklist', async () => {
            const i = {
                new_receiver: account6,
            }

            await u.assertRevert(ico.changeWallet(i.new_receiver, {
                from: root
            }));

            let wallet = await ico.wallet({from: root});
            eq(wallet, root);
        });

        it('Account4 transfer BUSD to get 1 KFIVE successfully. Make sure BUSD will transfer to Root', async () => {
            const value = web3.utils.toHex(4 * (10 ** tokenDecimalsBUSD));
            await ico.buyTokens(account4, value, { from: account4, value: value });

            const o = {
                ico_balance: 10 * (10 ** tokenDecimals) - (rate * 8),
                account4_balance_KFIVE: rate * 8,
                account4_balance_BUSD: 12 * (10 ** tokenDecimalsBUSD),
                account6_balance_BUSD: 6 * (10 ** tokenDecimalsBUSD),
                receiver_balance_BUSD: 36 * (10 ** tokenDecimalsBUSD),
                weiRaised: 8,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account4_balance_BUSD = await busd.balanceOf(account4, {
                from: root
            });
            eq(o.account4_balance_BUSD, account4_balance_BUSD.toString());

            let account6_balance_BUSD = await busd.balanceOf(account6, {
                from: root
            });
            eq(o.account6_balance_BUSD, account6_balance_BUSD.toString());

            let receiver_balance_BUSD = await busd.balanceOf(root, {
                from: root
            });
            eq(o.receiver_balance_BUSD, receiver_balance_BUSD.toString());

            let account4_balance_KFIVE = await kfive.balanceOf(account4, {
                from: root
            });
            eq(o.account4_balance_KFIVE, account4_balance_KFIVE.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });
    });

    describe('Presale with rate = 2631578947. 1 KFIVE = 3.8 BUSD', async () => {
        it('Redeem all the remaining token of the accounts', async () => {
            const owner = await kfive.owner({ from: root });
            const account1_balance = await kfive.balanceOf(account1, { from: root });
            const account2_balance = await kfive.balanceOf(account2, { from: root });
            const account3_balance = await kfive.balanceOf(account3, { from: root });
            const account4_balance = await kfive.balanceOf(account4, { from: root });
            const account5_balance = await kfive.balanceOf(account5, { from: root });
            const account6_balance = await kfive.balanceOf(account6, { from: root });
            const account7_balance = await kfive.balanceOf(account7, { from: root });
            const ico_balance = await kfive.balanceOf(ico.address, { from: root });
            const root_balance = await kfive.balanceOf(root, { from: root });
            const new_receiver_balance = await kfive.balanceOf(new_receiver, { from: root });

            await kfive.redeem(account1, account1_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account2, account2_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account3, account3_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account4, account4_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account5, account5_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account6, account6_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account7, account7_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(ico.address, ico_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(root, root_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(new_receiver, new_receiver_balance, OFFCHAIN, { from: owner });

            const owner_busd = await busd.owner({ from: root });
            const account1_balance_busd = await busd.balanceOf(account1, { from: root });
            const account2_balance_busd = await busd.balanceOf(account2, { from: root });
            const account3_balance_busd = await busd.balanceOf(account3, { from: root });
            const account4_balance_busd = await busd.balanceOf(account4, { from: root });
            const account5_balance_busd = await busd.balanceOf(account5, { from: root });
            const account6_balance_busd = await busd.balanceOf(account6, { from: root });
            const account7_balance_busd = await busd.balanceOf(account7, { from: root });
            const ico_balance_busd = await busd.balanceOf(ico.address, { from: root });
            const root_balance_busd = await busd.balanceOf(root, { from: root });
            const new_receiver_balance_busd = await busd.balanceOf(new_receiver, { from: root });

            await busd.redeem(account1, account1_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(account2, account2_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(account3, account3_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(account4, account4_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(account5, account5_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(account6, account6_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(account7, account7_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(root, root_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(ico.address, ico_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(new_receiver, new_receiver_balance_busd, OFFCHAIN, { from: owner_busd });
        });

        it('Start new ICO crowdsale and issue 10 KFIVE and BUSD to ICO', async () => {
            const soldToken = kfive.address;
            const acceptedToken = busd.address;
            const openingTime = Math.floor(new Date().getTime() / 1000 + 1);
            const closingTime = Math.floor(openingTime + 150);

            ico = await ICO.new(rate, wallet, soldToken, acceptedToken, cap, openingTime, closingTime, {
                from: root
            });

            const i = {
                to: ico.address,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                value_busd: web3.utils.toHex(10 * (10 ** tokenDecimalsBUSD)),
            }

            await kfive.issue(i.to, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(i.to, i.value_busd, OFFCHAIN, {
                from: root
            });

            const o = {
                ico_kfive_balance: 10 * (10 ** tokenDecimals),
                ico_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
            }

            let ico_kfive_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_kfive_balance, ico_kfive_balance.toString());

            let ico_busd_balance = await busd.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_busd_balance, ico_busd_balance.toString());
        });

        it('Issue (owner) BUSD token to Account1 and Account2', async () => {
            const i = {
                value: web3.utils.toHex(10 * (10 ** tokenDecimalsBUSD)),
            }

            await busd.issue(account1, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(account2, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(account3, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(account4, web3.utils.toHex(20 * (10 ** tokenDecimalsBUSD)), OFFCHAIN, {
                from: root
            });

            await busd.issue(account5, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(account6, i.value, OFFCHAIN, {
                from: root
            });

            const o = {
                account1_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
                account2_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
                account3_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
                account4_busd_balance: 20 * (10 ** tokenDecimalsBUSD),
                account5_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
                account6_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
            }

            let account1_busd_balance = await busd.balanceOf(account1, {
                from: root
            });
            eq(o.account1_busd_balance, account1_busd_balance.toString());

            let account2_busd_balance = await busd.balanceOf(account2, {
                from: root
            });
            eq(o.account2_busd_balance, account2_busd_balance.toString());

            let account3_busd_balance = await busd.balanceOf(account3, {
                from: root
            });
            eq(o.account3_busd_balance, account3_busd_balance.toString());

            let account4_busd_balance = await busd.balanceOf(account4, {
                from: root
            });
            eq(o.account4_busd_balance, account4_busd_balance.toString());

            let account5_busd_balance = await busd.balanceOf(account5, {
                from: root
            });
            eq(o.account5_busd_balance, account5_busd_balance.toString());

            let account6_busd_balance = await busd.balanceOf(account6, {
                from: root
            });
            eq(o.account6_busd_balance, account6_busd_balance.toString());
        });

        it('Approve BUSD to crowdsale', async () => {
            const i = {
                to: ico.address,
                value: web3.utils.toHex(10 * (10 ** tokenDecimalsBUSD)),
            }

            await busd.approve(i.to, i.value, {
                from: account1
            });

            await busd.approve(i.to, i.value, {
                from: account2
            });

            await busd.approve(i.to, i.value, {
                from: account3
            });

            await busd.approve(i.to, web3.utils.toHex(20 * (10 ** tokenDecimalsBUSD)), {
                from: account4
            });

            await busd.approve(i.to, i.value, {
                from: account5
            });

            await busd.approve(i.to, i.value, {
                from: account6
            });
        });

        it('Wait until crowdsale opens', async () => {
            await delay(1000);
        });

        it('Account1 transfer 4.3 BUSD to get 1 token', async () => {
            const value = web3.utils.toHex(4.3 * (10 ** tokenDecimalsBUSD));
            await ico.buyTokens(account1, value, { from: account1, value: value });

            const o = {
                ico_balance: 10 * (10 ** tokenDecimals) - (rate * 4),
                account1_balance_KFIVE: rate * 4,
                account1_balance_BUSD: 6 * (10 ** tokenDecimalsBUSD),
                root_balance_BUSD: 4 * (10 ** tokenDecimalsBUSD),
                weiRaised: 4,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account1_balance_BUSD = await busd.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance_BUSD, account1_balance_BUSD.toString());

            let root_balance_BUSD = await busd.balanceOf(root, {
                from: root
            });
            eq(o.root_balance_BUSD, root_balance_BUSD.toString());

            let account1_balance_KFIVE = await kfive.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance_KFIVE, account1_balance_KFIVE.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('Owner Destroy contract', async () => {
            await ico.destroySmartContract(account7, {
                from: root
            });

            const o = {
                ico_KFIVE_balance: 0,
                ico_BUSD_balance: 0,
                account7_KFIVE_balance: 10 * (10 ** tokenDecimals) - (rate * 4),
                account7_BUSD_balance: 10 * (10 ** tokenDecimalsBUSD),
            }

            let ico_KFIVE_balance = await kfive.balanceOf(ico.address, {from: root});
            eq(ico_KFIVE_balance.toString(), o.ico_KFIVE_balance);

            let ico_BUSD_balance = await busd.balanceOf(ico.address, {from: root});
            eq(ico_BUSD_balance.toString(), o.ico_BUSD_balance);

            let account7_KFIVE_balance = await kfive.balanceOf(account7, {from: root});
            eq(account7_KFIVE_balance.toString(), o.account7_KFIVE_balance);

            let account7_BUSD_balance = await busd.balanceOf(account7, {from: root});
            eq(account7_BUSD_balance.toString(), o.account7_BUSD_balance);
        });
    });

    describe('Presale with rate = 2631578947. 1 KFIVE = 3.8 BUSD. changeWallet Stage', async () => {
        it('Redeem all the remaining token of the accounts', async () => {
            const owner = await kfive.owner({ from: root });
            const account1_balance = await kfive.balanceOf(account1, { from: root });
            const account2_balance = await kfive.balanceOf(account2, { from: root });
            const account3_balance = await kfive.balanceOf(account3, { from: root });
            const account4_balance = await kfive.balanceOf(account4, { from: root });
            const account5_balance = await kfive.balanceOf(account5, { from: root });
            const account6_balance = await kfive.balanceOf(account6, { from: root });
            const account7_balance = await kfive.balanceOf(account7, { from: root });
            const ico_balance = await kfive.balanceOf(ico.address, { from: root });
            const root_balance = await kfive.balanceOf(root, { from: root });
            const new_receiver_balance = await kfive.balanceOf(new_receiver, { from: root });

            await kfive.redeem(account1, account1_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account2, account2_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account3, account3_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account4, account4_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account5, account5_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account6, account6_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account7, account7_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(ico.address, ico_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(root, root_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(new_receiver, new_receiver_balance, OFFCHAIN, { from: owner });

            const owner_busd = await busd.owner({ from: root });
            const account1_balance_busd = await busd.balanceOf(account1, { from: root });
            const account2_balance_busd = await busd.balanceOf(account2, { from: root });
            const account3_balance_busd = await busd.balanceOf(account3, { from: root });
            const account4_balance_busd = await busd.balanceOf(account4, { from: root });
            const account5_balance_busd = await busd.balanceOf(account5, { from: root });
            const account6_balance_busd = await busd.balanceOf(account6, { from: root });
            const ico_balance_busd = await busd.balanceOf(ico.address, { from: root });
            const root_balance_busd = await busd.balanceOf(root, { from: root });
            const new_receiver_balance_busd = await busd.balanceOf(new_receiver, { from: root });

            await busd.redeem(account1, account1_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(account2, account2_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(account3, account3_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(account4, account4_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(account5, account5_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(account6, account6_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(root, root_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(ico.address, ico_balance_busd, OFFCHAIN, { from: owner_busd });
            await busd.redeem(new_receiver, new_receiver_balance_busd, OFFCHAIN, { from: owner_busd });
        });

        it('Start new ICO crowdsale and issue 10 KFIVE to ICO', async () => {
            const soldToken = kfive.address;
            const acceptedToken = busd.address;
            const openingTime = Math.floor(new Date().getTime() / 1000 + 1);
            const closingTime = Math.floor(openingTime + 150);

            ico = await ICO.new(rate, wallet, soldToken, acceptedToken, cap, openingTime, closingTime, {
                from: root
            });

            const i = {
                to: ico.address,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            await kfive.issue(i.to, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(i.to, i.value, OFFCHAIN, {
                from: root
            });

            const o = {
                ico_kfive_balance: 10 * (10 ** tokenDecimals),
                ico_busd_balance: 10 * (10 ** tokenDecimals)
            }

            let ico_kfive_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_kfive_balance, ico_kfive_balance.toString());

            let ico_busd_balance = await busd.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_busd_balance, ico_busd_balance.toString());

            await busd.approve(ico.address, web3.utils.toHex(16 * (10 ** tokenDecimalsBUSD)), {
                from: account4
            });

            await busd.approve(ico.address, web3.utils.toHex(10 * (10 ** tokenDecimalsBUSD)), {
                from: account6
            });
        });

        it('Issue (owner) BUSB token to Account1 and Account2', async () => {
            const i = {
                value: web3.utils.toHex(10 * (10 ** tokenDecimalsBUSD)),
            }

            await busd.issue(account1, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(account2, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(account3, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(account4, web3.utils.toHex(20 * (10 ** tokenDecimalsBUSD)), OFFCHAIN, {
                from: root
            });

            await busd.issue(account5, i.value, OFFCHAIN, {
                from: root
            });

            await busd.issue(account6, i.value, OFFCHAIN, {
                from: root
            });

            const o = {
                account1_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
                account2_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
                account3_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
                account4_busd_balance: 20 * (10 ** tokenDecimalsBUSD),
                account5_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
                account6_busd_balance: 10 * (10 ** tokenDecimalsBUSD),
            }

            let account1_busd_balance = await busd.balanceOf(account1, {
                from: root
            });
            eq(o.account1_busd_balance, account1_busd_balance.toString());

            let account2_busd_balance = await busd.balanceOf(account2, {
                from: root
            });
            eq(o.account2_busd_balance, account2_busd_balance.toString());

            let account3_busd_balance = await busd.balanceOf(account3, {
                from: root
            });
            eq(o.account3_busd_balance, account3_busd_balance.toString());

            let account4_busd_balance = await busd.balanceOf(account4, {
                from: root
            });
            eq(o.account4_busd_balance, account4_busd_balance.toString());

            let account5_busd_balance = await busd.balanceOf(account5, {
                from: root
            });
            eq(o.account5_busd_balance, account5_busd_balance.toString());

            let account6_busd_balance = await busd.balanceOf(account6, {
                from: root
            });
            eq(o.account6_busd_balance, account6_busd_balance.toString());
        });

        it('Approve BUSD to crowdsale', async () => {
            const i = {
                to: ico.address,
                value: web3.utils.toHex(10 * (10 ** tokenDecimalsBUSD)),
            }

            await busd.approve(i.to, i.value, {
                from: account1
            });

            await busd.approve(i.to, i.value, {
                from: account2
            });

            await busd.approve(i.to, i.value, {
                from: account3
            });

            await busd.approve(i.to, web3.utils.toHex(20 * (10 ** tokenDecimalsBUSD)), {
                from: account4
            });

            await busd.approve(i.to, i.value, {
                from: account5
            });

            await busd.approve(i.to, i.value, {
                from: account6
            });
        });

        it('Wait until crowdsale opens', async () => {
            await delay(1000);
        });

        it('Account1 transfer 4.3 BUSD to get 1 token. Root (wallet) will receive BUSD', async () => {
            const value = web3.utils.toHex(4.3 * (10 ** tokenDecimalsBUSD));
            await ico.buyTokens(account1, value, { from: account1, value: value });

            const o = {
                ico_balance: 10 * (10 ** tokenDecimals) - (rate * 4),
                account1_balance_KFIVE: rate * 4,
                account1_balance_BUSD: 6 * (10 ** tokenDecimalsBUSD),
                root_balance_BUSD: 4 * (10 ** tokenDecimalsBUSD),
                weiRaised: 4,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account1_balance_BUSD = await busd.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance_BUSD, account1_balance_BUSD.toString());

            let root_balance_BUSD = await busd.balanceOf(root, {
                from: root
            });
            eq(o.root_balance_BUSD, root_balance_BUSD.toString());

            let account1_balance_KFIVE = await kfive.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance_KFIVE, account1_balance_KFIVE.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('Change receiver (root) to new_receiver', async () => {
            const i = {
                new_receiver: new_receiver,
            }

            await ico.changeWallet(i.new_receiver, {
                from: root
            });
        });

        it('Account2 transfer 7.6 BUSD to get 2 KFIVE. New_receiver (wallet) will receive BUSD', async () => {
            const value = web3.utils.toHex(7.6 * (10 ** tokenDecimalsBUSD));
            await ico.buyTokens(account2, value, { from: account2, value: value });

            const o = {
                ico_balance: 10 * (10 ** tokenDecimals) - (rate * 11),
                account2_balance_KFIVE: rate * 7,
                account2_balance_BUSD: 3 * (10 ** tokenDecimalsBUSD),
                root_balance_BUSD: 4 * (10 ** tokenDecimalsBUSD),
                new_receiver_balance_BUSD: 7 * (10 ** tokenDecimalsBUSD),
                weiRaised: 11,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account2_balance_BUSD = await busd.balanceOf(account2, {
                from: root
            });
            eq(o.account2_balance_BUSD, account2_balance_BUSD.toString());

            let root_balance_BUSD = await busd.balanceOf(root, {
                from: root
            });
            eq(o.root_balance_BUSD, root_balance_BUSD.toString());

            let new_receiver_balance_BUSD = await busd.balanceOf(new_receiver, {
                from: root
            });
            eq(o.new_receiver_balance_BUSD, new_receiver_balance_BUSD.toString());

            let account2_balance_KFIVE = await kfive.balanceOf(account2, {
                from: root
            });
            eq(o.account2_balance_KFIVE, account2_balance_KFIVE.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('Owner Destroy contract', async () => {
            await ico.destroySmartContract(account7, {
                from: root
            });

            const o = {
                ico_KFIVE_balance: 0,
                ico_BUSD_balance: 0,
                account7_KFIVE_balance: 10 * (10 ** tokenDecimals) - (rate * 11),
            }

            let ico_KFIVE_balance = await kfive.balanceOf(ico.address, {from: root});
            eq(ico_KFIVE_balance.toString(), o.ico_KFIVE_balance);

            let ico_BUSD_balance = await busd.balanceOf(ico.address, {from: root});
            eq(ico_BUSD_balance.toString(), o.ico_BUSD_balance);

            let account7_KFIVE_balance = await kfive.balanceOf(account7, {from: root});
            eq(account7_KFIVE_balance.toString(), o.account7_KFIVE_balance);
        });
    });
});