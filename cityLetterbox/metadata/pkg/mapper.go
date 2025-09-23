package model

import "cityletterbox.com/src/gen"

func MetadataToProto(m *Metadata) *gen.Metadata {
	return &gen.Metadata{
		Id: m.ID,
		Movie: &gen.MoviesStruct{
			Director:    m.Director,
			Title:       m.Title,
			Description: m.Description,
		},
	}
}

func MetadataFromProto(m *gen.Metadata) *Metadata {
	return &Metadata{
		ID:          m.Id,
		Title:       m.Movie.Title,
		Description: m.Movie.Description,
		Director:    m.Movie.Director,
	}
}
