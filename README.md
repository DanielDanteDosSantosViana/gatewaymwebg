# Gatewaymwebg
Gateway MWEBG to redirect requests to a given server.

[![Build Status](https://api.travis-ci.org/DanielDanteDosSantosViana/gatewaymwebg.svg)](https://api.travis-ci.org/DanielDanteDosSantosViana/gatewaymwebg.svg?branch=master)

## Instalation
```
   go install ../gatewaymwebg/
```

## Configuration
To configuration you need create one file in root directory with name config.toml.
example config file :
```
[[services]]
urlDest = "http://danielddsv.com.br/movie_store"
urlSrc = "/exemplo1"
host ="localhost"
method = "GET"
type = "json/application"
queryParam =""
```
##Note:
Description about a parameters in file configuration
```
urlDest -> Url that represent redirect to the destination
ulrSrc -> path parameter(key) that you need to send and redirect to urlDest
host -> host to Dest
method -> method http that represent service(urldest)
type -> type param that a service need

```

