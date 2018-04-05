package main

import (
	"fmt"

	"github.com/blendlabs/go-logger"
	"github.com/blendlabs/go-web"
)

func main() {
	log := logger.All()
	cfg := web.NewConfigFromEnv()

	app := web.NewFromConfig(cfg).WithLogger(log)
	app.GET("/", func(r *web.Ctx) web.Result {
		return r.Text().Result("foo")
	})

	log.Listen(logger.WebRequest, logger.DefaultListenerName, logger.NewWebRequestEventListener(func(wre *logger.WebRequestEvent) {
		fmt.Printf("Route: %s\n", wre.Route())
	}))

	log.SyncFatalExit(app.Start())
}
