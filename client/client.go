package main

import (
    "context"
    "grpc_go/proto/pb"
    "fmt"
    "time"
    "google.golang.org/grpc"
    "golang.org/x/exp/slog"

)

func main() {
    conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
    if err != nil {
        slog.Error(err.Error(), err)
        return
    }
    defer conn.Close()
    client := pb.NewFileCheckClient(conn)
    test := &pb.Req{Path: "some path"}
    time0 := time.Now()
    resp, err := client.Execute(context.Background(), test)
    if err != nil {
        slog.Error(err.Error(), err)
        return
    }
    fmt.Println(resp)
    fmt.Println("time cost: ", time.Now().Sub(time0))
}