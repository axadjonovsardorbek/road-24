package handler

import (
	"context"
	cp "auth/genproto/mobile"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// ServiceCreate handles the creation of a new service.
// @Summary Create service
// @Description Create a new service
// @Tags car
// @Accept json
// @Produce json
// @Param car body cp.CarCreateReq true "Car data"
// @Success 200 {object} string "Car created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/car/add [post]
func (h *Handler) CarCreate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)
	is_admin := claims.(jwt.MapClaims)["is_admin"].(string)

	if role == "driver" || role == "service" || is_admin == "false" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	var req cp.CarCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.StafId = id

	_, err := h.srvs.Car.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car created"})
}

// ServiceGetById handles the get a service.
// @Summary Get service
// @Description Get a service
// @Tags car
// @Accept json
// @Produce json
// @Param id query string true "Car ID"
// @Success 200 {object} cp.CarGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/car/{id} [get]
func (h *Handler) CarGetById(c *gin.Context) {
	by_id := cp.ById{}
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)


	if role == "driver"{
		id := claims.(jwt.MapClaims)["id"].(string)
		by_id.Id = id
	} else if role == "service" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	} else {
		by_id.Id = c.Query("id")
	}

	res, err := h.srvs.Car.GetById(context.Background(), &by_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get car", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ServiceGetAll handles getting all service.
// @Summary Get all service
// @Description Get all service
// @Tags car
// @Accept json
// @Produce json
// @Param page query integer false "Page"
// @Param model query string false "Model"
// @Param type query string false "Type"
// @Param year query integer false "Year"
// @Success 200 {object} cp.CarGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/car/all [get]
func (h *Handler) CarGetAll(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role == "driver" || role == "service" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	req := cp.CarGetAllReq{
		Filter: &cp.Filter{},
	}

	model := c.Query("model")
	typee := c.Query("type")
	yearStr := c.Query("year")
	var year int
	var err error
	if yearStr == "" {
		year = 0
	} else {
		year, err = strconv.Atoi(yearStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year parameter"})
			return
		}
	}

	pageStr := c.Query("page")
	var page int
	if pageStr == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
			return
		}
	}

	filter := cp.Filter{
		Page: int32(page),
	}

	req.Filter.Page = filter.Page
	req.Year = int32(year)
	req.Model = model
	req.Type = typee

	res, err := h.srvs.Car.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get cars", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ServiceUpdate handles updating an existing service.
// @Summary Update service
// @Description Update an existing service
// @Tags car
// @Accept json
// @Produce json
// @Param url query string false "Url"
// @Success 200 {object} string "Car updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Service not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /car/upload/image [put]
func (h *Handler) CarUpdate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role != "driver" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	number := claims.(jwt.MapClaims)["car_number"].(string)
	technical_passport := claims.(jwt.MapClaims)["technical_passport"].(string)

	req := cp.CarUpdateReq{
		ImageUrl:      c.Query("url"),
		CarNumber: number,
		TechnicalPassport: technical_passport,
	}

	_, err := h.srvs.Car.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update car", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car updated"})
}

// ServiceDelete handles deleting a service by ID.
// @Summary Delete service
// @Description Delete a service by ID
// @Tags car
// @Accept json
// @Produce json
// @Param id query string true "Car ID"
// @Success 200 {object} string "Car deleted"
// @Failure 400 {object} string "Invalid provider ID"
// @Failure 404 {object} string "Car not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/car/{id} [delete]
func (h *Handler) CarDelete(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)
	is_admin := claims.(jwt.MapClaims)["is_admin"].(string)

	if role == "driver" || role == "service" || is_admin == "false"{
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}
	
	id := &cp.ById{Id: c.Param("id")}

	_, err := h.srvs.Car.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete car", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted"})
}
