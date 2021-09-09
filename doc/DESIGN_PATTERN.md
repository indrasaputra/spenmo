## Design Pattern

This document lists down all design patterns and paradigms used to build the project.

### Clean Architecture

The code structure and layer follow [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). The innermost layer is located in [entity directory](../entity). The use cases layer is located in [service directory](../service). Any outer layers are located outside those two directories. 

### Clean Code

One of the property of clean code is concise and clear naming. This project is built with than mindset. Thus, all variables naming follow Clean Code and also [Effective Go](https://golang.org/doc/effective_go).

### S.O.L.I.D Principles

#### Single Responsibility

Each responsibility is only owned by a struct/class. It can be seen clearly in [service directory](../service).

#### Open-Closed

Not implemented due to small codebase and simple requirements.

#### Liskov Substitution

Not implemented due to small codebase and simple requirements.

#### Interface Segregation

Each interface is made granularly for a specific purpose. It can be seen in [service directory](../service).

#### Dependency Inversion

All structs that need any dependency use interface as their parameters. It can be seen in any constructor in this project.

### Builder

This pattern is used to wrap two or more services that work together. It can be seen in [builder directory](../internal/builder).

### Decorator

This pattern is used to add auxiliary capability to a certain use case. It can be seen in [interceptor directory](../internal/grpc/interceptor).

### Rate Limit

This pattern is used to prevent catastrophic call to the service. It can be seen in [interceptor directory](../internal/grpc/interceptor).
