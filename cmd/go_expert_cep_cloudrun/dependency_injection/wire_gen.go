// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependency_injection

import (
	"context"
	"github.com/aluferraz/go-expert-cep-cloudrun/internal/infra/web/webhandlers/get_temperature_handler"
	"github.com/aluferraz/go-expert-cep-cloudrun/internal/usecase/get_temperature"
	"net/http"
)

// Injectors from wire.go:

func NewTemperatureHandler(ctx *context.Context) *get_temperature_handler.WebGetTemperatureHandler {
	useCase := NewTemperatureUseCase(ctx)
	webGetTemperatureHandler := get_temperature_handler.NewGetTemperatureHandler(useCase)
	return webGetTemperatureHandler
}

// wire.go:

func NewTemperatureUseCase(ctx *context.Context) get_temperature.UseCase {
	return get_temperature.NewUseCase(ctx, &http.Client{})
}
