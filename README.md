# ifttt-notify [![Build Status][1]][2] [![Go Report Card][3]][4] [![License: MIT][5]][6]
[1]: https://img.shields.io/travis/seankhliao/ifttt-notify.svg?style=flat-square
[2]: https://travis-ci.org/seankhliao/ifttt-notify
[3]: https://goreportcard.com/badge/github.com/seankhliao/ifttt-notify?style=flat-square
[4]: https://goreportcard.com/report/github.com/seankhliao/ifttt-notify
[5]: https://img.shields.io/badge/License-MIT-blue.svg?longCache=true&style=flat-square
[6]: LICENSE

simple, flexible command line tool to trigger ifttt maker events

it does something similar to
```sh
curl "https://maker.ifttt.com/trigger/$(env IFTTT_EVENT)/with/key/$(env IFTTT_KEY)" \
-H "Content-Type: application/json" \
-d '{"value1": "'"'$1"'", "value2": "'"$2"'", "value3", "'"$3"'"}'
```

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
