# Blockchain Bootstrap Tool

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