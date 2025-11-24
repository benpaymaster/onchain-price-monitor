let latest = {
  spot: 0,
  oracle: 0,
  twap: 0,
  imbalancePct: 0,
  pegBps: 0
};

export function updateState(patch) {
  latest = { ...latest, ...patch };
}

export function getLatestState() {
  return latest;
}
