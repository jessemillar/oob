# oob
[![Man Hours](https://img.shields.io/endpoint?url=https%3A%2F%2Fmh.jessemillar.com%2Fhours%3Frepo%3Dhttps%3A%2F%2Fgithub.com%2Fjessemillar%2Foob.git)](https://jessemillar.com/r/man-hours) [![Go Report Card](https://goreportcard.com/badge/github.com/jessemillar/oob)](https://goreportcard.com/report/github.com/jessemillar/oob)

A cute Go library/CLI tool that changes vowels to “oob”.

## Library

### Usage

See [this Go Playground](https://play.golang.org/p/OLFdgmTvC54) for a minimal usage example.

## CLI

### Installation

> There is also a Windows binary available on [the Releases page](https://github.com/jessemillar/oob/releases)

```
curl -L https://github.com/jessemillar/oob/releases/latest/download/oob-linux-amd64 -o oob && chmod +x oob && mv oob /usr/local/bin
```

### Usage

```
oob "Hello there!"
Hooblloob thoobroob!

oob -f LICENSE
MOOBT
...

oob "Speak this text" | say
```
