const NLGinseng = artifacts.require("NLGinseng");

module.exports = function (deployer) {
    deployer.deploy(NLGinseng, "https://ipfs.io/ipfs/");
};