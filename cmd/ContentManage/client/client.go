package main

import (
	"ContentManage/api/operate"
	"context"
	"log"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func main() {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	client := operate.NewAppClient(conn)
	reply, err := client.UpdateContent(context.Background(), &operate.UpdateContentReq{
		Content: &operate.Content{
			Id:          3,
			Title:       "test content_manage create",
			VideoUrl:    "https://example.com/video.mp4",
			Author:      "wangyao",
			Description: "test update",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] CreateContent %v\n", reply)
}
