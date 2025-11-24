import { createPublicClient, http, parseAbi } from "viem";
import { mainnet } from "viem/chains";
import config from "./config.js";

// Chainlink aggregator ABI (minimal)
const CHAINLINK_ABI = parseAbi([
  "function latestRoundData() view returns (uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)"
]);

const client = createPublicClient({ 
  chain: mainnet, 
  transport: http(config.rpcUrl) 
});

/**
 * Get the latest price from a Chainlink price feed
 * @param {string} feedAddress - The address of the Chainlink price feed
 * @returns {Promise<number>} - The latest price as a number
 */
export async function chainlinkPrice(feedAddress) {
  try {
    const { answer } = await client.readContract({
      address: feedAddress,
      abi: CHAINLINK_ABI,
      functionName: "latestRoundData",
    });

    // Most Chainlink feeds use 8 decimals, but ETH/USD uses 8, DAI/USD uses 8, USDC/USD uses 8
    // Convert the answer (int256) to a number with proper decimal places
    return Number(answer) / 1e8;
  } catch (error) {
    console.error(`Error fetching Chainlink price from ${feedAddress}:`, error.message);
    throw error;
  }
}

/**
 * Get multiple Chainlink prices in parallel
 * @param {string[]} feedAddresses - Array of Chainlink feed addresses
 * @returns {Promise<number[]>} - Array of prices in the same order
 */
export async function chainlinkPrices(feedAddresses) {
  return Promise.all(feedAddresses.map(address => chainlinkPrice(address)));
}

/**
 * Get the latest round data from a Chainlink feed (including metadata)
 * @param {string} feedAddress - The address of the Chainlink price feed
 * @returns {Promise<object>} - Object with roundId, answer, startedAt, updatedAt, answeredInRound
 */
export async function chainlinkRoundData(feedAddress) {
  try {
    const result = await client.readContract({
      address: feedAddress,
      abi: CHAINLINK_ABI,
      functionName: "latestRoundData",
    });

    return {
      roundId: result[0],
      answer: Number(result[1]) / 1e8, // Convert to price
      startedAt: Number(result[2]),
      updatedAt: Number(result[3]),
      answeredInRound: result[4]
    };
  } catch (error) {
    console.error(`Error fetching Chainlink round data from ${feedAddress}:`, error.message);
    throw error;
  }
}