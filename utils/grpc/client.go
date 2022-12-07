package grpc

import (
	_ "github.com/mbobakov/grpc-consul-resolver"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service[T any] struct {
	Conn   *ggrpc.ClientConn
	Client T
}

func NewServiceInLocalHost[T any](
	serviceName string,
	newClientFunc func(ggrpc.ClientConnInterface) T,
) (*Service[T], error) {
	conn, err := ggrpc.Dial(
		"consul://127.0.0.1:8500/"+serviceName+"?healthy=true",
		// 负载均衡，指定round_robin策略
		ggrpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		ggrpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	client := newClientFunc(conn)
	return &Service[T]{
		Conn:   conn,
		Client: client,
	}, nil
}
