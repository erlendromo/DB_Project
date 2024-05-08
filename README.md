# DB_Project

DB-project for IDATG2204 (Datamodellering og databasesystemer)

## Authors

- Erlend Rømo
- Arthur Borger Thorkildsen
- Martin Morisbak
- Oskar Kjos

## Description

This application is a e-commerce web-application (designed as a RESTful JSON-API) for `Electromart`. The main point of this project is to design a database-schema, and normalize it to Boyce-Codd Normal Form (BCNF). Furthermore the integration of this database-schema into an application is required by the project-description. As frontend isn't part of the requirements, the group decided a RESTful JSON-API using go and postgresql would be appropriate.

## Endpoints

To see how each endpoint works, navigate to swagger-documentation on the application:
- If run locally: `http://localhost:8080/electromart/v1/swagger/`
- If run on vm: `http://<ip-to-vm>:8080/electromart/v1/swagger/`

## Prerequisites

Before deploying the application, make sure to have the following downloaded on the machine you are running.

- go -> https://go.dev/dl/
- postgresql -> https://www.postgresql.org/download/
- docker -> https://www.docker.com/products/docker-desktop/
- docker compose (linux) -> https://docs.docker.com/compose/install/linux/#install-the-plugin-manually

- FOR WINDOWS USERS
    - gnuwin32 (make) -> https://gnuwin32.sourceforge.net/packages/make.htm

- Optional
    - swag -> In terminal: `go install github.com/swaggo/swag/cmd/swag@latest`
        - If the swag binary isn't loaded properly in the terminal, the `make run` and `make swag` commands won't work. To fix possible issues with this, add the absolute path to where the swag-binary is located on your machine inside the `make runabsolute` and `make swagabsolute` commands in the Makefile.

## Deployment with Makefile

- `Configure .env file`
    - Before deploying, the example.env file will guide you to how a .env file used in this project can be setup. Configure this to satisfy your machines needs (all fields in example.env are required, but values can be amended).

- `Navigate to root-directory of project`
    - In a terminal of your choice, navigate to the root-directory of the project using commandline-tools (e.g. 'cd').

- `Make sure docker is running (docker desktop)`

- `Makefile deploy commands`
    - `make composeup` -> Runs a docker-container in the background, setting up postgres and the application.
    - `make composedown` -> Stops the docker-container.
    - `make applogs` -> Attaches the application-logs to the terminal.
    - `make dblogs` -> Attaches the database-logs to the terminal.

- `Open REST-application of choice (Postman, Thunderclient etc.)`
    - Local-url: `http://localhost:8080/electromart/v1/`
    - Vm-url: `http://<ip-to-vm>:8080/electromart/v1/`

- `Optional`
    - Open a browser on the URL selected above (Chrome, Edge, Firefox etc.)

## Other

- `Makefile development commands`
    - `make run` -> Configures swag documentation, and starts the container for the application (This is all you need to use).
    - `make stop` -> Stops the container, and removes database-data (N.B. Use this if you want to refresh database with only dummy-data, if not -> do `make composedown`).
    - `make swag` -> Append the swagger notation to the application.
    - `make deletedbdata` -> Removes directory containing database-data. (This will be auto-generated when running `make composeup`).
    - `make runabsolute` -> Same as `make run` but the path to the swag-binary file needs to be specified (in case swag is not set in $PATH).
    - `make swagabsolute` -> Same as `make swag`but the path to the swag-binary file needs to be specified (in case swag is not set in $PATH).