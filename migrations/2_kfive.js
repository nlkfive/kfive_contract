var KFIVE = artifacts.require('KFIVE');

module.exports = function (deployer, network, accounts) {
    deployer.deploy(KFIVE);
};
