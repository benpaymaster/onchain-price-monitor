async function fetchData() {
  try {
    const res = await fetch("http://localhost:3000/state");
    const data = await res.json();

    document.getElementById("spot").textContent = `$${data.spot.toFixed(2)}`;
    document.getElementById("oracle").textContent = `$${data.oracle.toFixed(2)}`;
    document.getElementById("twap").textContent = `$${data.twap.toFixed(2)}`;
    document.getElementById("imbalance").textContent = `${data.imbalancePct.toFixed(2)}%`;
    document.getElementById("peg").textContent = `${data.pegBps.toFixed(2)} bps`;
  } catch (e) {
    console.error("Update failed", e);
  }
}

setInterval(fetchData, 5000);
fetchData();
