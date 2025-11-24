import dotenv from "dotenv"; dotenv.config();
import config from "./config.js";
import { v3SpotPriceUSD } from "./uniswap-helpers.js";

async function main() {
  console.log("Starting Uniswap V3 price monitor…");

  while (true) {
    try {
      const price = await v3SpotPriceUSD(
        config.uniswap.v3Pools.WETH_USDC_500,
        config.tokens.WETH
      );

      console.log(`[${new Date().toISOString()}] WETH/USD: $${price.toFixed(2)}`);
    } catch (err) {
      console.error("Error:", err.message);
    }

    await new Promise(r => setTimeout(r, config.intervalMs));
  }
}

main();
