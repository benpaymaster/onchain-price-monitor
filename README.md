# Onchain Price Monitor & Sentinel

## Professional Summary

Onchain Price Monitor is a production-grade infrastructure monitoring and pricing engine. Originally built as a distributed Python system, it has been expanded to include a **Blockchain Sentinel** layer. This demonstrates the full-stack expertise required for top-tier trading firms: combining high-level data analytics with low-latency, systems-level "circuit breakers" that protect on-chain assets.

## Why This Project?

This project aligns with the requirements of firms like Jump Trading, Jane Street, or Wintermute. It highlights:

-   **Low-Latency Monitoring:** Using **Go** for direct, high-performance interaction with EVM-compatible chains.
    
-   **On-Chain Risk Management:** Solidity-based circuit breakers designed to pause trading during market anomalies or Oracle failures.
    
-   **Production Engineering:** Experience with Python, Docker, Kubernetes, Kafka, and Redis.
    
-   **Systems Integration:** Generating type-safe Go bindings from Solidity ABI to ensure robust, testable interactions between off-chain and on-chain layers.
    

## Tech Stack

-   **Languages:** Python 3.12 (Flask, pytest), Go 1.24 (Geth/abigen), Solidity (Foundry)
    
-   **Messaging & Cache:** Kafka (bitnami/kafka), Redis (redis:7)
    
-   **Persistence:** SQLite (local/dev), PostgreSQL (production-ready)
    
-   **Infrastructure:** Docker & docker-compose, Kubernetes (deploy/k8s)
    
-   **Blockchain:** Foundry/Anvil (local node), OpenZeppelin Access Control
    
-   **Observability:** Prometheus metrics, Structured Logging
    

## Features

-   **Blockchain Sentinel:** A Go-based monitoring agent that tracks price health and triggers an emergency `pause` via `TradingGuard.sol`.
    
-   **Automated Bindings:** CI/CD ready generation of Go interfaces for smart contracts.
    
-   **REST API:** Production-grade endpoints for pricing and system health.
    
-   **Distributed Processing:** Kafka producer/consumer integration for async monitoring.
    
-   **CI/CD:** Fully automated testing and linting via GitHub Actions.
    

## Demo & Execution

### 1\. Smart Contract Layer (Foundry)

To deploy the security layer and circuit breaker:

```Bash

    # Start local node
    anvil
    
    # Deploy contracts
    forge script script/Deploy.s.sol --rpc-url http://127.0.0.1:8545 --broadcast
```

### 2\. Blockchain Sentinel (Go)

To run the automated monitor and trigger the circuit breaker:

```Bash

    cd backend
    go mod tidy
    go run ./cmd/sentinel
```

### 3\. Distributed Infrastructure (Docker)

```Bash

    docker-compose up --build
```

#### Troubleshooting

-   **Permission Errors:** Run with `sudo docker-compose up --build`.
    
-   **Missing Bindings:** If `contracts/TradingGuard.go` is missing, regenerate via: `forge inspect TradingGuard abi --json > clean_abi.json && abigen --abi clean_abi.json --pkg contracts --type TradingGuard --out contracts/TradingGuard.go`
    

## Architecture

-   **Analytical Layer:** Python (Flask) handles complex pricing logic and REST interface.
    
-   **Security Layer:** Go Sentinel monitors on-chain conditions with minimal overhead.
    
-   **Execution Layer:** Solidity contracts enforce the "Kill Switch" logic on-chain.
    
-   **CI/CD:** GitHub Actions (.github/workflows/ci.yml) manages code quality.
    

## Agile & Collaboration

-   Issues, PR templates, and workflow samples included.
    
-   Code is documented with type hints and docstrings for team scalability.
    

## License

MIT

* * *

**Contact:** Reach out via LinkedIn or GitHub for collaboration or technical inquiries.