import fs from "fs";
import path from "path";
import dotenv from "dotenv";
dotenv.config();

const ADDR = JSON.parse(
  fs.readFileSync(path.join(process.cwd(), "scripts", "addresses.json"), "utf8")
).mainnet;

export default {
  rpcUrl: process.env.RPC_URL,
  mevWs: process.env.MEV_SHARE_WS,
  intervalMs: parseInt(process.env.INTERVAL_MS || "5000", 10),
  alertThresholdBps: parseInt(process.env.ALERT_THRESHOLD_BPS || "60", 10),
  twapWindowSecs: parseInt(process.env.TWAP_WINDOW_SECS || "300", 10),
  chainlink: ADDR.chainlink,
  tokens: ADDR.tokens,
  uniswap: ADDR.uniswap,
  routers: ADDR.routers.map(a => a.toLowerCase()),
  curve: ADDR.curve
};
