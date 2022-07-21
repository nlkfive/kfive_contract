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
    const admin = accounts[3]
    const seller = accounts[4]
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

    describe('AUCTION 0: Simple check functions: transfer/newOwnerCutPerMillion/Pause/Unpause/Blacklist', async () => {
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
                nftID: 2,
                nftURI: "2",
            }

            await nlgst.mint(account1, i.nftID, i.nftURI, {
                from: root
            });
        });

        it('(Account1) Create new auction. Cannot because NFT(1) has not approved to Auction Marketplace', async () => {
            const i = {
                nftID: 2,
                startPriceInWei: 1 * (10 ** tokenDecimals),
                startTime: Math.floor(new Date().getTime() / 1000 + 10),
                biddingEnd: Math.floor(new Date().getTime() / 1000 + 4000),
                minIncrement: 1 * (10 ** tokenDecimals),
            }

            await u.assertRevert(auctionMarketplace.createAuction(  nlgst.address,
                                                                    publicAuctionInfo(nlgst.address,
                                                                                        i.nftID,
                                                                                        i.startPriceInWei,
                                                                                        i.biddingEnd,
                                                                                        i.minIncrement,
                                                                                        Math.floor(new Date().getTime())),
                                                                    nftAssetInfo(nlgst.address,
                                                                                    i.nftID),
                                                                    i.nftID,
                                                                    i.startPriceInWei,
                                                                    i.startTime,
                                                                    i.biddingEnd,
                                                                    i.minIncrement, {
                from: account1
            }));
        });

        it('(Account1) Approve NFT to Auction Marketplace', async () => {
            const i = {
                nftID: 2,
            }

            await nlgst.approve(auctionMarketplace.address, i.nftID, {
                from: account1
            });
        });

        it('(Account1) Create new public auction. Cannot because have not updated Auction address in Storage', async () => {
            const i = {
                nftID: 2,
                startPriceInWei: 1 * (10 ** tokenDecimals),
                startTime: Math.floor(new Date().getTime() / 1000 + 10),
                biddingEnd: Math.floor(new Date().getTime() / 1000 + 4000),
                minIncrement: 1 * (10 ** tokenDecimals),
            }

            await u.assertRevert(auctionMarketplace.createAuction(  nlgst.address,
                                                                    publicAuctionInfo(nlgst.address,
                                                                                    i.nftID,
                                                                                    i.startPriceInWei,
                                                                                    i.biddingEnd,
                                                                                    i.minIncrement,
                                                                                    Math.floor(new Date().getTime())),
                                                                    nftAssetInfo(nlgst.address,
                                                                                i.nftID),
                                                                    i.nftID,
                                                                    i.startPriceInWei,
                                                                    i.startTime,
                                                                    i.biddingEnd,
                                                                    i.minIncrement, {
                from: account1
            }));
        });

        it('(Account1) Update AuctionMarketplace in the Storage. Cannot because only Admin can update', async () => {
            const i = {
                auctionMarketplaceAddress: auctionMarketplace.address
            }

            await u.assertRevert(storage.updatePublicAuctionMarketplace(i.auctionMarketplaceAddress, {
                from: account1
            }));
        });

        it('(Root) Update Auction Marketplace in the Storage', async () => {
            const i = {
                auctionMarketplaceAddress: auctionMarketplace.address
            }

            await storage.updatePublicAuctionMarketplace(i.auctionMarketplaceAddress, {
                from: root
            });
        });

        it('(Root) Create new auction. Cannot because only Owner of NFT can create auction', async () => {
            const i = {
                nftID: 2,
                startPriceInWei: 1 * (10 ** tokenDecimals),
                startTime: Math.floor(new Date().getTime() / 1000 + 10),
                biddingEnd: Math.floor(new Date().getTime() / 1000 + 4000),
                minIncrement: 1 * (10 ** tokenDecimals),
            }

            await u.assertRevert(auctionMarketplace.createAuction(  nlgst.address,
                                                                    publicAuctionInfo(nlgst.address,
                                                                                        i.nftID,
                                                                                        i.startPriceInWei,
                                                                                        i.biddingEnd,
                                                                                        i.minIncrement,
                                                                                        Math.floor(new Date().getTime())),
                                                                    nftAssetInfo(nlgst.address,
                                                                                    i.nftID),
                                                                    i.nftID,
                                                                    i.startPriceInWei,
                                                                    i.startTime,
                                                                    i.biddingEnd,
                                                                    i.minIncrement, {
                from: root
            }));
        });

        it('(Account1) Approve NFT1 to seller', async () => {
            await nlgst.approve(seller, 2, {
                from: account1
            });
        });

        it('(Seller) Create new auction NFT(1). Cannot, only owner', async () => {
            const i = {
                nftID: 2,
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

            await u.assertRevert(auctionMarketplace.createAuction( nlgst.address,
                                                    auctionId,
                                                    nftAsset,
                                                    i.nftID,
                                                    i.startPriceInWei,
                                                    i.startTime,
                                                    i.biddingEnd,
                                                    i.minIncrement, {
                from: seller
            }));
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

            const new_ownerCutPerMillion = await auctionMarketplace.ownerCutPerMillion();
            eq(new_ownerCutPerMillion.toString(), 1000);
        });

        it('(New_owner) Set new OwnerCutPerMillion to 200.000', async () => {
            const i = {
                new_ownerCutPerMillion: 200000,
            }

            await u.assertRevert(auctionMarketplace.setOwnerCutPerMillion(i.new_ownerCutPerMillion, {
                from: new_owner
            }));

            const new_ownerCutPerMillion = await auctionMarketplace.ownerCutPerMillion();
            eq(new_ownerCutPerMillion.toString(), 1000);
        });

        it('(New_owner) Grant ADMIN_ROLE to new_owner', async () => {
            await auctionMarketplace.grantRole("0x" + keccak256("ADMIN_ROLE"), new_owner, {
                from: new_owner
            });
        });

        it('(New_owner) Set new OwnerCutPerMillion to 200.000', async () => {
            const i = {
                new_ownerCutPerMillion: 200000,
            }

            await auctionMarketplace.setOwnerCutPerMillion(i.new_ownerCutPerMillion, {
                from: new_owner
            });

            const new_ownerCutPerMillion = await auctionMarketplace.ownerCutPerMillion();
            eq(new_ownerCutPerMillion.toString(), i.new_ownerCutPerMillion);
        });

        it('(Root) Pause Auction Marketplace. Cannot, not owner', async () => {
            await u.assertRevert(auctionMarketplace.pause({
                from: root
            }));

            const pauseStatus = await auctionMarketplace.paused();
            eq(pauseStatus, false);
        });

        it('(New_owner) Pause Auction Marketplace', async () => {
            await auctionMarketplace.pause({
                from: new_owner
            });

            const pauseStatus = await auctionMarketplace.paused();
            eq(pauseStatus, true);
        });

        it('(Root) UnPause Auction Marketplace. Cannot, not owner', async () => {
            await u.assertRevert(auctionMarketplace.unpause({
                from: root
            }));

            const pauseStatus = await auctionMarketplace.paused();
            eq(pauseStatus, true);
        });

        it('(New_owner) UnPause Auction Marketplace', async () => {
            await auctionMarketplace.unpause({
                from: new_owner
            });

            const pauseStatus = await auctionMarketplace.paused();
            eq(pauseStatus, false);
        });

        it('(Account1) Add account2 to blacklist. Cannot, only owner', async () => {
            const i = {
                evil: account2,
            }

            await u.assertRevert(auctionMarketplace.addBlackList(i.evil, {
                from: account1
            }));

            check_blacklisted = await auctionMarketplace.isBlackListed(account2, {
                from: root
            });
            eq(check_blacklisted, false);
        });

        it('(Root) Add account2 to blacklist. Cannot, only owner', async () => {
            const i = {
                evil: account2,
            }

            await u.assertRevert(auctionMarketplace.addBlackList(i.evil, {
                from: root
            }));

            check_blacklisted = await auctionMarketplace.isBlackListed(account2, {
                from: root
            });
            eq(check_blacklisted, false);
        });

        it('(Admin) Add account2 to blacklist. Cannot, only owner', async () => {
            const i = {
                evil: account2,
            }

            await u.assertRevert(auctionMarketplace.addBlackList(i.evil, {
                from: admin
            }));

            check_blacklisted = await auctionMarketplace.isBlackListed(account2, {
                from: root
            });
            eq(check_blacklisted, false);
        });

        it('(New_owner) Add account2 to blacklist', async () => {
            const i = {
                evil: account2,
            }

            await auctionMarketplace.addBlackList(i.evil, {
                from: new_owner
            });

            check_blacklisted = await auctionMarketplace.isBlackListed(account2, {
                from: root
            });
            eq(check_blacklisted, true);
        });

        it('(Account1) Remove blacklist. Cannot, only owner', async () => {
            const i = {
                evil: account2,
            }

            await u.assertRevert(auctionMarketplace.removeBlackList(i.evil, {
                from: account1
            }));

            check_blacklisted = await auctionMarketplace.isBlackListed(account2, {
                from: root
            });
            eq(check_blacklisted, true);
        });

        it('(Account2) Remove blacklist. Cannot, only owner', async () => {
            const i = {
                evil: account2,
            }

            await u.assertRevert(auctionMarketplace.removeBlackList(i.evil, {
                from: account2
            }));

            check_blacklisted = await auctionMarketplace.isBlackListed(account2, {
                from: root
            });
            eq(check_blacklisted, true);
        });

        it('(Root) Remove blacklist. Cannot, only owner', async () => {
            const i = {
                evil: account2,
            }

            await u.assertRevert(auctionMarketplace.removeBlackList(i.evil, {
                from: root
            }));

            check_blacklisted = await auctionMarketplace.isBlackListed(account2, {
                from: root
            });
            eq(check_blacklisted, true);
        });

        it('(New_owner) Remove blacklist', async () => {
            const i = {
                evil: account2,
            }

            await auctionMarketplace.removeBlackList(i.evil, {
                from: new_owner
            });

            check_blacklisted = await auctionMarketplace.isBlackListed(account2, {
                from: root
            });
            eq(check_blacklisted, false);
        });
    });
});