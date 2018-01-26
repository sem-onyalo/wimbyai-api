package main

import (
	"github.com/sem-onyalo/wimbyai-api/core/interactor"
	"github.com/sem-onyalo/wimbyai-api/service/request"
	"github.com/sem-onyalo/wimbyai-api/web"
)

func main() {
	configService := interactor.NewConfig()
	apiService := web.NewAPI(configService)
	apiService.Start(request.StartAPI{})
}
