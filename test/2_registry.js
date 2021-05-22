const TokenRegistry = artifacts.require('TokenRegistry');
const Tether = artifacts.require("BEP20USDT")
const Tokoin = artifacts.require("TokoinToken")

const eq = assert.equal
const u = require('./util.js')
let token_registry, tokoin, usdt;

contract('TokenRegistry', (accounts) => {

    const root = accounts[0]
    const account1 = accounts[1]
    const admin1 = accounts[4]
    const OFFCHAIN = web3.utils.fromAscii('1')

    before(async () => {

        token_registry = await TokenRegistry.new({
            from: root
        });

        usdt = await Tether.new({
            from: root
        });

        tokoin = await Tokoin.new({
            from: root
        });

    })


    describe('Add new token', async () => {

        it('add more admin', async () => {

            const i = {
                admin: admin1
            }

            await token_registry.addAdmin(i.admin, {
                from: root
            });

        })

        it('only admin', async () => {

            const i = {
                token: usdt.address,
                symbol: 'USDT',
                name: 'Tether USD',
                decimals: 6,
            }

            await u.assertRevert(token_registry.addNewToken(i.token, i.symbol, i.name, i.decimals, OFFCHAIN, {
                from: account1
            }));

            await token_registry.addNewToken(i.token, i.symbol, i.name, i.decimals, OFFCHAIN, {
                from: admin1
            });

        })

    })

    describe('Check token is existed or not', async () => {

        it('tokoin', async () => {

            const i = {
                token: tokoin.address,
            }

            const o = {
                result: false
            }

            let tx = await token_registry.tokenIsExisted(i.token, {
                from: account1
            });

            eq(o.result, tx);

        })

        it('usdt', async () => {

            const i = {
                token: usdt.address,
            }

            const o = {
                result: true
            }

            let tx = await token_registry.tokenIsExisted(i.token, {
                from: account1
            });

            eq(o.result, tx);

        })

    })

    describe('Get token information', async () => {

        it('tokoin', async () => {

            const i = {
                token: tokoin.address,
            }

            const o = {
                symbol: '',
                name: '',
                decimals: 0,
            }

            let tx = await token_registry.getTokenByAddr(i.token, {
                from: account1
            });

            eq(o.symbol, tx[0]);
            eq(o.name, tx[1]);
            eq(o.decimals, tx[2]);

        })

        it('usdt', async () => {

            const i = {
                token: usdt.address,
            }

            const o = {
                symbol: 'USDT',
                name: 'Tether USD',
                decimals: 6,
            }

            let tx = await token_registry.getTokenByAddr(i.token, {
                from: account1
            });

            eq(o.symbol, tx[0]);
            eq(o.name, tx[1]);
            eq(o.decimals, tx[2]);

        })

    })

})
