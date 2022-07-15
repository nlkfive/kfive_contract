const NLGST = artifacts.require("NLGST")
const KFIVE = artifacts.require("KFIVE")
const AuctionMarketplace = artifacts.require("PublicAuctionMarketplace")
const STORAGE = artifacts.require("MarketplaceStorage")

const eq = assert.equal
const u = require('./utils.js')
var nlgst, kfive, auctionMarketplace, storage;
const { soliditySha3 } = require("web3-utils");
const keccak256 = require('js-sha3').keccak256;
const tokenDecimals = 10;
const secret = "0x0000000000000000000000000000000000000000000000000000000000000001";
var biddingEndedAt = null;
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

function publicAuctionInfo(nftAddress, assetId, startPriceInWei, biddingEnd, minIncrement, timeNow) {
    return soliditySha3(
        { type: 'address', value: nftAddress },
        { type: 'uint256', value: assetId },
        { type: 'uint256', value: startPriceInWei },
        { type: 'uint256', value: biddingEnd },
        { type: 'uint256', value: minIncrement },
        { type: 'uint256', value: timeNow }
    );
}

function amountSellerReceiveAfterAuction(bidValue, OWNER_CUT_PER_MILLION) {
    return (bidValue - bidValue * OWNER_CUT_PER_MILLION / 1000000);
}

function amountContractOwnerReceiveAfterAuction(bidValue, OWNER_CUT_PER_MILLION) {
    return (bidValue * OWNER_CUT_PER_MILLION / 1000000);
}

