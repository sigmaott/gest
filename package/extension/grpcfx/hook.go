package grpcfx

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Params struct {
	fx.In
	GrpcPort        int               `name:"grpcPort"`
	GrpcServer      *grpc.Server      `name:"grpcServer"`
	RegisterServers []IGrpcController `group:"grpcControllers"`
}
type Result struct {
	fx.Out
	Controller IGrpcController `group:"grpcControllers"`
}

type IGrpcController interface {
	RegisterGrpcController(server *grpc.Server)
}

func InitGrpcController(server *grpc.Server, registerServers []IGrpcController) {
	for _, s := range registerServers {
		s.RegisterGrpcController(server)
	}

}

func RegisterGRPCHooks(
	lifecycle fx.Lifecycle,
	params Params,
) *grpc.Server {
	grpcSever := params.GrpcServer
	registerServers := params.RegisterServers

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					
					InitGrpcController(grpcSever, registerServers)

					lis, err := net.Listen("tcp", fmt.Sprintf(":%d", params.GrpcPort))
					if err != nil {
						log.Fatalf("failed to listen: %v", err)
					}
					log.Printf("starting gRPC server on %s", lis.Addr().String())
					if err = grpcSever.Serve(lis); err != nil {
						log.Fatalf("failed to int grpcSever: %v", err)
					}

				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				grpcSever.GracefulStop()
				return nil

			},
		})
	return grpcSever

}
