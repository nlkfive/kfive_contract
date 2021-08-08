const OrderMarketplace = artifacts.require("OrderMarketplace");
const KFIVE = artifacts.require("KFIVE");

module.exports = async function (deployer) {
    kfive = await KFIVE.deployed();
    await deployer.deploy(OrderMarketplace, kfive.address, 1000);
};