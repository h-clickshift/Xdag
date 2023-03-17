package main

import (
    "fmt"
    "context"
    pb "grpc_go/proto/pb"
    "golang.org/x/exp/slog"
    "net"
    
    "google.golang.org/grpc"

)


var (
    ADDR string = "localhost:8888"
    rds_addr string = "localhost:6379"
)

func main() {
    defer func() error{
        if p := recover(); p != nil {
            slog.Warn("Error,Quit!")
            err := error.New(fmt.Sprintf("defer %+v", p))
            return err
        }
        return nil
    }()
    slog.Info("grpc server is up...")
    lis, err := net.Listen("tcp", ADDR)
    if err != nil {
        slog.Error(err.Error(), err)
        panic(err)
    }

    rds := rds.RefisterRedis(rds_addr, "user", "", 0, 10)
    _, err = rds.Ping(context.Background()).Result()
    if err != nil {
        slog.Error("redis conn err", err)
        panic(err)
    }
    s := grpc.NewServer()
    service := &Service{}
    service.SetRedis(s)

    pb.RegisterFileCheckServer(s, service)
    err = s.Serve(lis)
    if err != nil {
        slog.Error(err.Error(), err)
    }
}

func MD5(filepath string) (string, error) {
    file, err := os.Open(filepath)
    if err != nil {
        return "", err
       }
    hash := md5.New()
    _, _ =io.Copy(hash, file)
    return hex.EncodeToString(hash.Sum(nil)), nil
}

type Service struct {
    rds_client *redis.Client
}

func(s *Service) SetRedis(redis_client *redis_Client) {
    s.rds_client = redis_client
}

func(s *Service) Execute(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
    response := pb.Resp{}
    cmd := s.rds_client.Get(ctx, req.Path)
    if err := cmd.Err(); err == redis.nil {
    slog.Info("Execute receive request", slog.String("Req->Path", req.Path))

    handler, err := os.Stat(req.Path)
    if os.IsNotExist(err) {
    slog.Warn("Not Exists", slog.String("Req->Path", req.Path))
    return &response, err
    }

    response.Path = req.Path
    if handler.IsDir() {
        tree := make([]string, 0)
        walk := func(path string, info os.FileInfo, err error) error {
            tree = append(tree, path)
            return nil
        }
        filepath.Walk(req.Path, walk)
        for _, item := range tree {
            response.Content = append(response.Content, item...)
        }
        hash := md5.Sum(response.Content)
        response.Md5 = fmt.Sprintf("%x", hash)
        response.IsDir = true
    } else { // file
        content, err := ioutil.ReadFile(response.Path)
        if err != nil {
            slog.Warn("Read Failed!", slog.String("Req->Path", req.Path))
        }
        response.Content = content
        response.Md5, _ = MD5(response.Path)
        response.IsDir = false
    }

    _hash := map[string]string{
        "path": response.Path,
        "md5": response.Md5,
        "IsDir": strconv.FromatBool(response.IsDir),
        "content": string(response.Content),
        }
    s.rds_client.HSet(ctx, response.Path, _hash)
    } else { // already in db
        _response, _ := s.rds_client.HGetAll(ctx, req.Path).Result()
        response.Path = _response["path"]
        response.Md5 = _response["md5"]
        response.IsDir, _ = strconv.ParseBool(_response["IsDir"])
        response.Content = []byte(_response["content"])
    }
    return &response. nil
}




