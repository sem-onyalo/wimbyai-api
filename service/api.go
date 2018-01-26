package service

import (
	"github.com/sem-onyalo/wimbyai-api/service/request"
)

// API is a boundary to the web api
type API interface {
	Start(request request.StartAPI)
}
