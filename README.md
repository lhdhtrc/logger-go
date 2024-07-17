## Go Logger
zap encapsulated log library

### How to use it?
`go get github.com/lhdhtrc/logger-go`

```go
package main

import loger "github.com/lhdhtrc/logger-go/pkg"

func main() {
	// The second parameter is to expose the logs, either locally or remotely
	instance := loger.New(&loger.ConfigEntity{}, nil)
}
```

### Finally
- If you feel good, click on star.
- If you have a good suggestion, please ask the issue.