# Go language course

[![Go Report Card](https://goreportcard.com/badge/github.com/RedHatOfficial/GoCourse)](https://goreportcard.com/report/github.com/RedHatOfficial/GoCourse)

This repository contains material for Go language course in Go [present format](https://godoc.org/golang.org/x/tools/present). To view the slides run the present command in the top level directory.

## Online versions of all slides

It's possible to view all slides online thanks to [talks.godoc.org](https://talks.godoc.org/) service:

1. [Introduction, Syntax, Formatting, Packages, Types](https://talks.godoc.org/github.com/RedHatOfficial/GoCourse/lesson1.slide)
1. [Go language fundamentals, the first part](https://talks.godoc.org/github.com/RedHatOfficial/GoCourse/lesson2.slide)
1. [Go language fundamentals, second part](https://talks.godoc.org/github.com/RedHatOfficial/GoCourse/lesson3.slide)
1. [Concurrency](https://talks.godoc.org/github.com/RedHatOfficial/GoCourse/lesson4.slide)
1. [Testing, CGo and tools](https://talks.godoc.org/github.com/RedHatOfficial/GoCourse/lesson5.slide)
1. [HTTP & Friends](https://talks.godoc.org/github.com/RedHatOfficial/GoCourse/lesson6.slide)
1. [External libraries and tools](https://talks.godoc.org/github.com/RedHatOfficial/GoCourse/lesson7.slide)
1. [Golang Testing Frameworks](https://talks.godoc.org/github.com/RedHatOfficial/GoCourse/testing.slide)

## Step by step guide to view the slides locally

This assumes that you have go compiler and git installed and on $PATH of your system.

```
git clone https://github.com/RedHatOfficial/GoCourse.git
cd GoCourse
go get golang.org/x/tools/cmd/present
present
```

Afterwards connect with you browser to the mentioned address. To terminate the server use Ctrl+C.

## Sharing slides with other people

It is possible to start service that serves slides to other computers via HTTP.
Please look at `present_and_share.sh` to see how it can be done.
