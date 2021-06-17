const NLGinseng = artifacts.require("NLGinseng")

const eq = assert.equal
const u = require('./utils.js')
var nlgst;

contract("NLGinseng", (accounts) => {

    const root = accounts[0]
    const account1 = accounts[1]
    const account2 = accounts[2]
    const evil = accounts[3]
    const new_owner = accounts[4]
    const account5 = accounts[5]
    const account6 = accounts[6]
    const new_admin = accounts[7]
    const account8 = accounts[8]
    const account9 = accounts[9]
    const OFFCHAIN = web3.utils.fromAscii('1')

    const baseURL = "https://ipfs.io/ipfs/";

    before(async () => {
        nlgst = await NLGinseng.new(baseURL, {
            from: root
        });
    });

    describe('Deployment stage', async () => {
        it('Token information', async () => {
            let o = {
                balance: 0,
                name: "Ngoc Linh Ginseng Token",
                symbol: "NLGT"
            }

            let owner_balance = await nlgst.balanceOf(root, {
                from: root
            });

            eq(o.balance, owner_balance.toString());
            eq(o.name, await nlgst.name.call());
            eq(o.symbol, await nlgst.symbol.call());
        });
    })
});
