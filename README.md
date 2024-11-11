# Go language course

[![Go Report Card](https://goreportcard.com/badge/github.com/RedHatOfficial/GoCourse)](https://goreportcard.com/report/github.com/RedHatOfficial/GoCourse)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/RedHatOfficial/GoCourse)

This repository contains material for Go language course in Go [present format](https://godoc.org/golang.org/x/tools/present). To view the slides run the present command in the top level directory.

## Step by step guide to view the slides locally

This assumes that you have go compiler and git installed and on $PATH of your system.

```shell
git clone https://github.com/RedHatOfficial/GoCourse.git
cd GoCourse
go run golang.org/x/tools/cmd/present
```

Afterwards connect with your browser to the mentioned address. To terminate the server use Ctrl+C.

## Sharing slides with other people

It is possible to start service that serves slides to other computers via HTTP.
```shell
go run golang.org/x/tools/cmd/present -http=:3999 -play=false
```
