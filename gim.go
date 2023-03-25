package gim

import (
	"fmt"

	"github.com/rivo/tview"
	_ "github.com/urfave/cli/v2"
)

func Run() {
	if false {
		test()
	}
}

func test() {
	// 導入しそうなパッケージの呼び出しテスト用
	test_tview()
}

func test_tview() {
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}

type Hello struct{}

func (h *Hello) Say(name string) (string, error) {
	return fmt.Sprintf("hello %s", name), nil
}
