package main

import (
	"flag"
	"log"

	"github.com/slowhigh/Umm/config"
	"github.com/slowhigh/Umm/internal/app"
	"github.com/slowhigh/Umm/pkg/logger"
)

func main() {
	log.Println("Starting service")

	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.Named(cfg.ServiceName)
	appLogger.Infof("CFG: %+v", cfg)
	appLogger.Fatal(app.NewApp(appLogger, *cfg).Run())
}
