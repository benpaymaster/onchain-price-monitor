# Onchain Price Monitor

![CI](https://img.shields.io/github/actions/workflow/status/benpaymaster/onchain-price-monitor/ci.yml?branch=main)
![Docker](https://img.shields.io/badge/docker-ready-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Professional Summary
Onchain Price Monitor is a production-grade infrastructure monitoring and pricing engine, built to showcase the skills and practices required for top-tier trading and infrastructure engineering roles. It demonstrates expertise in Python, distributed systems, cloud-native deployment, and collaborative Agile development.

## Why This Project?
This project was developed to align with the requirements of Jump Trading’s Production Engineering team. It highlights:
- Experience with Python, Docker, Kubernetes, Kafka, Redis, SQL
- Building robust, scalable, and testable systems
- Production readiness: logging, metrics, health checks, CI/CD
- Collaborative workflow and documentation

## Tech Stack
- **Python 3.12** (Flask, pytest, type hints)
- **Kafka** (bitnami/kafka)
- **Redis** (redis:7)
- **SQLite** (local, can be swapped for PostgreSQL)
- **Docker & docker-compose**
- **Kubernetes** (manifests in `deploy/k8s`)
- **GitHub Actions** (CI/CD)
- **Prometheus** (metrics)

## Features
- REST API for pricing and monitoring
- Kafka producer/consumer integration
- Redis cache for fast data access
- SQLite database for persistent storage
- Docker and Kubernetes orchestration
- CI/CD pipeline with GitHub Actions
- Structured logging and Prometheus metrics
- Health check endpoints
- Unit and integration tests
- Type hints and docstrings
- Agile workflow templates

## Demo
### Build and Run with Docker Compose
```bash
docker-compose up --build
```

#### Troubleshooting Docker Compose
- If you see a permissions error (e.g., `Permission denied`), run with `sudo`:
  ```bash
  sudo docker-compose up --build
  ```
- If you see `Could not find a required file. Name: index.js Searched in: /app/src`, make sure `frontend/react-dashboard/src/index.js` exists and is valid.
- For missing dependencies, run `npm install` in `frontend/react-dashboard` before building.
- For other errors, check that all services are correctly defined in `docker-compose.yaml` and that Docker is running.

### Run Tests
```bash
pytest backend/python/tests
```

### Example API Calls
```bash
curl 'http://localhost:5000/api/price?base_price=100&slippage=0.01'
curl 'http://localhost:5000/api/monitor'
curl 'http://localhost:5000/health'
curl 'http://localhost:5000/metrics'
```

## Architecture
- **Backend:** Python (Flask), Kafka, Redis, SQLite
- **Containers:** Docker, docker-compose
- **Orchestration:** Kubernetes manifests (deploy/k8s)
- **CI/CD:** GitHub Actions (.github/workflows/ci.yml)

## GitHub Workflow
To push your changes:
```bash
git add .
git commit -m "Fix docker-compose.yaml and update orchestration docs"
git push origin main
```

## Agile & Collaboration
- Issues, PR templates, and workflow samples included
- Code is documented and ready for team collaboration

## How to Contribute
1. Fork the repo and create a feature branch
2. Open an issue to discuss your idea or bug
3. Submit a pull request referencing the issue
4. Ensure all tests pass and code is documented
5. Participate in code reviews and discussions

## License
MIT

## Contact
For job applications, technical questions, or collaboration, please reach out via LinkedIn or email.

---

This project is designed to showcase production engineering skills for roles at top trading firms like Jump Trading.
