# Context-Aware Golang io Code

[![GoDoc](https://pkg.go.dev/badge/github.com/bassosimone/iox)](https://pkg.go.dev/github.com/bassosimone/iox) [![Build Status](https://github.com/bassosimone/iox/actions/workflows/go.yml/badge.svg)](https://github.com/bassosimone/iox/actions) [![codecov](https://codecov.io/gh/bassosimone/iox/branch/main/graph/badge.svg)](https://codecov.io/gh/bassosimone/iox)

The `iox` Go package contains context-aware extensions for the `io` package.

For example:

```go
import (
	"context"
	"bytes"
	"net/http"

	"github.com/bassosimone/iox"
	"github.com/bassosimone/runtimex"
)

// The iox package is context aware
var ctx context.Context

// We assume you somehow created the context possibly binding it to signals
// ...

// Create a request bound to an existing context
req := runtimex.PanicOnError1(http.NewRequestWithContext(ctx, "GET", "https://example.com", nil))

// Perform the round trip, whose maximum duration is context bounded
resp := runtimex.PanicOnError1(http.DefaultClient.Do(req))
defer resp.Body.Close()

// Use iox.CopyContext to read the body in such a way that, if the context is
// canceled, we immediately interrupt reading.
//
// CopyContext only closes resp.Body on context cancellation. On success,
// we rely on the `defer resp.Body.Close()` statement provided above.
buff := &bytes.Buffer{}
writer := iox.NewLockedWriteCloser(iox.NopWriteCloser(buff))
count, err := iox.CopyContext(ctx, writer, resp.Body)
```

See also the [example_test.go](example_test.go) example.

## Installation

To add this package as a dependency to your module:

```sh
go get github.com/bassosimone/iox
```

## Development

To run the tests:
```sh
go test -v .
```

To measure test coverage:
```sh
go test -v -cover .
```

## License

```
SPDX-License-Identifier: GPL-3.0-or-later
```

## History

Inspired by [ooni/probe-cli](https://github.com/ooni/probe-cli/blob/v3.20.1/internal/netxlite/iox.go).
