var RaceList = artifacts.require('RaceList');

module.exports = async function (deployer) {
    deployer.deploy(RaceList);
};