const OrderMarketplace = artifacts.require("OrderMarketplace");
const KFIVE = artifacts.require("KFIVE");
const MarketplaceStorage = artifacts.require("MarketplaceStorage");

module.exports = async function (deployer) {
  kfive = await KFIVE.deployed();
  marketplaceStorage = await MarketplaceStorage.deployed();
  await deployer.deploy(
    OrderMarketplace,
    kfive.address,
    marketplaceStorage.address,
    1000
  );
};
