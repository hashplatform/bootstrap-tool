# Blockchain Bootstrap Tool

[![The MIT License](https://img.shields.io/badge/license-MIT-orange.svg?style=flat-square)](http://opensource.org/licenses/MIT)
[![Travis](https://img.shields.io/travis/jackkdev/bootstrap-tool.svg?style=flat-square)](https://travis-ci.org/jackkdev/bootstrap-tool)
[![Go Report Card](https://goreportcard.com/badge/github.com/jackkdev/bootstrap-tool?style=flat-square)](https://goreportcard.com/report/github.com/jackkdev/bootstrap-tool)

A Go written bootstrapping tool in order to speed up the process and deployment of blockchain bootstraps.

### Features
* Easily configurable through using the JSON configuration file.
* Written in Go so it can easily be deployed on multiple platforms.
* Minimal resources are used. Can easily run on a Raspberry Pi.

### Planned Features
* API to view/initiate a bootstrap.
* Web interface to monitor the bootstrapping progress.

### Installing
```
cd $GOPATH/src/github.com
git clone https://github.com/jackkdev/bootstrap-tool.git
cd bootstrap-tool
dep ensure
```

### Usage
* Fill out the configuration file located at the base of the repository (`config.json`)
```json
{
  "coin": "coin-name",
  "directory": "bootstrap-directory",
  "destination": "bootstrap-archive-destination"
}
```

* Start a bootstrap by running
```go
go run cmd/bootstrap-tool/main.go
```

### Docker Builds
```
docker run --rm -it -v "$PWD":/go/src/bootstrap-tool -w /go/src/bootstrap-tool golang:1.12.4 ./build.sh 
```

Once its finished, check your destination path for an archived bootstrap.