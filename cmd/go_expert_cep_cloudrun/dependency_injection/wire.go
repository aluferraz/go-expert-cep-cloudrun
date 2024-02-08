//go:build wireinject
// +build wireinject

package dependency_injection

import (
	"context"
	"github.com/aluferraz/go-expert-cep-cloudrun/internal/infra/web/webhandlers/get_temperature_handler"
	"github.com/aluferraz/go-expert-cep-cloudrun/internal/usecase/get_temperature"
	"github.com/google/wire"
	"net/http"
)

func NewTemperatureUseCase(ctx *context.Context) get_temperature.UseCase {
	return get_temperature.NewUseCase(ctx, &http.Client{})
}
func NewTemperatureHandler(ctx *context.Context) *get_temperature_handler.WebGetTemperatureHandler {
	wire.Build(NewTemperatureUseCase, get_temperature_handler.NewGetTemperatureHandler)
	return &get_temperature_handler.WebGetTemperatureHandler{}
}

/*
var setSampleRepositoryDependency = wire.NewSet(
	database.SampleRepository,
	wire.Bind(new(entity.SampleRepositoryInterface), new(*database.SampleRepository)),
)

func NewListAllOrdersUseCase(db *sql.DB) *usecase.MyUseCase {
	wire.Build(
		setSampleRepositoryDependency,
		usecase.NewUseCase,
	)
	return &usecase.MyUseCase{}
}
*/
