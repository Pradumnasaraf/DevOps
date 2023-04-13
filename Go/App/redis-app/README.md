# Redis-App

## Tech Stack

The App is build with Fiber and Redis. I am using Docker to run the Redis Server

## App Workflow

The app contains two endpoints `/post` and `/get`.

First endpoint is used to create a user in Redis with the key and value both as the `username` which will sent visa POST request to the `/post` endpoint.

Second endpoint is used to check the user by sending a GET request to the `/get` endpoint with `username` as the parameter.