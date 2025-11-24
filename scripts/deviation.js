import dotenv from "dotenv";
dotenv.config();

import * as UH from "./uniswap-helpers.js";
import * as CH from "./curve-helpers.js";
import config from "./config.js";

async function main() {
  try {
    console.log("Reading Curve 3Pool state and Chainlink feeds...");

    const state = await CH.readThreePoolState();
    const [daiUsd, usdcUsd] = await Promise.all([
      UH.chainlinkPriceUSD ? UH.chainlinkPriceUSD(config.chainlink.DAI_USD) : UH.default?.chainlinkPriceUSD(config.chainlink.DAI_USD),
      UH.chainlinkPriceUSD ? UH.chainlinkPriceUSD(config.chainlink.USDC_USD) : UH.default?.chainlinkPriceUSD(config.chainlink.USDC_USD)
    ]);

    const devBps = CH.pegDeviationBps(daiUsd, usdcUsd);

    const out = {
      ts: new Date().toISOString(),
      pool: config.curve.threePool,
      virtualPrice: state.virtualPrice,
      daiBalance: state.daiBalance,
      usdcBalance: state.usdcBalance,
      imbalancePct: state.imbalancePct,
      daiUsd,
      usdcUsd,
      daiVsUsdcBps: devBps
    };

    console.log(JSON.stringify(out, null, 2));
    process.exit(0);
  } catch (e) {
    console.error("Error reading deviation:", e.stack || e.message || e);
    process.exit(1);
  }
}

main();