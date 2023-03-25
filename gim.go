package gim

import (
	"fmt"

	"github.com/rivo/tview"
	_ "github.com/urfave/cli/v2"

	"github.com/hashicorp/go-plugin"
	plug "github.com/xcd0/gim/go-plugin"
)

func Run() {
	if false {
		test()
	}
}

func test() {
	// 導入しそうなパッケージの呼び出しテスト用
	test_tview()
	test_plugin()
}

func test_tview() {
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}

func test_plugin() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plug.Handshake,
		Plugins: map[string]plugin.Plugin{
			"greeter": &plug.GreeterPlugin{Impl: &Hello{}},
		},

		GRPCServer: plugin.DefaultGRPCServer,
	})
}

type Hello struct{}

func (h *Hello) Say(name string) (string, error) {
	return fmt.Sprintf("hello %s", name), nil
}
