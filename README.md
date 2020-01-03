# Beer-Kong-Server

This repository provides code that is a part of the bigger solution for the Beer-Pong players. Application is divided by three services (this repo is including point 2 and 3): 
1. [Frontend](https://github.com/parostatkiem/beer-kong) 
2. Backend 
3. Prisma Server

Only Linux and macOS are the supported distributions for developing and testing the server if you prefer to develop on windows, keep in mind that this code was never run on it. 

## Prerequisites

1. [go](https://github.com/golang/go)
2. [prisma](https://github.com/prisma/prisma) (Only if you want to deploy the Prisma Server)

## Run

1. Clone repository:
   ```bash
   cd $GOPATH/src/github.com
   mkdir pPrecel && cd pPrecel
   git clone https://github.com/pPrecel/BeerKongServer.git
   ```
2. Run server:
   ```bash
   make go-run
   ```

## Prisma

You can simply deploy the prisma scheme on the your own prisma server. 

To do that you need to:

1. Edit `/pkg/prisma/prisma.yaml` to customize the prisma pattern to your use case
2. Deploy prisma:
   ```bash
   make prisma-deploy
   ```

## Graphql

Beer-Kong-Server use the [gqlgen](https://github.com/99designs/gqlgen) to generate the graphql handler with fully programmable resolvers

To regenerate handler:

1. Run code from github:
   ```bash
   make gqlgen-regenerate
   ```
