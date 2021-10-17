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
    return (bidValue - bidValue * OWNER_CUT_PER_MILLION / 1000000) * (10 ** tokenDecimals);
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

    describe('AUCTION 1: Simple auction', async () => {
        ////////////////////////////
        // Information
        // ISSUE 10 KFIVE to Account1, Account2, Account5, Account8
        // APPROVE 10 KFIVE to Auction Marketplace
        // UPDATE Auction Marketplace Address in Storage
        // MINT NFT(1) to Account1
        // Account1 CREATE AUCTION, startPrice = 1 KFIVE
        // Account2 BID: 0.1 KFIVE, deposit: 0.1 KFIVE, fake: false -> INVALID BID
        // Account5 BID:   2 KFIVE, deposit:   2 KFIVE, fake: false -> VALID BID
        // Account8 BID:   3 KFIVE, deposit:   3 KFIVE, fake: false -> VALID BID (*)
        // Reveal:
        // + Refund to Account5
        // Withdraw:
        // + Refund to Account2
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
            const account1_balance = await kfive.balanceOf(account1, { from: root });
            const account2_balance = await kfive.balanceOf(account2, { from: root });
            const account5_balance = await kfive.balanceOf(account5, { from: root });
            const account6_balance = await kfive.balanceOf(account6, { from: root });
            const account8_balance = await kfive.balanceOf(account8, { from: root });
            const auctionMarketplace_balance = await kfive.balanceOf(auctionMarketplace.address, { from: root });

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

        it('(Account2) Bid: 0.1 KFIVE; Deposit: 0.1 KFIVE; Fake: false. INVALID BID. Lower than startPrice', async () => {
            const i = {
                bidValue: 0.1 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 0.1 * (10 ** tokenDecimals),
                nftID: 1,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                getBlindedBid(i.bidValue, i.fake, i.secret),
                i.depositValue, {
                from: account2
            });
        });

        it('(Account5) Bid: 2 KFIVE; Deposit: 1 KFIVE; Fake: false. VALID BID. Higher than startPrice', async () => {
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

        it('(Account8) Bid: 3 KFIVE; Deposit: 3 KFIVE; Fake: false. VALID BID. Higher Bid than Account2', async () => {
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
                value_account5: 2 * (10 ** tokenDecimals),
                value_account8: 3 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
            }

            await u.assertRevert(auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account2, i.value_account5, i.value_account8],
                [i.fake, i.fake, i.fake],
                [i.secret, i.secret, i.secret], {
                from: root
            }));
        });

        it('Wait until Bidding end', async () => {
            await delay(120 * 1000);
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
                kfive_account2_balance_before: (10 - 0.1) * (10 ** tokenDecimals),
                kfive_account5_balance_before: (10 - 1) * (10 ** tokenDecimals),
                kfive_account8_balance_before: (10 - 3) * (10 ** tokenDecimals),
                kfive_market_balance_before_reveal: (4.1) * (10 ** tokenDecimals),
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

        it('Reveal Stage Auction 1. The Successful BID' + '\n' +
            '                       (Account8) Bid: 3 KFIVE; Deposit: 3 KFIVE; Fake: false', async () => {

            const i = {
                nftAddress: nlgst.address,
                nftID: 1,
                value_account2: 0.1 * (10 ** tokenDecimals),
                value_account5: 2 * (10 ** tokenDecimals),
                value_account8: 3 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account2, i.value_account5, i.value_account8],
                [i.fake, i.fake, i.fake],
                [i.secret, i.secret, i.secret], {
                from: root
            });
        });

        it('AFTER REVEAL Auction 1: Check balance of root, account1, account2, account5, account8 and auctionMarketplace' + '\n' +
            '                       Account2: 9.9 KFIVE. Because of Invalid Bid -> No refund' + '\n' +
            '                       Account5:  10 KFIVE. Because Valid Bid but not the highest Bid -> Refund' + '\n' +
            '                       Account8:   7 KFIVE. Because of Valid Bid and Highest Bid -> No refund' + '\n' +
            '                       AuMarket: 0.1 KFIVE', async () => {

            amount_seller_receive = amountSellerReceiveAfterAuction(3, OWNER_CUT_PER_MILLION);
            amount_contractOwner_receive = amountContractOwnerReceiveAfterAuction(3, OWNER_CUT_PER_MILLION);

            const o = {
                kfive_root_balance: (publicationFee1 + amount_contractOwner_receive) * (10 ** tokenDecimals),
                kfive_account1_balance: (issuedTokenAmount - publicationFee1 + amount_seller_receive).toString(),
                kfive_account2_balance: (10 - 0.1) * (10 ** tokenDecimals),
                kfive_account5_balance: 10 * (10 ** tokenDecimals),
                kfive_account8_balance: 7 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 0.1 * (10 ** tokenDecimals),
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

        it('Account8 should be the owner of NFT(1)', async () => {
            const o = {
                nft1_owner: account8,
            }

            const nft1_owner = await nlgst.ownerOf(1, { from: root })

            eq(o.nft1_owner, nft1_owner);
        });

        ////////////////////////////
        // Withdraw
        ////////////////////////////
        it('Wait until Reveal end', async () => {
            await delay(120 * 1000);
        });

        it('Withdraw Stage Auction 1. Only account2 can withdraw because of invalid Bid' + '\n' +
            '               Balance of Account2 after Withdraw: 10 KFIVE' + '\n' +
            '               Balance of AuMarket after Withdraw:  3 KFIVE', async () => {

            await auctionMarketplace.withdraw(auctionId, { from: account2 });

            const o = {
                kfive_account2_balance: 10 * (10 ** tokenDecimals),
                kfive_auctionMarketplace_balance: 7 * (10 ** tokenDecimals),
            }

            kfive_account2_balance = await kfive.balanceOf(account2, {from: root});
            eq(o.kfive_account2_balance, kfive_account2_balance.toString());

            kfive_auctionMarketplace_balance = await kfive.balanceOf(auctionMarketplace.address, {from: root});
            eq(o.kfive_auctionMarketplace_balance, kfive_auctionMarketplace_balance.toString());
        });
    });

    describe('AUCTION 2: Pause Auction / fake = true / Invalid Bid', async () => {
        ////////////////////////////
        // Information
        // ISSUE 10 KFIVE to Account1, Account2, Account5, Account6, Account8
        // APPROVE 10 KFIVE to Auction Marketplace
        // UPDATE Auction Marketplace Address in Storage
        // MINT NFT(2) to Account1
        // Account1 CREATE AUCTION, startPrice = 1 KFIVE
        // Account2 BID: 0.1 KFIVE, deposit: 0.1 KFIVE, fake: false -> INVALID BID
        // Account2 BID:   1 KFIVE, deposit:   1 KFIVE, fake: false -> VALID BID
        // Account2 BID:   3 KFIVE, deposit:   1 KFIVE, fake: false -> VALID BID (*)
        // Account5 BID:   2 KFIVE, deposit:   2 KFIVE, fake: false -> VALID BID
        // Account5 BID:   5 KFIVE, deposit:   5 KFIVE, fake: true -> INVALID BID
        // Account6 BID:   1 KFIVE, deposit:   8 KFIVE, fake: false ->  VALID BID
        // Account8 BID:   3 KFIVE, deposit:   2 KFIVE, fake: true -> INVALID BID
        // Reveal:
        // + Refund to Account2 (1 KFIVE), Account5 (2 KFIVE), Account6 (8 KFIVE)
        // Withdraw:
        // + Refund to Account5 (5 KFIVE), Account8 (2 KFIVE)
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
            const account1_balance = await kfive.balanceOf(account1, { from: root });
            const account2_balance = await kfive.balanceOf(account2, { from: root });
            const account5_balance = await kfive.balanceOf(account5, { from: root });
            const account6_balance = await kfive.balanceOf(account6, { from: root });
            const account8_balance = await kfive.balanceOf(account8, { from: root });
            const auctionMarketplace_balance = await kfive.balanceOf(auctionMarketplace.address, { from: root });

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

            await auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: account1
            });

            let o = {
                rootBalance: (publicationFee1 + publicationFee1).toString(),
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
        it('(Account2) Bid: 0.1 KFIVE; Deposit: 0.1 KFIVE; Fake: false. INVALID BID', async () => {
            const i = {
                bidValue: 0.1 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 0.1 * (10 ** tokenDecimals),
                nftID: 2,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(i.nftAddress, i.nftID), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            });
        });

        it('(Account2) Bid: 1 KFIVE; Deposit: 1 KFIVE; Fake: false. VALID BID', async () => {
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

        it('(Account2) Bid: 3 KFIVE; Deposit: 1 KFIVE; Fake: false. VALID BID', async () => {
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

        it('(Account5) Bid: 5 KFIVE; Deposit: 5 KFIVE; Fake: true. INVALID BID', async () => {
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

        it('(Account5) Bid: 2 KFIVE; Deposit: 2 KFIVE; Fake: false. VALID BID', async () => {
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

        it('(Account6) Bid: 1 KFIVE; Deposit: 8 KFIVE; Fake: false. VALID BID', async () => {
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

        it('(Account8) Bid: 3 KFIVE; Deposit: 2 KFIVE; Fake: true. INVALID BID', async () => {
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
            await delay(120 * 1000);
        });

        ////////////////////////////
        // Reveal
        ////////////////////////////
        it('BEFORE REVEAL Auction 3: Check balance of account2, account5, account6, account8 and auctionMarketplace', async () => {
            const o = {
                kfive_account2_balance_before: (10 - 2.1) * (10 ** tokenDecimals),
                kfive_account5_balance_before: (10 - 7) * (10 ** tokenDecimals),
                kfive_account6_balance_before: (10 - 8) * (10 ** tokenDecimals),
                kfive_account8_balance_before: (10 - 2) * (10 ** tokenDecimals),
                kfive_market_balance_before_reveal: 19.1 * (10 ** tokenDecimals),
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

        it('Reveal Stage Auction 2. The successful BID:' + '\n' +
        '                           (Account2) Bid: 3 KFIVE; Deposit: 1 KFIVE; Fake: false', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 2,
                auctionID: auctionId,
                value_account2: 1 * (10 ** tokenDecimals),
                value_account2_2: 3 * (10 ** tokenDecimals),
                value_account2_3: 0.1 * (10 ** tokenDecimals),
                value_account5_1: 2 * (10 ** tokenDecimals),
                value_account5_2: 5 * (10 ** tokenDecimals),
                value_account6: 3 * (10 ** tokenDecimals),
                value_account8: 3 * (10 ** tokenDecimals),
                fakeFalse: false,
                fakeTrue: true,
                secret: web3.utils.fromAscii("secret"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account2, i.value_account2_2, i.value_account2_3, i.value_account5_1, i.value_account5_2, i.value_account6, i.value_account8],
                [i.fakeFalse, i.fakeFalse, i.fakeFalse, i.fakeFalse, i.fakeTrue, i.fakeFalse, i.fakeTrue],
                [i.secret, i.secret, i.secret, i.secret, i.secret, i.secret, i.secret], {
                from: root
            })
        });

        it('AFTER REVEAL Auction 2: Check balance of account2, account5, account6, account8 and auctionMarketplace' + '\n' +
            '                       Account2 (1st): Valid Bid but not highest -> Refund 1 KFIVE' + '\n' +
            '                       Account2 (2nd): Valid Bid and Highest Bid -> No refund and account2 transfer 2 KFIVE to Auction' + '\n' +
            '                       Account2 (3nd): Invalid -> No refund' + '\n' +
            '                       Account5 (1st): Valid Bid but not highest -> Refund 2 KFIVE' + '\n' +
            '                       Account5 (2nd): Invalid Bid -> No refund' + '\n' +
            '                       Account6: Valid Bid but not highest -> Refund 8 KFIVE' + '\n' +
            '                       Account8: Invalid Bid -> No refund', async () => {

            amount_seller_receive = amountSellerReceiveAfterAuction(3, OWNER_CUT_PER_MILLION);
            amount_contractOwner_receive = amountContractOwnerReceiveAfterAuction(3, OWNER_CUT_PER_MILLION);

            const o = {
                kfive_root_balance: (0.1 + 0.1  + amount_contractOwner_receive) * (10 ** tokenDecimals),
                kfive_account1_balance: (issuedTokenAmount - publicationFee1 + amount_seller_receive).toString(),
                kfive_account2_balance: 6.9 * (10 ** tokenDecimals),
                kfive_account5_balance: 5 * (10 ** tokenDecimals),
                kfive_account6_balance: 10 * (10 ** tokenDecimals),
                kfive_account8_balance: 8 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 7.1 * (10 ** tokenDecimals),
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

        it('Account2 should be the owner of NFT(2)', async () => {
            const o = {
                nft2_owner: account2,
            }

            const nft2_owner = await nlgst.ownerOf(2, { from: root })

            eq(o.nft2_owner, nft2_owner);
        });

        ////////////////////////////
        // Withdraw
        ////////////////////////////
        it('Wait until Bidding end', async () => {
            await delay(120 * 1000);
        });

        it('Withdraw Stage. Only account2, account5 and account8 can withdraw because of invalid bid' + '\n' +
        '                   Balance of Account2 after Withdraw: 7 KFIVE' + '\n ' +
        '                   Balance of Account5 after Withdraw: 10 KFIVE' + '\n ' +
        '                   Balance of Account8 after Withdraw: 10 KFIVE' + '\n ' +
        '                   Balance of AuMarket after Withdraw: 3 KFIVE', async () => {

            await auctionMarketplace.withdraw(auctionId, { from: account2 });
            await auctionMarketplace.withdraw(auctionId, { from: account5 });
            await auctionMarketplace.withdraw(auctionId, { from: account8 });

            const o = {
                kfive_account2_balance: 7 * (10 ** tokenDecimals),
                kfive_account5_balance: 10 * (10 ** tokenDecimals),
                kfive_account8_balance: 10 * (10 ** tokenDecimals),
                kfive_auctionMarketplace_balance: 0 * (10 ** tokenDecimals),
            }

            kfive_account2_balance = await kfive.balanceOf(account2, {from: root});
            eq(o.kfive_account2_balance, kfive_account2_balance.toString());

            kfive_account5_balance = await kfive.balanceOf(account5, {from: root});
            eq(o.kfive_account5_balance, kfive_account5_balance.toString());

            kfive_account8_balance = await kfive.balanceOf(account8, {from: root});
            eq(o.kfive_account8_balance, kfive_account8_balance.toString());

            kfive_auctionMarketplace_balance = await kfive.balanceOf(auctionMarketplace.address, {from: root});
            eq(o.kfive_auctionMarketplace_balance, kfive_auctionMarketplace_balance.toString());
        });

        it('Account2 try to withdraw again. Cannot because there is no invalid Bid', async () => {
            await u.assertRevert(auctionMarketplace.withdraw(auctionId, { from: account2 }));
        });
    });

    describe('AUCTION 3: cancelAuction / changePublicationFee', async () => {
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
            const account1_balance = await kfive.balanceOf(account1, { from: root });
            const account2_balance = await kfive.balanceOf(account2, { from: root });
            const account5_balance = await kfive.balanceOf(account5, { from: root });
            const account6_balance = await kfive.balanceOf(account6, { from: root });
            const account8_balance = await kfive.balanceOf(account8, { from: root });
            const auctionMarketplace_balance = await kfive.balanceOf(auctionMarketplace.address, { from: root });

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
                nftID: 3,
                nftURI: "3",
            }

            await nlgst.mint(account1, i.nftID, i.nftURI, {
                from: root
            });
        });

        it('(Account1) Approve NFT to market', async () => {
            const i = {
                nftID: 3,
            }

            await nlgst.approve(auctionMarketplace.address, i.nftID, {
                from: account1
            });
        });

        ////////////////////////////
        // Create Auction
        // Seller cancel auction
        // Accounts can not BID
        ////////////////////////////
        it('(Account1) Create new auction NFT(3) successfully', async () => {
            const i = {
                nftID: 3,
                nftURI: "3",
                startPriceInWei: 1 * (10 ** tokenDecimals),
                biddingTime: Math.floor(new Date().getTime() / 1000 + 120),
                revealTime: Math.floor(new Date().getTime() / 1000 + 240),
            }

            await auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: account1
            });

            let o = {
                rootBalance: (publicationFee1 + publicationFee1 + publicationFee1).toString(),
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
                secret: web3.utils.fromAscii("secret"),
                depositValue: 5 * (10 ** tokenDecimals),
            }

            await u.assertRevert(auctionMarketplace.bidAuction(nftAssetInfo(nlgst.address, 3), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            }));
        });

        ////////////////////////////
        // Admin Change Publication Fee to 0.5 KFIVE
        // Create new auction
        // Account 2 BID
        // Canceller cancel auction
        // Cannot reveal??
        // Return bidded KFIVE?
        ////////////////////////////
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
                biddingTime: Math.floor(new Date().getTime() / 1000 + 120),
                revealTime: Math.floor(new Date().getTime() / 1000 + 240),
            }

            await auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: account1
            });

            let o = {
                rootBalance: (publicationFee1 + publicationFee1 + publicationFee1 + 0.5 * (10 ** tokenDecimals)).toString(),
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

        it('(Account2) Bid: 5 KFIVE; Deposit: 5 KFIVE; Fake: false. VALID BID', async () => {
            const i = {
                bidValue: 5 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 5 * (10 ** tokenDecimals),
                nftID: 3,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(i.nftAddress, i.nftID), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            });

            const o = {
                kfive_account2_balance: (10 - 5) * (10 ** tokenDecimals),
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

        it('(Root) Cancel Auction of NFT(3) again. Cannot because it has been cancelled', async () => {
            await u.assertRevert(auctionMarketplace.cancelAuction(nftAssetInfo(nlgst.address, 3), auctionId, {
                from: root
            }));
        });

        it('Reveal Stage Auction 3. Cannot reveal a cancelled auction', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 3,
                value_account2: 5 * (10 ** tokenDecimals),
                fakeFalse: false,
                secret: web3.utils.fromAscii("secret"),
            }

            await u.assertRevert(auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                auctionId,
                [i.value_account2],
                [i.fakeFalse],
                [i.secret], {
                from: root
            }));
        });

        it('Auction Marketplace should return bidded KFIVE', async () => {
            const o = {
                kfive_account2_balance: 10 * (10 ** tokenDecimals),
                kfive_market_balance: 0 * (10 ** tokenDecimals),
            }

            kfive_account2_balance = await kfive.balanceOf(account2, {from: root});
            eq(o.kfive_account2_balance, kfive_account2_balance.toString());

            kfive_market_balance = await kfive.balanceOf(auctionMarketplace.address, {from: root});
            eq(o.kfive_market_balance, kfive_market_balance.toString());
        });
    });

    describe('AUCTION 4: transferOwnership / changeOwnerCutPerMillion', async () => {
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
            const newOwner_balance = await kfive.balanceOf(new_owner, { from: root });
            const account1_balance = await kfive.balanceOf(account1, { from: root });
            const account2_balance = await kfive.balanceOf(account2, { from: root });
            const account5_balance = await kfive.balanceOf(account5, { from: root });
            const account6_balance = await kfive.balanceOf(account6, { from: root });
            const account8_balance = await kfive.balanceOf(account8, { from: root });
            const auctionMarketplace_balance = await kfive.balanceOf(auctionMarketplace.address, { from: root });

            await kfive.redeem(new_owner, newOwner_balance, OFFCHAIN, { from: owner });
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

        it('(Account1) Set new setOwnerCutPerMillion. Cannot because only owner and admin can change value', async () => {
            const i = {
                new_ownerCutPerMillion: 100000,
            }
            await u.assertRevert(auctionMarketplace.setOwnerCutPerMillion(i.new_ownerCutPerMillion, {
                from: account1
            }));
        });

        it('(Root) Set new OwnerCutPerMillion', async () => {
            const i = {
                new_ownerCutPerMillion: 100000,
            }
            await auctionMarketplace.setOwnerCutPerMillion(i.new_ownerCutPerMillion, {
                from: root
            });

            const new_ownerCutPerMillion = await auctionMarketplace.ownerCutPerMillion();
            eq(new_ownerCutPerMillion.toString(), i.new_ownerCutPerMillion);
        });

        it('(Root) Transfer Ownership to new_owner', async () => {
            const i = {
                new_owner: new_owner,
            }
            await auctionMarketplace.transferOwnership(i.new_owner, {
                from: root
            });

            let newOwner = await auctionMarketplace.owner({
                from: root
            })
            eq(i.new_owner, newOwner);
        });

        it('(New_owner) Set a new OwnerCutPerMillion (2.000.000). Cannot because ownerCutPerMillion must smaller than 1.000.000', async () => {
            const i = {
                new_ownerCutPerMillion: 2000000,
            }
            await u.assertRevert(auctionMarketplace.setOwnerCutPerMillion(i.new_ownerCutPerMillion, {
                from: new_owner
            }));
        });

        it('(New_owner) Set new OwnerCutPerMillion to 200.000. Cannot because new_owner does not have ADMIN_ROLE', async () => {
            const i = {
                new_ownerCutPerMillion: 200000,
            }
            await u.assertRevert(auctionMarketplace.setOwnerCutPerMillion(i.new_ownerCutPerMillion, {
                from: new_owner
            }));
        });

        it('(New_owner) Grant ADMIN_ROLE to new_owner', async () => {
            await auctionMarketplace.grantRole("0x" + keccak256("ADMIN_ROLE"), new_owner, {
                from: new_owner
            });
        });

        it('(New_owner) Set new OwnerCutPerMillion to 200.000 successfully', async () => {
            const i = {
                new_ownerCutPerMillion: 200000,
            }
            await auctionMarketplace.setOwnerCutPerMillion(i.new_ownerCutPerMillion, {
                from: new_owner
            });
        });

        it('(Owner) Mint and approve new NFT to (Account1)', async () => {
            const i = {
                nftID: 4,
                nftURI: "4",
            }

            await nlgst.mint(account1, i.nftID, i.nftURI, {
                from: root
            });

            await nlgst.approve(auctionMarketplace.address, i.nftID, {
                from: account1
            });
        });

        it('(Account1) Create new auction NFT(4) successfully', async () => {
            const i = {
                nftID: 4,
                nftURI: "4",
                startPriceInWei: 1 * (10 ** tokenDecimals),
                biddingTime: Math.floor(new Date().getTime() / 1000 + 120),
                revealTime: Math.floor(new Date().getTime() / 1000 + 240),
            }

            await auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: account1
            });

            let o = {
                rootBalance: (publicationFee1 + publicationFee1 + publicationFee1 + 0.5 * (10 ** tokenDecimals)).toString(),
                newOwnerBalance: (0.5 * (10 ** tokenDecimals)).toString(),
                account1Balance: (issuedTokenAmount - 0.5 * (10 ** tokenDecimals)).toString(),
                seller: account1,
                nftID: 4,
                nftAddress: nlgst.address,
            }

            const rootBalance = await kfive.balanceOf(root);
            eq(o.rootBalance, rootBalance.toString());

            const newOwnerBalance = await kfive.balanceOf(new_owner);
            eq(o.newOwnerBalance, newOwnerBalance.toString());

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

        it('(Account1) Transfer NFT(4) to (Account2). Cannot because NFT(4) is on the Auction Marketplace', async () => {
            const i = {
                nftID: 4,
                from: account1,
                to: account2
            }

            await u.assertRevert(nlgst.transferFrom(i.from, i.to, i.nftID, {
                from: i.from
            }));

            let o = {
                owner: account1
            }

            let owner = await nlgst.ownerOf(i.nftID, {
                from: root
            });
            eq(o.owner, owner);
        });

        it('(Root) Root can not call owner functions anymore', async () => {
            const i = {
                nftID: 4,
                nftAddress: nlgst.address,
            }
            await u.assertRevert(auctionMarketplace.transferOwnership(account1, {
                from: root
            }));

            await u.assertRevert(auctionMarketplace.pause({
                from: root
            }));

            await u.assertRevert(auctionMarketplace.unpause({
                from: root
            }));
        });

        it('(New_owner) New_owner call pause, unpause and cancel auction', async () => {
            const i = {
                nftID: 4,
                nftAddress: nlgst.address,
            }

            await auctionMarketplace.pause({
                from: new_owner
            });

            await u.assertRevert(auctionMarketplace.cancelAuction(nftAssetInfo(nlgst.address, 4), auctionId, {
                from: new_owner
            }));

            await auctionMarketplace.unpause({
                from: new_owner
            });

            await auctionMarketplace.cancelAuction(nftAssetInfo(nlgst.address, 4), auctionId, {
                from: new_owner
            });
        });

        it('(Account1) Create new auction NFT(4) again', async () => {
            const i = {
                nftID: 4,
                nftURI: "4",
                startPriceInWei: 1 * (10 ** tokenDecimals),
                biddingTime: Math.floor(new Date().getTime() / 1000 + 120),
                revealTime: Math.floor(new Date().getTime() / 1000 + 240),
            }

            await auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: account1
            });

            let o = {
                newOwnerBalance: (1 * (10 ** tokenDecimals)).toString(),
                account1Balance: (issuedTokenAmount - 1 * (10 ** tokenDecimals)).toString(),
                seller: account1,
                nftID: 4,
                nftAddress: nlgst.address,
            }

            const newOwnerBalance = await kfive.balanceOf(new_owner);
            eq(o.newOwnerBalance, newOwnerBalance.toString());

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


        it('(Account2) Bid: 5 KFIVE; Deposit: 5 KFIVE; Fake: false. VALID BID', async () => {
            const i = {
                bidValue: 5 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 5 * (10 ** tokenDecimals),
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(nlgst.address, 4), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            });
        });

        it('(Account5) Bid: 4 KFIVE; Deposit: 4 KFIVE; Fake: false. VALID BID', async () => {
            const i = {
                bidValue: 4 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 4 * (10 ** tokenDecimals),
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(nlgst.address, 4), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account5
            });
        });

        it('Wait until Bidding end', async () => {
            await delay(120 * 1000);
        });

        ////////////////////////////
        // Reveal
        ////////////////////////////
        it('BEFORE REVEAL Auction 4: Check balance of account2, account5 and auctionMarketplace', async () => {
            const o = {
                kfive_account2_balance_before: (10 - 5) * (10 ** tokenDecimals),
                kfive_account5_balance_before: (10 - 4) * (10 ** tokenDecimals),
                kfive_market_balance_before_reveal: 9 * (10 ** tokenDecimals),
            }

            kfive_account2_balance_before = await kfive.balanceOf(account2, {from: root});
            kfive_account5_balance_before = await kfive.balanceOf(account5, {from: root});
            kfive_market_balance_before_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account2_balance_before, kfive_account2_balance_before.toString());
            eq(o.kfive_account5_balance_before, kfive_account5_balance_before.toString());
            eq(o.kfive_market_balance_before_reveal, kfive_market_balance_before_reveal.toString());
        });

        it('(Root) Reveal Stage Auction 4. Cannot because root is not owner', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 4,
                auctionID: auctionId,
                value_account2: 5 * (10 ** tokenDecimals),
                value_account5: 4 * (10 ** tokenDecimals),
                fakeFalse: false,
                secret: web3.utils.fromAscii("secret"),
            }

            await u.assertRevert(auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account2, i.value_account5],
                [i.fakeFalse, i.fakeFalse],
                [i.secret, i.secret], {
                from: root
            }));
        });

        it('Reveal Stage Auction 4. The successful BID:' + '\n' +
        '                           (Account2) Bid: 5 KFIVE; Deposit: 5 KFIVE; Fake: false', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 4,
                auctionID: auctionId,
                value_account2: 5 * (10 ** tokenDecimals),
                value_account5: 4 * (10 ** tokenDecimals),
                fakeFalse: false,
                secret: web3.utils.fromAscii("secret"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account2, i.value_account5],
                [i.fakeFalse, i.fakeFalse],
                [i.secret, i.secret], {
                from: new_owner
            });
        });

        it('AFTER REVEAL Auction 4: Check balance of account2, account5, account6, account8 and auctionMarketplace' + '\n' +
            '                       Account2 (1st): Valid Bid and Highest Bid' + '\n' +
            '                       Account5 (2nd): Valid Bid but not highest -> Refund', async () => {

            amount_seller_receive = amountSellerReceiveAfterAuction(5, 200000);
            amount_contractOwner_receive = amountContractOwnerReceiveAfterAuction(5, 200000);

            const o = {
                kfive_newOwner_balance: (1 + amount_contractOwner_receive) * (10 ** tokenDecimals),
                kfive_account1_balance: (issuedTokenAmount - 1 * (10 ** tokenDecimals) + amount_seller_receive).toString(),
                kfive_account2_balance: 5 * (10 ** tokenDecimals),
                kfive_account5_balance: 10 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 0,
            }

            kfive_newOwner_balance = await kfive.balanceOf(new_owner, {from: root});
            kfive_account1_balance = await kfive.balanceOf(account1, {from: root});
            kfive_account2_balance = await kfive.balanceOf(account2, {from: root});
            kfive_account5_balance = await kfive.balanceOf(account5, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_newOwner_balance, kfive_newOwner_balance.toString());
            eq(o.kfive_account1_balance, kfive_account1_balance.toString());
            eq(o.kfive_account2_balance, kfive_account2_balance.toString());
            eq(o.kfive_account5_balance, kfive_account5_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });

        it('Account2 should be the owner of NFT(4)', async () => {
            const o = {
                nft4_owner: account2,
            }

            const nft4_owner = await nlgst.ownerOf(4, { from: root })

            eq(o.nft4_owner, nft4_owner);
        });
    });

    describe('AUCTION 5: Two accounts bid the same bidValue and deposit', async () => {
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
            const newOwner_balance = await kfive.balanceOf(new_owner, { from: root });
            const account1_balance = await kfive.balanceOf(account1, { from: root });
            const account2_balance = await kfive.balanceOf(account2, { from: root });
            const account5_balance = await kfive.balanceOf(account5, { from: root });
            const account6_balance = await kfive.balanceOf(account6, { from: root });
            const account8_balance = await kfive.balanceOf(account8, { from: root });
            const auctionMarketplace_balance = await kfive.balanceOf(auctionMarketplace.address, { from: root });

            await kfive.redeem(new_owner, newOwner_balance, OFFCHAIN, { from: owner });
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

        it('(Owner) Mint and approve new NFT to (Account1)', async () => {
            const i = {
                nftID: 5,
                nftURI: "5",
            }

            await nlgst.mint(account1, i.nftID, i.nftURI, {
                from: root
            });

            await nlgst.approve(auctionMarketplace.address, i.nftID, {
                from: account1
            });
        });

        it('(Account1) Create new auction NFT(5) successfully', async () => {
            const i = {
                nftID: 5,
                nftURI: "5",
                startPriceInWei: 1 * (10 ** tokenDecimals),
                biddingTime: Math.floor(new Date().getTime() / 1000 + 120),
                revealTime: Math.floor(new Date().getTime() / 1000 + 240),
            }

            await auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: account1
            });

            let o = {
                newOwnerBalance: (0.5 * (10 ** tokenDecimals)).toString(),
                account1Balance: (issuedTokenAmount - 0.5 * (10 ** tokenDecimals)).toString(),
                seller: account1,
                nftID: 5,
                nftAddress: nlgst.address,
            }

            const newOwnerBalance = await kfive.balanceOf(new_owner);
            eq(o.newOwnerBalance, newOwnerBalance.toString());

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

        it('(Account2) Bid: 4 KFIVE; Deposit: 4 KFIVE; Fake: false. VALID BID', async () => {
            const i = {
                bidValue: 4 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 4 * (10 ** tokenDecimals),
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(nlgst.address, 5), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            });
        });

        it('(Account5) Bid: 4 KFIVE; Deposit: 4 KFIVE; Fake: false. VALID BID', async () => {
            const i = {
                bidValue: 4 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 4 * (10 ** tokenDecimals),
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(nlgst.address, 5), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account5
            });
        });

        it('Wait until Bidding end', async () => {
            await delay(120 * 1000);
        });

        ////////////////////////////
        // Reveal
        ////////////////////////////
        it('BEFORE REVEAL Auction 5: Check balance of account2, account5 and auctionMarketplace', async () => {
            const o = {
                kfive_account2_balance_before: (10 - 4) * (10 ** tokenDecimals),
                kfive_account5_balance_before: (10 - 4) * (10 ** tokenDecimals),
                kfive_market_balance_before_reveal: 8 * (10 ** tokenDecimals),
            }

            kfive_account2_balance_before = await kfive.balanceOf(account2, {from: root});
            kfive_account5_balance_before = await kfive.balanceOf(account5, {from: root});
            kfive_market_balance_before_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account2_balance_before, kfive_account2_balance_before.toString());
            eq(o.kfive_account5_balance_before, kfive_account5_balance_before.toString());
            eq(o.kfive_market_balance_before_reveal, kfive_market_balance_before_reveal.toString());
        });

        it('Reveal Stage Auction 4. The successful BID:' + '\n' +
        '                           (Account2) Bid: 4 KFIVE; Deposit: 4 KFIVE; Fake: false', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 5,
                auctionID: auctionId,
                value_account2: 4 * (10 ** tokenDecimals),
                value_account5: 4 * (10 ** tokenDecimals),
                fakeFalse: false,
                secret: web3.utils.fromAscii("secret"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account2, i.value_account5],
                [i.fakeFalse, i.fakeFalse],
                [i.secret, i.secret], {
                from: new_owner
            });
        });

        it('AFTER REVEAL Auction 4: Check balance of account2, account5nd auctionMarketplace' + '\n' +
            '                       Account2 (1st): Valid Bid with same amount account5. Earlier bid' + '\n' +
            '                       Account5 (2nd): Valid Bid with same amount account2. Refund', async () => {

            amount_seller_receive = amountSellerReceiveAfterAuction(4, 200000);
            amount_contractOwner_receive = amountContractOwnerReceiveAfterAuction(4, 200000);

            const o = {
                kfive_newOwner_balance: (0.5 + amount_contractOwner_receive) * (10 ** tokenDecimals),
                kfive_account1_balance: (issuedTokenAmount - 0.5 * (10 ** tokenDecimals) + amount_seller_receive).toString(),
                kfive_account2_balance: 6 * (10 ** tokenDecimals),
                kfive_account5_balance: 10 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 0,
            }

            kfive_newOwner_balance = await kfive.balanceOf(new_owner, {from: root});
            kfive_account1_balance = await kfive.balanceOf(account1, {from: root});
            kfive_account2_balance = await kfive.balanceOf(account2, {from: root});
            kfive_account5_balance = await kfive.balanceOf(account5, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_newOwner_balance, kfive_newOwner_balance.toString());
            eq(o.kfive_account1_balance, kfive_account1_balance.toString());
            eq(o.kfive_account2_balance, kfive_account2_balance.toString());
            eq(o.kfive_account5_balance, kfive_account5_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });

        it('Account2 should be the owner of NFT(5)', async () => {
            const o = {
                nft5_owner: account2,
            }

            const nft5_owner = await nlgst.ownerOf(5, { from: root })

            eq(o.nft5_owner, nft5_owner);
        });
    });

    describe('AUCTION 6: Forget to reveal 1 bid', async () => {
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
            const newOwner_balance = await kfive.balanceOf(new_owner, { from: root });
            const account1_balance = await kfive.balanceOf(account1, { from: root });
            const account2_balance = await kfive.balanceOf(account2, { from: root });
            const account5_balance = await kfive.balanceOf(account5, { from: root });
            const account6_balance = await kfive.balanceOf(account6, { from: root });
            const account8_balance = await kfive.balanceOf(account8, { from: root });
            const auctionMarketplace_balance = await kfive.balanceOf(auctionMarketplace.address, { from: root });

            await kfive.redeem(new_owner, newOwner_balance, OFFCHAIN, { from: owner });
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

        it('(Owner) Mint and approve new NFT to (Account1)', async () => {
            const i = {
                nftID: 6,
                nftURI: "6",
            }

            await nlgst.mint(account1, i.nftID, i.nftURI, {
                from: root
            });

            await nlgst.approve(auctionMarketplace.address, i.nftID, {
                from: account1
            });
        });

        it('(Account1) Create new auction NFT(6) successfully', async () => {
            const i = {
                nftID: 6,
                nftURI: "6",
                startPriceInWei: 1 * (10 ** tokenDecimals),
                biddingTime: Math.floor(new Date().getTime() / 1000 + 120),
                revealTime: Math.floor(new Date().getTime() / 1000 + 240),
            }

            await auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: account1
            });

            let o = {
                newOwnerBalance: (0.5 * (10 ** tokenDecimals)).toString(),
                account1Balance: (issuedTokenAmount - 0.5 * (10 ** tokenDecimals)).toString(),
                seller: account1,
                nftID: 6,
                nftAddress: nlgst.address,
            }

            const newOwnerBalance = await kfive.balanceOf(new_owner);
            eq(o.newOwnerBalance, newOwnerBalance.toString());

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

        it('(Account2) Bid: 4 KFIVE; Deposit: 4 KFIVE; Fake: false. VALID BID', async () => {
            const i = {
                bidValue: 4 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 4 * (10 ** tokenDecimals),
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(nlgst.address, 6), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            });
        });

        it('(Account5) Bid: 5 KFIVE; Deposit: 5 KFIVE; Fake: false. VALID BID', async () => {
            const i = {
                bidValue: 5 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("secret"),
                depositValue: 5 * (10 ** tokenDecimals),
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(nlgst.address, 6), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account5
            });
        });

        it('Wait until Bidding end', async () => {
            await delay(120 * 1000);
        });

        ////////////////////////////
        // Reveal
        ////////////////////////////
        it('BEFORE REVEAL Auction 5: Check balance of account2, account5 and auctionMarketplace', async () => {
            const o = {
                kfive_account2_balance_before: (10 - 4) * (10 ** tokenDecimals),
                kfive_account5_balance_before: (10 - 5) * (10 ** tokenDecimals),
                kfive_market_balance_before_reveal: 9 * (10 ** tokenDecimals),
            }

            kfive_account2_balance_before = await kfive.balanceOf(account2, {from: root});
            kfive_account5_balance_before = await kfive.balanceOf(account5, {from: root});
            kfive_market_balance_before_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account2_balance_before, kfive_account2_balance_before.toString());
            eq(o.kfive_account5_balance_before, kfive_account5_balance_before.toString());
            eq(o.kfive_market_balance_before_reveal, kfive_market_balance_before_reveal.toString());
        });

        it('Reveal Stage Auction 6. The successful BID:' + '\n' +
        '                           (Account2) Bid: 5 KFIVE; Deposit: 5 KFIVE; Fake: false', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 6,
                auctionID: auctionId,
                value_account2: 4 * (10 ** tokenDecimals),
                value_account5: 5 * (10 ** tokenDecimals),
                fakeFalse: false,
                secret: web3.utils.fromAscii("secret"),
            }

            await u.assertRevert(auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account2],
                [i.fakeFalse],
                [i.secret], {
                from: new_owner
            }));

            const owner_nft6 = await nlgst.ownerOf(6, { from: root })
            eq(account1, owner_nft6);
        });
    });
});