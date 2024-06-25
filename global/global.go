package global

import (
	"go-ecommerce/pkg/logger"
	"go-ecommerce/pkg/setting"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
