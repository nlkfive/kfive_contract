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

    describe('AUCTION 6: Two accounts bid the same bidValue but different deposit value', async () => {
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
                OwnerBalance: (0.1 * (10 ** tokenDecimals)).toString(),
                account1Balance: (issuedTokenAmount - 0.1 * (10 ** tokenDecimals)).toString(),
                seller: account1,
                nftID: 6,
                nftAddress: nlgst.address,
            }

            const OwnerBalance = await kfive.balanceOf(root);
            eq(o.OwnerBalance, OwnerBalance.toString());

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

        it('(Account2) Bid: 4 KFIVE; Deposit: 2 KFIVE; Fake: false', async () => {
            const i = {
                bidValue: 4 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("1234"),
                depositValue: 5 * (10 ** tokenDecimals),
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(nlgst.address, 6), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            });
        });

        it('(Account5) Bid: 4 KFIVE; Deposit: 3 KFIVE; Fake: false', async () => {
            const i = {
                bidValue: 4 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("1234"),
                depositValue: 6 * (10 ** tokenDecimals),
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(nlgst.address, 6), auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account5
            });
        });

        it('Wait until Bidding end', async () => {
            while (Math.floor(new Date().getTime() / 1000 < biddingEndedAt)) {}
        });

        ////////////////////////////
        // Reveal
        ////////////////////////////
        it('BEFORE REVEAL Auction 5: Check balance of account2, account5 and auctionMarketplace', async () => {
            const o = {
                kfive_account2_balance_before: (10 - 5) * (10 ** tokenDecimals),
                kfive_account5_balance_before: (10 - 6) * (10 ** tokenDecimals),
                kfive_market_balance_before_reveal: 11 * (10 ** tokenDecimals),
            }

            kfive_account2_balance_before = await kfive.balanceOf(account2, {from: root});
            kfive_account5_balance_before = await kfive.balanceOf(account5, {from: root});
            kfive_market_balance_before_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account2_balance_before, kfive_account2_balance_before.toString());
            eq(o.kfive_account5_balance_before, kfive_account5_balance_before.toString());
            eq(o.kfive_market_balance_before_reveal, kfive_market_balance_before_reveal.toString());
        });

        it('(Account2) reveal bid', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 6,
                auctionID: auctionId,
                value_account2: 4 * (10 ** tokenDecimals),
                fakeFalse: false,
                secret: web3.utils.fromAscii("1234"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account2],
                [i.fakeFalse],
                [i.secret], {
                from: account2
            });

            const o = {
                kfive_account2_balance: 6 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 10 * (10 ** tokenDecimals),
            }

            kfive_account2_balance = await kfive.balanceOf(account2, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account2_balance, kfive_account2_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });

        it('(Account5) reveal bid. Refund account5', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 6,
                auctionID: auctionId,
                value_account5: 4 * (10 ** tokenDecimals),
                fakeFalse: false,
                secret: web3.utils.fromAscii("1234"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account5],
                [i.fakeFalse],
                [i.secret], {
                from: account5
            });

            const o = {
                kfive_account2_balance: 6 * (10 ** tokenDecimals),
                kfive_account5_balance: 10 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 4 * (10 ** tokenDecimals),
            }

            kfive_account2_balance = await kfive.balanceOf(account2, {from: root});
            kfive_account5_balance = await kfive.balanceOf(account5, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account2_balance, kfive_account2_balance.toString());
            eq(o.kfive_account5_balance, kfive_account5_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });

        it('Wait until Reveal end', async () => {
            while (Math.floor(new Date().getTime() / 1000 < revealEndedAt)) {}
        });

        it('AuctionEnd. Account1 call acutionEnd', async () => {
            await auctionMarketplace.auctionEnd(nlgst.address, 6, auctionId, { from: account1 });
        });

        it('Account2 should be the owner of NFT(6)', async () => {
            const o = {
                nft6_owner: account2,
            }

            const nft6_owner = await nlgst.ownerOf(6, { from: root })

            eq(o.nft6_owner, nft6_owner);
        });

        it('AFTER REVEAL Auction 6: Check balance of account2, account5nd auctionMarketplace' + '\n' +
        '                       Account2 (1st): Valid Bid with same amount account5. Refund because depositValue < account5' + '\n' +
        '                       Account5 (2nd): Valid Bid with same amount account2. Not Refund', async () => {

            amount_seller_receive = amountSellerReceiveAfterAuction(4, OWNER_CUT_PER_MILLION);
            amount_contractOwner_receive = amountContractOwnerReceiveAfterAuction(4, OWNER_CUT_PER_MILLION);

            const o = {
                kfive_Owner_balance: (0.104) * (10 ** tokenDecimals),
                kfive_account1_balance: (10 - 0.1 + amount_seller_receive) * (10 ** tokenDecimals),
                kfive_account2_balance: 6 * (10 ** tokenDecimals),
                kfive_account5_balance: 10 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 0,
            }

            kfive_Owner_balance = await kfive.balanceOf(root, {from: root});
            kfive_account1_balance = await kfive.balanceOf(account1, {from: root});
            kfive_account2_balance = await kfive.balanceOf(account2, {from: root});
            kfive_account5_balance = await kfive.balanceOf(account5, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_Owner_balance, kfive_Owner_balance.toString());
            eq(o.kfive_account1_balance, kfive_account1_balance.toString());
            eq(o.kfive_account2_balance, kfive_account2_balance.toString());
            eq(o.kfive_account5_balance, kfive_account5_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });
    });
});