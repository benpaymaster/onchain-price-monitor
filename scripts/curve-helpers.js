import { createPublicClient, http, parseAbi } from "viem";
import { mainnet } from "viem/chains";
import config from "./config.js";

// Minimal ABI for Curve 3pool (balances, coins, virtual price)
const CURVE_POOL_ABI = parseAbi([
  "function balances(uint256 i) view returns (uint256)",
  "function coins(uint256 i) view returns (address)",
  "function get_virtual_price() view returns (uint256)"
]);

const client = createPublicClient({ chain: mainnet, transport: http(config.rpcUrl) });

export async function readThreePoolState() {
  const pool = config.curve.threePool;
  const [daiBal, usdcBal, vprice, daiAddr, usdcAddr] = await Promise.all([
    client.readContract({ address: pool, abi: CURVE_POOL_ABI, functionName: "balances", args: [BigInt(config.curve.coins.DAI_INDEX)] }),
    client.readContract({ address: pool, abi: CURVE_POOL_ABI, functionName: "balances", args: [BigInt(config.curve.coins.USDC_INDEX)] }),
    client.readContract({ address: pool, abi: CURVE_POOL_ABI, functionName: "get_virtual_price" }),
    client.readContract({ address: pool, abi: CURVE_POOL_ABI, functionName: "coins", args: [BigInt(config.curve.coins.DAI_INDEX)] }),
    client.readContract({ address: pool, abi: CURVE_POOL_ABI, functionName: "coins", args: [BigInt(config.curve.coins.USDC_INDEX)] })
  ]);

  // Normalize to decimals (DAI 18, USDC 6)
  const dai = Number(daiBal) / 1e18;
  const usdc = Number(usdcBal) / 1e6;
  const total = dai + usdc;
  const imbalancePct = total > 0 ? ((dai - usdc) / total) * 100 : 0;

  return {
    virtualPrice: Number(vprice) / 1e18,
    daiBalance: dai,
    usdcBalance: usdc,
    imbalancePct,
    daiAddr,
    usdcAddr
  };
}

export function pegDeviationBps(daiUsd, usdcUsd) {
  // deviation of DAI from USDC in basis points
  const diff = (daiUsd - usdcUsd) / usdcUsd;
  return Math.round(diff * 1e4);
}