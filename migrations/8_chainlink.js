var ChainLink = artifacts.require('@chainlink/contracts/src/v0.8/Chainlink.sol');

module.exports = function (deployer) {
    deployer.deploy(ChainLink);
};