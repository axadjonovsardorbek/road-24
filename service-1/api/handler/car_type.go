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
// @Param car body cp.CarTypeCreateReq true "Car data"
// @Success 200 {object} string "Car created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/car/type/add [post]
func (h *Handler) CarTypeCreate(c *gin.Context) {
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

	var req cp.CarTypeCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err := h.srvs.CarType.Create(context.Background(), &req)

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
// @Param id query string false "Car ID"
// @Success 200 {object} cp.CarTypeGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/car/type/{id} [get]
func (h *Handler) CarTypeGetById(c *gin.Context) {
	by_id := cp.ById{}

	by_id.Id = c.Query("id")

	res, err := h.srvs.CarType.GetById(context.Background(), &by_id)
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
// @Success 200 {object} cp.CarTypeGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/car/type/all [get]
func (h *Handler) CarTypeGetAll(c *gin.Context) {
	req := cp.CarTypeGetAllReq{
		Filter: &cp.Filter{},
	}

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

	res, err := h.srvs.CarType.GetAll(context.Background(), &req)

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
// @Param id query string false "Car ID"
// @Success 200 {object} string "Car deleted"
// @Failure 400 {object} string "Invalid provider ID"
// @Failure 404 {object} string "Car not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/car/type/{id} [delete]
func (h *Handler) CarTypeDelete(c *gin.Context) {
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

	_, err := h.srvs.CarType.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete car", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted"})
}
