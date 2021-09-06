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
    const new_receiver = accounts[4]
    const account5 = accounts[3]
    const account6 = accounts[2]
    const account7 = accounts[1]
    const account8 = accounts[0]
    const OFFCHAIN = web3.utils.fromAscii('1')

    const rate = 1;
    const wallet = root;
    const cap = 54000 * (10 ** tokenDecimals);
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
                token: kfive.address,
                wallet: root,
                wei: 0,
                rate: 1,
            }

            let saleToken = await ico.token({
                from: root
            });
            eq(o.token, saleToken);

            let receiver = await ico.wallet({
                from: root
            });
            eq(o.wallet, receiver);

            let initial_wei = await ico.weiRaised({
                from: root
            });
            eq(o.wei, initial_wei);

            let rate = await ico.rate({
                from: root
            })
            eq(o.rate, rate);
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

    describe('Running ICO stage. 1 BNB = 0.01 KFIVE', async () => {
        it('Account1 transfer 10^10 wei to get 1 token. Cannot because crowdsale is not opening', async () => {
            const value = web3.utils.toWei(web3.utils.toBN('10'), 'nanoether');
            await u.assertRevert(ico.sendTransaction({from: account1, value: value}));
        });

        it('Wait until crowdsale opens', async () => {
            await delay(60000);
        });

        it('Account1 transfer wei to get 10.000 token', async () => {
            // Check receiver and account1 balance before transaction
            let receiver_eth_balance_before = await web3.eth.getBalance(wallet);
            let account1_eth_balance_before = await web3.eth.getBalance(account1);

            const value = web3.utils.toWei(web3.utils.toBN('100000'), 'nanoether');
            const transaction = await ico.sendTransaction({from: account1, value: value});

            const gasPrice = await web3.eth.getGasPrice();
            const gasUsed = transaction.receipt.gasUsed;
            const gasCost = web3.utils.toBN(gasPrice).mul(web3.utils.toBN(gasUsed));

            const o = {
                ico_balance: 44000 * (10 ** tokenDecimals),
                account1_balance: 10000 * (10 ** tokenDecimals),
                weiRaised: 10000 * (10 ** tokenDecimals),
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

            // Check receiver and account1 after transaction
            let receiver_eth_balance_after = await web3.eth.getBalance(wallet);
            eq(web3.utils.toBN(receiver_eth_balance_before).add(web3.utils.toBN(value)).toString(), receiver_eth_balance_after.toString());

            let account1_eth_balance_after = await web3.eth.getBalance(account1);
            eq((web3.utils.toBN(account1_eth_balance_before)).sub(web3.utils.toBN(value)).sub(web3.utils.toBN(gasCost)).toString(), account1_eth_balance_after.toString());
        });

        it('Account2 transfer wei to get 50.000 token. Cannot because crowdsale only has 44.000 left', async () => {
            const value = web3.utils.toWei(web3.utils.toBN('500000'), 'nanoether');
            await u.assertRevert(ico.sendTransaction({from: account2, value: value}));
        });

        it('Account2 transfer wei to get 10.000 token', async () => {
            let receiver_eth_balance_before = await web3.eth.getBalance(wallet);
            let account2_eth_balance_before = await web3.eth.getBalance(account2);

            const value = web3.utils.toWei(web3.utils.toBN('100000'), 'nanoether');
            const transaction = await ico.sendTransaction({from: account2, value: value});

            const gasPrice = await web3.eth.getGasPrice();
            const gasUsed = transaction.receipt.gasUsed;
            const gasCost = web3.utils.toBN(gasPrice).mul(web3.utils.toBN(gasUsed));

            const o = {
                ico_balance: 34000 * (10 ** tokenDecimals),
                account2_balance: 10000 * (10 ** tokenDecimals),
                weiRaised: 20000 * (10 ** tokenDecimals),
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

            let receiver_eth_balance_after = await web3.eth.getBalance(wallet);
            eq(web3.utils.toBN(receiver_eth_balance_before).add(web3.utils.toBN(value)).toString(), receiver_eth_balance_after.toString());

            let account2_eth_balance_after = await web3.eth.getBalance(account2);
            eq((web3.utils.toBN(account2_eth_balance_before)).sub(web3.utils.toBN(value)).sub(web3.utils.toBN(gasCost)).toString(), account2_eth_balance_after.toString());
        });

        it('Root transfer wei to get 10.000 token', async () => {
            let receiver_eth_balance_before = await web3.eth.getBalance(wallet);

            const value = web3.utils.toWei(web3.utils.toBN('100000'), 'nanoether');
            const transaction = await ico.sendTransaction({from: root, value: value});

            const gasPrice = await web3.eth.getGasPrice();
            const gasUsed = transaction.receipt.gasUsed;
            const gasCost = web3.utils.toBN(gasPrice).mul(web3.utils.toBN(gasUsed));

            const o = {
                ico_balance: 24000 * (10 ** tokenDecimals),
                root_balance: 10000 * (10 ** tokenDecimals),
                weiRaised: 30000 * (10 ** tokenDecimals),
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

            let receiver_eth_balance_after = await web3.eth.getBalance(wallet);
            eq((web3.utils.toBN(receiver_eth_balance_before)).sub(web3.utils.toBN(gasCost)).toString(), receiver_eth_balance_after);
        });

        it('Account3 transfer wei to get 20.000 token', async () => {
            let receiver_eth_balance_before = await web3.eth.getBalance(wallet);
            let account3_eth_balance_before = await web3.eth.getBalance(account3);

            const value = web3.utils.toWei(web3.utils.toBN('200000'), 'nanoether');
            const transaction = await ico.sendTransaction({from: account3, value: value});

            const gasPrice = await web3.eth.getGasPrice();
            const gasUsed = transaction.receipt.gasUsed;
            const gasCost = web3.utils.toBN(gasPrice).mul(web3.utils.toBN(gasUsed));

            const o = {
                ico_balance: 4000 * (10 ** tokenDecimals),
                account3_balance: 20000 * (10 ** tokenDecimals),
                weiRaised: 50000 * (10 ** tokenDecimals),
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

            let receiver_eth_balance_after = await web3.eth.getBalance(wallet);
            eq(web3.utils.toBN(receiver_eth_balance_before).add(web3.utils.toBN(value)).toString(), receiver_eth_balance_after.toString());

            let account3_eth_balance_after = await web3.eth.getBalance(account3);
            eq((web3.utils.toBN(account3_eth_balance_before)).sub(web3.utils.toBN(value)).sub(web3.utils.toBN(gasCost)).toString(), account3_eth_balance_after.toString());
        });

        it('Account4 transfer wei to get 5.000 token. Cannot because crowdsale does not have enough KFIVE', async () => {
            const value = web3.utils.toWei(web3.utils.toBN('50000'), 'nanoether');
            await u.assertRevert(ico.sendTransaction({from: account4, value: value}));
        });

        it('Change receiver (root) to new_receiver', async() => {
            const i = {
                new_receiver: new_receiver,
            }

            await ico.changeWallet(i.new_receiver, {
                from: root
            });
        });

        it('Account4 transfer wei to get 1.000 token', async () => {
            let receiver_eth_balance_before = await web3.eth.getBalance(new_receiver);
            let account4_eth_balance_before = await web3.eth.getBalance(account4);

            const value = web3.utils.toWei(web3.utils.toBN('10000'), 'nanoether');
            const transaction = await ico.sendTransaction({from: account4, value: value});

            const gasPrice = await web3.eth.getGasPrice();
            const gasUsed = transaction.receipt.gasUsed;
            const gasCost = web3.utils.toBN(gasPrice).mul(web3.utils.toBN(gasUsed));

            const o = {
                ico_balance: 3000 * (10 ** tokenDecimals),
                account4_balance: 1000 * (10 ** tokenDecimals),
                weiRaised: 51000 * (10 ** tokenDecimals),
            }

            let ico_balance = await kfive.balanceOf(ico.address, {
                from: root
            });
            eq(o.ico_balance, ico_balance.toString());

            let account4_balance = await kfive.balanceOf(account4, {
                from: root
            });
            eq(o.account4_balance, account4_balance.toString());

            let weiRaised = await ico.weiRaised({
                from: root
            });
            eq(o.weiRaised, weiRaised.toString());

            let receiver_eth_balance_after = await web3.eth.getBalance(new_receiver);
            eq(web3.utils.toBN(receiver_eth_balance_before).add(web3.utils.toBN(value)).toString(), receiver_eth_balance_after.toString());

            let account4_eth_balance_after = await web3.eth.getBalance(account4);
            eq((web3.utils.toBN(account4_eth_balance_before)).sub(web3.utils.toBN(value)).sub(web3.utils.toBN(gasCost)).toString(), account4_eth_balance_after.toString());
        });
    });

    describe('Closing ICO stage', async () => {
        it('Delay 80s', async () => {
            await delay(80000);
        });

        it('Account5 transfer wei to get 1.000 token (1)', async () => {
            const value = web3.utils.toWei(web3.utils.toBN('10000'), 'nanoether');
            await ico.sendTransaction({from: account5, value: value});
        });

        it('Account5 transfer wei to get 1.000 token (2). Cannot because crowdsale has closed', async () => {
            await delay(10000);
            const value = web3.utils.toWei(web3.utils.toBN('10000'), 'nanoether');
            await u.assertRevert(ico.sendTransaction({from: account5, value: value}));

            // Try to buy token again
            await u.assertRevert(ico.sendTransaction({from: account5, value: value}));
        });
    });
});