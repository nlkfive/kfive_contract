const NLGST = artifacts.require("NLGST")
const STORAGE = artifacts.require("MarketplaceStorage")

const NFT1 = require('./ipfs/NFT1.json');

global.fetch = require("node-fetch");

const IPFS = require('ipfs')

const eq = assert.equal
const u = require('./utils.js')
const keccak256 = require('js-sha3').keccak256;
const superAdminRole = "0x0000000000000000000000000000000000000000000000000000000000000000";
const adminRole = "0x" + keccak256("ADMIN_ROLE");
const minterRole = "0x" + keccak256("MINTER_ROLE");
const pauserRole = "0x" + keccak256("PAUSER_ROLE");

var nlgst, storage;

contract("NLGST", (accounts) => {

    const root = accounts[0]
    const account1 = accounts[1]
    const account2 = accounts[2]
    const evil = accounts[3]
    const new_owner = accounts[4]
    const admin = accounts[5]
    const account6 = accounts[6]
    const super_admin = accounts[7]
    const new_admin = accounts[8]
    const account9 = accounts[9]
    const OFFCHAIN = web3.utils.fromAscii('1')

    const baseURL = "https://ipfs.io/ipfs/";

    before(async () => {
        storage = await STORAGE.new({
            from: root
        });

        nlgst = await NLGST.new(storage.address, baseURL, {
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

        it('Give minter_role to account1 (account1): Cannot because account1 does not have permission for grant this role', async () => {
            await u.assertRevert(nlgst.grantRole(minterRole, account1, {
                from: account1
            }));
        });

        it('Mint new token to account1 (account1): Cannot because account1 does not have MINTER_ROLE', async () => {
            await u.assertRevert(nlgst.mint(account1, 2, "2", {
                from: account1
            }));
        });

        it('Give minter_role to account1 (root)', async () => {
            await nlgst.grantRole(minterRole, account1, {
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
            await nlgst.revokeRole(minterRole, account1, {
                from: root
            });
        });

        it('Mint new token (account1) to [3] account1. For future transaction', async () => {
            await nlgst.grantRole(minterRole, account1, {
                from: root
            });

            const i = {
                tokenId: 3,
                tokenURI: "3"
            }

            await nlgst.mint(account1, i.tokenId, i.tokenURI, {
                from: account1
            });

            let o = {
                account1_balance: 3,
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

        it('Mint again token (account1) to [3] account1. Failed because token id=3 is already minted', async () => {
            const i = {
                tokenId: 3,
                tokenURI: "3"
            }

            await u.assertRevert(nlgst.mint(account1, i.tokenId, i.tokenURI, {
                from: account1
            }));
        });

        it('Revoke minter_role to account1 (account2). Cannot because only owner can revokeRole', async () => {
            await u.assertRevert(nlgst.revokeRole(minterRole, account1, {
                from: account2
            }));
        });

        it('Revoke minter_role to account1 (root)', async () => {
            await nlgst.revokeRole(minterRole, account1, {
                from: root
            });
        });

        it('Mint new token (account1). Cannot because account1 does not have minter_role', async () => {
            await u.assertRevert(nlgst.mint(account1, 4, "4", {
                from: account1
            }));
        });
    });

    describe('Admin stage', async () => {
        it('Give admin_role to super_admin (root)', async () => {
            await nlgst.grantRole(superAdminRole, super_admin, {
                from: root
            });
        });

        it('Give admin_role to super_admin (super_admin)', async () => {
            await nlgst.grantRole(minterRole, super_admin, {
                from: root
            });
        });

        it('Give admin_role to super_admin (super_admin)', async () => {
            await nlgst.grantRole(pauserRole, super_admin, {
                from: root
            });
        });

        it('Give admin_role to new_admin (root)', async () => {
            await nlgst.grantRole(adminRole, new_admin, {
                from: root
            });
        });

        it('Revoke admin_role to new_admin (root)', async () => {
            await nlgst.revokeRole(adminRole, new_admin, {
                from: root
            });
        });

        it('Add evil to black list (super_admin). Cannot because onlyOwner can call', async () => {
            await u.assertRevert(nlgst.addBlackList(evil, {
                from: super_admin
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
                account1_balance: 2,
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

        it('Transfer token (account1) [4] to account2. Cannot because account1 does not own token_id=4', async () => {
            const i = {
                tokenId: 4,
                from: account1,
                to: account2
            }

            await u.assertRevert(nlgst.transferFrom(i.from, i.to, i.tokenId, {
                from: i.from
            }));
        });

        it('Transfer token (root) [1] from account1 to account2. Cannot because root does not own this token', async () => {
            const i = {
                tokenId: 1,
                from: account1,
                to: account2
            }

            await u.assertRevert(nlgst.transferFrom(i.from, i.to, i.tokenId, {
                from: root
            }));
        });

        it('Approve token (account1) [1] to root. For future transaction', async () => {
            const i = {
                tokenId: 1,
                to: root
            }

            await nlgst.approve(i.to, i.tokenId, {
                from: account1
            });
        });

        it('Transfer token (root) [1] from account1 to account2', async () => {
            const i = {
                tokenId: 1,
                from: account1,
                to: account2
            }

            await nlgst.transferFrom(i.from, i.to, i.tokenId, {
                from: root
            });

            let o = {
                account1_balance: 1,
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

        it('Transfer token (root) [1] from account2 to account1. Cannot because root does not own token [1] anymore', async () => {
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
        it('Pause (account1). Cannot because only pauser_role can call it', async () => {
            await u.assertRevert(nlgst.pause({
                from: account1
            }));
        });

        it('Give admin_role (root) to account1', async () => {
            await nlgst.grantRole(adminRole, account1, {
                from: root
            });
        });

        it('Pause (account1 with pauser_role).', async () => {
            await nlgst.pause({
                from: root
            });
        });

        it('Pause again (owner). Cannot because transaction has already been paused', async () => {
            await u.assertRevert(nlgst.pause({
                from: root
            }));
        });

        it('Transfer token (account2) to [2] account1. Cannot because cannot transfer while pausing', async () => {
            const i = {
                from: account2,
                to: account1,
                tokenId: 2,
            }

            await u.assertRevert(nlgst.transferFrom(i.from, i.to, i.tokenId, {
                from: root
            }));
        });

        it('Unpause (account2). Cannot because only pauser_role can call it', async () => {
            await u.assertRevert(nlgst.unpause({
                from: account2
            }));
        });

        it('Revoke admin_role and pauser_role (root) to account1', async () => {
            await nlgst.revokeRole(adminRole, account1, {
                from: root
            });

            await nlgst.revokeRole(pauserRole, account1, {
                from: root
            });
        });

        it('Unpause (account1). Cannot because account1 has no pauser_role anymore', async () => {
            await u.assertRevert(nlgst.unpause({
                from: account2
            }));
        });

        it('Unpause (owner)', async () => {
            await nlgst.unpause({
                from: root
            });
        });

        it('Unpause again (owner). Cannot because cannot unpause while constract is not pausing ', async () => {
            await u.assertRevert(nlgst.unpause({
                from: root
            }));
        });
    });

    describe('Blacklist stage', async () => {
        it('Add blacklist (account1). Cannot because only owner can call it', async () => {
            const i = {
                evil: evil,
            }
            await u.assertRevert(nlgst.addBlackList(i.evil, {
                from: account1
            }));

            await nlgst.grantRole(adminRole, account1, {
                from: root
            });

            await u.assertRevert(nlgst.addBlackList(i.evil, {
                from: account1
            }));
        });

        it('Transfer token (account1) to [3] evil. For future transaction', async () => {
            const i = {
                tokenId: 3,
                from: account1,
                to: evil,
            }

            await nlgst.transferFrom(i.from, i.to, i.tokenId, {
                from: account1
            });

            const o = {
                evil_balance: 1,
                account1_balance: 0,
                owner: evil,
            }

            let evil_balance = await nlgst.balanceOf(i.to, {
                from: root
            });
            eq(o.evil_balance, evil_balance.toString());

            let account1_balance = await nlgst.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance, account1_balance.toString());

            let owner = await nlgst.ownerOf(i.tokenId, {
                from: root
            });
            eq(o.owner, owner);
        });

        it('Add blacklist: Owner call', async () => {
            const i = {
                evil: evil,
            }
            await nlgst.addBlackList(i.evil, {
                from: root
            });
        });

        it('Evil cannot call transferFrom', async () => {
            const i = {
                tokenId: 3,
                from: evil,
                to: account1,
            }

            // evil try to call transfer and safeTransfer method
            await u.assertRevert(nlgst.transferFrom(i.from, i.to, i.tokenId, {
                from: evil
            }));

            await u.assertRevert(nlgst.safeTransferFrom(i.from, i.to, i.tokenId, {
                from: evil
            }));

            // evil try to approve token id = 3 to root and root call transferFrom + safeTransferFrom method
            await nlgst.approve(root, i.tokenId, {
                from: evil
            });

            await u.assertRevert(nlgst.transferFrom(evil, account1, i.tokenId, {
                from: root
            }));

            await u.assertRevert(nlgst.safeTransferFrom(evil, account1, i.tokenId, {
                from: root
            }));
        });

        it('Remove blacklist (account2 + account1 admin). Cannot because only owner can call it', async () => {
            const i = {
                evil: evil,
            }
            await u.assertRevert(nlgst.removeBlackList(i.evil, {
                from: account2
            }));

            await u.assertRevert(nlgst.removeBlackList(i.evil, {
                from: account1
            }));
        });

        it('Remove blacklist (Root)', async () => {
            const i = {
                evil: evil,
            }
            await nlgst.removeBlackList(i.evil, {
                from: root
            });
        });

        it('Evil can call transferFrom method', async () => {
            const i = {
                from: evil,
                to: account6,
                tokenId: 3,
            }

            // evil can call transferFrom method
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
    });

    describe('Transfer ownership / Revoke / Renounce', async () => {
        it('Transfer ownership (account1) to new_owner. Cannot: only owner can call transfer ownership', async () => {
            const i = {
                new_owner: new_owner,
            }

            await u.assertRevert(nlgst.transferOwnership(i.new_owner, {
                from: account1
            }));

            await u.assertRevert(nlgst.transferOwnership(i.new_owner, {
                from: account2
            }));
        });

        it('Renounce admin_role (root) to account1. Cannot because account1 should renouce itself', async () => {
            await u.assertRevert(nlgst.renounceRole(adminRole, account1, {
                from: root
            }));
        });

        it('Revoke admin_role (account2) to account1. Cannot because only owner can revoke', async () => {
            await u.assertRevert(nlgst.revokeRole(adminRole, account1, {
                from: account2
            }));
        });

        it('Renounce admin_role (account1)', async () => {
            await nlgst.renounceRole(adminRole, account1, {
                from: account1
            });
        });

        it('Renounce admin_role again (account1). Cannot because account1 has no admin role', async () => {
            await u.assertRevert(nlgst.renounceRole(adminRole, account1, {
                from: account1
            }));
        });

        it('Revoke admin_role (root) to account1. Cannot because account1 has no admin role', async () => {
            await u.assertRevert(nlgst.revokeRole(adminRole, account1, {
                from: root
            }));
        });

        it('Transfer ownership (root) to new_owner', async () => {
            const i = {
                new_owner: new_owner,
            }

            // For future transaction
            const e = {
                tokenId: 10,
                tokenURI: "10"
            }

            await nlgst.mint(root, e.tokenId, e.tokenURI, {
                from: root
            });

            await nlgst.transferOwnership(i.new_owner, {
                from: root
            });
        });

        it('Root (old owner) can not call transferOwnership, addBlacklist, removeBacklist, revokeRole, grantRole, mint, pause', async () => {
            const i = {
                to: account2,
            }

            await u.assertRevert(nlgst.transferOwnership(i.to, {
                from: root
            }));

            await u.assertRevert(nlgst.grantRole(minterRole, i.to, {
                from: root
            }));

            await u.assertRevert(nlgst.revokeRole(minterRole, i.to, {
                from: root
            }));

            await u.assertRevert(nlgst.addBlackList(i.to, {
                from: root
            }));

            await u.assertRevert(nlgst.removeBlackList(i.to, {
                from: root
            }));

            await u.assertRevert(nlgst.pause({
                from: root
            }));

            await u.assertRevert(nlgst.unpause({
                from: root
            }));

            await u.assertRevert(nlgst.mint(root, 11, "11", {
                from: root
            }));

            await nlgst.transferFrom(root, new_owner, 10, {
                from: root
            });
        });

        it('New owner can call addBlacklist, removeBacklist, grantRole, revokeRole', async () => {
            const i = {
                to: account9,
            }

            await nlgst.grantRole(minterRole, i.to, {
                from: new_owner
            });

            await nlgst.revokeRole(minterRole, i.to, {
                from: new_owner
            });

            await nlgst.addBlackList(i.to, {
                from: new_owner
            });

            await nlgst.removeBlackList(i.to, {
                from: new_owner
            });
        });
    });

    describe('Burn stage', async () => {
        it('Burn token (account1) [1]. Cannot because only owner can call it', async () => {
            const i = {
                tokenId: 1,
            }

            await u.assertRevert(nlgst.burn(i.tokenId, {
                from: account1
            }));

            await nlgst.grantRole(adminRole, account1, {
                from: new_owner
            });

            await u.assertRevert(nlgst.burn(i.tokenId, {
                from: account1
            }));

            await u.assertRevert(nlgst.grantRole(minterRole, account1, {
                from: account1
            }));

            await u.assertRevert(nlgst.burn(i.tokenId, {
                from: account1
            }));

            await u.assertRevert(nlgst.grantRole(pauserRole, account1, {
                from: account1
            }));

            await u.assertRevert(nlgst.burn(i.tokenId, {
                from: account1
            }));
        });

        it('Burn token (root) [1]. Cannot because root is not owner', async () => {
            const i = {
                tokenId: 1,
            }

            await u.assertRevert(nlgst.burn(i.tokenId, {
                from: root
            }));
        });

        it('Burn token (new_owner) [1]', async () => {
            const i = {
                tokenId: 1,
            }

            await nlgst.burn(i.tokenId, {
                from: new_owner
            });
        });

        it('Burn token again (new_owner) [1]. Cannot because token [1] has been burned', async () => {
            const i = {
                tokenId: 1,
            }

            await u.assertRevert(nlgst.burn(i.tokenId, {
                from: new_owner
            }));
        });
    });

    describe('Check NFT data', async () => {
        it('Add NFT1 data to IPFS', async () => {
            const node = await IPFS.create()

            const NFT1data = {
                name: "NFT NGOC LINH GINSENG 1",
                image: "https://ipfs.io/ipfs/QmRHiJcTYJCcxxuuRfmjqWPZ1nGamzKRBQhnosrHsXnGFj",
                dna: "https://ipfs.io/ipfs/QmNWmvmH4s7Df8VqZmY21NHRsSsm4otn4ercZxtU8SaKvz",
                explored_at: 1629199374,
                place: "Ngoc Linh mountain - Dak Na commune - Tu Mo Rong district - Kon Tum province",
                total_weight: 225,
                weight_unit: "gram",
                seeded_at: 94665600,
                type: "live plants",
                description: "grown and cared for at PoYan Garden in Tu Mo Rong district, Kon Tum province. It is expected that freeze-drying will be carried out on August 8, 2022 for preservation after in vitro conservation of seeds samples.",
                extra_data: {}
            };

            const file = {
                path: 'NFT1.json',
                content: JSON.stringify(NFT1data)
            };

            const result = await node.add(file)

            await node.stop()
        });

        it('Mint a tokenId (hash keccak(name)) and IPFS URI', async () => {
            const i = {
                tokenId: "0x" + keccak256("NFT NGOC LINH GINSENG 1"),
                tokenURI: "QmRPFQUjoVMxxaMfpEwDWByg9MSo35rER81iYuSDZCMAfB"
            }

            await nlgst.mint(root, i.tokenId, i.tokenURI, {
                from: new_owner
            });
        });

        it('Check NFT data of NFT NGOC LINH GINSENG 1. Get IPFS link through tokenId and compare IPFS data = original data', async () => {
            const i = {
                tokenId1: "0x" + keccak256("NFT NGOC LINH GINSENG 1"),
            }

            // Get tokenURI 1 by tokenId1
            let tokenuri1 = await nlgst.tokenURI(i.tokenId1, {
                from: new_owner
            });

            // Get CID of NFT1
            let CID_NFT1 = tokenuri1.substring(tokenuri1.length - 46);

            // Get NFT1 data from IPFS
            const node = await IPFS.create()
            let data_NFT1 = ''

            for await (const chunk of node.cat(CID_NFT1, { timeout: 10000 })) {
                data_NFT1 = JSON.parse(chunk)
            }

            // Compare IPFS data and original data
            assert.deepEqual(data_NFT1, NFT1);
        });
    });

    describe('Accident 1 - Hacker access new_owner account and change password', async () => {
        it('Admin should immediately pause transaction, revoke admin, minter, pauser role', async () => {
            await nlgst.pause({
                from: super_admin
            });

            await nlgst.revokeRole(adminRole, new_owner, {
                from: super_admin
            });
        });
    });

    describe('Accident 2 - Root transfer ownership to Evil (wrong account)', async () => {
        it('Unpause (admin). For future transaction', async () => {
            await nlgst.unpause({
                from: super_admin
            });

            await nlgst.grantRole(superAdminRole, new_owner, {
                from: super_admin
            });
        });

        it('Transfer ownership (new_owner) to evil', async () => {
            const i = {
                evil: evil,
            }

            await nlgst.transferOwnership(i.evil, {
                from: new_owner
            });
        });

        it('Admin should immediately pause transaction, revoke admin, minter, pauser role of Evil', async () => {
            await nlgst.pause({
                from: super_admin
            });

            await nlgst.grantRole(superAdminRole, evil, {
                from: super_admin
            });
        });
    });
});