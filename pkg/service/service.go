package service

import (
	"context"
	client "space-traders-playground/pkg/clients/openapi"
	"space-traders-playground/pkg/dom"
	errors "space-traders-playground/pkg/errors"
	repository "space-traders-playground/pkg/repository"
)

type Service interface {
	RegisterNewAgent(ctx context.Context, faction string, email, symbol string) (*client.Register201Response, error) 
}

type service struct {
	client *client.APIClient
	repository *repository.Repository
}

func NewService() Service {
	cfg := client.NewConfiguration()
	client := client.NewAPIClient(cfg)
	return &service{
		client: client,
	}
}

func (s *service) RegisterNewAgent(ctx context.Context, faction, email, symbol string) (*client.Register201Response, error) {
	
	f := client.FactionSymbols(faction)
	if !f.IsValid() {
		return nil, errors.ErrInvalidFaction
	}
	
	apiResp, _, err := s.client.DefaultAPI.Register(ctx).RegisterRequest(client.RegisterRequest{Email: &email, Faction: f, Symbol: symbol}).Execute()

	user := dom.User {
		Email: email,
		Symbol: symbol,
		Faction: faction,
	}

	//TODO add user to database

	return apiResp, err
}