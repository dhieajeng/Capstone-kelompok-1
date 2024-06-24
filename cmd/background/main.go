package main

import (
	"github.com/bloomingbug/depublic/configs"
	"github.com/bloomingbug/depublic/pkg/cache"
	"github.com/bloomingbug/depublic/pkg/processors"
)

func main() {
	cfg, err := configs.NewConfig(".env")
	checkError(err)
	redis := cache.InitCache(&cfg.Redis)

	process := processors.NewProcess(redis, *cfg)
	process.RunProcess()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
