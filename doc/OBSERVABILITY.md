## Observability

### Metrics

There are three main metrics we need to observe.

#### Throughput

![Spenmo throughput](https://user-images.githubusercontent.com/4661221/132625517-5329b784-6d27-479b-a6d4-eddec4449ac5.png)

#### Latency

![Spenmo latency](https://user-images.githubusercontent.com/4661221/132625726-939bdf8f-087a-42e8-b87d-cd0c1d37e8f9.png)

#### Error

The error can be in any form, such as error rate or error count.
The image below shows the count of each gRPC code, including the client error (invalid argument) and success response (OK).

![Spenmo error](https://user-images.githubusercontent.com/4661221/132625896-dc925a5a-7cdb-4ab7-babe-ec574ae6a5bc.png)

There are also default metrics from the package, such as the number of Goroutines and Garbage Collector latency.

![Spenmo goroutines](https://user-images.githubusercontent.com/4661221/132627297-6695e829-5723-4e06-a8d7-d7f694062219.png)

![Spenmo GC](https://user-images.githubusercontent.com/4661221/132627315-6bde9cdd-1b89-4d54-a8b2-9661a6f7b14a.png)


### Tracing

We also have tracing implemented in our service. Currently, it just traces in gRPC server/handler and doesn't propagate further.

![Spenmo Jager](https://user-images.githubusercontent.com/4661221/132626457-4669484e-5ba3-4a6d-9b7c-1ed5621d8fc0.png)

![Spenmo Jager](https://user-images.githubusercontent.com/4661221/132626628-51426815-2852-4878-b89e-b8c3e39c90cc.png)