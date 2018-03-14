# ifttt-notify [![GoDoc](https://godoc.org/github.com/seankhliao/ifttt-notify?status.svg)](https://godoc.org/github.com/seankhliao/ifttt-notify) [![Build Status](https://img.shields.io/travis/seankhliao/ifttt-notify.svg?style=flat-square)](https://travis-ci.org/seankhliao/ifttt-notify) [![Go Report Card](https://goreportcard.com/badge/github.com/seankhliao/ifttt-notify)](https://goreportcard.com/report/github.com/seankhliao/ifttt-notify)

# ifttt-notify
simple, flexible command line tool to trigger ifttt maker events

## Install
```sh
go get github.com/seankhliao/ifttt-notify
```

## Usage
```sh
export IFTTT_EVENT=event-name
export IFTTT_KEY=your-key
ifttt-notify arg1 arg2 arg3
```
or 
```sh
ifttt-notify -event event-name -key your-key arg1 arg2 arg3
```

## License
The MIT License (MIT) - see [`LICENSE`](LICENSE) for more details
