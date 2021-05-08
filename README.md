# rediskey

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)
![Tests Workflow](https://github.com/k-yomo/rediskey/workflows/Tests/badge.svg)
[![codecov](https://codecov.io/gh/k-yomo/rediskey/branch/main/graph/badge.svg)](https://codecov.io/gh/k-yomo/rediskey)
[![Go Report Card](https://goreportcard.com/badge/k-yomo/rediskey)](https://goreportcard.com/report/k-yomo/rediskey)

rediskey lets you organize redis key in an officially recommended naming convention with document-oriented way.

Officially recommended naming convention is following.
> -  Very short keys are often not a good idea. There is little point in writing "u1000flw" as a key if you can instead write "user:1000:followers". The latter is more readable and the added space is minor compared to the space used by the key object itself and the value object. While short keys will obviously consume a bit less memory, your job is to find the right balance.
> - Try to stick with a schema. For instance "object-type:id" is a good idea, as in "user:1000". Dots or dashes are often used for multi-word fields, as in "comment:1234:reply.to" or "comment:1234:reply-to".

More details are written in [https://redis.io/topics/data-types-intro#redis-keys](https://redis.io/topics/data-types-intro#redis-keys).

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
	key := redisKeyNameSpaceAuth.NewKey("session", "sess_id")

	ctx := context.Background()
	// key.String() => "auth:session:sess_id"
	err = redisClient.Set(ctx, key.String(), "value", 24 * time.Hour).Err()
	if err != nil {
		panic(err)
	}
}

```