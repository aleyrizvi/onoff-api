# Anagram API
---
This API was written for a test interview.

## Running and testing

The port listening for the api is `:9000`

If you are using docker, simply run `docker compose up`.

### OR

- Install [air](https://github.com/cosmtrek/air).
- run `air -c .air.conf`

## Test
There is one unit test written to make this task interesting.
To perform the test, simply run: 

```shell
go test ./...
```


## TODO
In future, I will add github actions to run tests on pull requests and generating and pushing docker containers to AWS registry.