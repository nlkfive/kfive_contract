const AuctionMarketplace = artifacts.require("AuctionMarketplace");
const KFIVE = artifacts.require("KFIVE");

module.exports = async function (deployer) {
    kfive = await KFIVE.deployed();
    await deployer.deploy(AuctionMarketplace, kfive.address, 1000);
};