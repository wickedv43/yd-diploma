package main

import (
	"os"
	"syscall"

	"github.com/samber/do/v2"
	"github.com/wickedv43/yd-diploma/internal/config"
	"github.com/wickedv43/yd-diploma/internal/logger"
	"github.com/wickedv43/yd-diploma/internal/server"
	"github.com/wickedv43/yd-diploma/internal/storage"
)

func main() {
	// provide part
	i := do.New()

	do.Provide(i, server.NewServer)
	do.Provide(i, config.NewConfig)
	do.Provide(i, logger.NewLogger)

	//storage
	do.Provide(i, storage.NewPostgresStorage)

	do.MustInvoke[*logger.Logger](i)
	do.MustInvoke[*server.Server](i).Start()

	i.ShutdownOnSignals(syscall.SIGTERM, os.Interrupt)
}
