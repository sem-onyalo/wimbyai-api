package service

import (
	"github.com/sem-onyalo/wimbyai-api/service/request"
	"github.com/sem-onyalo/wimbyai-api/service/response"
)

// Config is a boundary to application configuration operations
type Config interface {
	GetValue(request request.GetConfigValue) response.GetConfigValue
}
