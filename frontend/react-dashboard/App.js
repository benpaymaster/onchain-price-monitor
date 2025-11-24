import React, { useEffect, useState } from 'react';
import axios from 'axios';

function App() {
  const [price, setPrice] = useState(null);
  const [monitor, setMonitor] = useState(null);
  const [health, setHealth] = useState(null);

  useEffect(() => {
    axios.get('/api/price?base_price=100&slippage=0.01').then(res => setPrice(res.data));
    axios.get('/api/monitor').then(res => setMonitor(res.data));
    axios.get('/health').then(res => setHealth(res.data));
  }, []);

  return (
    <div style={{ fontFamily: 'Arial', padding: 20 }}>
      <h1>Onchain Price Monitor Dashboard</h1>
      <section>
        <h2>Price</h2>
        {price ? <pre>{JSON.stringify(price, null, 2)}</pre> : 'Loading...'}
      </section>
      <section>
        <h2>Monitor Data</h2>
        {monitor ? <pre>{JSON.stringify(monitor, null, 2)}</pre> : 'Loading...'}
      </section>
      <section>
        <h2>Health</h2>
        {health ? <pre>{JSON.stringify(health, null, 2)}</pre> : 'Loading...'}
      </section>
    </div>
  );
}

export default App;
