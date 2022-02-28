# redis-golang-poc
![build workflow](https://github.com/go-redis/redis/actions/workflows/build.yml/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-redis/redis/v8)](https://pkg.go.dev/github.com/go-redis/redis/v8?tab=doc)
[![Documentation](https://img.shields.io/badge/redis-documentation-informational)](https://redis.uptrace.dev/)


Basic Golang Code to connect to Redis with defined functions to GET, SET and SET Key with EXPR expiry.


## Prerequisites
go-redis supports 2 last Go versions and requires a Go version with
[modules](https://github.com/golang/go/wiki/Modules) support. So make sure to initialize a Go
module:

```shell
go mod init github.com/my/repo
```

And then install go-redis/v8 (note _v8_ in the import; omitting it is a popular mistake):

```shell
go get github.com/go-redis/redis/v8
```

## Base Code

```go
import (
    "context"
    "github.com/go-redis/redis/v8"
    "fmt"
)

var ctx = context.Background()

func ExampleClient() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := rdb.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
    // Output: key value
    // key2 does not exist
}
```
