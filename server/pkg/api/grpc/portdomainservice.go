package grpc

import (
	"context"

	protos "github.com/lrotermund/port-challenge/server/pkg/api/portservice/v1"
	"github.com/rs/zerolog/log"
)

type PortDomainService struct {
	protos.UnimplementedPortDomainServer
}

func NewPortDomainService() *PortDomainService {
	return &PortDomainService{}
}

func (a *PortDomainService) CreatePort(
	ctx context.Context,
	cpr *protos.CreatePortRequest,
) (*protos.CreatePortResponse, error) {
	log.Info().Msgf("Handle CreatePort: PortID %s", cpr.Port.PortID)

	return &protos.CreatePortResponse{PortID: cpr.Port.PortID}, nil
}

func (a *PortDomainService) ReadPort(
	ctx context.Context,
	rpr *protos.ReadPortRequest,
) (*protos.ReadPortResponse, error) {
	log.Info().Msgf("Handle ReadPort: PortID %s", rpr.PortID)

	return &protos.ReadPortResponse{Port: &protos.Port{
		PortID: rpr.PortID,
	}}, nil
}
