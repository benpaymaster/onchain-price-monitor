import dotenv from "dotenv"; dotenv.config();
import config from "./config.js";
import { v3SpotPriceUSD } from "./uniswap-helpers.js";
import { chainlinkPrice } from "./chainlink-helpers.js";
import { updateState } from "./state-cache.js";   // <-- NEW

// Rolling 5-minute buffer
const prices = [];
const WINDOW_MS = config.twapWindowSecs * 1000; // default 300 seconds (5m)

function pruneOld() {
  const cutoff = Date.now() - WINDOW_MS;
  while (prices.length && prices[0].t < cutoff) prices.shift();
}

function twap() {
  if (prices.length === 0) return null;
  const sum = prices.reduce((acc, p) => acc + p.price, 0);
  return sum / prices.length;
}

function bps(a, b) {
  return ((a - b) / b) * 10000;
}

async function main() {
  console.log("Starting Uniswap vs Chainlink (Spot + TWAP) monitor…");

  while (true) {
    try {
      const spot = await v3SpotPriceUSD(
        config.uniswap.v3Pools.WETH_USDC_500,
        config.tokens.WETH
      );

      const oracle = await chainlinkPrice(config.chainlink.ETH_USD);

      // Save spot price with timestamp
      prices.push({ t: Date.now(), price: spot });
      pruneOld();
      const twapPrice = twap();

      const devSpot = bps(spot, oracle);
      const devTwap = twapPrice ? bps(twapPrice, oracle) : 0;

      const ts = new Date().toISOString();
      const flagSpot = Math.abs(devSpot) > config.alertThresholdBps ? "⚠️" : "✅";
      const flagTwap = Math.abs(devTwap) > config.alertThresholdBps ? "⚠️" : "✅";

      console.log(
        `${ts} | Spot: $${spot.toFixed(2)} (${devSpot.toFixed(2)} bps ${flagSpot}) | ` +
        `TWAP(5m): ${twapPrice ? `$${twapPrice.toFixed(2)} (${devTwap.toFixed(2)} bps ${flagTwap})` : "warming…"} | ` +
        `Oracle: $${oracle.toFixed(2)}`
      );

      // ✅ UPDATE UI STATE HERE
      updateState({
        spot,
        oracle,
        twap: twapPrice ?? null
      });

    } catch (err) {
      console.error("Error:", err.message);
    }

    await new Promise(r => setTimeout(r, config.intervalMs));
  }
}

main();
