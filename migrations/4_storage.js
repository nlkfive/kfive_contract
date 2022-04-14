const MarketplaceStorage = artifacts.require("MarketplaceStorage");

module.exports = async function (deployer) {
  await deployer.deploy(MarketplaceStorage);
};
