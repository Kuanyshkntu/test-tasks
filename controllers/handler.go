package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"test_tasks/models"
)

func (c *Controller) GetPeople(reqCtx *gin.Context) {

	query := reqCtx.Request.URL.Query()

	params := map[string]string{
		"code":       query.Get("code"),
		"name":       query.Get("name"),
		"task_group": query.Get("task_group"),
		"is_active":  query.Get("is_active"),
		"limit":      query.Get("limit"),
		"page":       query.Get("page"),
	}

	res, err := c.Svc.GetPeople(params)
	if err != nil {
		log.Fatal(err)
		reqCtx.JSON(http.StatusInternalServerError, models.MessageType{
			Message:     "Ошибка при отправке",
			Description: err.Error(),
		})
		return
	}

	reqCtx.JSON(http.StatusOK, res)

	return
}

func (c *Controller) AddPerson(reqCtx *gin.Context) {

	req := new(models.Person)
	err := reqCtx.BindJSON(req)
	if err != nil {
		log.Fatal(err)
		reqCtx.AbortWithStatusJSON(http.StatusBadRequest, models.MessageType{Message: err.Error()})
		return
	}

	err = c.Svc.AddPerson(req)
	if err != nil {
		log.Fatal(err)
		reqCtx.JSON(http.StatusBadRequest, models.MessageType{Message: err.Error()})
		return
	}

	reqCtx.JSON(http.StatusCreated, nil)

	return
}

func (c *Controller) UpdatePerson(reqCtx *gin.Context) {

	var req models.Person
	err := reqCtx.ShouldBindJSON(&req)
	if err != nil {
		log.Fatal(err)
		reqCtx.AbortWithStatusJSON(http.StatusBadRequest, models.MessageType{Message: err.Error()})
		return
	}

	req.ID, err = strconv.Atoi(reqCtx.Param("id"))
	if err != nil {
		log.Fatal(err)
		reqCtx.AbortWithStatusJSON(http.StatusBadRequest, models.MessageType{Message: err.Error()})
		return
	}

	err = c.Svc.UpdatePerson(req)
	if err != nil {
		log.Fatal(err)
		reqCtx.JSON(http.StatusBadRequest, models.MessageType{Message: err.Error()})
		return
	}

	reqCtx.JSON(http.StatusOK, nil)

	return
}

func (c *Controller) DeletePersonByID(reqCtx *gin.Context) {

	id := reqCtx.Param("id")

	err := c.Svc.DeletePersonByID(id)
	if err != nil {
		log.Fatal(err)
		reqCtx.JSON(http.StatusBadRequest, models.MessageType{Message: err.Error()})
		return
	}

	reqCtx.JSON(http.StatusOK, nil)

	return
}
