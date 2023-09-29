# Go-Utils

Go-Utils is an integrated Golang utility library designed to provide common functionalities and utility functions to accelerate and simplify the development process of Golang applications.

## Features

- Provides a rich set of utility functions and handy tools that cover various common programming tasks and functionalities.
- Modular design, easy to use, and convenient to integrate into existing Golang projects.
- Carefully crafted and optimized for high performance and efficiency.
- Continuously updated and maintained to adapt to evolving development needs.

## Installation

Use the go get command to install Go-Utils:

```bash
go get -u github.com/cc-hearts/go-utils
```

## Usage

Here are some example code snippets demonstrating the usage of Go-Utils:

```go
package main

import "github.com/cc-hearts/go-utils/lib/validator"

func main() {
	if validator.IsValidJson(`{"name": "cc heart"}`) == true {
		// ...
    }
}

```

## LICENSE

[MIT](./License)