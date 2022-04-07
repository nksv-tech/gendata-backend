package service

import (
	context "context"

	v1 "github.com/gen-data/gendata-server/pkg/gendata/v1"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGenDataService)

func NewGenDataService() *GenDataService {
	return &GenDataService{}
}

type GenDataService struct {
	v1.UnimplementedGenDataServiceServer
}

func (s *GenDataService) Gen(ctx context.Context, request *v1.GenRequest) (*v1.GenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GenDataService) PredefinedSettings(ctx context.Context, request *v1.PredefinedSettingsRequest) (*v1.PredefinedSettingsResponse, error) {
	return &v1.PredefinedSettingsResponse{}, nil
}
