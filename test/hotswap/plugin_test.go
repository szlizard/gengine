package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/edwingeng/hotswap"
	"github.com/edwingeng/slog"
	"github.com/szlizard/gengine/builder"
	"github.com/szlizard/gengine/context"
	"github.com/szlizard/gengine/engine"
)

var logger = slog.NewDevelopmentConfig().MustBuild()
var swapper = hotswap.NewPluginManagerSwapper("/Users/nickolast/code/yiguoji/gengine/test/hotswap", hotswap.WithLogger(logger))

func Test_hotswap2_plugin(t *testing.T) {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }

	// var logger = slog.NewDevelopmentConfig().MustBuild()
	// swapper := hotswap.NewPluginManagerSwapper(dir, hotswap.WithLogger(logger))
	pluginDetails, err := swapper.LoadPlugins(nil)
	if err != nil {
		panic(fmt.Sprintf("failed to load plugin: %v", err))
	}
	fmt.Printf("Plugin loaded with details: %v\n", pluginDetails)

	swapper.Current().FindPlugin("hotswap").InvokeFunc("SaveLive", 1)
	// swapper.Current().InvokeEach("SaveLive", 1)
}

func Test_hotswap_plugin_with_gengine(t *testing.T) {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }

	// var logger = slog.NewDevelopmentConfig().MustBuild()
	// swapper := hotswap.NewPluginManagerSwapper(dir, hotswap.WithLogger(logger))
	dc := context.NewDataContext()
	pluginDetails, err := swapper.LoadPlugins(nil)
	if err != nil {
		panic(fmt.Sprintf("failed to load plugin: %v", err))
	}
	fmt.Printf("Plugin loaded with details: %v\n", pluginDetails)

	// currentPlugins := swapper.Current().Plugins()
	// // currentPlugins.InvokeFunc("SaveLive", 1)
	// for _, p := range currentPlugins {
	// 	fmt.Println("plugin name: ", p.Name)
	// 	dc.Add("m", swapper.Current().FindPlugin("hotswap").InvokeFunc)
	// }
	dc.Add("m", swapper.Current().FindPlugin("hotswap").InvokeFunc)

	dc.Add("println", fmt.Println)
	ruleBuilder := builder.NewRuleBuilder(dc)
	err = ruleBuilder.BuildRuleFromString(`
	rule "1"
	begin
	//this method is defined in plugin
	err = m("SaveLive", 2)

	if isNil(err) {
	   println("err is nil")
	}
	end
	`)

	if err != nil {
		panic(err)
	}
	gengine := engine.NewGengine()
	err = gengine.Execute(ruleBuilder, false)

	if err != nil {
		panic(err)
	}
}

func Test_hotswap_plugin_with_pool(t *testing.T) {
	rule := `
	rule "1"
	begin
	//this method is defined in plugin
	err = m("SaveLive", 3)

	if isNil(err) {
	   println("err is nil")
	}
	end`

	apis := make(map[string]interface{})
	apis["println"] = fmt.Println

	// dir, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }

	// var logger = slog.NewDevelopmentConfig().MustBuild()
	// swapper := hotswap.NewPluginManagerSwapper(dir, hotswap.WithLogger(logger))
	pluginDetails, err := swapper.LoadPlugins(nil)
	if err != nil {
		panic(fmt.Sprintf("failed to load plugin: %v", err))
	}
	fmt.Printf("Plugin loaded with details: %v\n", pluginDetails)

	// currentPlugins := swapper.Current().Plugins()
	// // currentPlugins.InvokeFunc("SaveLive", 1)
	// for _, p := range currentPlugins {
	// 	fmt.Println("plugin name: ", p.Name)
	// 	apis["m"] = p.InvokeFunc
	// }
	apis["m"] = swapper.Current().FindPlugin("hotswap").InvokeFunc

	pool, e := engine.NewGenginePool(1, 2, 1, rule, apis)
	if e != nil {
		panic(e)
	}

	data := make(map[string]interface{})
	e, _ = pool.Execute(data, true)
	if e != nil {
		panic(e)
	}

	//twice execute
	e, _ = pool.Execute(data, true)
	if e != nil {
		panic(e)
	}
}

func Test_hotswap_file(t *testing.T) {
	files := "plugin_superman.so"
	dir, file := filepath.Split(files)
	println(dir, file, filepath.Base(files), path.Ext(files))
	s, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	println(s)
}
