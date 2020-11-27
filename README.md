# Movie Service

## Overview

Movie API using [OMDB](http://www.omdbapi.com/)

## Setup

### Without docker

- Copy and paste `.env.sample` to `.env`. Set your database credential
- **Dont forget to set your OMDB Key**
- Import the `init/01.sql` in your database to create the table
- Run `go run main.go`

### With docker

- Run your docker desktop
- Run `make compose_up`. Docker will automatically run the app and the database.
- Run `make compose_down` to stop all.

## Test

- Run `make test` to run the unit test or `make testcoverage` to run the unit test and displaying the coverage percentage.
