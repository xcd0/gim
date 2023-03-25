package gim

// プラグインはpluginディレクトリに1つにつき1ファイルで<pluginの名前>.goを実装し、
// <pluginの名前>.soをビルドするスクリプトをplugin/makefileに記述する。

var (
	plugins = [][]string{
		[]string{"plug/plugin.so", "F"},
	}
)

type PluginList struct {
	LoadedPlugin      Plugins                            // 正しく読み込まれたプラグイン一覧
	PluginFunctionMap map[string]map[string]func() error // プラグイン名(.soファイルの名前)→関数名
}
type Plugins []Plugin
type Plugin struct {
	ID   int
	Name string
	// TODO: プラグインの機能を呼ぶ関数を呼べるようにこの辺にいい感じで何か追加する
	Functions map[string]func() error
}

func LoadPlugin() {
	pluginList := PluginList{}

	for _, p := range plugins {
		pluginList.Load(p[0], p[1:]...)
	}

	// 登録した関数を呼び出す例
}

func (plugins *PluginList) Load(pluginSoPath string, functionName ...string) error {
	/*
		// ビルド済みプラグインをロード
		p, err := plugin.Open(pluginSoPath)
		if err != nil {
			return err
		}
		// ロードできたらそのプラグイン専用のPlugin構造体を作成してLoadedPluginに追加
		plugins.LoadedPlugin = append(plugins.LoadedPlugin, Plugin{})
			for _, f := range functionName {
				if err := func(plugins *PluginList, p *plugin.Plugin, funcname string) error {
					f, err := p.Lookup(funcname)
					if err != nil {
						return err
					}
					// プラグインで実装している関数を構造体に登録する
					plugins.LoadedPlugin.back().Functions[funcname] = f.(func())
					return nil
				}(plugins, p, f); err != nil {
					return err
				}
			}
	*/
	return nil
}

func (p *Plugins) back() *Plugin { return &(*p)[len(*p)-1] }

func (l *PluginList) Call(pluginName string, funcName string) error {
	return l.LoadedPlugin.back().Functions["F"]() // prints "Hello, world"
}
