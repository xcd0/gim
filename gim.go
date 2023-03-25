package gim

import (
	"github.com/rivo/tview"
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

func createBox(title string) {
	box := tview.NewBox().SetBorder(true).SetTitle(title)
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
