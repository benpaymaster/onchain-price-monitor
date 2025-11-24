import express from "express";
import { getLatestState } from "./state-cache.js";

const app = express();

app.get("/state", (req, res) => {
  res.json(getLatestState());
});

const PORT = 3000;
app.listen(PORT, () => {
  console.log(`UI API running at http://localhost:${PORT}/state`);
});
