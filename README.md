# go-wunderground
A go package for interacting with the Wunderground weather API.

[![GoDoc](https://godoc.org/github.com/ZachtimusPrime/go-wunderground/wunderground?status.svg)](https://godoc.org/github.com/ZachtimusPrime/go-wunderground/wunderground)
[![Build Status](https://travis-ci.org/ZachtimusPrime/go-wunderground.svg?branch=master)](https://travis-ci.org/ZachtimusPrime/go-wunderground)
[![Coverage Status](https://coveralls.io/repos/github/ZachtimusPrime/go-wunderground/badge.svg?branch=master)](https://coveralls.io/github/ZachtimusPrime/go-wunderground?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ZachtimusPrime/go-wunderground)](https://goreportcard.com/report/github.com/ZachtimusPrime/go-wunderground)

## Table of Contents ##

* [Installation](#installation)
* [Usage](#usage)

## Installation ##

```bash
go get "github.com/ZachtimusPrime/go-wunderground/wunderground"
```

## Usage ##

Construct a new Wunderground client to pull weather data for a specific city.

For example:

```go
package main

import (
        "github.com/ZachtimusPrime/go-wunderground/wunderground"
)

func main() {

		// Create new Wunderground client
		client := wunderground.NewClient(nil, "TN", "Nashville", {your-API-key})

		// Get current weather data
        weather, err := client.GetWeather()
        if err != nil {
            log.Print(err)
        }
}

```
