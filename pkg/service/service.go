package service

import (
	"context"
	client "space-traders-playground/pkg/clients/openapi"
	"space-traders-playground/pkg/dom"
	errors "space-traders-playground/pkg/errors"
	"space-traders-playground/pkg/repository/mongo"

	"go.uber.org/zap"
)

type Service interface {
	RegisterNewAgent(ctx context.Context, faction string, email, symbol string) (*client.Register201Response, error)
}

type service struct {
	client     *client.APIClient
	logger     *zap.Logger
	repository mongo.Repository
}

func NewService(logger *zap.Logger, repo mongo.Repository) Service {
	cfg := client.NewConfiguration()
	client := client.NewAPIClient(cfg)

	return &service{
		client:     client,
		repository: repo,
		logger:     logger,
	}
}

func (s *service) RegisterNewAgent(ctx context.Context, faction, email, symbol string) (*client.Register201Response, error) {

	f := client.FactionSymbols(faction)
	if !f.IsValid() {
		return nil, errors.ErrInvalidFaction
	}

	apiResp, _, err := s.client.DefaultAPI.Register(ctx).RegisterRequest(client.RegisterRequest{Email: &email, Faction: f, Symbol: symbol}).Execute()
	if err != nil {
		s.logger.Error(errors.ErrFailedAPIRequest.Error(), zap.String("type", "RegisterNewAgent"), zap.Error(err))
		return nil, errors.ErrFailedAPIRequest
	}

	user := dom.User{
		Email:   email,
		Symbol:  symbol,
		Faction: faction,
	}

	//TODO add user to database
	err = s.repository.AddUser(ctx, user)

	return apiResp, err
}
