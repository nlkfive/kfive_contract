const Marketplace = artifacts.require("Marketplace");
const KFIVE = artifacts.require("KFIVE");

module.exports = async function (deployer) {
    kfive = await KFIVE.deployed();
    await deployer.deploy(Marketplace, kfive.address, 1000);
};