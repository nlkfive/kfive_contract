const NLGST = artifacts.require("NLGST")
const KFIVE = artifacts.require("KFIVE")
const AuctionMarketplace = artifacts.require("AuctionMarketplace")
const STORAGE = artifacts.require("MarketplaceStorage")

const eq = assert.equal
const u = require('./utils.js')
var nlgst, kfive, auctionMarketplace, storage;
const { soliditySha3 } = require("web3-utils");
const keccak256 = require('js-sha3').keccak256;
const tokenDecimals = 10;
const secret = "0x0000000000000000000000000000000000000000000000000000000000000001";
var biddingEndedAt = null;
var revealEndedAt = null;
var amount_seller_receive, amount_contractOwner_receive;

function getBlindedBid(value, fake, secret) {
    return soliditySha3(
        { type: 'uint256', value: value },
        { type: 'bool', value: fake },
        { type: 'bytes32', value: secret }
    );
}

function nftAssetInfo(nftAddress, assetId) {
    return soliditySha3(
        { type: 'address', value: nftAddress },
        { type: 'uint256', value: assetId }
    );
}

function amountSellerReceiveAfterAuction(bidValue, OWNER_CUT_PER_MILLION) {
    return (bidValue - bidValue * OWNER_CUT_PER_MILLION / 1000000);
}

function amountContractOwnerReceiveAfterAuction(bidValue, OWNER_CUT_PER_MILLION) {
    return (bidValue * OWNER_CUT_PER_MILLION / 1000000);
}

