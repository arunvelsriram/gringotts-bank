# Gringotts Bank

Simple Microservices app to demonstarte Distributed Tracing.

## Start Dependent Apps (Database, Open Telemetry Collector, Jaeger)

```shell
docker-compose up
```

## Seed Data

```shell
PGPASSWORD='postgres' psql -h localhost -p 25432 -U postgres -f data/customer/data.sql
PGPASSWORD='postgres' psql -h localhost -p 25432 -U postgres -f data/payment/data.sql
redis-cli -p 16379 < data/recommendation/data.redis
```

## Running

### Run Services Individually

```shell
make run/<service-name>

make run/frontend
make run/recommendation
make run/customer
make run/payment
```

### Run All Services

Requires [GNU Parallel](https://savannah.gnu.org/projects/parallel/) to start services parallely.

```shell
make run
```
