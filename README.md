# Hexagonal Architecture Template
### Golang 1.20

## Description
This repository was created in order to provide a template for developers to quickly bootstrap a new service based on the principles of:
- Hexagonal Architecture
- Clean Code
- SOLID
- Elements from Clean Architecture

## Expanding the architecture
There are guidelines when developing using this pattern that should be followed as described bellow:

- App domain **dictates** the way ports communicate (in simple words, **never import infra dependencies (models) into domain**)
- Do not pollute layers with unnecessary info. Each layer should have its own model which is being transformed on the Port side to the relevant model.

### Expanding the Structure
- **cmd/app/main**: Where the bootstraping takes place
- **deploy/build**: holds dockerization files (compose & docker files)
- **deploy/charts**: holds the helm charts for deployment
- **internal/config**: holds the configuration of the service
- **internal/core/models**: models representing the core domain (DTOs)
- **internal/core/ports**: holds the ports(interfaces) that used to communicate with various components
- **internal/core/service**: holds the services that interact and orchestrate business logic
- **internal/infra/grpc**: holds gRPC component with its services / handlers
- **internal/infra/http**: holds HTTP component with its services / handlers
- **internal/infra/repository**: holds any component related to storage (sql, s3 etc) Implements repository pattern.
- **migrations**: holds the database migrations
- **testing**: holds the testing initialization for databases and other components needed in integration testing
- **testing/sql**: holds the sql integration testing files
- **testing/sql/seeds**: seed files for the sql integration testing to populate the tables with initial data
- **scripts**: any sh scripts usually related to deployment


## Offerings

- [x] HTTP configuration with (Gin)[https://github.com/gin-gonic/gin]
- [x] Database configuration for postgres with (sqlx)[https://github.com/jmoiron/sqlx]
- [x] Dynamic ENV VARs configuration based on (Viper)[https://github.com/spf13/viper]
- [x] Helm templates for deployment


## Development Guidelines F.A.Q
What should i add in the `internal/core` folder?
- The `internal/core` folder is the heart of the application. It should hold the business logic of the application and should be the only component that is aware of the application domain. It should not be aware of any other component in the application.
- The `internal/core` folder should hold the following components:
    - **models**: The models that represent the application domain
    - **ports**: The ports that are being used to communicate with other components
    - **services**: The services that are responsible for orchestrating the business logic of the application
  
What is the difference between a service and a handler?
- A service is a component that is responsible for the business logic of the application. It is the component that is being called by the handler and is responsible for orchestrating the business logic.

What is the difference between a service and a repository?
- A service is a component that is responsible for the business logic of the application. It is the component that is being called by the repository and is responsible for orchestrating the business logic.

How do I add a new adapter?
In order to add a new adapter you need to follow the following steps:
1. Create a new folder under `internal/infra` with the name of the adapter