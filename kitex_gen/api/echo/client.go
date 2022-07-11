// Code generated by Kitex v1.9.2. DO NOT EDIT.

package echo

import (
	"code.byted.org/kite/kitex/byted"
	"code.byted.org/kite/kitex/client"
	"code.byted.org/kite/kitex/client/callopt"
	"context"
	"github.com/daidai21/kitex_pr_issue433/kitex_gen/api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Echo(ctx context.Context, req *api.Request, callOptions ...callopt.Option) (r *api.Response, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	config := byted.NewClientConfig()
	config.DestService = destService

	options = append(options, byted.ClientSuiteWithConfig(serviceInfo(), config))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kEchoClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kEchoClient struct {
	*kClient
}

func (p *kEchoClient) Echo(ctx context.Context, req *api.Request, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Echo(ctx, req)
}

// NewClientWithBytedConfig creates a client for the service defined in IDL.
func NewClientWithBytedConfig(destService string, config *byted.ClientConfig, opts ...client.Option) (Client, error) {
	if config == nil {
		config = byted.NewClientConfig()
	}
	config.DestService = destService

	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, byted.ClientSuiteWithConfig(serviceInfo(), config))
	options = append(options, opts...)
	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kEchoClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClientWithBytedConfig creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClientWithBytedConfig(destService string, config *byted.ClientConfig, opts ...client.Option) Client {
	kc, err := NewClientWithBytedConfig(destService, config, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}
