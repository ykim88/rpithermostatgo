package api

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	bind(*gin.Engine)
}

type Api interface {
	Up()
}

func New(controllers ...Controller) Api {
	router := gin.Default()

	for _, c := range controllers {
		c.bind(router)
	}

	return &api{router: router, controllers: controllers}
}

type api struct {
	router      *gin.Engine
	controllers []Controller
}

func (a *api) Up() {
	a.router.Run()
}
