package client

import (
	"context"

	"github.com/elevenyi11/gostudy/seckill/pb/proto"
	"github.com/elevenyi11/gostudy/seckill/pkg/discover"
	"github.com/elevenyi11/gostudy/seckill/pkg/loadbalance"
	"github.com/opentracing/opentracing-go"
)

type OAuthClient interface {
	CheckToken(ctx context.Context, tracer opentracing.Tracer, request *proto.CheckTokenRequest) (*proto.
		CheckTokenResponse, error)
}

type OAuthClientImpl struct {
	manager     ClientManager
	serviceName string
	loadBalance loadbalance.LoadBalance
	tracer      opentracing.Tracer
}

func (impl *OAuthClientImpl) CheckToken(ctx context.Context, tracer opentracing.Tracer, request *proto.CheckTokenRequest) (*proto.
	CheckTokenResponse, error) {
	response := new(proto.CheckTokenResponse)
	if err := impl.manager.DecoratorInvoke("/proto.OAuthService/CheckToken", "token_check", tracer, ctx, request,
		response); err == nil {
		return response, nil
	} else {
		return nil, err
	}
}

func NewOAuthClient(serverName string, lb loadbalance.LoadBalance, tracer opentracing.Tracer) (OAuthClient, error) {
	if serverName == "" {
		serverName = "oauth"
	}
	if lb == nil {
		lb = defaultLoadBalance
	}

	return &OAuthClientImpl{
		manager: &DefaultClientManager{
			serviceName:     serverName,
			loadBalance:     lb,
			discoveryClient: discover.ConsulService,
			logger:          discover.Logger,
		},
		serviceName: serverName,
		loadBalance: lb,
		tracer:      tracer,
	}, nil
}
