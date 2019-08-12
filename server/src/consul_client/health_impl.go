package consul_client

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// HealthImpl 健康检查实现
type HealthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (h HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h HealthImpl)Watch(req *grpc_health_v1.HealthCheckRequest, watch grpc_health_v1.Health_WatchServer) error {
	fmt.Println(req)
	fmt.Println(watch)
	return nil
}