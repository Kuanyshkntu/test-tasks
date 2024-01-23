package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test_tasks/service"
)

type Controller struct {
	Svc service.Service
}

func NewHandler(svc service.Service) *Controller {
	return &Controller{
		Svc: svc,
	}
}

func (h *Controller) GetRouter() (*gin.Engine, error) {
	router := gin.New()
	if e := router.SetTrustedProxies(nil); e != nil {
		log.Fatal(e)
	}
	router.Use(gin.Logger())
	router.Use(
		gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, err any) {
			c.JSON(http.StatusInternalServerError, map[string]any{"message": fmt.Sprint(err)})
			c.Abort()
		}),
	)

	router.GET("/test-tasks/people", h.GetPeople)
	router.POST("/test-tasks/people", h.AddPerson)
	router.PUT("/test-tasks/people/:id", h.UpdatePerson)
	router.DELETE("/test-tasks/people/:id", h.DeletePersonByID)

	return router, nil
}
