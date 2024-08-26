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
// @Param car body cp.CarServiceCreateReq true "Car data"
// @Success 200 {object} string "Car created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/car/service/add [post]
func (h *Handler) CarServiceCreate(c *gin.Context) {
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

	var req cp.CarServiceCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err := h.srvs.CarService.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service created"})
}

// ServiceGetById handles the get a service.
// @Summary Get service
// @Description Get a service
// @Tags car
// @Accept json
// @Produce json
// @Param id query string true "Car ID"
// @Success 200 {object} cp.CarServiceGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/car/service/{id} [get]
func (h *Handler) CarServiceGetById(c *gin.Context) {
	by_id := cp.ById{}

	by_id.Id = c.Query("id")

	res, err := h.srvs.CarService.GetById(context.Background(), &by_id)
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
// @Param car_number query string false "CarNumber"
// @Param service_type query string false "ServiceType"
// @Success 200 {object} cp.CarServiceGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/car/service/all [get]
func (h *Handler) CarServiceGetAll(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	req := cp.CarServiceGetAllReq{
		Filter: &cp.Filter{},
	}

	car_number := c.Query("car_number")
	service_type := c.Query("service_type")

	pageStr := c.Query("page")
	var page int
	var err error
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
	req.CarNumber = car_number
	req.ServiceType = service_type

	if role == "driver"{
		number := claims.(jwt.MapClaims)["car_number"].(string)
		req.CarNumber = number
	}

	res, err := h.srvs.CarService.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get cars", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
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
// @Router /admin/car/service/{id} [delete]
func (h *Handler) CarServiceDelete(c *gin.Context) {
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
	
	id := &cp.ById{Id: c.Query("id")}

	_, err := h.srvs.CarService.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete car", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted"})
}
