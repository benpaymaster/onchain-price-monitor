import { createPublicClient, http, parseAbi, getAddress } from "viem";
import { mainnet } from "viem/chains";
import config from "./config.js";

// Create RPC client
export const client = createPublicClient({
  chain: mainnet,
  transport: http(config.rpcUrl),
});

// Uniswap V3 pool ABI (minimal)
const V3_POOL_ABI = parseAbi([
  "function slot0() view returns (uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)",
  "function token0() view returns (address)",
  "function token1() view returns (address)"
]);

// Convert sqrtPriceX96 → price
function decodePrice(sqrtPriceX96) {
  // (sqrtPriceX96^2 / 2^192)
  return Number((BigInt(sqrtPriceX96) * BigInt(sqrtPriceX96)) >> 192n) / 1e6;
}

// Returns WETH/USD price from Uniswap V3 pool
export async function v3SpotPriceUSD(poolAddress, wethAddress) {
  const { sqrtPriceX96 } = await client.readContract({
    address: poolAddress,
    abi: V3_POOL_ABI,
    functionName: "slot0",
  });

  const token0 = await client.readContract({
    address: poolAddress,
    abi: V3_POOL_ABI,
    functionName: "token0",
  });

  const weth = getAddress(wethAddress);
  const token0IsWeth = getAddress(token0) === weth;

  const raw = decodePrice(sqrtPriceX96);

  return token0IsWeth ? (1 / raw) : raw;
}
