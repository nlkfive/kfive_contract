const NLGinseng = artifacts.require("NLGinseng")
const KFIVE = artifacts.require("KFIVE")
const Marketplace = artifacts.require("Marketplace")

const eq = assert.equal
const u = require('./utils.js')
var nlgst, kfive;
const { soliditySha3 } = require("web3-utils");

function getBlindedBid(value, fake, secret) {
    return soliditySha3(
        value, fake, secret
    );
}

contract("NLGinseng", (accounts) => {
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
    const issuedTokenAmount = 5000 * (10 ** tokenDecimals);
    const publicationFee1 = 0.1 * (10 ** tokenDecimals);

    before(async () => {
        kfive = await KFIVE.new({
            from: root
        });
        nlgst = await NLGinseng.new(BASE_URL, {
            from: root
        });
        marketplace = await Marketplace.new(kfive.address, OWNER_CUT_PER_MILLION);
    });

    describe('Deployment stage', async () => {
        it('Marketplace information', async () => {
            let o = {
                acceptedToken: kfive.address
            }

            const acceptedToken = marketplace.acceptedToken();
            eq(acceptedToken, o.acceptedToken());
        });

        it('Set publication fee (0.1 KFIVE)', async () => {
            let i = {
                fee: publicationFee1
            }
            await marketplace.setPublicationFee(i.fee);

            let o = {
                fee: publicationFee1.toString()
            }
            fee = await marketplace.publicationFeeInWei();
            eq(fee.toString(), o.fee);
        })
    });

    beforeEach(async () => {
        const issuedToken = web3.utils.toHex(issuedTokenAmount);
        await kfive.issue(account1, issuedToken, OFFCHAIN, { from: root });
        await kfive.issue(account2, issuedToken, OFFCHAIN, { from: root });
    });

    afterEach(async () => {
        const issuedToken = web3.utils.toHex(issuedTokenAmount);
        await kfive.redeem(account1, issuedToken, OFFCHAIN, { from: root });
        await kfive.redeem(account2, issuedToken, OFFCHAIN, { from: root });
    });


    describe('Auction 1: Complete auction', async () => {
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
            await nlgst.approve(marketplace.address, nftID, {
                from: account1
            });
        });

        it('(Account1) Approve KFIVE to market (for paying publication fee - 10 KFIVE)', async () => {
            var allowance = web3.utils.toHex(10 * (10 ** tokenDecimals));
            await kfive.approve(marketplace.address, allowance, {
                from: account1
            });
        });

        it('(Account1) Create new auction', async () => {
            await marketplace.createAuction(nlgst.address, nftID, startPriceInWei, biddingTime, revealTime, { from: account1 });

            let o = {
                rootBalance: publicationFee1.toString(),
                account1Balance: (issuedTokenAmount - publicationFee1).toString(),
                seller: account1
            }

            const rootBalance = await kfive.balanceOf(root);
            eq(o.rootBalance, rootBalance.toString());

            const account1Balance = await kfive.balanceOf(account1);
            eq(o.account1Balance, account1Balance.toString());

            const lastestEvent = await marketplace.getPastEvents("AuctionCreated");
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
            await marketplace.bidAuction(nlgst.address, 1, auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
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
            await marketplace.bidAuction(nlgst.address, 1, auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
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
            await marketplace.bidAuction(nlgst.address, 1, auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
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
            await marketplace.bidAuction(nlgst.address, 1, auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
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
            await marketplace.bidAuction(nlgst.address, 1, auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
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
            await u.assertRevert(marketplace.bidAuction(nlgst.address, 1, auctionId, getBlindedBid(i.bidValue, i.fake, i.secret), i.depositValue, {
                from: account2
            }));
        });

        ////////////////////////////
        // Reveal
        ////////////////////////////
    });
});