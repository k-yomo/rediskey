# rediskey

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)
![Tests Workflow](https://github.com/k-yomo/rediskey/workflows/Tests/badge.svg)
[![codecov](https://codecov.io/gh/k-yomo/rediskey/branch/main/graph/badge.svg)](https://codecov.io/gh/k-yomo/rediskey)
[![Go Report Card](https://goreportcard.com/badge/k-yomo/rediskey)](https://goreportcard.com/report/k-yomo/rediskey)

rediskey lets you organize redis key.

## example
```go
package main

import (
	"context"
	"github.com/k-yomo/rediskey"
	"github.com/uopeople-jp/uopeople-x/backend/pkg/redisutil"
	redis "pkg/mod/github.com/go-redis/redis/v8"
	"time"
)

var redisKeyNameSpaceAuth = rediskey.NewNamespace("auth", nil)

func main()  {
	redisClient, err := redisutil.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	if err != nil {
		panic(err)
	}
	key := rediskey.NewKey("session", "sess_id", redisKeyNameSpaceAuth)

	ctx := context.Background()
	err = redisClient.Set(ctx, key.String(), "value", 24 * time.Hour).Err()
	if err != nil {
		panic(err)
	}
}

```