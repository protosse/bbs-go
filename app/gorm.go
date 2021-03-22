package app

import (
	"bbs-go/common/config"
	"bbs-go/models"
	"bbs-go/services"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GormServer struct {
	Config *config.Config
}

func NewGormServer(config *config.Config) *GormServer {
	server := &GormServer{
		Config: config,
	}
	return server
}

func (s *GormServer) Connect() (err error) {
	gormConf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   s.Config.DB.Prefix,
			SingularTable: true,
		},
	}

	return services.OpenDB(s.Config.DB.Conn, gormConf, 10, 20, models.Models...)
}
