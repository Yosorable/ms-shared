package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"syscall"

	"github.com/Yosorable/ms-shared/utils"
	"github.com/Yosorable/ms-shared/utils/consul"

	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/health"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// 本地运行rpc、consul，consul使用默认端口
func RunRpcServerInLocalHost(
	serviceName string,
	grpcServer *ggrpc.Server,
	gracefulStopFuncs ...func(),
) error {
	// 随机分配端口
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	port := lis.Addr().(*net.TCPAddr).Port

	// 健康检查
	healthcheck := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthcheck)

	consul, err := consul.NewConsul("127.0.0.1:8500")
	if err != nil {
		return err
	}
	consul.RegisterService(serviceName, "127.0.0.1", port)

	quit := make(chan os.Signal, 1)
	go func() {
		log.Printf("[%s] server listening at %v", serviceName, lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("[%s] failed to serve: %v", serviceName, err)
			quit <- syscall.SIGINT
		}
	}()

	gracefulStopFuncs = append(gracefulStopFuncs, func() {
		consul.Deregister(fmt.Sprintf("%s-%s-%d", serviceName, "127.0.0.1", port))
		grpcServer.GracefulStop()
	})

	utils.WaitForGracefullyStop(quit, gracefulStopFuncs...)

	return nil
}
