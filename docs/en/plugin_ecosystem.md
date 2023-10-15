# Overview

This article introduces the plugin ecosystem of tRPC.

# Maturity classification

In terms of maturity, tRPC plugins are currently divided into the following 4 levels:

| Maturity   | Description                             |
| -------- | -------------------------------- |
| stable   | Already used on a large scale |
| on trial | It has been tested and has been used in certain businesses |
| tested   | It has been tested and has not been used in business yet. |
| archived | Archived, no longer maintained, and not recommended for use |

# Plugin ecosystem in each language

The plugin types of tRPC mainly include codec, filter, naming, config, metrics, tracing, logging, telemetry, etc.

## tRPC-Cpp plugin ecosystem 

| Type  |   Plugin   |  Maturity  |  Doc |
| -------- | -------- | :------: | :---------------------------------------------------------------------------------: |
| Codec  | HTTP |  stable  | [link](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/http_protocol_service.md) |
| Codec  | gRPC |  stable  | [link](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/grpc_protocol_service.md) |
| Serialiazation  | Protobuf |  stable  | [link](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/serialization.md) |
| Serialiazation  | JSON |  stable  | [link](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/serialization.md) |
| Serialiazation  | text |  stable  | [link](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/serialization.md) |
| Serialiazation  | binary |  stable  | [link](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/serialization.md) |
| Compressor  | Gzip |  stable  | [link](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/compression.md) |
| Compressor  | lz4 |  stable  | [link](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/compression.md) |
| Compressor  | Snappy |  stable  | [link](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/compression.md) |
| Compressor  | zlib |  stable  | [link](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/compression.md) |
| naming  | MeshPolaris |  tested  | [link](https://github.com/trpc-group/cpp-naming-polarismesh/blob/main/README.md) |
| config  | etcd |  tested  | [link](https://github.com/trpc-group/cpp-config-etcd/blob/main/README.md) |
| metrics  | Prometheus |  stable  | [link](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/prometheus_metrics.md) |
| tracing  | Jaeger |  stable  | [link](https://github.com/trpc-group/cpp-tracing-jaeger/blob/main/README_zh.md) |
| logging  | CLS |  stable  | [link](https://github.com/trpc-group/cpp-logging-cls/blob/main/README_zh.md) |
| telemetry  | OpenTelemetry |  stable  | [link](https://github.com/trpc-group/cpp-telemetry-opentelemetry/blob/main/README_zh.md) |

## tRPC-Go plugin ecosystem

| Type  |   Plugin   |  Maturity  |  Doc |
| -------- | -------- | :------: | :---------------------------------------------------------------------------------: |
| Codec  | HTTP |  stable  | [link](https://github.com/trpc-group/trpc-go/tree/main/http) |
| Codec  | gRPC |  stable  | [link](https://github.com/trpc-ecosystem/go-codec/tree/main/grpc) |
| naming  | MeshPolaris |  tested  | [link](https://github.com/trpc-ecosystem/go-naming-polarismesh) |
| config  | etcd |  tested  | [link](https://github.com/trpc-ecosystem/go-config-etcd) |
| metrics  | Prometheus |  stable  | [link](https://github.com/trpc-ecosystem/go-metrics-prometheus) |
| tracing  | Jaeger |  stable  | [link](https://github.com/trpc-ecosystem/go-opentracing-jaeger) |
| logging  | CLS |  stable  | [link](https://github.com/trpc-ecosystem/go-log-cls) |
| telemetry  | OpenTelemetry |  stable  | [link](https://github.com/trpc-ecosystem/go-opentelementry) |
| filter | debuglog | stable | [link](https://github.com/trpc-ecosystem/go-filter/tree/main/debuglog) |
| filter | degrade | stable | [link](https://github.com/trpc-ecosystem/go-filter/tree/main/degrade) |
| filter | filterextensions | stable | [link](https://github.com/trpc-ecosystem/go-filter/tree/main/filterextensions) |
| filter | hystrix | stable | [link](https://github.com/trpc-ecosystem/go-filter/tree/main/hystrix) |
| filter | jwt | stable | [link](https://github.com/trpc-ecosystem/go-filter/tree/main/jwt) |
| filter | masking | stable | [link](https://github.com/trpc-ecosystem/go-filter/tree/main/masking) |
| filter | mock（client mock）| stable | [link](https://github.com/trpc-ecosystem/go-filter/tree/main/mock) |
| filter | recover ）| stable | [link](https://github.com/trpc-ecosystem/go-filter/tree/main/recovery) |
| filter | referer | stable | [link](https://github.com/trpc-ecosystem/go-filter/tree/main/referer) |
| filter | slime | stable | [link](https://github.com/trpc-ecosystem/go-filter/tree/main/slime) |
| filter | transinfo-blocker | stable | [link](https://github.com/trpc-ecosystem/go-filter/tree/main/transinfo-blocker) |
| filter | tvar | stable | [link](https://github.com/trpc-ecosystem/go-filter/tree/main/tvar) |
| filter | validation | stable | [link](https://github.com/trpc-ecosystem/go-filter/tree/main/validation) |
