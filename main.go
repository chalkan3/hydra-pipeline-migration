package main

import (
	IoC "fastshop.com.br/create_pipelines/infra/provider"
)

type mainvars struct {
	provider *IoC.Provider
}

var vars *mainvars

func init() {
	vars = &mainvars{
		provider: IoC.NewIocProvider(),
	}
}

func main() {
	err := vars.provider.Run()

	if err != nil {
		panic(err)
	}
}
