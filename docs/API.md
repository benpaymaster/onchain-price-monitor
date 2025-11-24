# API Documentation

## Authentication
- `POST /login` — Get JWT token (body: `{ "username": "user", "password": "pass" }`)

## Endpoints
- `GET /api/price?base_price=100&slippage=0.01` — Get calculated price (JWT required)
- `GET /api/monitor` — Get monitoring data (JWT required)
- `GET /health` — Health check
- `GET /metrics` — Prometheus metrics

## Example Usage
```bash
# Get JWT token
curl -X POST http://localhost:5000/login -H "Content-Type: application/json" -d '{"username":"user","password":"pass"}'

# Use token for protected endpoints
curl http://localhost:5000/api/price?base_price=100&slippage=0.01 -H "Authorization: Bearer <TOKEN>"
```
