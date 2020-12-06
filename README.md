# YSON

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)

## About <a name = "about"></a>

A simple CLI tool to convert a YAML file into JSON. Written in Go.

## Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to have Go installed in order to build/install this tool.

<a href="https://golang.org/dl/">Install Go</a>


### Build

```
go build ./cmd/yson
```

### Test

```
go test yson.com/yson...
```

### Install

From the project folder run:

```
go install yson.com/yson
```
Make sure ```GOPATH/bin``` is in your ```$PATH```. To check your ```GOPATH``` run:

```
go env GOPATH
```

## Usage <a name = "usage"></a>

```
./yson <filename>.yaml
```
If you have the executable in your $PATH:
```
yson <filename>.yaml
```
You can also use a pipe to pass yaml data:

```
cat file.yaml | yson
```
### Options
[--raw] Prints the raw string instead of the default 'pretty' printed version.

```
yson --raw <filename>.yaml
```
