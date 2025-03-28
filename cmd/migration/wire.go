//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-nunu/nunu-layout-advanced/internal/repository"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)

func newApp(*viper.Viper, *log.Logger) (*Migrate, func(), error) {
	panic(wire.Build(
		RepositorySet,
		NewMigrate,
	))
}
