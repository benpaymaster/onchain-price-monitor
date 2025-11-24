import dotenv from "dotenv"; dotenv.config();
import config from "./config.js";
import { readThreePoolState, pegDeviationBps } from "./curve-helpers.js";

// Import chainlink price helpers if they exist
let chainlinkPrice;
try {
  const chainlinkModule = await import("./chainlink-helpers.js");
  chainlinkPrice = chainlinkModule.chainlinkPrice;
} catch (err) {
  console.warn("Chainlink helpers not found, using mock prices");
  chainlinkPrice = async () => 1.0; // Mock function
}

function bps(a, b) {
  return ((a - b) / b) * 10000;
}

async function main() {
  console.log("Starting Curve 3Pool monitor…");
  console.log(`Monitoring pool: ${config.curve.threePool}`);
  console.log(`Alert threshold: ${config.alertThresholdBps} bps\n`);

  while (true) {
    try {
      // Get Curve pool state
      const poolState = await readThreePoolState();
      
      // Get oracle prices for DAI and USDC
      let daiPrice, usdcPrice;
      try {
        [daiPrice, usdcPrice] = await Promise.all([
          chainlinkPrice(config.chainlink.DAI_USD),
          chainlinkPrice(config.chainlink.USDC_USD)
        ]);
      } catch (err) {
        console.warn("Using fallback prices due to oracle error:", err.message);
        daiPrice = 1.0;
        usdcPrice = 1.0;
      }

      // Calculate peg deviation
      const pegDev = pegDeviationBps(daiPrice, usdcPrice);
      
      // Calculate pool imbalance vs expected
      const expectedRatio = daiPrice / usdcPrice; // Should be ~1.0
      const actualRatio = poolState.daiBalance / poolState.usdcBalance;
      const ratioDeviation = bps(actualRatio, expectedRatio);

      const ts = new Date().toISOString();
      const pegFlag = Math.abs(pegDev) > config.alertThresholdBps ? "⚠️" : "✅";
      const imbalanceFlag = Math.abs(poolState.imbalancePct) > 5 ? "⚠️" : "✅"; // 5% threshold
      const ratioFlag = Math.abs(ratioDeviation) > config.alertThresholdBps ? "⚠️" : "✅";

      console.log(
        `${ts} | Virtual Price: ${poolState.virtualPrice.toFixed(6)} | ` +
        `DAI/USDC Peg: ${pegDev.toFixed(2)} bps ${pegFlag} | ` +
        `Pool Imbalance: ${poolState.imbalancePct.toFixed(2)}% ${imbalanceFlag} | ` +
        `Ratio Dev: ${ratioDeviation.toFixed(2)} bps ${ratioFlag}`
      );

      console.log(
        `  Balances: DAI ${poolState.daiBalance.toFixed(0)} | USDC ${poolState.usdcBalance.toFixed(0)} | ` +
        `Prices: DAI $${daiPrice.toFixed(4)} | USDC $${usdcPrice.toFixed(4)}`
      );

      // Alert on significant deviations
      if (Math.abs(pegDev) > config.alertThresholdBps || 
          Math.abs(poolState.imbalancePct) > 10 || 
          Math.abs(ratioDeviation) > config.alertThresholdBps) {
        console.log("🚨 ALERT: Significant deviation detected!");
      }

    } catch (err) {
      console.error("Error:", err.message);
    }

    await new Promise(r => setTimeout(r, config.intervalMs));
  }
}

main().catch(console.error);