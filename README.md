# Beer-Kong-Server

This repository provides code that is a part of the bigger solution for the Beer-Pong players. Application is divided by three services (this repo is including point 2 and 3): 

1. [Frontend](https://github.com/parostatkiem/beer-kong) 
2. Backend Server
3. Prisma Server

Only Linux and macOS are the supported distributions for developing and testing the Servers if you prefer to develop on windows, keep in mind that this code was never run on it. 

## Prerequisites

1. [go](https://github.com/golang/go)
2. [Prisma](https://github.com/prisma/prisma)
3. [docker](https://www.docker.com/) (only for local installation)
4. [docker-compose](https://docs.docker.com/compose/install/) (only for local installation)

## Run

To run the Backend Server you should deploy the Prisma Server first on your cloud or locally

### Prisma

You can simply deploy the Prisma Scheme on the your own Prisma Server. 

To do that you need to do one installation procedures described below:

Lokally:

1. Run the Prisma Server locally:
   ```bash
   make prisma-local
   ```

2. Reach the Prisma Server on address described in the configuration file

>**NOTE:** You can customize the Prisma configuration file to your use case. It's located in pkg/prisma/prisma-local.yml

Heroku:

1. Learn about deploying the Prisma on the heroku from [this movie](https://www.youtube.com/watch?v=b2ofz3XxR14&feature=emb_title)

2. Deploy the Prisma:
   ```bash
   make prisma-deploy
   ```

3. Reach the Prisma Server on address described in the configuration file

>**NOTE:** You can customize the Prisma configuration file to your use case. It's located in pkg/prisma/prisma.yml

### Server

locally:

1. Clone repository:
   ```bash
   cd $GOPATH/src/github.com
   mkdir pPrecel && cd pPrecel
   git clone https://github.com/pPrecel/BeerKongServer.git
   cd BeerKongServer
   ```

2. Run Server:
   ```bash
   make go-run
   ```

3. Reach the Prisma Server on `http://localhost` and port described in running log

Heroku:

1. Read [documentation](https://devcenter.heroku.com/categories/go-support) and deploy your Backend Server on Heroku

### Configuration

You can configure the Server and the Prisma Server via system environments set before running/deploying

Server:

| Parameter | Description | Default value |
|-----------|-------------|---------------|
|PORT|Server Port|80|
|PRISMA_ENDPOINT|Prisma Server Endpoint|http://localhost:4466/beer-kong-Prisma/beer-kong|
|PRISMA_SECRET|Prisma Server Secret|PrismaSecret|

Prisma

| Parameter | Description | Default value |
|-----------|-------------|---------------|
|PRISMA_SECRET|Prisma Server Secret|PrismaSecret|

>**NOTE:** PRISMA_SECRET should be same for Backend Server and Prisma Server

## Generated Clients

### Graphql

Beer-Kong-Server use the [gqlgen](https://github.com/99designs/gqlgen) to generate the graphql handler with fully programmable resolvers

To regenerate handler:

1. Run code from github:
   ```bash
   make gqlgen-regenerate
   ```

### Prisma

Backend Server user the Prisma to generate the [Prisma Client](https://www.prisma.io/docs/prisma-client/setup/generating-the-client-GO-r3c3/) to the Prisma Server

1. Regenerate Client:
   ```bash
   make Prisma-generate
   ``` 
