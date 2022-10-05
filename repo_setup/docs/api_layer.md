# API

## Architecutre

- **v1, v2**...: represent various versions of API's
  - **routers**: all the routes
  - **controllers**: operations that need to performed to serve the request go here for all the specific routes
- **models**: will hold the data models
- **middlewares**: will hold route-middlewares
- **repository**: will hold all the operations that are to be performed on database using the data-model
- **cache**: caching required or frequently used data
- **wrappers**: will hold response/request wrappers and converters
- **mappers**: `x.go` includes all mappers for `x_to_y`, `x_to_z`...

```mermaid
graph TD

    A(REST Server) --> B(gin)
    B --> C(routers)
    C --> |authentication layer| D(controllers)
    D --> E(wrappers)
    E --> D
    D --> F(repository)
    F --> I(cache)
    I --> F
    F --> G(models) --> H(database)
    G --> L(mappers)
    D --> J(file system)
    J --> D
    B --> K(static)
    K --> J
```
