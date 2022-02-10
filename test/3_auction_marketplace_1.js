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
    const delay = ms => new Promise(res => setTimeout(res, ms));

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

    describe('AUCTION 1: Simple auction. Single Bid / Account', async () => {
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
            await kfive.approve(i.to, i.amount, { from: account8 });
        });

        it('(Owner) Mint new NFT to (Account1)', async () => {
            const i = {
                nftID: 1,
                nftURI: "1",
            }

            await nlgst.mint(account1, i.nftID, i.nftURI, {
                from: root
            });
        });

        it('(Account1) Create new auction. Cannot because NFT(1) has not approved to Auction Marketplace', async () => {
            const i = {
                nftID: 1,
                nftURI: "1",
                startPriceInWei: 1 * (10 ** tokenDecimals),
                biddingTime: Math.floor(new Date().getTime() / 1000 + 4000),
                revealTime: Math.floor(new Date().getTime() / 1000 + 8000),
            }

            await u.assertRevert(auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: account1
            }));
        });

        it('(Account1) Approve NFT to Auction Marketplace', async () => {
            const i = {
                nftID: 1,
            }

            await nlgst.approve(auctionMarketplace.address, i.nftID, {
                from: account1
            });
        });

        it('(Account1) Create new auction. Cannot because have not updated Auction address in Storage', async () => {
            const i = {
                nftID: 1,
                nftURI: "1",
                startPriceInWei: 1 * (10 ** tokenDecimals),
                biddingTime: Math.floor(new Date().getTime() / 1000 + 4000),
                revealTime: Math.floor(new Date().getTime() / 1000 + 8000),
            }

            await u.assertRevert(auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: account1
            }));
        });

        it('(Account1) Update AuctionMarketplace in the Storage. Cannot because only Admin can update', async () => {
            const i = {
                auctionMarketplaceAddress: auctionMarketplace.address
            }

            await u.assertRevert(storage.updateAuctionMarketplace(i.auctionMarketplaceAddress, {
                from: account1
            }));
        });

        it('(Root) Update OrderMarketplace in the Storage', async () => {
            const i = {
                auctionMarketplaceAddress: auctionMarketplace.address
            }

            await storage.updateAuctionMarketplace(i.auctionMarketplaceAddress, {
                from: root
            });
        });

        it('(Root) Create new auction. Cannot because only Owner of NFT can create auction', async () => {
            const i = {
                nftID: 1,
                nftURI: "1",
                startPriceInWei: 1 * (10 ** tokenDecimals),
                biddingTime: Math.floor(new Date().getTime() / 1000 + 4000),
                revealTime: Math.floor(new Date().getTime() / 1000 + 8000),
            }

            await u.assertRevert(auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: root
            }));
        });

        it('(Account1) Create new auction NFT(1) successfully', async () => {
            const i = {
                nftID: 1,
                nftURI: "1",
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
                rootBalance: publicationFee1.toString(),
                account1Balance: (issuedTokenAmount - publicationFee1).toString(),
                seller: account1,
                nftID: 1,
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

        ////////////////////////////
        // Bidding
        ////////////////////////////
        it('(Account2) Bid wrong nftID', async () => {
            const i = {
                bidValue: 0.1 * (10 ** tokenDecimals),
                fake: false,
                depositValue: 0.1 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await u.assertRevert(auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, secret),
                i.depositValue, {
                from: account2
            }));
        });

        it('(Account2) Bid: 11 KFIVE; Deposit: 11 KFIVE; Fake: false. Cannot because BID value + deposit over balance', async () => {
            const i = {
                bidValue: 11 * (10 ** tokenDecimals),
                fake: false,
                depositValue: 11 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await u.assertRevert(auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, secret),
                i.depositValue, {
                from: account2
            }));
        });

        it('(Account2) Bid: 11 KFIVE; Deposit: 20 KFIVE; Fake: false. Cannot because BID value over balance', async () => {
            const i = {
                bidValue: 11 * (10 ** tokenDecimals),
                fake: false,
                depositValue: 20 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await u.assertRevert(auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, secret),
                i.depositValue, {
                from: account2
            }));
        });

        it('(Account5) Bid: 2 KFIVE; Deposit: 4 KFIVE; Fake: false.', async () => {
            const i = {
                bidValue: 2 * (10 ** tokenDecimals),
                fake: false,
                depositValue: 4 * (10 ** tokenDecimals),
                nftID: 1,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, secret),
                i.depositValue, {
                from: account5
            });
        });

        it('(Account5) Bid: 2 KFIVE; Deposit: 1 KFIVE; Fake: false. VALID BID. Cannot because BID value + Deposit same as previous, already bidded', async () => {
            const i = {
                bidValue: 2 * (10 ** tokenDecimals),
                fake: false,
                depositValue: 4 * (10 ** tokenDecimals),
                nftID: 1,
                nftAddress: nlgst.address,
            }

            await u.assertRevert(auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, secret),
                i.depositValue, {
                from: account5
            }));
        });

        it('(Account8) Bid: 3 KFIVE; Deposit: 4 KFIVE; Fake: false', async () => {
            const i = {
                bidValue: 3 * (10 ** tokenDecimals),
                fake: false,
                depositValue: 4 * (10 ** tokenDecimals),
                nftID: 1,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, secret),
                i.depositValue, {
                from: account8
            });
        });

        it('Reveal Stage Auction 1. Cannot because bidding time has not expired', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 1,
                value_account5: 2 * (10 ** tokenDecimals),
                fake: false,
            }

            await u.assertRevert(auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account5],
                [i.fake],
                [secret], {
                from: account5
            }));
        });

        it('Wait until Bidding end', async () => {
            while (Math.floor(new Date().getTime() / 1000 < biddingEndedAt)) {}
        });

        it('(Account8) Bid: 8 KFIVE; Deposit: 8 KFIVE; Fake: false. VALID BID. Cannot BID because bidding time has expired', async () => {
            const i = {
                bidValue: 3 * (10 ** tokenDecimals),
                fake: false,
                depositValue: 5 * (10 ** tokenDecimals),
                nftID: 1,
                nftAddress: nlgst.address,
            }

            await u.assertRevert(auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, secret),
                i.depositValue, {
                from: account8
            }));
        });

        ////////////////////////////
        // Reveal
        ////////////////////////////
        it('BEFORE REVEAL Auction 1: Check balance of account2, account5, account8 and auctionMarketplace', async () => {
            const o = {
                kfive_account2_balance_before: (10) * (10 ** tokenDecimals),
                kfive_account5_balance_before: (10 - 4) * (10 ** tokenDecimals),
                kfive_account8_balance_before: (10 - 4) * (10 ** tokenDecimals),
                kfive_market_balance_before_reveal: (8) * (10 ** tokenDecimals),
            }

            kfive_account2_balance_before = await kfive.balanceOf(account2, {from: root});
            kfive_account5_balance_before = await kfive.balanceOf(account5, {from: root});
            kfive_account8_balance_before = await kfive.balanceOf(account8, {from: root});
            kfive_market_balance_before_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account2_balance_before, kfive_account2_balance_before.toString());
            eq(o.kfive_account5_balance_before, kfive_account5_balance_before.toString());
            eq(o.kfive_account8_balance_before, kfive_account8_balance_before.toString());
            eq(o.kfive_market_balance_before_reveal, kfive_market_balance_before_reveal.toString());
        });

        // REVEAL
        it('(Account2) Reveal Bid. Cannot because Account2 does not bid', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 1,
                value_account5: 2 * (10 ** tokenDecimals),
                fake: false,
            }

            await u.assertRevert(auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account5],
                [i.fake],
                [secret], {
                from: account2
            }));
        });

        it('(Account5) Reveal Bid', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 1,
                value_account5: 2 * (10 ** tokenDecimals),
                fake: false,
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account5],
                [i.fake],
                [secret], {
                from: account5
            });

            const o = {
                kfive_account5_balance: 8 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 6 * (10 ** tokenDecimals),
            }

            kfive_account5_balance = await kfive.balanceOf(account5, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account5_balance, kfive_account5_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });

        it('(Account5) Reveal Bid of account5 again with the same value. Cannot because value has been revealled', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 1,
                value_account5: 2 * (10 ** tokenDecimals),
                fake: false,
            }

            await u.assertRevert(auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account5],
                [i.fake],
                [secret], {
                from: account5
            }));
        });

        it('(Account8) Reveal Stage Auction 1', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 1,
                value_account8: 3 * (10 ** tokenDecimals),
                fake: false,
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account8],
                [i.fake],
                [secret], {
                from: account8
            });

            const o = {
                kfive_account8_balance: 7 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 5 * (10 ** tokenDecimals),
            }

            kfive_account8_balance = await kfive.balanceOf(account8, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account8_balance, kfive_account8_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });

        it('Wait until Reveal end', async () => {
            while (Math.floor(new Date().getTime() / 1000 < revealEndedAt)) {}
        });

        ////////////////////////////
        // Withdraw
        ////////////////////////////
        it('AuctionEnd. Account1 call auctionEnd', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 1,
            }

            await auctionMarketplace.auctionEnd(i.nftAddress, i.nftID, auctionId, { 
                from: account1 });
        });

        it('Account8 should be the owner of NFT(1)', async () => {
            const o = {
                nft1_owner: account8,
            }

            const nft1_owner = await nlgst.ownerOf(1, { from: root })
            eq(o.nft1_owner, nft1_owner);
        });

        it('Check balance of each accounts' + '\n' +
        '               Root: 0.103 KFIVE. PublicationFee 0.1 KFIVE + ownerCutPerMillion 0.003 KFIVE' + '\n' +
        '               Account1: 12.897 KFIVE. PublicationFee 0.1 KFIVE + nft1Amount 2.997 KFIVE' + '\n' +
        '               Account5:  6 KFIVE. Already withdraw' + '\n' +
        '               Account8:  7 KFIVE. Valid Bid and Highest Bid -> No refund' + '\n' +
        '               AuMarket:  2 KFIVE', async () => {

            amount_seller_receive = amountSellerReceiveAfterAuction(3, OWNER_CUT_PER_MILLION);
            amount_contractOwner_receive = amountContractOwnerReceiveAfterAuction(3, OWNER_CUT_PER_MILLION);

            const o = {
                kfive_root_balance: (0.103) * (10 ** tokenDecimals),
                kfive_account1_balance: (10 - 0.1 + amount_seller_receive) * (10 ** tokenDecimals),
                kfive_account2_balance: 10 * (10 ** tokenDecimals),
                kfive_account5_balance: 8 * (10 ** tokenDecimals),
                kfive_account8_balance: 7 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 2 * (10 ** tokenDecimals),
            }

            kfive_root_balance = await kfive.balanceOf(root, {from: root});
            kfive_account1_balance = await kfive.balanceOf(account1, {from: root});
            kfive_account2_balance = await kfive.balanceOf(account2, {from: root});
            kfive_account5_balance = await kfive.balanceOf(account5, {from: root});
            kfive_account8_balance = await kfive.balanceOf(account8, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_root_balance, kfive_root_balance.toString());
            eq(o.kfive_account1_balance, kfive_account1_balance.toString());
            eq(o.kfive_account2_balance, kfive_account2_balance.toString());
            eq(o.kfive_account5_balance, kfive_account5_balance.toString());
            eq(o.kfive_account8_balance, kfive_account8_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });

        it('Withdraw Stage Auction 1. Account2 withdraw. Cannot because account2 has no bid', async () => {
            await u.assertRevert(auctionMarketplace.withdraw(auctionId, { from: account2 }));
        });

        it('Withdraw Stage Auction 1. Account5 withdraw', async () => {
            await auctionMarketplace.withdraw(auctionId, { from: account5 });

            const o = {
                kfive_account5_balance: 10 * (10 ** tokenDecimals),
                kfive_auctionMarketplace_balance: 0,
            }

            kfive_account5_balance = await kfive.balanceOf(account5, {from: root});
            eq(o.kfive_account5_balance, kfive_account5_balance.toString());

            kfive_account1_balance = await kfive.balanceOf(account1, {from: root});

            kfive_auctionMarketplace_balance = await kfive.balanceOf(auctionMarketplace.address, {from: root});
            eq(o.kfive_auctionMarketplace_balance, kfive_auctionMarketplace_balance.toString());
        });

        it('Withdraw Stage Auction 1. Account8 withdraw. Cannot because account8 has the highest BidValue', async () => {
            await u.assertRevert(auctionMarketplace.withdraw(auctionId, { from: account8 }));
        });
    });
});