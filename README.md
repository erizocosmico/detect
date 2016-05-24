# detect
[![Build Status](https://travis-ci.org/mvader/detect.svg?branch=master)](https://travis-ci.org/mvader/detect) [![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org) [![GoDoc](https://godoc.org/github.com/mvader/detect?status.svg)](http://godoc.org/github.com/mvader/detect)

Golang library to detect the device platform given an user agent.

```go
package main

import (
        "fmt"

        "github.com/mvader/detect"
)

func main() {
        var userAgent = "Mozilla/5.0 (iPhone; U; CPU like Mac OS X; en) AppleWebKit/420.1 (KHTML, like Gecko) Version/3.0 Mobile/4A102 Safari/419"

        platform := detect.Platform(userAgent)
        if platform == detect.Android {
                fmt.Println("Android platform detected!")
        } else {
                fmt.Printf("Platform %s detected instead :(\n", platform)
        }

        if detect.IsMobile(userAgent) {
                fmt.Println("Is mobile!")
        }
}
```
