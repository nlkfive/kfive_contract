const NLGST = artifacts.require("NLGST")
const KFIVE = artifacts.require("KFIVE")
const AuctionMarketplace = artifacts.require("AuctionMarketplace")
const OrderMarketplace = artifacts.require("OrderMarketplace")

const eq = assert.equal
const u = require('./utils.js')
var nlgst, kfive, auctionMarketplace, orderMarketplace;
const { soliditySha3 } = require("web3-utils");

function getBlindedBid(value, fake, secret) {
    return soliditySha3(
        value, fake, secret
    );
}

contract("NLGST", (accounts) => {
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

    const BASE_URL = "https://ipfs.io/ipfs/";
    const OWNER_CUT_PER_MILLION = 1000;
    const tokenDecimals = 10;
    const issuedTokenAmount = 10 * (10 ** tokenDecimals);
    const publicationFee1 = 0.1 * (10 ** tokenDecimals);
    const delay = ms => new Promise(res => setTimeout(res, ms));

    before(async () => {
        kfive = await KFIVE.new({
            from: root
        });
        nlgst = await NLGST.new(BASE_URL, {
            from: root
        });
        auctionMarketplace = await AuctionMarketplace.new(kfive.address, OWNER_CUT_PER_MILLION);

        orderMarketplace = await OrderMarketplace.new(kfive.address, OWNER_CUT_PER_MILLION, {
            from: root
        });
    });

    describe('Deployment stage', async () => {
        it('Order Marketplace information', async () => {
            let o = {
                acceptedToken: kfive.address,
                ownerCutPerMillion: 1000,
                owner: root,
                publicationFeeInWei: 0,
            }

            const acceptedToken = await orderMarketplace.acceptedToken();
            eq(acceptedToken, o.acceptedToken);

            const ownerCutPerMillion = await orderMarketplace.ownerCutPerMillion();
            eq(ownerCutPerMillion.toString(), o.ownerCutPerMillion);

            const owner = await orderMarketplace.owner();
            eq(owner, o.owner);

            const publicationFeeInWei = await orderMarketplace.publicationFeeInWei();
            eq(publicationFeeInWei.toString(), o.publicationFeeInWei);
        });

        it('Auction Marketplace information', async () => {
            let o = {
                acceptedToken: kfive.address
            }

            const acceptedToken = await auctionMarketplace.acceptedToken();
            eq(acceptedToken, o.acceptedToken);
        });

        it('Set publication fee (0.1 KFIVE)', async () => {
            let i = {
                fee: publicationFee1
            }
            await auctionMarketplace.setPublicationFee(i.fee);

            let o = {
                fee: publicationFee1.toString()
            }
            fee = await auctionMarketplace.publicationFeeInWei();
            eq(fee.toString(), o.fee);
        })
    });

    beforeEach(async () => {
        const issuedToken = web3.utils.toHex(issuedTokenAmount);
        const owner = await kfive.owner({ from: root });
        await kfive.issue(account1, issuedToken, OFFCHAIN, { from: owner });
        await kfive.issue(account2, issuedToken, OFFCHAIN, { from: owner });

        const o = {
            account1_balance: 10 * (10 ** tokenDecimals),
            account2_balance: 10 * (10 ** tokenDecimals),
        }

        let account1_balance = await kfive.balanceOf(account1, { from: root });
        eq(o.account1_balance, account1_balance.toString());

        let account2_balance = await kfive.balanceOf(account2, { from: root });
        eq(o.account2_balance, account2_balance.toString());
    });

    afterEach(async () => {
        const owner = await kfive.owner({ from: root });
        const account1_balance = await kfive.balanceOf(account1, { from: root });
        const account2_balance = await kfive.balanceOf(account2, { from: root });
        await kfive.redeem(account1, account1_balance, OFFCHAIN, { from: owner });
        await kfive.redeem(account2, account2_balance, OFFCHAIN, { from: owner });
    });

    describe('Order 1: ORDER, EXECUTE STAGE', async () => {
        ////////////////////////////
        // ONE
        ////////////////////////////
        it('(Account1) Create order NFT (1). Cannot because NFT (1) does not exist', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 1,
                priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                expiresAt: Math.floor(new Date().getTime() / 1000 + 300),
            }
            await u.assertRevert(orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                from: account1
            }));
        });

        it('(Owner) Mint new NFT (1) to (Account1)', async () => {
            const i = {
                nftID: 1,
                nftURI: "1",
            }
            await nlgst.mint(account1, i.nftID, i.nftURI, {
                from: root
            });

            const o = {
                nft_owner: account1,
            }

            let nft_owner = await nlgst.ownerOf(i.nftID, {
                from: root
            })
            eq(o.nft_owner, nft_owner);
        });

        it('(Account1) Create order NFT (1). Cannot because NFT (1) has not been approved to market', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 1,
                priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                expiresAt: Math.floor(new Date().getTime() / 1000 + 300),
            }
            await u.assertRevert(orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                from: account1
            }));
            // Root try to create order
            await u.assertRevert(orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                from: root
            }));
        });

        it('(Root) Approve NFT (1) to market. Cannot because only owner of NFT (1) can approve', async () => {
            const i = {
                to: orderMarketplace.address,
                nftID: 1,
            }
            await u.assertRevert(nlgst.approve(i.to, i.nftID, {
                from: root
            }));

            await u.assertRevert(nlgst.approve(i.to, i.nftID, {
                from: account2
            }));

            // Account1 approve NFT(1) to root
            await nlgst.approve(root, i.nftID, {
                from: account1
            });
            // Root try to approve NFT again to market
            await u.assertRevert(nlgst.approve(i.to, i.nftID, {
                from: root
            }));
        });

        it('(Account1) Approve NFT (1) to market', async () => {
            const i = {
                to: orderMarketplace.address,
                nftID: 1,
            }
            await nlgst.approve(i.to, i.nftID, {
                from: account1
            });

            const o = {
                getApproved: orderMarketplace.address,
            }

            let getApproved = await nlgst.getApproved(i.nftID, {
                from: root
            })
            eq(o.getApproved, getApproved);
        });

        it('(Account2) Create order NFT (1). Cannot because only root and owner of NFT (1) can create', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 1,
                priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                expiresAt: Math.floor(new Date().getTime() / 1000 + 300),
            }
            await u.assertRevert(orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                from: account2
            }));
        });

        it('(Account1) Create order NFT (1) with wrong expireAt time. Cannot because time must be more than 1 minute in the future', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 1,
                priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                expiresAt1: Math.floor(new Date().getTime() / 1000),
                expiresAt2: Math.floor(new Date().getTime() / 1000 - 60),
                expiresAt3: Math.floor(new Date().getTime() / 1000 + 50),
            }

            await u.assertRevert(orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt1, {
                from: account1
            }));
            await u.assertRevert(orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt2, {
                from: account1
            }));
            await u.assertRevert(orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt3, {
                from: account1
            }));
        });

        it('(Account1) Create order NFT (1)', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 1,
                priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                expiresAt: Math.floor(new Date().getTime() / 1000 + 300),
            }
            await orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                from: account1
            });
        });

        it('(Account2) approve 10 KFIVE token to Market', async () => {
            const i = {
                amount: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                to: orderMarketplace.address,
            }

            await kfive.approve(i.to, i.amount, {
                from: account2
            });

            const o = {
                approvedTokenAmount: 10 * (10 ** tokenDecimals),
            }

            let approvedTokenAmount = await kfive.allowance(account2, orderMarketplace.address, {
                from: root
            })
            eq(o.approvedTokenAmount, approvedTokenAmount.toString());
        });

        it('(Account2) Execute order', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 1,
                price: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }
            await orderMarketplace.executeOrder(i.nftAddress, i.assetId, i.price, {
                from: account2
            });

            const o = {
                account1_balance: 20 * (10 ** tokenDecimals) - (10 * OWNER_CUT_PER_MILLION / 1000000) * (10 ** tokenDecimals),
                account2_balance: 0,
                nft_owner: account2,
            }

            let account1_balance = await kfive.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance, account1_balance.toString());

            let account2_balance = await kfive.balanceOf(account2, {
                from: root
            });
            eq(o.account2_balance, account2_balance.toString());

            let nft_owner = await nlgst.ownerOf(i.assetId, {
                from: root
            })
            eq(o.nft_owner, nft_owner);
        });

        it('(Account5) Execute order again. Cannot because NFT (1) has been sold', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 1,
                price: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }
            await u.assertRevert(orderMarketplace.executeOrder(i.nftAddress, i.assetId, i.price, {
                from: account2
            }));
        });
    });

    describe('Order 2: CANCEL, PAUSE STAGE', async () => {
        ////////////////////////////
        // TWO
        ////////////////////////////
        it('(Owner) Mint new NFT (2) to (Account2)', async () => {
            const i = {
                nftID: 2,
                nftURI: "2",
            }
            await nlgst.mint(account2, i.nftID, i.nftURI, {
                from: root
            });

            const o = {
                nft_owner: account2,
            }

            let nft_owner = await nlgst.ownerOf(i.nftID, {
                from: root
            })
            eq(o.nft_owner, nft_owner);
        });

        it('(Account2) Approve NFT (2) to market', async () => {
            const i = {
                to: orderMarketplace.address,
                nftID: 2,
            }
            await nlgst.approve(i.to, i.nftID, {
                from: account2
            });

            const o = {
                getApproved: orderMarketplace.address,
            }

            let getApproved = await nlgst.getApproved(i.nftID, {
                from: root
            })
            eq(o.getApproved, getApproved);
        });

        it('(Account2) Create order NFT (2)', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 2,
                priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                expiresAt: Math.floor(new Date().getTime() / 1000 + 65),
            }
            await orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                from: account2
            });
        });

        it('(Account1) approve 10 KFIVE token to Market', async () => {
            const i = {
                amount: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                to: orderMarketplace.address,
            }

            await kfive.approve(i.to, i.amount, {
                from: account1
            });

            const o = {
                approvedTokenAmount: 10 * (10 ** tokenDecimals),
            }

            let approvedTokenAmount = await kfive.allowance(account1, orderMarketplace.address, {
                from: root
            })
            eq(o.approvedTokenAmount, approvedTokenAmount.toString());
        });

        it('Wait 60s', async () => {
            await delay(65 * 1000);
        });

        it('(Account1) Execute order NFT(2). Cannot because over expire time', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 1,
                price: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }
            await u.assertRevert(orderMarketplace.executeOrder(i.nftAddress, i.assetId, i.price, {
                from: account1
            }));
        });

        it('(Account2) Create order NFT (2)', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 2,
                priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                expiresAt: Math.floor(new Date().getTime() / 1000 + 65),
            }
            await orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                from: account2
            });
        });

        it('(Account1) cancel order NFT(2). Cannot because only root and owner can cancel order', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 2,
            }
            await u.assertRevert(orderMarketplace.cancelOrder(i.nftAddress, i.assetId, {
                from: account1
            }));
        });

        it('(Root) cancel order NFT(2)', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 2,
            }
            await orderMarketplace.cancelOrder(i.nftAddress, i.assetId, {
                from: root
            });
        });

        it('(Account1) Execute order NFT(2). Cannot because NFT(2) has been cancelled', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 1,
                price: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }
            await u.assertRevert(orderMarketplace.executeOrder(i.nftAddress, i.assetId, i.price, {
                from: account1
            }));
        });

        it('(Account2) Create order NFT (2) again', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 2,
                priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                expiresAt: Math.floor(new Date().getTime() / 1000 + 300),
            }
            await nlgst.approve(orderMarketplace.address, i.assetId, {
                from: account2
            });

            await orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                from: account2
            });
        });

        it('(Owner) Pause', async () => {
            await orderMarketplace.pause({
                from: root
            });
        });

        it('(Account1) Create order NFT (1). Cannot because transaction is pausing', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 1,
                priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                expiresAt: Math.floor(new Date().getTime() / 1000 + 300),
            }
            await u.assertRevert(orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                from: account1
            }));
        });

        it('(Account1) Execute order NFT (2). Cannot because transaction is pausing', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 2,
                priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }
            await u.assertRevert(orderMarketplace.executeOrder(i.nftAddress, i.assetId, i.priceInWei, {
                from: account1
            }));
        });

        it('(Account2 + Root) Cancel order NFT (2). Cannot because transaction is pausing', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 2,
                priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }
            await u.assertRevert(orderMarketplace.executeOrder(i.nftAddress, i.assetId, i.priceInWei, {
                from: account2
            }));

            // Root try to cancel order
            await u.assertRevert(orderMarketplace.executeOrder(i.nftAddress, i.assetId, i.priceInWei, {
                from: root
            }));
        });

        it('(Owner) UnPause', async () => {
            await orderMarketplace.unpause({
                from: root
            });
        });

        it('(Account1) Execute order NFT (2) successfully', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 2,
                priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
            }
            await orderMarketplace.executeOrder(i.nftAddress, i.assetId, i.priceInWei, {
                from: account1
            });

            const o = {
                account1_balance: 0,
                account2_balance: 20 * (10 ** tokenDecimals) - (10 * OWNER_CUT_PER_MILLION / 1000000) * (10 ** tokenDecimals),
                nft_owner: account1,
            }

            let account1_balance = await kfive.balanceOf(account1, {
                from: root
            });
            eq(o.account1_balance, account1_balance.toString());

            let account2_balance = await kfive.balanceOf(account2, {
                from: root
            });
            eq(o.account2_balance, account2_balance.toString());

            let nft_owner = await nlgst.ownerOf(i.assetId, {
                from: root
            })
            eq(o.nft_owner, nft_owner);
        });
    });

    describe('Order 3: New OwnerCutPerMillion and transferOwnership STAGE', async () => {
        ////////////////////////////
        // THREE
        ////////////////////////////
        it('(Account1) Set new setOwnerCutPerMillion. Cannot because only owner can change value', async () => {
            const i = {
                new_ownerCutPerMillion: 100000,
            }
            await u.assertRevert(orderMarketplace.setOwnerCutPerMillion(i.new_ownerCutPerMillion, {
                from: account1
            }));
        });

        it('(Root) Set new OwnerCutPerMillion', async () => {
            const i = {
                new_ownerCutPerMillion: 100000,
            }
            await orderMarketplace.setOwnerCutPerMillion(i.new_ownerCutPerMillion, {
                from: root
            });

            const new_ownerCutPerMillion = await orderMarketplace.ownerCutPerMillion();
            eq(new_ownerCutPerMillion.toString(), i.new_ownerCutPerMillion);
        });

        it('(Root) Set new OwnerCutPerMillion = 2.000.000. Cannot because OwnerCutPerMillion must be < 1.000.000', async () => {
            const i = {
                new_ownerCutPerMillion: 2000000,
            }
            await u.assertRevert(orderMarketplace.setOwnerCutPerMillion(i.new_ownerCutPerMillion, {
                from: root
            }));
        });

        it('(Account1) Create order NFT (2)', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 2,
                priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                expiresAt: Math.floor(new Date().getTime() / 1000 + 300),
            }

            await nlgst.approve(orderMarketplace.address, i.assetId, {
                from: account1
            });

            await orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                from: account1
            });
        });

        it('(Root) Transfer Ownership to new_owner', async () => {
            const i = {
                new_owner: new_owner,
            }
            await orderMarketplace.transferOwnership(i.new_owner, {
                from: root
            });

            let newOwner = await orderMarketplace.owner({
                from: root
            })
            eq(i.new_owner, newOwner);
        });

        it('(Root) Root can not call owner functions anymore', async () => {
            const i = {
                nftID: 2,
                nftAddress: nlgst.address,
            }
            await u.assertRevert(orderMarketplace.transferOwnership(account1, {
                from: root
            }));

            await u.assertRevert(orderMarketplace.cancelOrder(i.nftAddress, i.nftID, {
                from: root
            }));

            await u.assertRevert(orderMarketplace.pause({
                from: root
            }));

            await u.assertRevert(orderMarketplace.unpause({
                from: root
            }));
        });

        it('(New_owner) New_owner call pause, unpause and cancel order', async () => {
            const i = {
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await orderMarketplace.pause({
                from: new_owner
            });

            await u.assertRevert(orderMarketplace.cancelOrder(i.nftAddress, i.nftID, {
                from: new_owner
            }));

            await orderMarketplace.unpause({
                from: new_owner
            });

            await orderMarketplace.cancelOrder(i.nftAddress, i.nftID, {
                from: new_owner
            });
        });

        it('(Account2) Execute order NFT (2)', async () => {
            let account1_balance_before = await kfive.balanceOf(account1, { from: root });
            let account2_balance_before = await kfive.balanceOf(account2, { from: root });

            const i = {
                nftAddress: nlgst.address,
                assetId: 2,
                priceInWei: web3.utils.toHex(5 * (10 ** tokenDecimals)),
                expiresAt: Math.floor(new Date().getTime() / 1000 + 300),
            }

            await nlgst.approve(orderMarketplace.address, i.assetId, {
                from: account1
            });

            await kfive.approve(orderMarketplace.address, i.priceInWei, {
                from: account2
            });

            await orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                from: account1
            });

            await orderMarketplace.executeOrder(i.nftAddress, i.assetId, i.priceInWei, {
                from: account2
            });

            const o = {
                account1_balance_increase: (5 * (10 ** tokenDecimals)) - ((5 * 100000 / 1000000) * (10 ** tokenDecimals)),
                account2_balance_decrease: (5 * (10 ** tokenDecimals)),
                nft_owner: account2,
            }

            let account1_balance_after = web3.utils.toBN(o.account1_balance_increase).add(web3.utils.toBN(account1_balance_before)).toString();
            let account1_balance = await kfive.balanceOf(account1, {
                from: root
            });
            eq(account1_balance_after, account1_balance.toString());

            let account2_balance_after = web3.utils.toBN(account2_balance_before).sub(web3.utils.toBN(o.account2_balance_decrease)).toString();
            let account2_balance = await kfive.balanceOf(account2, {
                from: root
            });
            eq(account2_balance_after, account2_balance.toString());

            let nft_owner = await nlgst.ownerOf(i.assetId, {
                from: root
            })
            eq(o.nft_owner, nft_owner);
        });
    });

    describe('Order 4: PUBLICATION FEE STAGE', async () => {
        it('(root - account1) Set publication fee (0.1 KFIVE). Cannot because root, account1 is not owner', async () => {
            let i = {
                fee: publicationFee1
            }
            await u.assertRevert(orderMarketplace.setPublicationFee(i.fee, { from: root }));
            await u.assertRevert(orderMarketplace.setPublicationFee(i.fee, { from: account1 }));
        });

        it('(new_owner) Set publication fee (0.1 KFIVE)', async () => {
            let i = {
                fee: publicationFee1
            }
            await orderMarketplace.setPublicationFee(i.fee, { from: new_owner });

            let o = {
                fee: publicationFee1.toString()
            }
            fee = await orderMarketplace.publicationFeeInWei();
            eq(fee.toString(), o.fee);
        });

        it('(Account2) Create - Execute order NFT (1). Only balance of account2 - 0.1 KFIVE when create order successfully', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 1,
                priceInWei: web3.utils.toHex(5 * (10 ** tokenDecimals)),
                expiresAt: Math.floor(new Date().getTime() / 1000 + 300),
            }

            await kfive.approve(orderMarketplace.address, issuedTokenAmount, {
                from: account1
            });

            await kfive.approve(orderMarketplace.address, issuedTokenAmount, {
                from: account2
            });

            await nlgst.approve(orderMarketplace.address, i.assetId, {
                from: account2
            });

            // CREATE ORDER. Should subtract 0.1 KFIVE publication fee after create order
            await orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                from: account2
            });

            const o = {
                account2_balance_create: 10 * (10 ** tokenDecimals) - 0.1 * (10 ** tokenDecimals),
            }

            let account2_balance_create = await kfive.balanceOf(account2, { from: root });
            eq(o.account2_balance_create, account2_balance_create.toString());

            // EXECUTE ORDER. Should not subtract 0.1 KFIVE publication fee after execute order
            await orderMarketplace.executeOrder(i.nftAddress, i.assetId, i.priceInWei, {
                from: account1
            });

            const e = {
                account1_balance_execute: 5 * (10 ** tokenDecimals),
                account2_balance_execute: 15 * (10 ** tokenDecimals) - 0.1 * (10 ** tokenDecimals) - ((5 * 100000 / 1000000) * (10 ** tokenDecimals)),
                nft_owner: account1,
            }

            let account1_balance_execute = await kfive.balanceOf(account1, {
                from: root
            });
            eq(e.account1_balance_execute, account1_balance_execute.toString());

            let account2_balance_execute = await kfive.balanceOf(account2, {
                from: root
            });
            eq(e.account2_balance_execute, account2_balance_execute.toString());

            let nft_owner = await nlgst.ownerOf(i.assetId, {
                from: root
            })
            eq(e.nft_owner, nft_owner);
        });

        it('(Account1) Create - Cancel order NFT (1). Only balance of account1 - 0.1 KFIVE when create order successfully', async () => {
            const i = {
                nftAddress: nlgst.address,
                assetId: 1,
                priceInWei: web3.utils.toHex(5 * (10 ** tokenDecimals)),
                expiresAt: Math.floor(new Date().getTime() / 1000 + 300),
            }

            await kfive.approve(orderMarketplace.address, issuedTokenAmount, {
                from: account1
            });

            await nlgst.approve(orderMarketplace.address, i.assetId, {
                from: account1
            });

            // Create order
            await orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                from: account1
            });

            // Check balance after create order. Should subtract 0.1 KFIVE publication fee after create order
            const o = {
                account1_balance_create: 10 * (10 ** tokenDecimals) - 0.1 * (10 ** tokenDecimals),
            }

            let account1_balance_create = await kfive.balanceOf(account1, { from: root });
            eq(o.account1_balance_create, account1_balance_create.toString());

            // Cancel order
            await orderMarketplace.cancelOrder(i.nftAddress, i.assetId, {
                from: account1
            });

            // Check balance after cancel. Should not subtract 0.1 KFIVE publication fee after cancel order
            const e = {
                account1_balance_execute: 10 * (10 ** tokenDecimals) - 0.1 * (10 ** tokenDecimals),
            }

            let account1_balance_execute = await kfive.balanceOf(account1, { from: root });
            eq(e.account1_balance_execute, account1_balance_execute.toString());
        });

        describe('Order 5: BURN and CREATE ORDER STAGE', async () => {
            it('(New_owner) Burn NFT (1)', async () => {
                const i = {
                    assetId: 1,
                }

                await nlgst.burn(i.assetId, {
                    from: root
                });
            });

            it('(Account1) Create order NFT (1). Cannot because NFT (1) has been burned', async () => {
                const i = {
                    nftAddress: nlgst.address,
                    assetId: 1,
                    priceInWei: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                    expiresAt: Math.floor(new Date().getTime() / 1000 + 300),
                }
                await u.assertRevert(orderMarketplace.createOrder(i.nftAddress, i.assetId, i.priceInWei, i.expiresAt, {
                    from: account1
                }));
            });
        });
    });

    describe('Auction 2: Complete auction', async () => {
        const nftID = 1;
        const nftURI = "1";
        const startPriceInWei = 0.1 * (10 ** tokenDecimals); // 0.1 KFIVE
        const biddingTime = 60; // 1 minute
        const revealTime = 60; // 1 minute
        var auctionId;
        var auctionBiddingEnd;
        var auctionRevealEnd;

        it('(Owner) Mint new NFT to (Account1)', async () => {
            await nlgst.mint(account1, nftID, nftURI, {
                from: root
            });
        });

        it('(Account1) Approve NFT to market', async () => {
            await nlgst.approve(auctionMarketplace.address, nftID, {
                from: account1
            });
        });

        it('(Account1) Approve KFIVE to market (for paying publication fee - 10 KFIVE)', async () => {
            var allowance = web3.utils.toHex(10 * (10 ** tokenDecimals));
            await kfive.approve(auctionMarketplace.address, allowance, {
                from: account1
            });
        });

        it('(Account1) Create new auction', async () => {
            await auctionMarketplace.createAuction(nlgst.address, nftID, startPriceInWei, biddingTime, revealTime, { from: account1 });

            let o = {
                rootBalance: publicationFee1.toString(),
                account1Balance: (issuedTokenAmount - publicationFee1).toString(),
                seller: account1
            }

            const rootBalance = await kfive.balanceOf(root);
            eq(o.rootBalance, rootBalance.toString());

            const account1Balance = await kfive.balanceOf(account1);
            eq(o.account1Balance, account1Balance.toString());

            const lastestEvent = await auctionMarketplace.getPastEvents("AuctionCreated");
            auctionId = lastestEvent[0].returnValues.id;
            auctionBiddingEnd = parseInt(lastestEvent[0].returnValues.biddingEnd);
            revealEnd = parseInt(lastestEvent[0].returnValues.revealEnd);

            const seller = lastestEvent[0].returnValues.seller;
            eq(o.seller, seller);

            const startPrice = lastestEvent[0].returnValues.startPriceInWei;
            eq(startPrice, startPriceInWei);
        });

        ////////////////////////////
        // Bidding
        ////////////////////////////
        it('(Account2) 1st Bid auction (0.01 KFIVE) - Lower than start price', async () => {
            let i = {
                bidValue: 0.01 * (10 ** tokenDecimals),
                fake: false,
                secret: 0x123,
                depositValue: 0.01 * (10 ** tokenDecimals)
            }
            await auctionMarketplace.bidAuction(nlgst.address, 1, auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            });
        });

        it('(Account2) 2nd Bid auction (0.2 KFIVE) - Higher than start price', async () => {
            let i = {
                bidValue: 0.2 * (10 ** tokenDecimals),
                fake: false,
                secret: 0x123,
                depositValue: 0.2 * (10 ** tokenDecimals)
            }
            await auctionMarketplace.bidAuction(nlgst.address, 1, auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            });
        });

        it('(Account2) 3rd Bid auction (0.3 KFIVE / 0.03 KFIVE) - (Bid higher / Deposit lower) than start price (0.1 KFIVE)', async () => {
            let i = {
                bidValue: 0.3 * (10 ** tokenDecimals),
                fake: false,
                secret: 0x123,
                depositValue: 0.03 * (10 ** tokenDecimals)
            }
            await auctionMarketplace.bidAuction(nlgst.address, 1, auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            });
        });

        it('(Account2) 4th Bid auction (0.3 KFIVE) - Higher than the highest bid (2nd bid - 0.2 KFIVE)', async () => {
            let i = {
                bidValue: 0.3 * (10 ** tokenDecimals),
                fake: false,
                secret: 0x123,
                depositValue: 0.3 * (10 ** tokenDecimals),
            }
            await auctionMarketplace.bidAuction(nlgst.address, 1, auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            });
        });

        it('(Account2) 5th Bid auction (0.25 KFIVE / 0.4 KFIVE) - (Bid lower / Deposit higher) than the highest bid (4th bid - 0.25 KFIVE)', async () => {
            let i = {
                bidValue: 0.25 * (10 ** tokenDecimals),
                fake: false,
                secret: 0x123,
                depositValue: 0.4 * (10 ** tokenDecimals)
            }
            await auctionMarketplace.bidAuction(nlgst.address, 1, auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            });
        });

        it('(Account2) 6th Bid auction (0.3 KFIVE) - After bidding stage end', async () => {
            // Wait until bidding end
            assert.isTrue(Math.floor(Date.now()) < auctionBiddingEnd * 1000);
            sleep(auctionBiddingEnd * 1000 - Math.floor(Date.now()));

            let i = {
                bidValue: 0.9 * (10 ** tokenDecimals),
                fake: false,
                secret: 0x123,
                depositValue: 0.9 * (10 ** tokenDecimals),
            }
            await u.assertRevert(auctionMarketplace.bidAuction(nlgst.address, 1, auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            }));
        });

        ////////////////////////////
        // Reveal
        ////////////////////////////
    });
});