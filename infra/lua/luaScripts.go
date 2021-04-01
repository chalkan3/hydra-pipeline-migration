package luascript

import (
	"net/http"

	scriptsfunctions "fastshop.com.br/create_pipelines/infra/lua/scriptsFunctions"

	"github.com/cjoudrey/gluahttp"
	"github.com/kohkimakimoto/gluayaml"
	lua "github.com/yuin/gopher-lua"
)

// GoLua is a script template
type GoLua struct {
	LuaState            *lua.LState
	LuaFunction         map[string]lua.LGFunction
	migrateLuaFunctions *scriptsfunctions.MigrationFunction
}

// RunLua MainLua
func (golua *GoLua) RunLua(folder string) {

	defer golua.LuaState.Close()
	golua.AppendFunction("migrate", golua.migrateLuaFunctions.Migrate)

	golua.LuaState.PreloadModule("yaml", gluayaml.Loader)
	golua.LuaState.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)
	golua.LuaState.PreloadModule("golangFunctions", golua.Loader)

	if err := golua.LuaState.DoFile("./scripts/migrations/" + folder + "/main.lua"); err != nil {
		panic(err)
	}

}

// AppendFunction is a append function
func (golua *GoLua) AppendFunction(functionName string, newFunction lua.LGFunction) {
	golua.LuaFunction[functionName] = newFunction
}

// Loader is a loader of lua function
func (golua *GoLua) Loader(L *lua.LState) int {
	tb := L.NewTable()
	L.SetFuncs(tb, golua.LuaFunction)
	L.Push(tb)
	return 1
}

// NewLGoLua Ioc
func NewLGoLua() *GoLua {
	return &GoLua{
		LuaState:            lua.NewState(),
		LuaFunction:         map[string]lua.LGFunction{},
		migrateLuaFunctions: scriptsfunctions.NewMigrationFunction(),
	}
}
