# ⚔️ Joute is a `JSON-RPC over HTTP -> HTTP` proxy (gateway?)

*Joute is currently WIP.*

*Though the mision of this utility is much more ambitious, for now it only acts like a proxy for installing utility JSON-RPC headers.*

## Problem to be solved
Sometimes you just need to provide JSON-RPC API over HTTP but for some reasons you are unable to do this. Joute aims to be an adapter which accepts JSON-RPC requests from clients and transfers them to a target API as ordinary JSON payloads.

The result will look like this:

![Problem](assets/Problem.drawio.svg)

## Features for now
Acts like a proxy and sets following HTTP headers according to request:
- `X-Rpc-Method`: `$.method`
- `X-Rpc-Id`: `$.id`

## Roadmap
- [x] Utility headers (MVP I guess...)
- [] By-Method target endpoint selection
- [] Gateway functionality (By-Method routing?)
- [] JSON-RPC request -> JSON request (extract params)
- [] JSON response -> JSON-RPC response (result and error insertion)

## Building

### Binary

To build yourself use Go 1.23 and GNU make (optional).

With make it is enough to:
```
$ git clone https://github.com/amricko0b/joute.git
$ cd joute
$ make
```

### Docker image

Joute also comes as docker image.

```
$ docker pull amricko0b/joute
```

As to WIP it is only `latest` tag available.