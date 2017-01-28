# go-wunderground
A go package for interacting with the Wunderground weather API.

[![GoDoc](https://godoc.org/github.com/ZachtimusPrime/go-wunderground/wunderground?status.svg)](https://godoc.org/github.com/ZachtimusPrime/go-wunderground/wunderground)
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
