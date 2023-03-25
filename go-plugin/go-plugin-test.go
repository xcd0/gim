package plug

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/xcd0/gim/proto/greeter"
	"google.golang.org/grpc"
)

type Greeter interface {
	Say(name string) (string, error)
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Greeter
	greeter.UnimplementedGreeterServer
}

func (m *GRPCServer) Say(ctx context.Context, r *greeter.Request) (*greeter.Reply, error) {
	msg, err := m.Impl.Say(r.Name)
	if err != nil {
		return nil, err
	}
	return &greeter.Reply{Message: msg}, nil
}

type GRPCClient struct{ client greeter.GreeterClient }

func (m *GRPCClient) Say(name string) (string, error) {
	r := &greeter.Request{Name: name}
	res, err := m.client.Say(context.Background(), r)
	if err != nil {
		return "", err
	}
	return res.Message, nil
}

type GreeterPlugin struct {
	plugin.Plugin
	Impl Greeter
}

func (p *GreeterPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	greeter.RegisterGreeterServer(s, &GRPCServer{Impl: p.Impl}) // TODO: impl
	return nil
}

func (p *GreeterPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: greeter.NewGreeterClient(c)}, nil
}

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "GREETER_PLUGIN",
	MagicCookieValue: "greeter",
}

var PluginMap = map[string]plugin.Plugin{
	"greeter": &GreeterPlugin{},
}
