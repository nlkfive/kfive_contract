const NLGinseng = artifacts.require("NLGinseng")

const eq = assert.equal
const u = require('./utils.js')
const keccak256 = require('js-sha3').keccak256;

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
                symbol: "NLGST"
            }

            let owner_balance = await nlgst.balanceOf(root, {
                from: root
            });

            eq(o.balance, owner_balance.toString());
            eq(o.name, await nlgst.name.call());
            eq(o.symbol, await nlgst.symbol.call());
        });
    });

    describe('Mint stage', async () => {
        it('Mint new token (account1): Failed because account1 does not have minter_role', async () => {
            await u.assertRevert(nlgst.mint(account1, 1, "1", {
                from: account1
            }));
        });

        it('Mint new token (owner) to [1] account1', async () => {
            const i = {
                tokenId: 1,
                tokenURI: "1"
            }

            await nlgst.mint(account1, i.tokenId, i.tokenURI, {
                from: root
            });

            let o = {
                account1_balance: 1,
                tokenURI: baseURL + i.tokenURI,
                owner: account1
            }
            let account1_balance = await nlgst.balanceOf(account1, {
                from: account1
            });
            eq(o.account1_balance, account1_balance.toString());

            let tokenURI = await nlgst.tokenURI(i.tokenId, {
                from: account1
            });
            eq(o.tokenURI, tokenURI);

            let owner = await nlgst.ownerOf(i.tokenId, {
                from: root
            });
            eq(o.owner, owner);
        });

        it('Give minter_role to account1 (account1): Failed because account1 does not have admin role for grant this role', async () => {
            await u.assertRevert(nlgst.grantRole("0x" + keccak256("MINTER_ROLE"), account1, {
                from: account1
            }));
        });

        it('Give minter_role to account1 (root)', async () => {
            await nlgst.grantRole("0x" + keccak256("MINTER_ROLE"), account1, {
                from: root
            });
        });

        it('Mint new token (account1) to [2] account1', async () => {
            const i = {
                tokenId: 2,
                tokenURI: "2"
            }

            await nlgst.mint(account1, i.tokenId, i.tokenURI, {
                from: account1
            });

            let o = {
                account1_balance: 2,
                tokenURI: baseURL + i.tokenURI,
                owner: account1
            }
            let account1_balance = await nlgst.balanceOf(account1, {
                from: account1
            });
            eq(o.account1_balance, account1_balance.toString());

            let tokenURI = await nlgst.tokenURI(i.tokenId, {
                from: account1
            });
            eq(o.tokenURI, tokenURI);

            let owner = await nlgst.ownerOf(i.tokenId, {
                from: root
            });
            eq(o.owner, owner);
        });

        it('Revoke minter_role to account1 (root)', async () => {
            await nlgst.revokeRole("0x" + keccak256("MINTER_ROLE"), account1, {
                from: root
            });
        });

        it('Mint new token (account1): Failed because account1 does not have minter_role', async () => {
            await u.assertRevert(nlgst.mint(account1, 1, "1", {
                from: account1
            }));
        });
    });

    describe('Transfer stage', async () => {
        it('Transfer token (account1) [2] to account2', async () => {
            const i = {
                tokenId: 2,
                from: account1,
                to: account2
            }

            await nlgst.transferFrom(i.from, i.to, i.tokenId, {
                from: i.from
            });

            let o = {
                account1_balance: 1,
                account2_balance: 1,
                owner: account2
            }
            let account1_balance = await nlgst.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance, account1_balance.toString());

            let account2_balance = await nlgst.balanceOf(account2, {
                from: root
            });
            eq(o.account2_balance, account2_balance.toString());

            let owner = await nlgst.ownerOf(i.tokenId, {
                from: root
            });
            eq(o.owner, owner);
        });

        it('Transfer (root): token [1] from account1 to account2: Failed because root does not own this token', async () => {
            const i = {
                tokenId: 1,
                from: account1,
                to: account2
            }

            await u.assertRevert(nlgst.transferFrom(i.from, i.to, i.tokenId, {
                from: root
            }));
        });

        it('Approve (account1): token [1] to root for future transaction', async () => {
            const i = {
                tokenId: 1,
                to: root
            }

            await nlgst.approve(i.to, i.tokenId, {
                from: account1
            });
        });

        it('Transfer (root): token [1] from account1 to account2', async () => {
            const i = {
                tokenId: 1,
                from: account1,
                to: account2
            }

            await nlgst.transferFrom(i.from, i.to, i.tokenId, {
                from: root
            });

            let o = {
                account1_balance: 0,
                account2_balance: 2,
                owner: account2
            }
            let account1_balance = await nlgst.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance, account1_balance.toString());

            let account2_balance = await nlgst.balanceOf(account2, {
                from: root
            });
            eq(o.account2_balance, account2_balance.toString());

            let owner = await nlgst.ownerOf(i.tokenId, {
                from: root
            });
            eq(o.owner, owner);
        });

        it('Transfer (root): token [1] from account2 to account1', async () => {
            const i = {
                tokenId: 1,
                from: account2,
                to: account1
            }

            await u.assertRevert(nlgst.safeTransferFrom(i.from, i.to, i.tokenId, {
                from: root
            }));
        });
    });

    describe('Pausable stage', async () => {
        it('Pause: Failed because ony owner can call it', async () => {
            await u.assertRevert(nlgst.pause({
                from: account1
            }));
        })

        it('Pause: Onwer call', async () => {
            await nlgst.pause({
                from: root
            });
        });

        it('Pause: Failed because cannot pause while already paused', async () => {
            await u.assertRevert(nlgst.pause({
                from: root
            }));
        })

        it('TransferFrom: Failed because cannot transfer while pausing', async () => {
            const i = {
                from: account2,
                to: account1,
                tokenId: 2,
            }

            await u.assertRevert(nlgst.transferFrom(i.from, i.to, i.tokenId, {
                from: root
            }));
        });

        it('Unpause (account1): Failed because only owner can call it', async () => {
            await u.assertRevert(nlgst.unpause({
                from: account1
            }));
        });

        it('Unpause (owner)', async () => {
            await nlgst.unpause({
                from: root
            });
        })

        it('Unpause: failed because cannot unpause while constract is not pausing ', async () => {
            await u.assertRevert(nlgst.unpause({
                from: root
            }));
        });
    })

    describe('Blacklist stage', async () => {
        it('Add blacklist (account1): failed because only owner can call it', async () => {
            const i = {
                evil: evil,
            }
            await u.assertRevert(nlgst.addBlackList(i.evil, {
                from: account1
            }));
        })

        it('Add blacklist: Owner call', async () => {
            const i = {
                evil: evil,
            }
            await nlgst.addBlackList(i.evil, {
                from: root
            });
        })

        it('Evil cannot call transferFrom', async () => {
            const i = {
                tokenId: 1,
                to: evil,
            }

            await nlgst.transferFrom(i.to, i.tokenId, {
                from: root
            });

            const o = {
                evil_balance: 1,
                owner: evil,
            }

            let evil_balance = await nlgst.balanceOf(i.to, {
                from: evil
            });
            eq(o.evil_balance, evil_balance.toString());

            // evil try to call transfer method
            await u.assertRevert(nlgst.transferFrom(evil, account1, i.value, {
                from: evil
            }));

            await nlgst.approve(root, "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", {
                from: evil
            });
            // evil try to call transfer from method
            await u.assertRevert(nlgst.transferFrom(evil, account1, i.value, {
                from: root
            }));
        });

        it('Remove blacklist (account1): Failed because only owner can call it', async () => {
            const i = {
                evil: evil,
            }
            await u.assertRevert(nlgst.removeBlackList(i.evil, {
                from: account1
            }));
        })

        it('Remove blacklist', async () => {
            const i = {
                evil: evil,
            }
            await nlgst.removeBlackList(i.evil, {
                from: root
            });
        })

        it('Evil can call transfer from method', async () => {
            const i = {
                from: evil,
                to: account6,
                tokenId: 1,
            }

            // evil can call transfer from method
            await nlgst.transferFrom(i.from, i.to, i.tokenId, {
                from: i.from
            });

            const o = {
                account6_balance: 1,
                evil_balance: 0,
            }

            let account6_balance = await nlgst.balanceOf(i.to, {
                from: root
            });
            eq(o.account6_balance, account6_balance.toString());
        });

        it('Destroy hacker (account6) funds: Failed because this is not black funds account', async () => {
            const i = {
                to: account6,
                tokenId: 2,
            }
            await nlgst.transferFrom(i.from, i.to, i.value, {
                from: root
            });
            await u.assertRevert(nlgst.destroyBlackFunds(i.to, {
                from: root
            }));
        })

        it('Destroy hacker (evil) funds: Failed because only owner can destroy', async () => {
            const i = {
                evil: evil,
                tokenId: 3
            }
            await nlgst.transferFrom(i.evil, i.value, {
                from: root
            });

            let evil_balance = await nlgst.balanceOf(i.evil, {
                from: root
            });
            const o = {
                evil_balance: 2
            }
            eq(o.evil_balance, evil_balance.toString());

            await nlgst.addBlackList(i.evil, {
                from: root
            });
            await u.assertRevert(nlgst.destroyBlackFunds(i.evil, {
                from: account1,
            }));
        });

        it('Destroy hacker (evil) funds: Failed because only owner can destroy', async () => {
            const i = {
                evil: evil,
            }

            const o = {
                evil_address: evil,
                evil_balance: 0,
                evil_funds: 2,
            }

            let tx = await nlgst.destroyBlackFunds(i.evil, {
                from: root,
            });
            let evil_balance = await nlgst.balanceOf(i.evil, {
                from: i.evil
            });

            eq(o.evil_balance, evil_balance);
            eq(o.evil_address, await oc(tx, "DestroyedBlackFunds", "_blackListedUser"))
            eq(o.evil_funds, await oc(tx, "DestroyedBlackFunds", "_balance"))
        });
    });
});
