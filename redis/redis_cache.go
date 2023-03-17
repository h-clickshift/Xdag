package rds

import (
    "time"
)

func RegisterRedis(addr string, user string, passwd string, db int, dialTimeout time.Duration) {
    return redis.NewClient(&redis.Options{
        Addr: addr,
        Password: passwd,
        DB:db,
        DialTimeout: dialTimeout,
    })
}