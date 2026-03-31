# Project Title

A tool for managing your recipes.

![Project Image/Badge Example](https://img.shields.io) <!-- Add relevant badges, e.g., using Shields.io. -->

## Description
Potential names:
* Comforting Food Hub
* Open Recipe Box
<!-- Go into more depth about what your project does, the problem it solves, and why a user should care. You can add a list of key features here as well. -->

## Getting Started

### Dependencies

<!-- List any prerequisites, libraries, operating system versions, or tools needed before installation or execution. -->

* Container management software (e.g. [podman](https://podman.io/docs/installation) or [docker](https://docs.docker.com/get-started/get-docker/))
* A postgres client, or a db migration tool of your choosing (e.g. [Go Migrate](https://github.com/golang-migrate/migrate))

### Local Development


<!-- Provide clear, step-by-step instructions on how to download and install your project locally. Include code blocks for commands. -->

Install a container management software

Clone this repository

Create a `.env` file in the repository with the following lines
```
DATABASE_URL="postgres://postgres:{password}@database:5432/postgres"
TEMPLATE_DIRECTORY="/usr/src/app/templates"
```

Create a `secrets/` folder in the repository, and add a file `db_password.txt` containing only the default password you'd like to set for your database

Run your container-compose command to build and run the container
For docker:
`docker-compose up --build`
For podman:
`podman-compose up --build`

If you have postgres intalled on your machine, you can connect to the database and run the migration script to setup the database
`psql -h localhost -U postgres -d postgres -f internal/db/migrations/20260131231321_initial_db_setup.up.sql`
if not, you can use an external migration tool like go migrate with the same .sql file as above

After all that, you should be able to connect to the server at http://localhost:8080

<!-- Old steps
Install the container solution of your choice docker podman and set up a postgres image

Create a new postgres database for this project

Install a migration tool (go-migrate)

run
```bash
migrate -path internal/db/migrations -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose up
```

edit the .evn file locally to include
```bash
DATABASE_URL={YOUR_CONTAINER_POSTGRES_URL_HERE}
TEMPLATE_DIRECTORY={ABSOLUTE_PATH_TO_TEMPLATE_DIRECTORY_IN_PROJECT}
```
-->
