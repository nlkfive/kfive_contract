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
var biddingEndedAt = null;
var revealEndedAt = null;
var amount_seller_receive, amount_contractOwner_receive;

function getBlindedBid(value, fake, secret) {
    return soliditySha3(
        value, fake, secret
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
        ////////////////////////////
        // Information
        // Account1 CREATE AUCTION, startPrice = 1 KFIVE
        // Bid:
        // Account2 BID: 0.1 KFIVE, deposit: 0.1 KFIVE, fake: false
        // Account5 BID:   2 KFIVE, deposit:   2 KFIVE, fake: false
        // Account8 BID:   3 KFIVE, deposit:   3 KFIVE, fake: false
        // Reveal:
        // + Account5 reveal
        // + Account8 reveal
        // Withdraw:
        // + Account2 withdraw
        // + Account5 withdraw
        ////////////////////////////

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

        after(async () => {
            const owner = await kfive.owner({ from: root });
            const root_balance = await kfive.balanceOf(root, { from: root });
            const account1_balance = await kfive.balanceOf(account1, { from: root });
            const account2_balance = await kfive.balanceOf(account2, { from: root });
            const account5_balance = await kfive.balanceOf(account5, { from: root });
            const account6_balance = await kfive.balanceOf(account6, { from: root });
            const account8_balance = await kfive.balanceOf(account8, { from: root });
            const auctionMarketplace_balance = await kfive.balanceOf(auctionMarketplace.address, { from: root });

            await kfive.redeem(root, root_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account1, account1_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account2, account2_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account5, account5_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account6, account6_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account8, account8_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(auctionMarketplace.address, auctionMarketplace_balance, OFFCHAIN, { from: owner });
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
                biddingTime: Math.floor(new Date().getTime() / 1000 + 120),
                revealTime: Math.floor(new Date().getTime() / 1000 + 240),
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
        it('(Account2) Bid: 5 KFIVE; Deposit: 5 KFIVE; Fake: false. Cannot because Bid wrong nftID', async () => {
            const i = {
                bidValue: 0.1 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 0.1 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await u.assertRevert(auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, i.secret),
                i.depositValue, {
                from: account2
            }));
        });

        it('(Account2) Bid: 11 KFIVE; Deposit: 11 KFIVE; Fake: false. Cannot because BID value + deposit over balance', async () => {
            const i = {
                bidValue: 0.1 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 0.1 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await u.assertRevert(auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, i.secret),
                i.depositValue, {
                from: account2
            }));
        });

        it('(Account2) Bid: 11 KFIVE; Deposit: 5 KFIVE; Fake: false. Cannot because BID value over balance', async () => {
            const i = {
                bidValue: 0.1 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 0.1 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await u.assertRevert(auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, i.secret),
                i.depositValue, {
                from: account2
            }));
        });

        it('(Account2) Bid: 0.1 KFIVE; Deposit: 0.1 KFIVE; Fake: false. Cannot because BidValue is LOWER than startPrice', async () => {
            const i = {
                bidValue: 0.1 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 0.1 * (10 ** tokenDecimals),
                nftID: 1,
                nftAddress: nlgst.address,
            }

            await u.assertRevert(auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, i.secret),
                i.depositValue, {
                from: account2
            }));
        });

        it('(Account5) Bid: 2 KFIVE; Deposit: 1 KFIVE; Fake: false. VALID BID. BidValue is HIGHER than startPrice', async () => {
            const i = {
                bidValue: 2 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 1 * (10 ** tokenDecimals),
                nftID: 1,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, i.secret),
                i.depositValue, {
                from: account5
            });
        });

        it('(Account5) Bid: 2 KFIVE; Deposit: 1 KFIVE; Fake: false. VALID BID. Cannot because BID value + Deposit same as previous, already bidded', async () => {
            const i = {
                bidValue: 2 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 1 * (10 ** tokenDecimals),
                nftID: 1,
                nftAddress: nlgst.address,
            }

            await u.assertRevert(auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, i.secret),
                i.depositValue, {
                from: account5
            }));
        });

        it('(Account8) Bid: 3 KFIVE; Deposit: 3 KFIVE; Fake: false. VALID BID. BidValue is HIGHER than Account2', async () => {
            const i = {
                bidValue: 3 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 3 * (10 ** tokenDecimals),
                nftID: 1,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, i.secret),
                i.depositValue, {
                from: account8
            });
        });

        it('Reveal Stage Auction 1. Cannot because bidding time has not expired', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 1,
                value_account2: 0.1 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
            }

            await u.assertRevert(auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account2],
                [i.fake],
                [i.secret], {
                from: account2
            }));
        });

        it('Wait until Bidding end', async () => {
            while (Math.floor(new Date().getTime() / 1000 < biddingEndedAt)) {}
        });

        it('(Account8) Bid: 8 KFIVE; Deposit: 8 KFIVE; Fake: false. VALID BID. Cannot BID because bidding time has expired', async () => {
            const i = {
                bidValue: 3 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 3 * (10 ** tokenDecimals),
                nftID: 1,
                nftAddress: nlgst.address,
            }

            await u.assertRevert(auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, i.secret),
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
                kfive_account5_balance_before: (10 - 1) * (10 ** tokenDecimals),
                kfive_account8_balance_before: (10 - 3) * (10 ** tokenDecimals),
                kfive_market_balance_before_reveal: (4) * (10 ** tokenDecimals),
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
        it('(Account2) Reveal Bid of account5. Cannot because only account5 can reveal its bid', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 1,
                value_account5: 2 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
            }

            await u.assertRevert(auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account5],
                [i.fake],
                [i.secret], {
                from: account2
            }));
        });

        it('(Root) Reveal Bid of account5. Cannot because only account5 can reveal its bid', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 1,
                value_account5: 2 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
            }

            await u.assertRevert(auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account5],
                [i.fake],
                [i.secret], {
                from: root
            }));
        });

        it('(Account5) Reveal Bid of account5', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 1,
                value_account5: 2 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account5],
                [i.fake],
                [i.secret], {
                from: account5
            });

            const o = {
                kfive_account5_balance: 9 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 4 * (10 ** tokenDecimals),
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
                secret: web3.utils.fromAscii("secret"),
            }

            await u.assertRevert(auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account5],
                [i.fake],
                [i.secret], {
                from: account5
            }));
        });

        it('(Account8) Reveal Stage Auction 1', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 1,
                value_account8: 3 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account8],
                [i.fake],
                [i.secret], {
                from: account8
            });

            const o = {
                kfive_account8_balance: 7 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 4 * (10 ** tokenDecimals),
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
        '               Account2: 10 KFIVE. Invalid Bid' + '\n' +
        '               Account5: 10 KFIVE. Already withdraw' + '\n' +
        '               Account8:  7 KFIVE. Valid Bid and Highest Bid -> No refund' + '\n' +
        '               AuMarket:  0 KFIVE', async () => {

            amount_seller_receive = amountSellerReceiveAfterAuction(3, OWNER_CUT_PER_MILLION);
            amount_contractOwner_receive = amountContractOwnerReceiveAfterAuction(3, OWNER_CUT_PER_MILLION);

            const o = {
                kfive_root_balance: (0.103) * (10 ** tokenDecimals),
                kfive_account1_balance: (10 - 0.1 + amount_seller_receive) * (10 ** tokenDecimals),
                kfive_account2_balance: 9 * (10 ** tokenDecimals),
                kfive_account5_balance: 10 * (10 ** tokenDecimals),
                kfive_account8_balance: 7 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 1 * (10 ** tokenDecimals),
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
                kfive_auctionMarketplace_balance: 3 * (10 ** tokenDecimals),
            }

            kfive_account5_balance = await kfive.balanceOf(account5, {from: root});
            eq(o.kfive_account5_balance, kfive_account5_balance.toString());

            kfive_account1_balance = await kfive.balanceOf(account1, {from: root});
            console.log("kfive_account1_balance", kfive_account1_balance);

            kfive_auctionMarketplace_balance = await kfive.balanceOf(auctionMarketplace.address, {from: root});
            eq(o.kfive_auctionMarketplace_balance, kfive_auctionMarketplace_balance.toString());
        });

        it('Withdraw Stage Auction 1. Account8 withdraw. Cannot because account8 has the highest BidValue', async () => {
            await u.assertRevert(auctionMarketplace.withdraw(auctionId, { from: account8 }));
        });
    });

    describe('AUCTION 2: Multiple Bids/Account + Pause Auction / fake = true / Invalid Bid', async () => {
        ////////////////////////////
        // Information
        // Account1 CREATE AUCTION, startPrice = 1 KFIVE
        // Bid:
        // Account2 BID:   1 KFIVE, deposit:   1 KFIVE, fake: false -> VALID BID (1)
        // Account2 BID:   3 KFIVE, deposit:   1 KFIVE, fake: false -> VALID BID (2)(*)
        // Account5 BID:   2 KFIVE, deposit:   2 KFIVE, fake: false -> VALID BID (3)
        // Account5 BID:   5 KFIVE, deposit:   5 KFIVE, fake: true -> INVALID BID (4)
        // Account6 BID:   1 KFIVE, deposit:   8 KFIVE, fake: false ->  VALID BID (5)
        // Account8 BID:   3 KFIVE, deposit:   2 KFIVE, fake: true -> INVALID BID (6)
        // Reveal:
        // + Account5 reveal (4)
        // + Account2 reveal (2)
        // + Account2 reveal (1) -> Refund
        // + Account5 reveal (3) -> Refund
        // + Account6 reveal (5) -> Refund
        // + Account8 reveal (6) -> Refund
        // Withdraw:
        // + Account5 withdraw (4)
        // + Account8 withdraw (6)
        ////////////////////////////

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

        after(async () => {
            const owner = await kfive.owner({ from: root });
            const root_balance = await kfive.balanceOf(root, { from: root });
            const account1_balance = await kfive.balanceOf(account1, { from: root });
            const account2_balance = await kfive.balanceOf(account2, { from: root });
            const account5_balance = await kfive.balanceOf(account5, { from: root });
            const account6_balance = await kfive.balanceOf(account6, { from: root });
            const account8_balance = await kfive.balanceOf(account8, { from: root });
            const auctionMarketplace_balance = await kfive.balanceOf(auctionMarketplace.address, { from: root });

            await kfive.redeem(root, root_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account1, account1_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account2, account2_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account5, account5_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account6, account6_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(account8, account8_balance, OFFCHAIN, { from: owner });
            await kfive.redeem(auctionMarketplace.address, auctionMarketplace_balance, OFFCHAIN, { from: owner });
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

        it('(Owner) Mint new NFT to (Account1)', async () => {
            const i = {
                nftID: 2,
                nftURI: "2",
            }

            await nlgst.mint(account1, i.nftID, i.nftURI, {
                from: root
            });
        });

        it('(Account1) Approve NFT to market', async () => {
            const i = {
                nftID: 2,
            }

            await nlgst.approve(auctionMarketplace.address, i.nftID, {
                from: account1
            });
        });

        it('(Account1) Create new auction NFT(2) successfully', async () => {
            const i = {
                nftID: 2,
                nftURI: "2",
                startPriceInWei: 1 * (10 ** tokenDecimals),
                biddingTime: Math.floor(new Date().getTime() / 1000 + 120),
                revealTime: Math.floor(new Date().getTime() / 1000 + 240),
            }

            biddingEndedAt = i.biddingTime;
            revealEndedAt = i.revealTime;

            await auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: account1
            });

            let o = {
                rootBalance: (publicationFee1).toString(),
                account1Balance: (issuedTokenAmount - publicationFee1).toString(),
                seller: account1,
                nftID: 2,
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
        it('(Account2) Bid: 1 KFIVE; Deposit: 1 KFIVE; Fake: false. BidValue is EQUAL to startPrice', async () => {
            const i = {
                bidValue: 1 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 1 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(i.nftAddress, i.nftID), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            });
        });

        it('(Owner) Pause Auction Marketplace', async () => {
            await auctionMarketplace.pause({
                from: root
            });
        });

        it('(Account2) Bid: 3 KFIVE; Deposit: 1 KFIVE; Fake: false. Cannot because contract is pausing', async () => {
            const i = {
                bidValue: 3 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 1 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await u.assertRevert(auctionMarketplace.bidAuction(nftAssetInfo(i.nftAddress, i.nftID), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            }));
        });

        it('(Pauser) UnPause Auction Marketplace. Cannot because pauser has no PAUSER_ROLE', async () => {
            await u.assertRevert(auctionMarketplace.pause({
                from: pauser
            }));
        });

        it('(Root) Grant role PAUSER_ROLE to (Pauser)', async () => {
            await auctionMarketplace.grantRole("0x" + keccak256("PAUSER_ROLE"), pauser, {
                from: root
            })
        });

        it('(Pauser) UnPause Auction Marketplace successfully', async () => {
            await auctionMarketplace.unpause({
                from: pauser
            });
        });

        it('(Account2) Bid: 3 KFIVE; Deposit: 1 KFIVE; Fake: false. BidValue is HIGHER than previous bid (account2)', async () => {
            const i = {
                bidValue: 3 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 1 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(i.nftAddress, i.nftID), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            });
        });

        it('(Account5) Bid: 5 KFIVE; Deposit: 5 KFIVE; Fake: true. BidValue is Highest but fake is true', async () => {
            const i = {
                bidValue: 5 * (10 ** tokenDecimals),
                fake: true,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 5 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(i.nftAddress, i.nftID), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account5
            });
        });

        it('(Account5) Bid: 2 KFIVE; Deposit: 2 KFIVE; Fake: false. BidValue is HIGHER than startPrice but LOWER than account2', async () => {
            const i = {
                bidValue: 2 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 2 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(i.nftAddress, i.nftID), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account5
            });
        });

        it('(Account6) Bid: 1 KFIVE; Deposit: 8 KFIVE; Fake: false. BidValue is EQUAL to startPrice, but deposit is the HIGHEST', async () => {
            const i = {
                bidValue: 1 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 8 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(i.nftAddress, i.nftID), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account6
            });
        });

        it('(Account8) Bid: 3 KFIVE; Deposit: 2 KFIVE; Fake: true. BidValue is EQUAL to account2, but fake is true', async () => {
            const i = {
                bidValue: 3 * (10 ** tokenDecimals),
                fake: true,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 2 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(i.nftAddress, i.nftID), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account8
            });
        });

        it('Wait until Bidding end', async () => {
            while (Math.floor(new Date().getTime() / 1000 < biddingEndedAt)) {}
        });

        ////////////////////////////
        // Reveal
        ////////////////////////////
        it('BEFORE REVEAL Auction 2: Check balance of account2, account5, account6, account8 and auctionMarketplace', async () => {
            const o = {
                kfive_account2_balance_before: (10 - 2) * (10 ** tokenDecimals),
                kfive_account5_balance_before: (10 - 7) * (10 ** tokenDecimals),
                kfive_account6_balance_before: (10 - 8) * (10 ** tokenDecimals),
                kfive_account8_balance_before: (10 - 2) * (10 ** tokenDecimals),
                kfive_market_balance_before_reveal: 19 * (10 ** tokenDecimals),
            }

            kfive_account2_balance_before = await kfive.balanceOf(account2, {from: root});
            kfive_account5_balance_before = await kfive.balanceOf(account5, {from: root});
            kfive_account6_balance_before = await kfive.balanceOf(account6, {from: root});
            kfive_account8_balance_before = await kfive.balanceOf(account8, {from: root});
            kfive_market_balance_before_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account2_balance_before, kfive_account2_balance_before.toString());
            eq(o.kfive_account5_balance_before, kfive_account5_balance_before.toString());
            eq(o.kfive_account6_balance_before, kfive_account6_balance_before.toString());
            eq(o.kfive_account8_balance_before, kfive_account8_balance_before.toString());
            eq(o.kfive_market_balance_before_reveal, kfive_market_balance_before_reveal.toString());
        });

        it('(Account5) reveal bids with fake = true. Account gets the refund', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 2,
                auctionID: auctionId,
                value_account5_2: 5 * (10 ** tokenDecimals),
                fakeFalse: false,
                fakeTrue: true,
                secret: web3.utils.fromAscii("secret"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account5_2],
                [i.fakeTrue],
                [i.secret], {
                from: account5
            });

            const o = {
                kfive_account5_balance: 8 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 14 * (10 ** tokenDecimals),
            }

            kfive_account5_balance = await kfive.balanceOf(account5, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account5_balance, kfive_account5_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });

        it('(Account2) reveal bids', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 2,
                auctionID: auctionId,
                value_account2_1: 3 * (10 ** tokenDecimals),
                value_account2_2: 1 * (10 ** tokenDecimals),
                fakeFalse: false,
                fakeTrue: true,
                secret: web3.utils.fromAscii("secret"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account2_1, i.value_account2_2],
                [i.fakeFalse, i.fakeFalse],
                [i.secret, i.secret], {
                from: account2
            });

            const o = {
                kfive_account2_balance: 9 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 13 * (10 ** tokenDecimals),
            }

            kfive_account2_balance = await kfive.balanceOf(account2, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account2_balance, kfive_account2_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });

        it('(Account5) reveal bids', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 2,
                auctionID: auctionId,
                value_account5_1: 2 * (10 ** tokenDecimals),
                fakeFalse: false,
                fakeTrue: true,
                secret: web3.utils.fromAscii("secret"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account5_1],
                [i.fakeFalse],
                [i.secret], {
                from: account5
            });

            const o = {
                kfive_account5_balance: 10 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 11 * (10 ** tokenDecimals),
            }

            kfive_account5_balance = await kfive.balanceOf(account5, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account5_balance, kfive_account5_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });

        it('(Account6) reveal bids', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 2,
                auctionID: auctionId,
                value_account6: 1 * (10 ** tokenDecimals),
                fakeFalse: false,
                fakeTrue: true,
                secret: web3.utils.fromAscii("secret"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account6],
                [i.fakeFalse],
                [i.secret], {
                from: account6
            });

            const o = {
                kfive_account6_balance: 10 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 3 * (10 ** tokenDecimals),
            }

            kfive_account6_balance = await kfive.balanceOf(account6, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account6_balance, kfive_account6_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });

        it('(Account6) reveal non-existed bids. Cannot because the bid is not existed', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 2,
                auctionID: auctionId,
                value_account6: 10 * (10 ** tokenDecimals),
                fakeFalse: false,
                fakeTrue: true,
                secret: web3.utils.fromAscii("secret"),
            }

            await u.assertRevert(auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account6],
                [i.fakeFalse],
                [i.secret], {
                from: account6
            }));
        });

        it('(Account8) reveal bids', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 2,
                auctionID: auctionId,
                value_account8: 3 * (10 ** tokenDecimals),
                fakeFalse: false,
                fakeTrue: true,
                secret: web3.utils.fromAscii("secret"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account8],
                [i.fakeTrue],
                [i.secret], {
                from: account8
            });

            const o = {
                kfive_account8_balance: 10 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 1 * (10 ** tokenDecimals),
            }

            kfive_account8_balance = await kfive.balanceOf(account6, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account8_balance, kfive_account8_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });

        it('Wait until Reveal end', async () => {
            while (Math.floor(new Date().getTime() / 1000 < revealEndedAt)) {}
        });
        
        it('AuctionEnd. Account1 call acutionEnd', async () => {
            await auctionMarketplace.auctionEnd(nlgst.address, 1, auctionId, { from: account1 });
        });

        it('Account2 should be the owner of NFT(2)', async () => {
            const o = {
                nft2_owner: account2,
            }

            const nft2_owner = await nlgst.ownerOf(2, { from: root })

            eq(o.nft2_owner, nft2_owner);
        });

        it('AFTER REVEAL Auction 2: Check balance of account2, account5, account6, account8 and auctionMarketplace' + '\n' +
        '                           Root: PublicationFee 0.1 + 0.003 (Auction2)' + '\n' +
        '                           Account1: PublicationFee 0.1 + 2.997 (Auction2)' + '\n' +
        '                           Account2 (1st): Invalid Bid but not highest -> Refund 0.1 KFIVE' + '\n' +
        '                           Account2 (2nd): Valid Bid but not highest -> Refund' + '\n' +
        '                           Account2 (3nd): Valid Bid and Highest Bid -> No refund and transfer 2 more KFIVE to Auction -> Refund' + '\n' +
        '                           Account5 (1st): Valid Bid but not highest -> Refund 2 KFIVE' + '\n' +
        '                           Account5 (2nd): Invalid Bid -> Refund' + '\n' +
        '                           Account6: Valid Bid but not highest -> Refund 8 KFIVE' + '\n' +
        '                           Account8: Invalid Bid -> Refund', async () => {

            amount_seller_receive = amountSellerReceiveAfterAuction(3, OWNER_CUT_PER_MILLION);
            amount_contractOwner_receive = amountContractOwnerReceiveAfterAuction(3, OWNER_CUT_PER_MILLION);

            const o = {
                kfive_root_balance: (0.1 + amount_contractOwner_receive) * (10 ** tokenDecimals),
                kfive_account1_balance: (10 - 0.1 + amount_seller_receive) * (10 ** tokenDecimals),
                kfive_account2_balance: 7 * (10 ** tokenDecimals),
                kfive_account5_balance: 10 * (10 ** tokenDecimals),
                kfive_account6_balance: 10 * (10 ** tokenDecimals),
                kfive_account8_balance: 10 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 0,
            }

            kfive_root_balance = await kfive.balanceOf(root, {from: root});
            kfive_account1_balance = await kfive.balanceOf(account1, {from: root});
            kfive_account2_balance = await kfive.balanceOf(account2, {from: root});
            kfive_account5_balance = await kfive.balanceOf(account5, {from: root});
            kfive_account6_balance = await kfive.balanceOf(account6, {from: root});
            kfive_account8_balance = await kfive.balanceOf(account8, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_root_balance, kfive_root_balance.toString());
            eq(o.kfive_account1_balance, kfive_account1_balance.toString());
            eq(o.kfive_account2_balance, kfive_account2_balance.toString());
            eq(o.kfive_account5_balance, kfive_account5_balance.toString());
            eq(o.kfive_account6_balance, kfive_account6_balance.toString());
            eq(o.kfive_account8_balance, kfive_account8_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });
    });
});