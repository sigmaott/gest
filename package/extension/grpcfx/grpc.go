package grpcfx

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func ForRoot(uri string, options ...grpc.ServerOption) fx.Option {
	return fx.Module("grpcfx",
		fx.Provide(
			fx.Annotate(
				func() *grpc.Server {

					return NewGrpcServer(options...)
				},
				fx.ResultTags(`name:"grpcServer"`)),
		),
		fx.Provide(
			fx.Annotate(
				func() string {
					return uri
				},
				fx.ResultTags(`name:"grpcUri"`))),

		fx.Provide(RegisterGRPCHooks))
}

func NewGrpcServer(options ...grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(options...)
}

func AsRoute(f any, annotation ...fx.Annotation) any {
	annotation = append(annotation, fx.As(new(IGrpcController)),
		fx.ResultTags(`group:"grpcControllers"`))
	return fx.Annotate(
		f,
		annotation...,
	)
}
