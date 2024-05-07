# DB_Project

DB-project for IDATG2204 (Datamodellering og databasesystemer)

## Authors

- Erlend Rømo
- Arthur Borger Thorkildsen
- Martin Morisbak
- Oskar Kjos

## Description

TODO write something here...

## Endpoints

To see how each endpoint works, navigate to swagger-documentation on the application:
- If run locally: `http://localhost:8080/electromart/v1/swagger`
- If run on vm: `http://<ip-to-vm>:8080/electromart/v1/swagger`

## Prerequisites

Before deploying the application, make sure to have the following downloaded on the machine you are running.

- go -> https://go.dev/dl/
- postgresql -> https://www.postgresql.org/download/
- docker -> https://www.docker.com/products/docker-desktop/
- docker compose (linux) -> https://docs.docker.com/compose/install/linux/#install-the-plugin-manually

- FOR WINDOWS USERS
    - gnuwin32 (make) -> https://gnuwin32.sourceforge.net/packages/make.htm

## Deployment with Makefile

- `Configure .env file`
    - Before deploying, the example.env file will guide you to how a .env file used in this project can be setup. Configure this to satisfy your machines needs (all fields in example.env are required, but values can be amended).

- `Navigate to root-directory of project`
    - In a terminal of your choice, navigate to the root-directory of the project using commandline-tools (e.g. 'cd').

- `Makefile commands`
    - `make composeup` -> Runs a docker-container in the background, setting up postgres and the application.
    - `make composedown` -> Stops the docker-container.
    - `make applogs` -> Attaches the application-logs to the terminal.
    - `make dblogs` -> Attaches the database-logs to the terminal.
    - `make deletedbdata` -> Removes directory containing database-data. (This will be auto-generated when running `make composeup`).

- `Open REST-application of choice (Postman, Thunderclient etc.)`
    - Local-url: `http://localhost:8080/electromart/v1/`
    - Vm-url: `http://<ip-to-vm>:8080/electromart/v1/`

- `Optional`
    - Open a browser on the URL's above (Chrome, Edge, Firefox etc.)