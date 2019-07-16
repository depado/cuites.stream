<h1 align="center">cuites.stream</h1>
<h2 align="center">
  <img src="front/public/fox.gif" alt="mascot" height="50px">
  
  [![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/uses-badges.svg)](https://forthebadge.com)

  ![Go Version](https://img.shields.io/badge/Go%20Version-latest-brightgreen.svg)
  [![Go Report Card](https://goreportcard.com/badge/github.com/Depado/cuites.stream)](https://goreportcard.com/report/github.com/Depado/cuites.stream)
  [![Build Status](https://drone.depa.do/api/badges/Depado/cuites.stream/status.svg)](https://drone.depa.do/Depado/cuites.stream)
  [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Depado/cuites.stream/blob/master/LICENSE)
  [![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://saythanks.io/to/Depado)

  Site for the Cuites project
</h2>

## Backend

### Install

```sh
$ git clone 
```

### Setup

For the backend to run, you'll need a SoundCloud client ID. You can find this
ID by opening the network tab on SoundCloud (while connected) and check for the
`client_id` parameter in any Ajax request. Once you have that, you'll need the
set of user ID associated. Create a `conf.yml` file next to the binary (or the 
`main.go` if you're using `go run`) containing the following information:

```yaml
client_id: yourclientID
user_ids: [17771323, 93734268, 20836701, 153939520, 39713634]
server:
  host: 127.0.0.1
  port: 8081
  cors:
    origins:
      - http://127.0.0.1:8080
      - http://localhost:8080
```

The `server` section defines how the server should start (port, host) and the 
associated CORS configuration. In the example above, we're explicitly allowing
our `localhost:8080` to connect, but we can also set `cors.all: true` to allow
any origin to connect.

### Usage

```
Backend app that will aggregate playlists

Usage:
  cuites.stream <options> [flags]

Flags:
  -i, --client_id string              Define the SoundCloud Client ID
      --conf string                   configuration file to use
      --debug                         Enable or disable debug mode
  -h, --help                          help for cuites.stream
      --log.format string             one of text or json (default "text")
      --log.level string              one of debug, info, warn, error or fatal (default "info")
      --log.line                      enable filename and line in logs
      --server.cors.all               defines that all origins are allowed
      --server.cors.disabled          disable CORS completely
      --server.cors.expose strings    array of exposed headers
      --server.cors.headers strings   array of allowed headers (default [Origin,Authorization,Content-Type])
      --server.cors.methods strings   array of allowed method when cors is enabled (default [GET,PUT,POST,DELETE,OPTION,PATCH])
      --server.cors.origins strings   array of allowed origins (overwritten if all is active)
      --server.host string            host on which the server should listen (default "127.0.0.1")
      --server.mode string            server mode can be either 'debug', 'test' or 'release' (default "release")
      --server.port int               port on which the server should listen (default 8080)
      --user_ids strings              SoundCloud users to search the playlists from
```

