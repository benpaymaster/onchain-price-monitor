// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import "forge-std/Script.sol";
import "../contracts/TradingGuard.sol";

contract DeployTradingGuard is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        new TradingGuard(
            0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266, // Admin
            0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266  // Bot
        );

        vm.stopBroadcast();
    }
}