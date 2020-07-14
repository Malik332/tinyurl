# TinyURL - Make Personal URL Shortening Service Easy

[![Build Status](https://travis-ci.com/tinyurl/tinyurl.svg?branch=master)](https://travis-ci.com/tinyurl/tinyurl)  [![Go Report Card](https://goreportcard.com/badge/github.com/tinyurl/tinyurl)](https://goreportcard.com/report/github.com/tinyurl/tinyurl)  [![GoDoc](https://godoc.org/github.com/tinyurl/tinyurl?status.svg)](https://godoc.org/github.com/tinyurl/tinyurl)

<p align="center">
  <a href="https://adolphlwq.xyz" target="_blank">
    <img src="assets/tinyurl.gif" width="700px">
    <br>
    Live Demo
  </a>
</p>

<p align="center">a url shorten web service written by Golang, Vue and Gin.</p>

## Requisites
- Golang(1.12+)
- MySQL/Sqlite3
- make
- Docker

## Quick Start
1. clone project to **GOPATH**
```
git clone https://github.com/tinyurl/tinyurl.git $GOPATH/src/github.com/tinyurl/tinyurl
```
2. build binary
```
make
```
3. change config in `default.properties`
4. run binary
```
./tinyurl -config default.properties
```
5. open index.html in `frontend/` with broswer
6. default Swagger API url is `http://0.0.0.0:8877/swagger/index.html`

## TODOs
- [X] validate input url format
- [X] improve random generate string algorithm
    - [X] use math/rand.Read instead math/rand.Intn func
- [X] use logrus replace golang log lib
- [X] reserch [wrk](https://github.com/wg/wrk)
- [X] add test case
- [ ] Backend
  - [X] data type support multi database(index in sender)
  - [X] adjust short path generating algorithm
  - [X] Swagger for api management
  - [ ] custom short url
  - [ ] API rate
  - [ ] Admin account?
  - [ ] count each url parse time (high concurrent situation)
- [ ] Frontend
  - [ ] qrcode support

## Reference
- [GitHub-YOURLS](https://github.com/YOURLS/YOURLS)
- [知乎-短 URL 系统是怎么设计的？](https://www.zhihu.com/question/29270034)
- [GitHub/Ourls](https://github.com/takashiki/Ourls)
- [GitHub/uriuni](https://github.com/dchest/uniuri)