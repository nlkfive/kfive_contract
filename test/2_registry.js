const TokenRegistry = artifacts.require('TokenRegistry');
const KFIVE = artifacts.require("KFIVE")

const eq = assert.equal
const u = require('./utils.js')
let token_registry, kfive;

contract('TokenRegistry', (accounts) => {

    const root = accounts[0]
    const account1 = accounts[1]
    const admin1 = accounts[4]
    const OFFCHAIN = web3.utils.fromAscii('1')

    before(async () => {
        token_registry = await TokenRegistry.new({
            from: root
        });

        kfive = await KFIVE.new({
            from: root
        });
    })


    describe('Add new token', async () => {
        it('Add an admin', async () => {
            const i = {
                admin: admin1
            }

            await token_registry.addAdmin(i.admin, {
                from: root
            });
        })

        it('Register new token: Only admin', async () => {
            const i = {
                token: kfive.address,
                symbol: 'KFIVE',
                name: 'KFIVE',
                decimals: 10,
            }

            await u.assertRevert(token_registry.addNewToken(i.token, i.symbol, i.name, i.decimals, OFFCHAIN, {
                from: account1
            }));

            await token_registry.addNewToken(i.token, i.symbol, i.name, i.decimals, OFFCHAIN, {
                from: admin1
            });
        })
    })

    describe('Check if token is existed', async () => {
        it('KFIVE token', async () => {
            const i = {
                token: kfive.address,
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
        it('KFIVE', async () => {
            const i = {
                token: kfive.address,
            }

            const o = {
                symbol: 'KFIVE',
                name: 'KFIVE',
                decimals: 10,
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
