# Hades

## Project overview
This repository is a small Go library providing utilities and helper functions for website or backend api development (Gin framework is my favorite). It includes features:

- send email by SMTP: mail.go
- html template functions: html.go
- token generation and validation: token.go
- price formatting: price.go
- spam detection: spam.go


## Install
Clone or fetch the package in module mode (modules are recommended for Go 1.17+):

```
go get github.com/sunzhongwei/hades@latest
```


## Upgrade to latest version
To upgrade to the latest version of the library, run:

```
go get -u github.com/sunzhongwei/hades
```


## Usage
Import and use in your project:

```go
import "github.com/sunzhongwei/hades"

// Example: call an exported function from the library (adjust to actual exported symbols)
// result := hades.SomeFunction(args)
```

## License
MIT License. See the LICENSE file for details.
