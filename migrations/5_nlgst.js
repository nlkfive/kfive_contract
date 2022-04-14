const NLGST = artifacts.require("NLGST");
const MarketplaceStorage = artifacts.require("MarketplaceStorage");

module.exports = async function (deployer) {
  marketplaceStorage = await MarketplaceStorage.deployed();
  deployer.deploy(NLGST, marketplaceStorage.address, "https://ipfs.io/ipfs/");
};