contract("Public Auction Marketplace", (accounts) => {
    const root = accounts[0]
    const account1 = accounts[1]
    const account2 = accounts[2]
    const beneficary = accounts[3]
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
        kfive = await KFIVE.new({from: root});
        storage = await STORAGE.new({from: root});
        nlgst = await NLGST.new(storage.address, BASE_URL, {from: root});
        auctionMarketplace = await AuctionMarketplace.new(kfive.address, storage.address, root, OWNER_CUT_PER_MILLION, {
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
        it('Public Auction Marketplace information', async () => {
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

    describe('AUCTION 3: Simple public auction marketplace', async () => {
        var auctionId;
        var auctionStartTime;
        var auctionBiddingEnd;
        var nftAsset;

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

        it('(Owner) Mint new NFT to (Account1)', async () => {
            const i = {
                nftID: 1,
                nftURI: "1",
            }

            await nlgst.mint(account1, i.nftID, i.nftURI, {
                from: root
            });
        });

        it('(Account1) Approve NFT to Auction Marketplace', async () => {
            const i = {
                nftID: 1,
            }

            await nlgst.approve(auctionMarketplace.address, i.nftID, {
                from: account1
            });
        });

        it('(Root) Update OrderMarketplace in the Storage', async () => {
            const i = {
                auctionMarketplaceAddress: auctionMarketplace.address
            }

            await storage.updatePublicAuctionMarketplace(i.auctionMarketplaceAddress, {
                from: root
            });
        });

        it('(Account1) Create new auction NFT(1) successfully', async () => {
            const i = {
                nftID: 1,
                startPriceInWei: 1 * (10 ** tokenDecimals),
                startTime: Math.floor(new Date().getTime() / 1000 + 15),
                biddingEnd: Math.floor(new Date().getTime() / 1000 + 100),
                minIncrement: 1 * (10 ** tokenDecimals),
            }

            auctionStartTime = i.startTime;
            biddingEndedAt = i.biddingEnd;
            auctionId = publicAuctionInfo(nlgst.address,
                                          i.nftID,
                                          i.startPriceInWei,
                                          i.startTime,
                                          i.biddingEnd,
                                          i.minIncrement,
                                          Math.floor(new Date().getTime()));
            nftAsset = nftAssetInfo(nlgst.address, i.nftID);

            await auctionMarketplace.createAuction( nlgst.address,
                                                    auctionId,
                                                    nftAsset,
                                                    i.nftID,
                                                    i.startPriceInWei,
                                                    i.startTime,
                                                    i.biddingEnd,
                                                    i.minIncrement, {
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

            const lastestEvent = await auctionMarketplace.getPastEvents("PublicAuctionCreatedSuccessful");

            assetOwner = lastestEvent[0].returnValues.assetOwner;
            nftAddress = lastestEvent[0].returnValues.nftAddress;
            publicAuctionId = parseInt(lastestEvent[0].returnValues.publicAuctionId);
            assetId = parseInt(lastestEvent[0].returnValues.assetId);
            startTime = parseInt(lastestEvent[0].returnValues.startTime);
            auctionBiddingEnd = parseInt(lastestEvent[0].returnValues.biddingEnd);
            startPriceInWei = parseInt(lastestEvent[0].returnValues.startPriceInWei);
            minIncrement = parseInt(lastestEvent[0].returnValues.minIncrement);

            eq(account1, assetOwner);
            eq(o.nftAddress, nftAddress);
            eq(auctionId, publicAuctionId);
            eq(o.nftID, assetId);
            eq(auctionStartTime, startTime);
            eq(biddingEndedAt, auctionBiddingEnd);
            eq(startPriceInWei, i.startPriceInWei);
            eq(minIncrement, i.minIncrement);
        });

        it('Wait until Auction Stage starts', async () => {
            while (Math.floor(new Date().getTime() / 1000 < auctionStartTime + 1)) {}
        });

        it('(Account2) Bid: 3 KFIVE', async () => {
            const i = {
                bidValue: 3 * (10 ** tokenDecimals)
            }

            await auctionMarketplace.bid(nftAsset, auctionId, i.bidValue, {
                from: account2
            });

            let account2_balance = await kfive.balanceOf(account2, { from: root });
            eq(7 * (10 ** tokenDecimals), account2_balance.toString());
        });

        it('(Account5) Bid: 4 KFIVE', async () => {
            const i = {
                bidValue: 4 * (10 ** tokenDecimals)
            }

            await auctionMarketplace.bid(nftAsset, auctionId, i.bidValue, {
                from: account5
            });

            let account5_balance = await kfive.balanceOf(account5, { from: root });
            eq(6 * (10 ** tokenDecimals), account5_balance.toString());

            let account2_balance = await kfive.balanceOf(account2, { from: root });
            eq(10 * (10 ** tokenDecimals), account2_balance.toString());
        });

        it('(Account5) Bid: 6 KFIVE', async () => {
            const i = {
                bidValue: 6 * (10 ** tokenDecimals)
            }

            await auctionMarketplace.bid(nftAsset, auctionId, i.bidValue, {
                from: account5
            });

            let account5_balance = await kfive.balanceOf(account5, { from: root });
            eq(4 * (10 ** tokenDecimals), account5_balance.toString());
        });

        it('(Account2) Bid: 9 KFIVE', async () => {
            const i = {
                bidValue: 9 * (10 ** tokenDecimals),
                amount: web3.utils.toHex(10 * (10 ** tokenDecimals)),
                to: auctionMarketplace.address,
            }

            await kfive.approve(i.to, i.amount, { from: account2 });

            await auctionMarketplace.bid(nftAsset, auctionId, i.bidValue, {
                from: account2
            });

            let account2_balance = await kfive.balanceOf(account2, { from: root });
            eq(1 * (10 ** tokenDecimals), account2_balance.toString());

            let account5_balance = await kfive.balanceOf(account5, { from: root });
            eq(10 * (10 ** tokenDecimals), account5_balance.toString());
        });

        it('Wait until Bidding Stage end', async () => {
            while (Math.floor(new Date().getTime() / 1000 < biddingEndedAt + 1)) {}
        });

        it('(Account2) Receive reward', async () => {
            await auctionMarketplace.receiveReward(nlgst.address, auctionId, nftAsset, 1, {
                from: account2
            });

            const lastestEvent = await auctionMarketplace.getPastEvents("GrantAuctionRewardSuccessful");
            console.log(lastestEvent);

            let account2_balance = await kfive.balanceOf(account2, { from: root });
            eq(1 * (10 ** tokenDecimals), account2_balance.toString());

            let new_owner = await nlgst.ownerOf(1, {from: root});
            eq(account2, new_owner);
        });
    });
});