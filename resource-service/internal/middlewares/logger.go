package middlewares

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func UnaryLoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	start := time.Now()
	resp, err := handler(ctx, req)

	log.Printf(
		"method=%s duration=%s error=%v",
		info.FullMethod,
		time.Since(start),
		err,
	)

	return resp, err
}
