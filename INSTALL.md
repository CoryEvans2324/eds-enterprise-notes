# How to install annd setup Enterprise Notes locally

## Install docker & docker-compose

Install instructions on docs.docker.com
- [docker](https://docs.docker.com/engine/install/)
- [docker-compose](https://docs.docker.com/compose/install/)

Note:

When installing docker on linux, your user will need to be added to the docker group.

Run `usermod -aG docker $USER`

## Build the image for the web container
From the projects root directory run:

`docker-compose build`

## Deploy the stack
From the projects root directory run:

`docker-compose up -d`

For simplicity, the configuration files are already predefined and included in this repository. The passwords/secrets are insecure and should not be used in production.

Since this project is just for educational purposes and only deployed locally, there shouldn't be any issue with this.

## For the database unit tests
A user in postgres needs to be setup with the username `testing` and the password `testing`. Make sure the 'Can login` flag is set on the new user.

A database needs to be created named `testing` owned by the postgres user `testing`.

These credentials can be changed in the `database/database_test.go` file in the function `createConfig()`.

## Setting up the database for the first time.
Run `go run cmd/databaseInit/databaseInit.go`

## Accessing the services
The web service (the go application) is listening on port 8000 for http requests.
This can be changed in the `docker-compose.yml` file by mapping a different port to port 80 in the container:

```docker
services:
  ...
  web:
    ...
	ports:
	  - "<port_of_choice>:80"
```

Postgresql is running on the standard port `5432`.

PgAdmin is running on port `5050`. The default credentials for signing in are defined in the `docker-compose.yml` file.

## To stop the containers
`docker-compose stop` stops everything.

`docker-compose down` stops everything and **deletes** all containers and their data.