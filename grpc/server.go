package grpc

import (
	"aronimage/lib"
	pt "aronimage/protocols"
	"context"
)

type Server struct {
	Builder lib.AronImageBuilder
}

func (s *Server) ProcessingImage(ctx context.Context, req *pt.ProcessingImageRequest) (*pt.ProcessingImageResponse, error) {
	return nil, nil
}

func (s *Server) SetupConfiguration(ctx context.Context, req *pt.SetupConfigurationRequest) (*pt.SetupConfigurationResponse, error) {
	return nil, nil
}

func (s *Server) GetImageListPath(ctx context.Context, req *pt.ImageListPathRequest) (*pt.ImageListPathResponse, error) {
	return nil, nil
}

func NewServer() *Server {
	var server Server

	return &server
}
