package service

import (
	"context"

	"github.com/ron96G/go-api-hackathon/internal/api"
)

var _ api.StrictServerInterface = (*Service)(nil)

type Service struct {
}

func NewService(ctx context.Context) *Service {
	return &Service{}
}

// AddPet implements api.StrictServerInterface.
func (s *Service) AddPet(ctx context.Context, request api.AddPetRequestObject) (api.AddPetResponseObject, error) {
	panic("unimplemented")
}

// DeletePet implements api.StrictServerInterface.
func (s *Service) DeletePet(ctx context.Context, request api.DeletePetRequestObject) (api.DeletePetResponseObject, error) {
	panic("unimplemented")
}

// GetPetById implements api.StrictServerInterface.
func (s *Service) GetPetById(ctx context.Context, request api.GetPetByIdRequestObject) (api.GetPetByIdResponseObject, error) {
	panic("unimplemented")
}

// ListPets implements api.StrictServerInterface.
func (s *Service) ListPets(ctx context.Context, request api.ListPetsRequestObject) (api.ListPetsResponseObject, error) {
	panic("unimplemented")
}

// UpdatePet implements api.StrictServerInterface.
func (s *Service) UpdatePet(ctx context.Context, request api.UpdatePetRequestObject) (api.UpdatePetResponseObject, error) {
	panic("unimplemented")
}
