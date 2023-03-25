package helloplug

import (
	"fmt"
	"os"
	"os/exec"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	plug "github.com/xcd0/gim/go-plugin"
)

func Run() {
	pluginName := os.Getenv("GREETER_PLUGIN")
	if pluginName == "" {
		fmt.Println("no plugin")
		return
	}

	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(
		&plugin.ClientConfig{
			HandshakeConfig: plug.Handshake,
			Plugins:         plug.PluginMap,
			Cmd:             exec.Command("sh", "-c", os.Getenv("GREETER_PLUGIN")),
			AllowedProtocols: []plugin.Protocol{
				plugin.ProtocolGRPC,
			},
			Logger: hclog.New(&hclog.LoggerOptions{
				Output: hclog.DefaultOutput,
				Level:  hclog.Error, // デフォルトで hclog.Trace
				Name:   "plugin",
			}),
		},
	)
	defer client.Kill()

	rpcClient, _ := client.Client()
	raw, _ := rpcClient.Dispense("greeter")
	say, _ := raw.(plug.Greeter)

	msg, _ := say.Say("gopher")
	fmt.Println(msg)
}
