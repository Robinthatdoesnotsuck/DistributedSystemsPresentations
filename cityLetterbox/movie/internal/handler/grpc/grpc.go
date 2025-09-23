package grpc_handler

import (
	"context"

	model "cityletterbox.com/metadata/pkg"
	"cityletterbox.com/movie/internal/controller/movie"
	"cityletterbox.com/src/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedMovieServiceServer
	ctrl *movie.Controller
}

func New(ctrl *movie.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetMovieDetails(ctx context.Context, req *gen.GetMovieDetailsRequest) (*gen.GetMovieDetailsResponse, error) {
	m, err := h.ctrl.Get(ctx, req.MovieId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetMovieDetailsResponse{MovieDetails: &gen.MovieDetails{Metadata: model.MetadataToProto(&m.Metadata), Rating: *m.Rating}}, nil
}
