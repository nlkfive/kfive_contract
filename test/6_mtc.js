const MTC = artifacts.require("MTC")
const KFIVE = artifacts.require("KFIVE")

const eq = assert.equal
const u = require('./utils.js')
const oc = u.oc;
const tokenDecimals = 10;
const tokenCap = 1116400000;
const tokenCapKfive = 1080000;
var mtc, kfive;

contract("MTC", (accounts) => {

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

    before(async () => {
        kfive = await KFIVE.new({
            from: root
        });
        mtc = await MTC.new(kfive.address, {
            from: root
        });
    });

    describe('Deployment stage', async () => {
        it('Token information', async () => {
            let o = {
                balance: 0,
                name: 'Maitreya coin',
                symbol: 'MTC',
                decimals: tokenDecimals,
                // kfiveToken: kfive.address,
            }

            let owner_balance = await mtc.balanceOf(root, {
                from: root
            });

            eq(o.balance, owner_balance.toString());
            eq(o.name, await mtc.name.call());
            eq(o.symbol, await mtc.symbol.call());
            eq(o.decimals, await mtc.decimals.call());
        });
    });

    describe('Issue / Redeem', async () => {
        it('Root issue all MTC to account owner', async () => {
            const balance = tokenCap * (10 ** tokenDecimals);
            await mtc.issue(root, balance.toString(), { // Issue all token to root
                from: root
            });
            owner_balance = await mtc.balanceOf(root, {
                from: root
            });
            eq(balance.toString(), owner_balance.toString());
        });

        it('Root issue 10 mtc to account 1 (Cannot because all mtc issued to owner)', async () => {
            const value = web3.utils.toHex(10 * (10 ** tokenDecimals));
            await u.assertRevert(mtc.issue(account1, value, {
                from: root
            }));
        });

        it('Redeem 5 mtc from account 1 (owner call) (Cannot because account1 has no mtc)', async () => {
            const value = web3.utils.toHex(5 * (10 ** tokenDecimals));
            await u.assertRevert(mtc.redeem(account1, value, {
                from: root
            }));
        });

        it('Redeem 5 mtc from root (account 1 call) - Only owner can call this function', async () => {
            const value = web3.utils.toHex(5 * (10 ** tokenDecimals));
            await u.assertRevert(mtc.redeem(root, value, {
                from: account1
            }))
        });

        it('Redeem 10 token from root', async () => {
            const i = {
                redeemer: root,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }
            await mtc.redeem(i.redeemer, i.value, {
                from: root
            });

            const o = {
                owner_balance: (tokenCap - 10) * (10 ** tokenDecimals),
                total_supply: (tokenCap - 10) * (10 ** tokenDecimals),
            }

            let total_supply = await mtc.totalSupply({
                from: root
            });
            eq(o.total_supply, total_supply.toString());

            let owner_balance = await mtc.balanceOf(root, {
                from: root
            });
            eq(o.owner_balance, owner_balance.toString());
        });

        it('Issue 11 mtc to account1 (Cannot because there are only 10 mtc remaining)', async () => {
            const value = web3.utils.toHex(11 * (10 ** tokenDecimals));
            await u.assertRevert(mtc.issue(account1, value, {
                from: root
            }));
        })

        it('Issue 10 mtc to account1 (total supply will be reached)', async () => {
            const i = {
                issuer: account1,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }
            await mtc.issue(i.issuer, i.value, {
                from: root
            });

            const o = {
                owner_balance: (tokenCap - 10) * (10 ** tokenDecimals),
                issuer_balance: 10 * (10 ** tokenDecimals),
                total_supply: tokenCap * (10 ** tokenDecimals),
            }
            let total_supply = await mtc.totalSupply({
                from: root
            });
            eq(o.total_supply, total_supply.toString());

            let owner_balance = await mtc.balanceOf(root, {
                from: root
            });
            eq(o.owner_balance, owner_balance.toString());

            let issuer_balance = await mtc.balanceOf(i.issuer, {
                from: root
            });
            eq(o.issuer_balance, issuer_balance.toString());
        });
    });

    describe('Transfer stage', async () => {
        it('Transfer By Admin (only admin / owner) account 1 -> account 2: 10 MTC (Failed because account 1 is not an admin', async () => {
            const i = {
                from: account1,
                to: account2,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            await u.assertRevert(mtc.transferByAdmin(i.from, i.to, i.value, {
                from: account1
            }));
        });

        it('Transfer By Admin (only admin / owner) account 1 -> account 2: 10 MTC', async () => {
            const i = {
                from: account1,
                to: account2,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            await mtc.transferByAdmin(i.from, i.to, i.value, {
                from: root
            });

            const o = {
                account1_balance: 0,
                account2_balance: 10 * (10 ** tokenDecimals)
            }
            let account2_balance = await mtc.balanceOf(i.to, {
                from: root
            });
            eq(o.account2_balance, account2_balance.toString());

            let account1_balance = await mtc.balanceOf(i.from, {
                from: root
            });
            eq(o.account1_balance, account1_balance.toString());
        });

        it('Transfer: account 2 -> account 1: 10 MTC', async () => {
            const i = {
                to: account1,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            await mtc.transfer(i.to, i.value, {
                from: account2
            });

            const o = {
                account1_balance: 10 * (10 ** tokenDecimals),
                account2_balance: 0
            }
            let account2_balance = await mtc.balanceOf(account2, {
                from: root
            });
            eq(o.account2_balance, account2_balance.toString());

            let account1_balance = await mtc.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance, account1_balance.toString());
        });

        it('Transfer from: Root use allowance from account 1 transfer account1 -> account2 10 MTC, Cannot because account 1 has not approved to root', async () => {
            const i = {
                from: account1,
                to: account2,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }
            await u.assertRevert(i.from, i.to, i.value, {
                from: root
            });
        });

        it('Approve: account1 approve 10 MTC to root', async () => {
            var allowance = web3.utils.toHex(10 * (10 ** tokenDecimals));
            await mtc.approve(root, allowance, {
                from: account1
            });

            const o = {
                account1_allowance: web3.utils.toHex(10 * (10 ** tokenDecimals))
            }

            account1_allowance = await mtc.allowance(account1, root, {
                from: account1
            });
            eq(o.account1_allowance, web3.utils.toHex(account1_allowance));
        });

        // using the allowance mechanism. `amount` is then deducted from the caller\'s allowance.
        it('Transfer from: Root use allowance from account 1 transfer account1 -> account2 10 MTC', async () => {
            const i = {
                from: account1,
                to: account2,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            await mtc.transferFrom(i.from, i.to, i.value, {
                from: root
            });

            const o = {
                account1_balance: 0,
                account2_balance: 10 * (10 ** tokenDecimals),
            }

            let account1_balance = await mtc.balanceOf(i.from, {
                from: root
            });
            eq(o.account1_balance, account1_balance.toString());

            let account2_balance = await mtc.balanceOf(i.to, {
                from: root
            });
            eq(o.account2_balance, account2_balance.toString());
        });
    });

    describe('Pausable stage', async () => {
        it('Pause: Failed because ony owner can call it', async () => {
            await u.assertRevert(mtc.pause({
                from: account1
            }));
        });

        it('Pause: Onwer call', async () => {
            await mtc.pause({
                from: root
            });
        });

        it('Pause: Failed because cannot pause while already paused', async () => {
            await u.assertRevert(mtc.pause({
                from: root
            }));
        });

        it('Transfer: Failed because cannot transfer while pausing', async () => {
            const i = {
                to: account1,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            await u.assertRevert(mtc.transfer(i.to, i.value, {
                from: account2
            }));
        });

        it('TransferFrom: Failed because cannot transfer while pausing', async () => {
            const i = {
                from: account2,
                to: account1,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            await u.assertRevert(mtc.transferFrom(i.from, i.to, i.value, {
                from: root
            }));
        });

        it('Unpause (account1): Failed because only owner can call it', async () => {
            await u.assertRevert(mtc.unpause({
                from: account1
            }));
        });

        it('Unpause (owner)', async () => {
            await mtc.unpause({
                from: root
            });
        });

        it('Unpause: failed because cannot unpause while constract is not pausing ', async () => {
            await u.assertRevert(mtc.unpause({
                from: root
            }));
        });
    });

    describe('Blacklist stage', async () => {
        it('Add blacklist (account1): failed because only owner can call it', async () => {
            const i = {
                evil: evil,
            }
            await u.assertRevert(mtc.addBlackList(i.evil, {
                from: account1
            }));
        })

        it('Add blacklist: Owner call', async () => {
            const i = {
                evil: evil,
            }
            await mtc.addBlackList(i.evil, {
                from: root
            });
        });

        it('Evil cannot call transfer + transfer from', async () => {
            const i = {
                to: evil,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            await mtc.transfer(i.to, i.value, {
                from: root
            });

            const o = {
                evil_balance: 10 * (10 ** tokenDecimals),
            }

            let evil_balance = await mtc.balanceOf(i.to, {
                from: evil
            });
            eq(o.evil_balance, evil_balance.toString());

            // evil try to call transfer method
            await u.assertRevert(mtc.transfer(account1, i.value, {
                from: evil
            }));

            await mtc.approve(root, "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", {
                from: evil
            });
            // evil try to call transfer from method
            await u.assertRevert(mtc.transferFrom(evil, account1, i.value, {
                from: root
            }));
        });

        it('Remove blacklist (account1): Failed because only owner can call it', async () => {
            const i = {
                evil: evil,
            }
            await u.assertRevert(mtc.removeBlackList(i.evil, {
                from: account1
            }));
        });

        it('Remove blacklist', async () => {
            const i = {
                evil: evil,
            }
            await mtc.removeBlackList(i.evil, {
                from: root
            });
        });

        it('Evil can call transfer + transfer from method', async () => {
            const i = {
                from: evil,
                to: account6,
                value: web3.utils.toHex(5 * (10 ** tokenDecimals)),
            }

            // evil can call transfer method
            await mtc.transfer(i.to, i.value, {
                from: i.from
            });

            // evil can call transfer from method
            await mtc.transferFrom(i.from, i.to, i.value, {
                from: root
            });

            const o = {
                account6_balance: 10 * (10 ** tokenDecimals),
            }

            let account6_balance = await mtc.balanceOf(i.to, {
                from: root
            });
            eq(o.account6_balance, account6_balance.toString());
        });

        it('Destroy hacker (account6) funds: Failed because this is not black funds account', async () => {
            const i = {
                to: account6,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }
            await mtc.transfer(i.to, i.value, {
                from: root
            });
            await u.assertRevert(mtc.destroyBlackFunds(i.to, {
                from: root
            }));
        });

        it('Destroy hacker (evil) funds: Failed because only owner can destroy', async () => {
            const i = {
                evil: evil,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals))
            }
            await mtc.transfer(i.evil, i.value, {
                from: root
            });

            let evil_balance = await mtc.balanceOf(i.evil, {
                from: root
            });
            const o = {
                evil_balance: 10 * (10 ** tokenDecimals)
            }
            eq(o.evil_balance, evil_balance.toString());

            await mtc.addBlackList(i.evil, {
                from: root
            });
            await u.assertRevert(mtc.destroyBlackFunds(i.evil, {
                from: account1,
            }));
        });

        it('Destroy hacker (evil) funds', async () => {
            const i = {
                evil: evil,
            }

            const o = {
                evil_address: evil,
                evil_balance: 0,
                evil_funds: 10 * (10 ** tokenDecimals),
            }

            let tx = await mtc.destroyBlackFunds(i.evil, {
                from: root,
            });
            let evil_balance = await mtc.balanceOf(i.evil, {
                from: i.evil
            });

            eq(o.evil_balance, evil_balance);
            eq(o.evil_address, await oc(tx, "DestroyedBlackFunds", "_blackListedUser"))
            eq(o.evil_funds, await oc(tx, "DestroyedBlackFunds", "_balance"))
        });
    });

    describe('Admin stage', async () => {
        it('Add a new admin,  Cannot because only owner can add admin', async () => {
            const i = {
                new_admin: new_admin,
                max_issuing_per_times: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                max_total_issuing_token: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            await u.assertRevert(mtc.addAdmin(i.new_admin, i.max_issuing_per_times, i.max_total_issuing_token, {
                from: account1
            }));
        });

        it('Cannot add admin with max_total_issuing_token > cap value', async () => {
            const i = {
                new_admin: new_admin,
                max_issuing_per_times: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                max_total_issuing_token: web3.utils.toHex((tokenCap + 1) * (10 ** tokenDecimals)),
            }

            await u.assertRevert(mtc.addAdmin(i.new_admin, i.max_issuing_per_times, i.max_total_issuing_token, {
                from: account1
            }));
        });

        it('Owner add admin', async () => {
            const i = {
                new_admin: new_admin,
                max_issuing_token_per_time: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                max_total_issuing_token: web3.utils.toHex(100 * (10 ** tokenDecimals)),
            }

            const o = {
                status: true,
                max_issuing_token_per_time: 10 * (10 ** tokenDecimals),
                max_total_issuing_token: 100 * (10 ** tokenDecimals),
                remaining_issuing_token: 100 * (10 ** tokenDecimals),
            }

            await mtc.addAdmin(i.new_admin, i.max_issuing_token_per_time, i.max_total_issuing_token, {
                from: root
            });

            let a = await mtc.admin(i.new_admin)

            eq(o.status, a.status);
            eq(o.max_issuing_token_per_time, a.maxIssuingTokenPerTime.toString());
            eq(o.max_total_issuing_token, a.maxTotalIssuingToken.toString());
            eq(o.remaining_issuing_token, a.remainingIssuingToken.toString());
        });

        it('Only owner can remove admin', async () => {
            const i = {
                new_admin: new_admin,
            }

            await u.assertRevert(mtc.removeAdmin(i.new_admin, {
                from: account1
            }));
        });

        it('Owner remove admin', async () => {
            const i = {
                new_admin: new_admin,
            }

            await mtc.removeAdmin(i.new_admin, {
                from: root
            });

            let a = await mtc.admin(i.new_admin)
            eq(false, a.status);
        });
    });

    describe('Issue by admin stage', async () => {
        it('Issue by admin: root (also be an Admin) 10 Token to account 8 - Total supply is reached', async () => {
            i = {
                issuer: account8,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }
            await u.assertRevert(mtc.issueByAdmin(i.issuer, i.value, {
                from: root
            }))
        });

        it('Issue by admin: account1 (Admin) 6 Token to account 8 - Maximum issuing token per time is only 5', async () => {
            let i = {
                new_admin: account1,
                max_issuing_token_per_time: web3.utils.toHex(5 * (10 ** tokenDecimals)),
                max_total_issuing_token: web3.utils.toHex(100 * (10 ** tokenDecimals)),
            }

            await mtc.addAdmin(i.new_admin, i.max_issuing_token_per_time, i.max_total_issuing_token, {
                from: root
            });

            i = {
                issuer: account8,
                value: web3.utils.toHex(6 * (10 ** tokenDecimals)),
            }
            await u.assertRevert(mtc.issueByAdmin(i.issuer, i.value, {
                from: account1
            }))
        });

        it('Issue By Admin: account2 isn\'t an Admin', async () => {
            i = {
                issuer: account8,
                value: web3.utils.toHex(11 * (10 ** tokenDecimals)),
            }
            await u.assertRevert(mtc.issueByAdmin(i.issuer, i.value, {
                from: account2
            }))
        });

        it('Redeem owner address for next step issue token', async () => {

            const i = {
                redeemer: root,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            const o = {
                total_supply: (tokenCap - 10) * (10 ** tokenDecimals),
            }

            await mtc.redeem(root, i.value, {
                from: root
            });

            let total_supply = await mtc.totalSupply({
                from: root
            });

            eq(o.total_supply, total_supply.toString());
        });

        it('Issue by admin', async () => {
            const i = {
                issuer: account9,
                value: web3.utils.toHex(5 * (10 ** tokenDecimals)),
            }

            const o = {
                issuer_balance: 5 * (10 ** tokenDecimals),
                total_supply: (tokenCap - 5) * (10 ** tokenDecimals),
                admin_remaining_issuing: 95 * (10 ** tokenDecimals),
            }

            await mtc.issueByAdmin(i.issuer, i.value, {
                from: account1
            });

            let a = await mtc.admin(account1)
            eq(o.admin_remaining_issuing, a.remainingIssuingToken.toString());

            let total_supply = await mtc.totalSupply({
                from: root
            });

            let issuer_balance = await mtc.balanceOf(i.issuer, {
                from: root
            });

            eq(o.total_supply, total_supply.toString());
            eq(o.issuer_balance, issuer_balance.toString());
        });

        it('Issue by admin: cannot issue total value greater than remaining issue value (100 - 5 = 95 Token)', async () => {

            const i = {
                issuer: account9,
                value: web3.utils.toHex(96 * (10 ** tokenDecimals)),
            }

            await u.assertRevert(mtc.issueByAdmin(i.issuer, i.value, {
                from: account1
            }));
        });
    });

    describe('Transfer ownership', async () => {
        it('only owner can call transfer ownership', async () => {
            const i = {
                new_owner: new_owner,
            }

            await u.assertRevert(mtc.transferOwnership(i.new_owner, {
                from: account1
            }));
        });

        it('transfer ownership', async () => {
            const i = {
                new_owner: new_owner,
            }

            await mtc.transferOwnership(i.new_owner, {
                from: root
            });
        });

        it('root cannot call onlyOwner methods anymore', async () => {
            const i = {
                new_owner: new_owner,
                from: root,
                to: account5,
                value: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }

            await u.assertRevert(mtc.transferOwnership(i.new_owner, {
                from: root
            }));

            await u.assertRevert(mtc.transferByAdmin(i.from, i.to, i.value, {
                from: root
            }));

            await u.assertRevert(mtc.redeem(root, i.value, {
                from: root
            }));

            await u.assertRevert(mtc.issue(i.to, i.value, {
                from: root
            }));
        });
    });

    describe('Transfer KFIVE receiver', async () => {
        it('Account 1 transfer KFIVE_receiver to new KfiveReceiver, Cannot because account1 is not owner', async () => {
            const i = {
                new_kfiveReceiver: kfive.address,
            }
            await u.assertRevert(mtc.transferKfiveReceiver(i.new_kfiveReceiver, {
                from: account1
            }));
        });

        it('Owner transfer KFIVE_receiver to new KfiveReceiver, Cannot transfer to address(0)', async () => {
            const i = {
                new_kfiveReceiver: '0x0000000000000000000000000000000000000000',
            }
            await u.assertRevert(mtc.transferKfiveReceiver(i.new_kfiveReceiver, {
                from: new_owner
            }));
        });

        it('Owner transfer KFIVE_receiver to new KfiveReceiver', async () => {
            const i = {
                new_kfiveReceiver: root,
            }
            await mtc.transferKfiveReceiver(i.new_kfiveReceiver, {
                from: new_owner
            });
        });

        it('Owner transfer KFIVE_receiver to account2', async () => {
            const i = {
                new_kfiveReceiver: account2,
            }
            await mtc.transferKfiveReceiver(i.new_kfiveReceiver, {
                from: new_owner
            });
        });
    });

    describe('Update KFIVE Address', async () => {
        it('Account1 (admin) update KfiveAddress, Cannot because account1 is not owner', async () => {
            const i = {
                new_kfiveAddress: kfive.address,
            }
            await u.assertRevert(mtc.updateKfiveAddress(i.new_kfiveAddress, {
                from: account1
            }));
        });

        it('Account2 update KfiveAddress, Cannot because account2 is not owner', async () => {
            const i = {
                new_kfiveAddress: kfive.address,
            }
            await u.assertRevert(mtc.updateKfiveAddress(i.new_kfiveAddress, {
                from: account2
            }));
        });

        it('Owner update KfiveAddress, Cannot because account2 is not a deployed contract', async () => {
            const i = {
                new_kfiveAddress: account2,
            }
            await u.assertRevert(mtc.updateKfiveAddress(i.new_kfiveAddress, {
                from: new_owner
            }));
        });

        it('Owner update KfiveAddress', async () => {
            const i = {
                new_kfiveAddress: kfive.address,
            }
            await mtc.updateKfiveAddress(i.new_kfiveAddress, {
                from: new_owner
            });
        });
    });

    describe('Exchange From KFIVE', async () => {
        before(async () => {
            kfive = await KFIVE.new({
                from: root
            });
            mtc = await MTC.new(kfive.address, {
                from: root
            });
        });

        it('Account5 approve contract MTC can transfer KFIVE', async () => {
            var allowance = web3.utils.toHex(1080 * (10 ** tokenDecimals));
            await kfive.approve(mtc.address, allowance, {
                from: account5
            });

            const o2 = {
                root_allowance: web3.utils.toHex(1080 * (10 ** tokenDecimals))
            }

            root_allowance = await kfive.allowance(account5, mtc.address, {
                from: account5
            });
            eq(o2.root_allowance, web3.utils.toHex(root_allowance));
        });

        it('Account5 transfer 1 Kfive and receive 1080 MTC, Cannot because account5 does not have any KFIVE', async () => {
            const value = 1 * (10 ** tokenDecimals);
            await u.assertRevert(mtc.exchangeFromKfive(value, {
                from: account5
            }));
        });

        it('Account6 transfer 1 Kfive and receive 1080 MTC, Cannot because account6 does not approve MTC contract', async () => {
            const balance = 1 * (10 ** tokenDecimals);
            await kfive.issue(account6, balance.toString(), OFFCHAIN, {
                from: root
            });
            await u.assertRevert(mtc.exchangeFromKfive(balance, {
                from: account6
            }));
        });

        it('Account6 also approve contract MTC can transfer KFIVE', async () => {
            var allowance = web3.utils.toHex(1080 * (10 ** tokenDecimals));
            await kfive.approve(mtc.address, allowance, {
                from: account6
            });

            const o2 = {
                root_allowance: web3.utils.toHex(1080 * (10 ** tokenDecimals))
            }

            root_allowance = await kfive.allowance(account6, mtc.address, {
                from: account6
            });
            eq(o2.root_allowance, web3.utils.toHex(root_allowance));
        });

        it('Issue 1 kfive to account5', async () => {
            const balance = 1 * (10 ** tokenDecimals);
            await kfive.issue(account5, balance.toString(), OFFCHAIN, {
                from: root
            });
            account5_balance = await kfive.balanceOf(account5, {
                from: root
            });
            eq(balance.toString(), account5_balance.toString());
        });

        it('Account5 transfer 1 Kfive and receive 1080 MTC', async () => {
            const value = 1 * (10 ** tokenDecimals);
            await mtc.exchangeFromKfive(value, {
                from: account5
            });

            const o = {
                account5_mtc_balance: 1080 * (10 ** tokenDecimals),
            }

            let account5_mtc_balance = await mtc.balanceOf(account5, {
                from: root
            });
            eq(o.account5_mtc_balance, account5_mtc_balance.toString());

            const o2 = {
                account5_kfive_balance: 0,
                root_kfive_balance: 1 * (10 ** tokenDecimals),
            }

            let account5_kfive_balance = await kfive.balanceOf(account5, {
                from: root
            });
            eq(o2.account5_kfive_balance, account5_kfive_balance.toString());

            let root_kfive_balance = await kfive.balanceOf(root, {
                from: root
            });
            eq(o2.root_kfive_balance, root_kfive_balance.toString());
        });

        it('Account6 transfer 5 Kfive and receive 1080*5 MTC, Cannot because account6 only has 1 KFIVE', async () => {
            const value = 5 * (10 ** tokenDecimals);
            await u.assertRevert(mtc.exchangeFromKfive(value, {
                from: account6
            }));
        });

        it('Account6 transfer 1 Kfive and receive 1080 MTC', async () => {
            const value = 1 * (10 ** tokenDecimals);
            await mtc.exchangeFromKfive(value, {
                from: account6
            });

            const o = {
                account6_mtc_balance: 1080 * (10 ** tokenDecimals),
            }

            let account6_mtc_balance = await mtc.balanceOf(account5, {
                from: root
            });
            eq(o.account6_mtc_balance, account6_mtc_balance.toString());

            const o2 = {
                account6_kfive_balance: 0,
                root_kfive_balance: 2 * (10 ** tokenDecimals),
            }

            let account6_kfive_balance = await kfive.balanceOf(account5, {
                from: root
            });
            eq(o2.account6_kfive_balance, account6_kfive_balance.toString());

            let root_kfive_balance = await kfive.balanceOf(root, {
                from: root
            });
            eq(o2.root_kfive_balance, root_kfive_balance.toString());
        });
    });
});