[![Go Report Card](https://goreportcard.com/badge/github.com/diegohce/riib)](https://goreportcard.com/report/github.com/diegohce/riib)
[![GPLv3 license](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://github.com/diegohce/riib/blob/master/LICENSE)

# riib (Request it in Background)
Agent for asynchronous request for web applications.

Riib is not finished yet, and this readme lacks a lot of detail (it's still a work in progress).

# Protocols
Response for every protocol

```
{
    "<ticket>": "<remote response>"
}
```

## HTTP Protocol
HTTP POST expecting json body:

```
{
    "method": "<POST | GET>",
    "remote: "http[s]://remote_host:port/some/method",
    "callback_url": "http[s]://remote_host:port/where/to/get/the/response",
    "data": "whatever you want to send to <remote>"
}
```

## Raw Protocol
HTTP POST expecting json body:
```
{
    "remote: "remote_host:port",
    "callback_url": "http[s]://remote_host:port/where/to/get/the/response",
    "data": "whatever you want to send to <remote>"
}
```

# Dependencies
```
go get github.com/julienschmidt/httprouter
```
