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

## Usage <a name = "usage"></a>

```
./yson <filename>.yaml
```
If you have the executable in your $PATH:
```
yson <filename>.yaml
```


By default, the converted string is pretty printed. If you want the raw string, use the ```--raw``` flag:
```
yson --raw <filename>.yaml
```
