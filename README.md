[![Build Status](https://travis-ci.com/DragonSSS/cloud-audition.svg?token=thAs5cG2qtRSo4WkQp6Z&branch=master)](https://travis-ci.com/DragonSSS/cloud-audition)

# Cloud Audition

## Intro

Cloud audition is the basic REST web service that manages messages (CRUD operatons on message) and detects if message is palindrome.

## Architecture

The web service is basically separated into two main parts: handlers and messagemap.

The handlers are the standard Go web http request handlers which are responsible to create/get/update/delete message:
[handlers](https://github.com/DragonSSS/cloud-audition/blob/master/handlers/handlers.go) that includes the handlers for POST/GET/PUT/DELETE HTTP methods.

The data model handled by the handlers is a struct of Message:

```go
type Message struct {
	Msg string `json:"msg"`
}
```

The messagemap is the mutex map with methods that is used to store and manage internal messages data:
[messagemap](https://github.com/DragonSSS/cloud-audition/blob/master/messagemap/messagemap.go) that supports to create/update/get/delete message in mutex map.

The key in the map is an integer id, the value is a struct of InternalMessage:

```go
type InternalMessage struct {
	ID           int    `json:"id"`
	Msg          string `json:"msg"`
	IsPalindrome bool   `json:"isPalindrome"`
}
```

The key id generation is using an integer variable with mutex to get unique id for new message by plus one.

## APIs

The Go swaggo plugin is used to generate API doc with annotations in the code.

The yaml of swagger api file is at [swagger.yaml](https://github.com/DragonSSS/cloud-audition/blob/master/docs/swagger.yaml). Please review it with [online swagger editor](https://editor.swagger.io/)

## Infrastructure

[Makefile](https://github.com/DragonSSS/cloud-audition/blob/master/Makefile) takes care of:

* Run code linter with golangci-lint docker
* Download dependencies using Go modules
* Build binary into build/bin/
* Build docker image using [Dockerfile](https://github.com/DragonSSS/cloud-audition/blob/master/Dockerfile)
* Run unit test and generate test coverage report into coverage.out
* Generate swagger API doc
* Clean the compiled binary

The repo is integrated with Travis CI pipeline with [travis.yml](https://github.com/DragonSSS/cloud-audition/blob/master/.travis.yml), which supports stages:

* Lint (make lint)
* Unit test (make test)
* Build binary (make build)
* Deploy the local built image on k8s cluster, check health of web service
  * Install k8s cluster on the fly using k8s Kind (running k8s cluster into container)
  * Build cloud audition docker image locally
  * Load local built image into k8s cluster
  * Deploy cloud audition app
  * Check the log of running pod
  * Deploy a k8s service for cloud audition app
  * Forward local traffic into the k8s service by kubectl port-forward
  * Check health of app by curl

## Improvement

This project is the basic implementation to reach the requirement, so there are places can be improved later:

* Current unit test coverage could be improved by adding more http request failure cases
* Change messagemap into NoSQL DB that fits into key-value date structure
* Add integration test with k8s cluster in CI pipeline
