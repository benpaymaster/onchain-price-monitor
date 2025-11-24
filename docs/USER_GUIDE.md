# User Guide

## Setup
1. Clone the repository
2. Install Docker and Docker Compose
3. Run `docker-compose up --build`
4. Access services:
   - Backend API: http://localhost:5000
   - React Dashboard: http://localhost:3000
   - Grafana: http://localhost:3001 (admin/admin)

## Running Tests
- Run `pytest backend/python/tests`

## Troubleshooting
- Check container logs with `docker-compose logs`
- Ensure all services are running (`docker ps`)
- For database issues, check MongoDB and Redis containers

## Contributing
- Fork the repo, create a feature branch, submit a PR
- Use issue and PR templates for collaboration