contract("Auction Marketplace", (accounts) => {
    const root = accounts[0]
    const account1 = accounts[1]
    const account2 = accounts[2]
    const pauser = accounts[3]
    const admin = accounts[4]
    const account5 = accounts[5]
    const account6 = accounts[6]
    const new_owner = accounts[7]
    const account8 = accounts[8]
    const canceller = accounts[9]
    const OFFCHAIN = web3.utils.fromAscii('1')

    const BASE_URL = "https://ipfs.io/ipfs/";
    const OWNER_CUT_PER_MILLION = 1000;
    const issuedTokenAmount = 10 * (10 ** tokenDecimals);
    const publicationFee1 = 0.1 * (10 ** tokenDecimals);

    before(async () => {
        kfive = await KFIVE.new( {from: root} );

        storage = await STORAGE.new( {from: root} );

        nlgst = await NLGST.new(storage.address, BASE_URL, {from: root});

        auctionMarketplace = await AuctionMarketplace.new(kfive.address, storage.address, OWNER_CUT_PER_MILLION, {
            from: root
        });
    });

    afterEach(function() {
        const {currentTest} = this
        if (currentTest.state === 'failed') {
          currentTest.parent.pending = true
        }
    });

    describe('Deployment stage', async () => {
        it('Auction Marketplace information', async () => {
            let o = {
                acceptedToken: kfive.address,
                marketplaceStorage: storage.address,
                owner: root,
                ownerCutPerMillion: 1000,
            }

            const acceptedToken = await auctionMarketplace.acceptedToken();
            eq(acceptedToken, o.acceptedToken);

            const marketplaceStorage = await auctionMarketplace.marketplaceStorage();
            eq(marketplaceStorage, o.marketplaceStorage);

            const owner = await auctionMarketplace.owner();
            eq(owner, o.owner);

            const ownerCutPerMillion = await auctionMarketplace.ownerCutPerMillion();
            eq(ownerCutPerMillion, o.ownerCutPerMillion);
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
        });
    });

    describe('AUCTION 3: cancelAuction / changePublicationFee / admin', async () => {
        var auctionId;
        var auctionBiddingEnd;
        var auctionRevealEnd;

        before(async () => {
            const issuedToken = web3.utils.toHex(issuedTokenAmount);

            await kfive.issue(account1, issuedToken, OFFCHAIN, { from: root });
            await kfive.issue(account2, issuedToken, OFFCHAIN, { from: root });
            await kfive.issue(account5, issuedToken, OFFCHAIN, { from: root });
            await kfive.issue(account6, issuedToken, OFFCHAIN, { from: root });
            await kfive.issue(account8, issuedToken, OFFCHAIN, { from: root });

            const o = {
                account1_balance: 10 * (10 ** tokenDecimals),
                account2_balance: 10 * (10 ** tokenDecimals),
                account5_balance: 10 * (10 ** tokenDecimals),
                account6_balance: 10 * (10 ** tokenDecimals),
                account8_balance: 10 * (10 ** tokenDecimals),
                auctionMarketplace_balance: 0,
            }

            let account1_balance = await kfive.balanceOf(account1, { from: root });
            eq(o.account1_balance, account1_balance.toString());

            let account2_balance = await kfive.balanceOf(account2, { from: root });
            eq(o.account2_balance, account2_balance.toString());

            let account5_balance = await kfive.balanceOf(account5, { from: root });
            eq(o.account5_balance, account5_balance.toString());

            let account6_balance = await kfive.balanceOf(account6, { from: root });
            eq(o.account6_balance, account6_balance.toString());

            let account8_balance = await kfive.balanceOf(account8, { from: root });
            eq(o.account8_balance, account8_balance.toString());

            let auctionMarketplace_balance = await kfive.balanceOf(auctionMarketplace.address, { from: root });
            eq(o.auctionMarketplace_balance, auctionMarketplace_balance.toString());
        });
        
        it('Approve 10 KFIVE to Auction market', async () => {
            const i = {
                amount: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                to: auctionMarketplace.address,
            }

            await kfive.approve(i.to, i.amount, { from: account1 });
            await kfive.approve(i.to, i.amount, { from: account2 });
            await kfive.approve(i.to, i.amount, { from: account5 });
            await kfive.approve(i.to, i.amount, { from: account6 });
            await kfive.approve(i.to, i.amount, { from: account8 });
        });

        it('(Owner) Mint and approve new NFT to (Account1)', async () => {
            const i = {
                nftID: 3,
                nftURI: "3",
            }

            await nlgst.mint(account1, i.nftID, i.nftURI, {
                from: root
            });

            await nlgst.approve(auctionMarketplace.address, i.nftID, {
                from: account1
            });
        });

        it('(Root) Update OrderMarketplace in the Storage', async () => {
            const i = {
                auctionMarketplaceAddress: auctionMarketplace.address
            }

            await storage.updateAuctionMarketplace(i.auctionMarketplaceAddress, {
                from: root
            });
        });

        it('(Account1) Create new auction NFT(3) successfully', async () => {
            const i = {
                nftID: 3,
                nftURI: "3",
                startPriceInWei: 1 * (10 ** tokenDecimals),
                biddingTime: Math.floor(new Date().getTime() / 1000 + 65),
                revealTime: Math.floor(new Date().getTime() / 1000 + 130),
            }

            biddingEndedAt = i.biddingTime;
            revealEndedAt = i.revealTime;

            await storage.updateAuctionMarketplace(auctionMarketplace.address, {
                from: root
            });

            await auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: account1
            });

            let o = {
                rootBalance: (publicationFee1).toString(),
                account1Balance: (issuedTokenAmount - publicationFee1).toString(),
                seller: account1,
                nftID: 3,
                nftAddress: nlgst.address,
            }

            const rootBalance = await kfive.balanceOf(root);
            eq(o.rootBalance, rootBalance.toString());

            const account1Balance = await kfive.balanceOf(account1);
            eq(o.account1Balance, account1Balance.toString());

            const lastestEvent = await auctionMarketplace.getPastEvents("AuctionCreated");
            auctionId = lastestEvent[0].returnValues.auctionId;
            auctionBiddingEnd = parseInt(lastestEvent[0].returnValues.biddingEnd);
            auctionRevealEnd = parseInt(lastestEvent[0].returnValues.revealEnd);

            const seller = lastestEvent[0].returnValues.seller;
            eq(o.seller, seller);

            const startPrice = lastestEvent[0].returnValues.startPriceInWei;
            eq(startPrice, i.startPriceInWei);

            const nftID = lastestEvent[0].returnValues.assetId;
            eq(o.nftID, nftID);

            const nftAddress = lastestEvent[0].returnValues.nftAddress;
            eq(o.nftAddress, nftAddress);
        });

        it('(Account2) Cancel Auction of NFT(3). Cannot because only seller or canceller can cancel', async () => {
            await u.assertRevert(auctionMarketplace.cancelAuction(nftAssetInfo(nlgst.address, 3), auctionId, {
                from: account2
            }));
        });

        it('(Account1) Cancel Auction of NFT(3) successfully', async () => {
            await auctionMarketplace.cancelAuction(nftAssetInfo(nlgst.address, 3), auctionId, {
                from: account1
            });
        });

        it('(Account2) Bid: 5 KFIVE; Deposit: 5 KFIVE; Fake: false. VALID BID. Cannot BID because Auction has been cancelled', async () => {
            const i = {
                bidValue: 5 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("1234"),
                depositValue: 5 * (10 ** tokenDecimals),
            }

            await u.assertRevert(auctionMarketplace.bidAuction(nftAssetInfo(nlgst.address, 3), auctionId, getBlindedBid(i.bidValue, i.fake, secret), i.depositValue, {
                from: account2
            }));
        });

        it('(Admin) Set publication fee (0.5 KFIVE). Cannot because admin does not have ADMIN_ROLE', async () => {
            let i = {
                fee: 0.5 * (10 ** tokenDecimals),
            }
            await u.assertRevert(auctionMarketplace.setPublicationFee(i.fee, {
                from: admin
            }));
        });

        it('(Root) Grant ADMIN_ROLE to admin', async () => {
            await auctionMarketplace.grantRole("0x" + keccak256("ADMIN_ROLE"), admin, {
                from: root
            });
        });

        it('(Admin) Set publication fee (0.5 KFIVE) successfully', async () => {
            let i = {
                fee: 0.5 * (10 ** tokenDecimals),
            }

            await auctionMarketplace.setPublicationFee(i.fee, {
                from: admin
            });

            let o = {
                fee: 0.5 * (10 ** tokenDecimals),
            }
            fee = await auctionMarketplace.publicationFeeInWei();
            eq(fee.toString(), o.fee);
        });

        it('(Account1) Create new auction NFT(3) again successfully', async () => {
            const i = {
                nftID: 3,
                nftURI: "3",
                startPriceInWei: 1 * (10 ** tokenDecimals),
                biddingTime: Math.floor(new Date().getTime() / 1000 + 65),
                revealTime: Math.floor(new Date().getTime() / 1000 + 130),
            }

            biddingEndedAt = i.biddingTime;
            revealEndedAt = i.revealTime;

            await auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: account1
            });

            let o = {
                rootBalance: (publicationFee1 + 0.5 * (10 ** tokenDecimals)).toString(),
                account1Balance: (issuedTokenAmount - publicationFee1 - 0.5 * (10 ** tokenDecimals)).toString(),
                seller: account1,
                nftID: 3,
                nftAddress: nlgst.address,
            }

            const rootBalance = await kfive.balanceOf(root);
            eq(o.rootBalance, rootBalance.toString());

            const account1Balance = await kfive.balanceOf(account1);
            eq(o.account1Balance, account1Balance.toString());

            const lastestEvent = await auctionMarketplace.getPastEvents("AuctionCreated");
            auctionId = lastestEvent[0].returnValues.auctionId;
            auctionBiddingEnd = parseInt(lastestEvent[0].returnValues.biddingEnd);
            auctionRevealEnd = parseInt(lastestEvent[0].returnValues.revealEnd);

            const seller = lastestEvent[0].returnValues.seller;
            eq(o.seller, seller);

            const startPrice = lastestEvent[0].returnValues.startPriceInWei;
            eq(startPrice, i.startPriceInWei);

            const nftID = lastestEvent[0].returnValues.assetId;
            eq(o.nftID, nftID);

            const nftAddress = lastestEvent[0].returnValues.nftAddress;
            eq(o.nftAddress, nftAddress);
        });

        it('(Account2) Bid: 5 KFIVE; Deposit: 6 KFIVE; Fake: false', async () => {
            const i = {
                bidValue: 5 * (10 ** tokenDecimals),
                fake: true,
                secret: web3.utils.fromAscii("1234"),
                depositValue: 6 * (10 ** tokenDecimals),
                nftID: 3,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(i.nftAddress, i.nftID), auctionId, getBlindedBid(i.bidValue, i.fake, secret), i.depositValue, {
                from: account2
            });

            const o = {
                kfive_account2_balance: (10 - 6) * (10 ** tokenDecimals),
            }

            kfive_account2_balance = await kfive.balanceOf(account2, {from: root});
            eq(o.kfive_account2_balance, kfive_account2_balance.toString());
        });

        it('(Canceller) Cancel Auction of NFT(3). Cannot because has no CANCEL_ROLE', async () => {
            await u.assertRevert(auctionMarketplace.cancelAuction(nftAssetInfo(nlgst.address, 3), auctionId, {
                from: canceller
            }));
        });

        it('(Root) Grant CANCEL_ROLE to admin', async () => {
            await auctionMarketplace.grantRole("0x" + keccak256("CANCEL_ROLE"), canceller, {
                from: root
            });
        });

        it('(Canceller) Cancel Auction of NFT(3) successfully', async () => {
            await auctionMarketplace.cancelAuction(nftAssetInfo(nlgst.address, 3), auctionId, {
                from: canceller
            });
        });

        it('Wait until Bidding end', async () => {
            while (Math.floor(new Date().getTime() / 1000 < biddingEndedAt)) {}
        });

        it('(Root) Cancel Auction of NFT(3) again. Cannot because it has been cancelled', async () => {
            await u.assertRevert(auctionMarketplace.cancelAuction(nftAssetInfo(nlgst.address, 3), auctionId, {
                from: root
            }));
        });

        it('(Account2) reveal to withdraw', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 3,
                value_account2: 5 * (10 ** tokenDecimals),
                fakeFalse: true,
                secret: web3.utils.fromAscii("1234"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account2],
                [i.fakeFalse],
                [secret], {
                from: account2
            });

            const account2_balance = await kfive.balanceOf(account2, { from: root })
            eq(account2_balance.toString(), 10 * (10 ** tokenDecimals));
        });

        it('Account1 should be the owner of NFT(1)', async () => {
            const o = {
                nft1_owner: account1,
            }

            const nft1_owner = await nlgst.ownerOf(3, { from: root })

            eq(o.nft1_owner, nft1_owner);
        });
    });
});