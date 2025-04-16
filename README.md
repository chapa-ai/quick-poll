![Go](https://img.shields.io/github/languages/top/chapa-ai/quick-poll)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/chapa-ai/quick-poll)
![GitHub Repository Size](https://img.shields.io/github/repo-size/chapa-ai/quick-poll)
![Github Open Issues](https://img.shields.io/github/issues/chapa-ai/quick-poll)
![GitHub last commit](https://img.shields.io/github/last-commit/chapa-ai/quick-poll)

# Quick Poll API

API server for creating polls, voting, and collecting poll statistics in real time.

## Solution Notes

- :book: Standard Go project layout
- :whale: Dockerized with `docker-compose`
- :cloud: PostgreSQL for storing polls and votes
- :satellite: Kafka for real-time vote processing

## Key Assumptions & Trade-offs

### Core Design Decisions
- **Real-time Processing**: Chose Kafka for vote streaming to enable horizontal scaling, trading initial complexity for future scalability
- **ACID Compliance**: PostgreSQL ensures data integrity but requires careful connection pooling for high write loads

### Implementation Trade-offs
- **Development Speed vs Safety**: Skipped authentication in MVP to accelerate development (security debt)
- **Simplicity vs Performance**: Used sync DB writes for votes instead of batch processing (lower throughput but simpler code)
- **Local Focus**: Optimized for Docker-based development over cloud-native patterns (easier onboarding)

## HOWTO
- :hammer_and_wrench: Clean up dependencies with `make tidy`
- :whale: Run app in Docker Compose with `make run` or `docker-compose up --build`
- :computer: API Endpoints:
    - `POST /polls` — Create a new poll
    - `POST /polls/{id}/vote?option=...` — Vote for an option in a poll
    - `GET /polls/{id}/results` — Get poll results

## Future Enhancements

### Core System Upgrades
- :whale: Unit tests for the API endpoints
- :shield: **Auth Framework** - OAuth2/JWT integration with rate limiting
- :zap: Caching Layer - Redis-backed results cache with TTL invalidation
- :traffic_light: Monitoring Stack - Prometheus/Grafana dashboard with custom metrics

## Example

1. **Create Poll (POST /polls)**
    - Example request:
      ```
      curl -X POST http://localhost:9999/polls \
        -H "Content-Type: application/json" \
        -d '{"question":"Who are you?","options":["yes","no"]}'
      ```

2. **Vote (POST /polls/{id}/vote?option=yes)**
    - Example request:
      ```
      curl -X POST "http://localhost:9999/polls/721fc570-0304-4728-b1cf-3f29cab53a10/vote?option=yes"
      ```

3. **Get Results (GET /polls/{id}/results)**
    - Example request:
      ```
      curl http://localhost:9999/polls/721fc570-0304-4728-b1cf-3f29cab53a10/results
      ```
---
