var RaceList = artifacts.require('RaceList');
var ChainLink = artifacts.require('@chainlink/contracts/src/v0.8/Chainlink.sol');

module.exports = async function (deployer) {
    chainlink = await ChainLink.deployed();
    const oracle = "0x77a5310E41F0B9FE35E95239Fa5624390fadFbBA";
    const jobId = "0x2063653033393033306163363934393763626239363135623962316437623336";
    const fee = "40000000000000000";
    deployer.deploy(RaceList, chainlink.address, oracle, jobId, fee);
};