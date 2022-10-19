package collie

import (
	"github.com/shinjuwu/collie/cluster"
	"github.com/shinjuwu/collie/conf"
	"github.com/shinjuwu/collie/console"
	"github.com/shinjuwu/collie/log"
	"github.com/shinjuwu/collie/module"
	"os"
	"os/signal"
)

func Run(mods ...module.Module) {
	// logger
	if conf.LogLevel != "" {
		logger, err := log.New("")
		if err != nil {
			panic(err)
		}
		log.Export(logger)
		defer logger.Close()
	}

	log.Info("github.com/shinjuwu/collie starting up")

	// module
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}
	module.Init()

	// cluster
	cluster.Init()

	// console
	console.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Info("github.com/shinjuwu/collie closing down (signal: %v)", sig)
	console.Destroy()
	cluster.Destroy()
	module.Destroy()
}
