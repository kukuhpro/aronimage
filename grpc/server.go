package grpc

import (
	"aronimage/lib"
	pt "aronimage/protocols"
	"context"
)

type Server struct {
	Builder lib.AronImageBuilder
}

/**
 * Processing image for handle multiple process
 * @param ctx context.Context
 * @param req *pt.ProcessingImageRequest
 * @return *pt.ProcessingImageResponse | error
 */
func (s *Server) ProcessingImage(ctx context.Context, req *pt.ProcessingImageRequest) (*pt.ProcessingImageResponse, error) {
	var response pt.ProcessingImageResponse

	aron := s.Builder.WithBase64(req.Base64String).WithFileName(req.Imagename).Build(req.Name)
	if req.Prefixpath != "" {
		aron.PrefixPath = req.Prefixpath
	}

	err := aron.ProcessImage()

	if err == nil {
		response.Status = 1
	} else {
		response.Status = 99
	}

	return &response, err
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
