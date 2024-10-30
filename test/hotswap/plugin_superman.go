package main

import (
	"strconv"

	"github.com/edwingeng/hotswap/vault"
)

const (
	pluginName = "superman"
)

var (
	CompileTimeString string
)

func OnLoad(data interface{}) error {
	// g.Logger.Infof("<%s.%s> OnLoad", pluginName, CompileTimeString)
	return nil
}

func OnInit(sharedVault *vault.Vault) error {
	// g.Logger.Infof("<%s.%s> OnInit", pluginName, CompileTimeString)
	return nil
}

func OnFree() {
	// g.Logger.Infof("<%s.%s> OnFree", pluginName, CompileTimeString)
}

func Export() interface{} {
	// g.Logger.Infof("<%s.%s> Export", pluginName, CompileTimeString)
	return nil
}

func Import() interface{} {
	// g.Logger.Infof("<%s.%s> Import", pluginName, CompileTimeString)
	return nil
}

func InvokeFunc(name string, params ...interface{}) (interface{}, error) {
	switch name {
	case "SaveLive":
		var repeat int = 0
		if params[0] != nil {
			switch params[0].(type) {
			case int:
				repeat = params[0].(int)
			case int64:
				repeat = int(params[0].(int64))
			}
		}
		println("execute finished... repeat:" + strconv.Itoa(repeat))
	}
	return nil, nil
}

func Reloadable() bool {
	// g.Logger.Infof("<%s.%s> Reloadable", pluginName, CompileTimeString)
	return true
}
