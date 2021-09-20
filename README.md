# dotenv-validator

![CI](https://github.com/duffn/dotenv-validator/actions/workflows/ci.yml/badge.svg) [![codecov](https://codecov.io/gh/duffn/dotenv-validator/branch/main/graph/badge.svg?token=LMF5XXIA8A)](https://codecov.io/gh/duffn/dotenv-validator) [![Go Report Card](https://goreportcard.com/badge/github.com/duffn/dotenv-validator)](https://goreportcard.com/report/github.com/duffn/dotenv-validator) [![Go Reference](https://pkg.go.dev/badge/github.com/duffn/dotenv-validator.svg)](https://pkg.go.dev/github.com/duffn/dotenv-validator)

`dotenv-validator` is a simple package that validates your environment based upon a formatted `.env.sample` file. Based upon [dotenv_validator](https://github.com/fastruby/dotenv_validator) from [FastRuby.io](https://github.com/fastruby?type=source).

## Usage

### Installation

```
go get github.com/duffn/dotenv-validator
```

### Configuring your environment variables

Tell `dotenv-validator` how you expect your environment variables by commenting them in your `.env.sample` file.

```
VAR1=testing # required
VAR2=123 # format=int
VAR3=notrequired
VAR4=bob@bobloblaw.com # required,format=email
```

### Formats

- `int`, `integer`
- `float`
- `str`, `string` (this is always true!)
- `email` (by [`mail.ParseAddress`](https://pkg.go.dev/net/mail#ParseAddress))
- `url` (by [`url.ParseRequestURI`](https://pkg.go.dev/net/url#ParseRequestURI))
- anything else is treated as a regular expression!

### `.env.sample`

```
VAR1=bob # required
VAR2=lobloaw # required,format=str
VAR3=1.3415 # format=float
VAR4=ABCDEF # format=[A-Z]+
VAR5=notrequired
```

### Running

```go
package main

import (
	"fmt"

	validator "github.com/duffn/dotenv-validator"
)

func main() {
	err := validator.Validate()
	fmt.Println(err)
}
```

Choose your sample file.

```go
package main

import (
	"fmt"

	validator "github.com/duffn/dotenv-validator"
)

func main() {
	err := validator.ValidateWithFilename("env_sample")
	fmt.Println(err)
}
```

## License

[MIT](https://opensource.org/licenses/MIT)
