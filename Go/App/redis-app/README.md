# Redis-App

## Tech Stack

The App is build with Fiber and Redis. I am using Docker to run the Redis Server

## App Workflow

The app contains two endpoints `/post` and `/get`.

First endpoint is used to create a user in Redis with the key and value both as the `username` which will sent visa POST request to the `/post` endpoint.

Second endpoint is used to check the user by sending a GET request to the `/get` endpoint with `username` as the parameter.

Note: The data added by the POST request get expired after 3600 seconds. So, the GET request will return error. If you remove the expiration, head over to `main.go` chnage the `time.Second` to `0` in the following line.

```go
    err = rdb.Set(ctx, user.UserName, user.UserName, 60*time.Second).Err()
```
