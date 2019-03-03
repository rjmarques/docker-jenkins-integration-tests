# docker-jenkins-integration-tests

Me trying out Docker and Jenkins. The aim was to be able to run golang integration tests in a dynamic dev environment, where the tests would hit a DB running inside of docker. Additionally, I wanted a CI (jenkins) to be able to run this, so I added a jenkins file to this project

**This project depends on having docker and docker-compose available locally, as well as golang. Make sure both are available locally before proceeding**

## Start environment

`docker-compose up --build`

The new postgres db should be available at localhost:5432. It is initialized with `init.sql` available in the db folder

## Run unit tests

`go test ./...`

## Tear down environment

`docker-compose down --rmi all`
