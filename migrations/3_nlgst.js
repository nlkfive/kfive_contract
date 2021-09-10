const NLGST = artifacts.require("NLGST");

module.exports = function (deployer) {
    deployer.deploy(NLGST, "https://ipfs.io/ipfs/");
};