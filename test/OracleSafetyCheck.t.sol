// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../contracts/OracleSafetyCheck.sol";

contract OracleSafetyCheckTest is Test {
    OracleSafetyCheck public safetyCheck;
    
    // Mainnet Contract Addresses
    address constant CHAINLINK_ETH_USD = 0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419;
    address constant UNISWAP_V3_ETH_USDC_POOL = 0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640;
    
    uint32 constant INTERVAL = 600; // 10 minutes

    function setUp() public {
        // Forking Mainnet to get real oracle and pool data
        vm.createSelectFork("https://eth-mainnet.g.alchemy.com/v2/YOUR_ALCHEMY_RPC_URL");
        
        safetyCheck = new OracleSafetyCheck(
            CHAINLINK_ETH_USD, 
            UNISWAP_V3_ETH_USDC_POOL, 
            INTERVAL
        );
    }

    /// @notice Test that the deviation check passes under normal market conditions
    function test_CheckDeviation_Success() public view {
        (bool isSafe, uint256 cPrice, uint256 uPrice) = safetyCheck.isPriceSafe();
        
        console.log("Chainlink Price: ", cPrice);
        console.log("Uniswap TWAP:    ", uPrice);
        
        assertTrue(isSafe, "Deviation should be < 2% in normal markets");
    }

    /// @notice Test that the check fails if we manipulate the data (Simulating an exploit)
    function test_CheckDeviation_Fail_On_Manipulation() public {
        // We use vm.mockCall to simulate a massive price crash in Chainlink
        // This simulates what would happen during a flash loan or oracle exploit
        vm.mockCall(
            CHAINLINK_ETH_USD,
            abi.encodeWithSelector(AggregatorV3Interface.latestRoundData.selector),
            abi.encode(uint80(1), int256(500 * 1e8), uint256(0), uint256(0), uint80(1)) // ETH at $500
        );

        (bool isSafe,,) = safetyCheck.isPriceSafe();
        assertFalse(isSafe, "Should flag unsafe if Chainlink crashes to $500 while Uniswap is normal");
    }
}