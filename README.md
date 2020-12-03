# Movie Service

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=vinbyte_movies&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=vinbyte_movies) [![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=vinbyte_movies&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=vinbyte_movies) [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=vinbyte_movies&metric=security_rating)](https://sonarcloud.io/dashboard?id=vinbyte_movies)

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
