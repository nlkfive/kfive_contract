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
                rootBalance: 0,
                newOwnerBalance: (0.1 * (10 ** tokenDecimals)).toString(),
                account1Balance: (issuedTokenAmount - 0.1 * (10 ** tokenDecimals)).toString(),
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
                biddingTime: Math.floor(new Date().getTime() / 1000 + 65),
                revealTime: Math.floor(new Date().getTime() / 1000 + 130),
            }

            biddingEndedAt = i.biddingTime;
            revealEndedAt = i.revealTime;

            await auctionMarketplace.createAuction(nlgst.address, i.nftID, i.startPriceInWei, i.biddingTime, i.revealTime, {
                from: account1
            });

            let o = {
                newOwnerBalance: (0.2 * (10 ** tokenDecimals)).toString(),
                account1Balance: (issuedTokenAmount - 0.2 * (10 ** tokenDecimals)).toString(),
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
                secret: web3.utils.fromAscii("1234"),
                depositValue: 5 * (10 ** tokenDecimals),
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(nlgst.address, 4), auctionId, getBlindedBid(i.bidValue, i.fake, secret), i.depositValue, {
                from: account2
            });
        });

        it('(Account5) Bid: 4 KFIVE; Deposit: 4 KFIVE; Fake: false. VALID BID', async () => {
            const i = {
                bidValue: 4 * (10 ** tokenDecimals),
                fake: false,
                secret: web3.utils.fromAscii("1234"),
                depositValue: 4 * (10 ** tokenDecimals),
            }

            await auctionMarketplace.bidAuction(nftAssetInfo(nlgst.address, 4), auctionId, getBlindedBid(i.bidValue, i.fake, secret), i.depositValue, {
                from: account5
            });
        });

        it('Wait until Bidding end', async () => {
            while (Math.floor(new Date().getTime() / 1000 < biddingEndedAt)) {}
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

        it('(Account5) Reveal bid', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 4,
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
                [secret], {
                from: account5
            });

            const o = {
                kfive_account5_balance: 6 * (10 ** tokenDecimals),
                kfive_market_balance_after_reveal: 9 * (10 ** tokenDecimals),
            }

            kfive_account5_balance = await kfive.balanceOf(account5, {from: root});
            kfive_market_balance_after_reveal = await kfive.balanceOf(auctionMarketplace.address, {from: root});

            eq(o.kfive_account5_balance, kfive_account5_balance.toString());
            eq(o.kfive_market_balance_after_reveal, kfive_market_balance_after_reveal.toString());
        });

        it('Wait until Reveal end', async () => {
            while (Math.floor(new Date().getTime() / 1000 < revealEndedAt)) {}
        });

        ////////////////////////////
        // Withdraw
        ////////////////////////////
        it('AuctionEnd. Account1 call acutionEnd', async () => {
            await auctionMarketplace.auctionEnd(nlgst.address, 4, auctionId, { from: account1 });
        });

        it('Account5 is the owner of NFT(4)', async () => {
            const o = {
                nft5_owner: account5,
            }

            const nft5_owner = await nlgst.ownerOf(4, { from: root })

            eq(o.nft5_owner, nft5_owner);
        });

        it('(Account2) Reveal bid after auctionEnd to receive bidded KFIVE', async () => {
            const i = {
                nftAddress: nlgst.address,
                nftID: 4,
                auctionID: auctionId,
                value_account2: 5 * (10 ** tokenDecimals),
                fakeFalse: false,
                secret: web3.utils.fromAscii("1234"),
            }

            await auctionMarketplace.revealBid(
                nftAssetInfo(i.nftAddress, i.nftID),
                i.auctionID,
                [i.value_account2],
                [i.fakeFalse],
                [secret], {
                from: account2
            });

            const account2_balance = await kfive.balanceOf(account2, { from: root })
            eq(account2_balance.toString(), 10 * (10 ** tokenDecimals));
        });

        it('AFTER REVEAL Auction 4: Check balance of account2, account5, account6, account8 and auctionMarketplace' + '\n' +
        '                       Account2 (1st): Valid Bid and Highest Bid' + '\n' +
        '                       Account5 (2nd): Valid Bid but not highest -> Refund', async () => {

        amount_seller_receive = amountSellerReceiveAfterAuction(4, 200000);
        amount_contractOwner_receive = amountContractOwnerReceiveAfterAuction(4, 200000);

        const o = {
            kfive_newOwner_balance: (0.2 + amount_contractOwner_receive) * (10 ** tokenDecimals),
            kfive_account1_balance: (10 - 0.2 + amount_seller_receive) * (10 ** tokenDecimals),
            kfive_account2_balance: 10 * (10 ** tokenDecimals),
            kfive_account5_balance: 6 * (10 ** tokenDecimals),
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
    });
});