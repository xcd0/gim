# gim

vim っぽいテキストエディタを実装してみる。

## ふわっとした展望
* コマンドラインの処理は[urfave/cli](https://github.com/urfave/cli)で片付ける。便利。
* CUIの見た目とかは[tcel](https://github.com/gdamore/tcell)が良さそう。
	* 直で使わないでwrapper使った方が良さそう。
* [goreleaser](https://github.com/goreleaser/goreleaser)でリリース作業を簡単にしたい。yamlの設定ファイル書けば終わりでお手軽。
* できるだけTDDしたい。
* githubのciを使いたい。
* プラグインとか使えるようにしたい。何に使うかは不明だけれども。
	* vimのプラグインがそのまま動いたらいいなあ。
		* [vim-jp/go-vimlparser](https://github.com/vim-jp/go-vimlparser)とかあるのでなんとかなったり？
	* neovimみたいにluaとかでプラグインかけるようにしたい。
	* 大きめの機能拡張には [urfave/cli](https://github.com/hashicorp/go-plugin) が良さそう。

## TDD


