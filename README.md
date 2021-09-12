# Spenmo

[![Go Report Card](https://goreportcard.com/badge/github.com/indrasaputra/spenmo)](https://goreportcard.com/report/github.com/indrasaputra/spenmo)
[![Workflow](https://github.com/indrasaputra/spenmo/workflows/Test/badge.svg)](https://github.com/indrasaputra/spenmo/actions)
[![codecov](https://codecov.io/gh/indrasaputra/spenmo/branch/main/graph/badge.svg?token=TF36qAeLI0)](https://codecov.io/gh/indrasaputra/spenmo)
[![Maintainability](https://api.codeclimate.com/v1/badges/9db3e7da6dad95be8c35/maintainability)](https://codeclimate.com/github/indrasaputra/spenmo/maintainability)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=indrasaputra_spenmo&metric=alert_status)](https://sonarcloud.io/dashboard?id=indrasaputra_spenmo)
[![Go Reference](https://pkg.go.dev/badge/github.com/indrasaputra/spenmo.svg)](https://pkg.go.dev/github.com/indrasaputra/spenmo)

## Owner

[Indra Saputra](https://github.com/indrasaputra)

## Tasks

1. The first task can be seen in [tree.go](cmd/tree/tree.go) and [tree_test.go](cmd/tree/tree_test.go) for the test.

2. The second task can be seen in [sequence.go](cmd/sequence/sequence.go) and [sequence_test.go](cmd/sequence/sequence_test.go) for the test.

3. The third task can be seen in [schema](db/schema/SCHEMA.md)

4. The fourth task is in this entire repository.

## Assumptions

Read [Assumptions](doc/ASSUMPTIONS.md)

## API

### gRPC

The API documentation can be seen in proto files (`*.proto`) in directory [proto](proto/indrasaputra/spenmo/v1/spenmo.proto).

### RESTful JSON

The API is automatically generated in OpenAPIv2 format when generating gRPC codes.
The generated files are stored in directory [openapiv2](openapiv2) in JSON format (`*.json`).
To see the RESTful API documentation, do the following:
- Open the generated json file(s), such as [spenmo.swagger.json](openapiv2/proto/indrasaputra/spenmo/v1/spenmo.swagger.json)
- Copy the content
- Open [https://editor.swagger.io/](https://editor.swagger.io/)
- Paste the content in [https://editor.swagger.io/](https://editor.swagger.io/)

## How to Run

- For complete prerequisites, read [Prerequisites](doc/PREREQUISITES.md). Otherwise, for example, you prefer running the service using docker, then skip this step.
- Then, read [How to Run](doc/HOW_TO_RUN.md).

## Code Map and Design Pattern

- Read [Code Map](doc/CODE_MAP.md).
- Read [Design Pattern](doc/DESIGN_PATTERN.md).

## Testing

### Unit Test

```
$ make test.unit
```

### Integration Test

[godog](https://github.com/cucumber/godog/#install) is mandatory to perform integration test.

To run the integration test, make sure you already run the application successfully. Follow [How to Run](doc/HOW_TO_RUN.md) for the guideline, including performing database migration to setup and seed the database.

The integration test doesn't focus on testing the rate limit functionality. It only focuses on whether the API can perform the job well.
Therefore, please adjust the rate limit configuration to accept many requests.
`RATE_LIMIT_PER_SECOND=100` and `RATE_BURST_PER_SECOND=100` should suffice.

When application is running, then run command to execute integration test.

```
$ make test.integration
```

You can also set the server URL, in case your default server is not localhost.

```
$ SERVER_URL=http://spenmo:8081/v1/users/cards make test.integration
```

## Observability

The application already emits necessary telemetry. If application's dependencies are run using [docker compose](doc/HOW_TO_RUN.md#docker), then monitoring is [provided by default](docker-compose.yaml). Otherwise, you have to provide them.

The observability is implemented at the handler/controller level. So, every request will be observed. For example, there will be four golden signals (throughput, latency, error rate, and saturation), traces, and logs for each endpoint. If needed, we can implement custom monitoring on block of code we need to observe.

These are stacks used as monitoring system.

| Observability    | Stack                                      | Address                                           |
| ---              | ---                                        | ---                                               |
| Metrics          | [Prometheus](https://prometheus.io/)       | [http://localhost:9090](http://localhost:9090)    |
| Visualization    | [Grafana](https://grafana.com/)            | [http://localhost:3000](http://localhost:3000)    |
| Tracing          | [Jaeger](https://www.jaegertracing.io/)    | [http://localhost:16686](http://localhost:16686)  |
| Log              | [Zap](https://github.com/uber-go/zap)      | Stdout                                            |

Visit [OBSERVABILITY](doc/OBSERVABILITY.md) for observability result example.
