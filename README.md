# Gringotts Bank

Simple Microservices app to demonstarte Distributed Tracing.

## Start Dependent Apps (Databases, Cache, Open Telemetry Collector, Grafana Tempo Stack)

```shell
docker-compose up
```

## Seed Data

```shell
PGPASSWORD='postgres' psql -h localhost -p 25432 -U postgres -f data/customer/data.sql
PGPASSWORD='postgres' psql -h localhost -p 25432 -U postgres -f data/payment/data.sql
redis-cli -p 16379 < data/recommendation/data.redis
```
