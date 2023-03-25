# gim

vim っぽいテキストエディタを実装してみる。


## ふわふわ仕様

* 位置付け
	* とりあえずテキストエディタを作ってみたいので作っているに過ぎない趣味アプリケーション。
	* ある程度使えるようになったら自分が使う分にはプラグインを導入しなくても使えるシングルバイナリなvimとして使いたい。

* 実装予定機能
	* vimっぽい動作
		* モード
			* Normal
			* Insert
			* Visual
			* Terminal

	* メジャーなvimプラグインと同等の機能の組み込み
		* VimFiler
		* PreVim
		* EasyAlign

	* 最初に実装する機能  
		とりあえず日常使用に困らない機能を実装

		* [ ] テキスト編集画面実装
		* [ ] 新規テキストファイル作成
		* [ ] 既存テキストファイル編集
		* [ ] テキストファイル保存
		* [ ] 文字入力
		* [ ] 文字削除
		* [ ] カーソル移動
		* [ ] 画面スクロール機能
		* [ ] コピペ機能
		* [ ] undo
		* [ ] redo

	* とりあえず実装する予定機能
		* [ ] 画面分割
		* [ ] タブ
		* [ ] 自動フォーマット
		* [ ] 自動インデント
		* [ ] 検索
		* [ ] 置換
		* [ ] 補完
		* [ ] マクロ
		* [ ] リマップ
		* [ ] 文字コード変換
		* [ ] 制御文字表示
		* [ ] 設定ファイル

	* いつか実装したい機能
		* [ ] プラグイン機能
		* [ ] バイナリ編集
		* [ ] tmuxのような端末多重化機能の実装
		* [ ] gitに都合が良い機能の実装
		* [ ] githubに都合が良い機能の実装

## 実装方針

上の実装予定機能表から作業する機能を一つ取り上げTDDで開発する。  
最初はテキストエディタとして最低限必要な機能を実装する。
[TDDについては別に書いた。](./tdd.md)  
TODOリストはせっかくなのでgithubの機能を使用する。  
https://github.com/users/xcd0/projects/2/views/1  
TDDの1サイクルをissue1つと関連させる。

## ブランチ運用方針
安定版を保持するmaster、開発版を保持するdevelp、機能拡張を行うfeature-xxxの３種類のブランチで運用する。  
featureブランチはissueごとに作成する。  
featureブランチの名前はissueの番号を後ろに付与する。

## LICENSE

vimに肖ってGPL3にしています。(vimはGPL互換のVIM LICENSE)  
ウガンダの恵まれない子供たちへの寄付については以下のページを参照してください。  
https://github.com/vim/vim/blob/master/runtime/doc/uganda.txt
当然ですが、使用させていただいている各ライブラリは各ライブラリのライセンスです。

## ふわっとした展望

* とりあえず達成
	* [goreleaser](https://github.com/goreleaser/goreleaser)でリリース
	* github actionの使用
* これからする
	* できるだけTDDしたい。
	* コマンドラインの処理は[urfave/cli](https://github.com/urfave/cli)で片付ける。便利。
	* CUIの見た目とかは[tcel](https://github.com/gdamore/tcell)が良さそう。
		* 直で使わないでwrapper使った方が良さそう。
		* [rivo/tview](https://github.com/rivo/tview)を使う。
* この先
	* プラグインとか使えるようにしたい。何に使うかは不明だけれども。
		* vimのプラグインがそのまま動いたらいいなあ。
			* [vim-jp/go-vimlparser](https://github.com/vim-jp/go-vimlparser)とかあるのでなんとかなったり？
		* neovimみたいにluaとかでプラグインかけるようにしたい。
		* 大きめの機能拡張には [urfave/cli](https://github.com/hashicorp/go-plugin) が良さそう。