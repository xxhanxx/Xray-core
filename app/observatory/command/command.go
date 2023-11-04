//go:build !confonly
// +build !confonly

package command

import (
	"context"

	"github.com/xxhanxx/Xray-core/app/observatory"
	"github.com/xxhanxx/Xray-core/common"
	core "github.com/xxhanxx/Xray-core/core"
	"github.com/xxhanxx/Xray-core/features/extension"
	"google.golang.org/grpc"
)

type service struct {
	UnimplementedObservatoryServiceServer
	v *core.Instance

	observatory extension.Observatory
}

func (s *service) GetOutboundStatus(ctx context.Context, request *GetOutboundStatusRequest) (*GetOutboundStatusResponse, error) {
	resp, err := s.observatory.GetObservation(ctx)
	if err != nil {
		return nil, err
	}
	retdata := resp.(*observatory.ObservationResult)
	return &GetOutboundStatusResponse{
		Status: retdata,
	}, nil
}

func (s *service) Register(server *grpc.Server) {
	RegisterObservatoryServiceServer(server, s)
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, cfg interface{}) (interface{}, error) {
		s := core.MustFromContext(ctx)
		sv := &service{v: s}
		err := s.RequireFeatures(func(Observatory extension.Observatory) {
			sv.observatory = Observatory
		})
		if err != nil {
			return nil, err
		}
		return sv, nil
	}))
}
