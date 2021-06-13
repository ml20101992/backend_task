# Backend Task
# General
According to the provided task I have implemented an Go based microservice that processes and saves the MP4 files sent over the Nats and an NodeJS app that loads the file and sends it to Nats and finally recieves the path on which the selected file is saved.

### Structure
The structure of this task is divided into three sections:
- app - contains the NodeJS app
- nats - contains docker compose file required to start up nats
- processing_service - Go microservice

Additionally, there is a `setup.sh` script that can be used to install NodeJS and Go dependencies.

## Running the Task
The following steps are required for running this task
- navigate to nats folder and run `docker-compose -d`
- once nats is up in docker, navigate to `processing_service` and run `go run main.go`
- lastly, navigate to `app` folder and run `nodejs index.js PATH_TO_FILE`

The processing service does not stop consuming until stopped, so there is no need to restart it in order to process a new file.