const ICO = artifacts.require("KFIVECrowdsale")
const KFIVE = artifacts.require("KFIVE")

const eq = assert.equal
const u = require('./utils.js')
var kfive;

const tokenCap = 1080000;
const tokenDecimals = 10;

contract("NLGinseng", (accounts) => {
    const root = accounts[9]
    const account1 = accounts[8]
    const account2 = accounts[7]
    const account3 = accounts[6]
    const account4 = accounts[5]
    const account5 = accounts[4]
    const account6 = accounts[3]
    const account7 = accounts[2]
    const account8 = accounts[1]
    const account9 = accounts[0]
    const OFFCHAIN = web3.utils.fromAscii('1')

    const rate = 1;
    const wallet = root;
    const cap = 540000000000000;
    const openingTime = Math.floor(new Date().getTime() / 1000 + 60);
    const closingTime = Math.floor(openingTime + 150);
    const delay = ms => new Promise(res => setTimeout(res, ms));

    before(async () => {
        kfive = await KFIVE.new({
            from: root
        });

        const token = kfive.address;

        ico = await ICO.new(rate, wallet, token, cap, openingTime, closingTime, {
            from: root
        });
    });

    describe('Deployment stage', async () => {
        it('Token information', async () => {
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

        it('ICO information', async () => {
            let o = {
                token: kfive.address
            }

            let saleToken = await ico.token({
                from: root
            });

            eq(o.token, saleToken);
        });
    });

    describe('Pre ICO stage', async () => {
        it('Issue (owner) 54.000 KFIVE token to Crowdsale contract', async () => {
            const i = {
                to: ico.address,
                value: web3.utils.toHex((5/100 * tokenCap) * (10 ** tokenDecimals)),
            }

            await kfive.issue(i.to, i.value, OFFCHAIN, {
                from: root
            });

            const o = {
                ico_balance: 54000 * (10 ** tokenDecimals),
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());
        });
    });

    describe('Running ICO stage', async () => {
        it('Account1 transfer wei to get 1 token. Cannot because crowdsale has not started', async () => {
            await u.assertRevert(ico.sendTransaction({from: account1, value: '10000000000'}));
        });

        it('Account1 transfer wei to get 10.000 token', async () => {
            await delay(60000);
            await ico.sendTransaction({from: account1, value: '100000000000000'});

            const o = {
                ico_balance: 44000 * (10 ** tokenDecimals),
                account1_balance: 10000 * (10 ** tokenDecimals),
                weiRaised: 100000000000000,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account1_balance = await kfive.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance, account1_balance.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('Account2 transfer wei to get 50.000 token. Cannot because crowdsale only has 44.000 left', async () => {
            await u.assertRevert(ico.sendTransaction({from: account2, value: '500000000000000'}));
        });

        it('Account2 transfer wei to get 10.000 token', async () => {
            await ico.sendTransaction({from: account2, value: '100000000000000'});

            const o = {
                ico_balance: 34000 * (10 ** tokenDecimals),
                account2_balance: 10000 * (10 ** tokenDecimals),
                weiRaised: 200000000000000,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account2_balance = await kfive.balanceOf(account2, {
                from: root
            });
            eq(o.account2_balance, account2_balance.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('Root transfer wei to get 10.000 token', async () => {
            await ico.sendTransaction({from: root, value: '100000000000000'});

            const o = {
                ico_balance: 24000 * (10 ** tokenDecimals),
                root_balance: 10000 * (10 ** tokenDecimals),
                weiRaised: 300000000000000,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let root_balance = await kfive.balanceOf(account2, {
                from: root
            });
            eq(o.root_balance, root_balance.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('Account3 transfer wei to get 20.000 token', async () => {
            await ico.sendTransaction({from: account3, value: '200000000000000'});

            const o = {
                ico_balance: 40000000000000,
                account3_balance: 20000 * (10 ** tokenDecimals),
                weiRaised: 500000000000000,
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account3_balance = await kfive.balanceOf(account3, {
                from: root
            });
            eq(o.account3_balance, account3_balance.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());
        });

        it('Account4 transfer wei to get 5.000 token. Cannot because crowdsale does not have enough KFIVE', async () => {
            await u.assertRevert(ico.sendTransaction({from: account4, value: '50000000000000'}));
        });
    });

    describe('Closing ICO stage', async () => {
        it('Account4 transfer wei to get 4.000 KFIVE. Cannot because the crowdsale has been closed', async () => {
            await delay(100000);
            await u.assertRevert(ico.sendTransaction({from: account4, value: '40000000000000'}));
        });
    });
});