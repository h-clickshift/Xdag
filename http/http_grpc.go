package main

import (
    "context"
    gw "grpc_go/proto/pb"
)

var (
    addr = "localhost:8888"
)

func run_http() error {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()
    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}
    err := gw.RegisterFileCheckHandlerFromEndpoint(ctx, mux, addr, opts)
    if err != nil {
        slog.Error("Error Register Http Handler", err)
        return err
    }
    return http.ListenAndServe(":8889", mux)
}

func main() {
    flag.Parse()
    if err := run_http(); err != nil {
    slog.Error("Error Register Http Handler", err)
    }
}