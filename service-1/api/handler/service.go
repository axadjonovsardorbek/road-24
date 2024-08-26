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
// @Tags service
// @Accept json
// @Produce json
// @Param car body cp.ServiceCreateReq true "Service data"
// @Success 200 {object} string "Service created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/service/add [post]
func (h *Handler) ServiceCreate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role != "superadmin" && role != "service" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	var req cp.ServiceCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err := h.srvs.Service.Create(context.Background(), &req)

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
// @Tags service
// @Accept json
// @Produce json
// @Param id query string false "Service ID"
// @Success 200 {object} cp.ServiceGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/service/{id} [get]
func (h *Handler) ServiceGetById(c *gin.Context) {
	by_id := cp.ById{}

	by_id.Id = c.Query("id")

	res, err := h.srvs.Service.GetById(context.Background(), &by_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get service", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ServiceGetAll handles getting all service.
// @Summary Get all service
// @Description Get all service
// @Tags service
// @Accept json
// @Produce json
// @Param page query integer false "Page"
// @Param address query string false "Address"
// @Success 200 {object} cp.ServiceGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/service/all [get]
func (h *Handler) ServiceGetAll(c *gin.Context) {
	req := cp.ServiceGetAllReq{
		Filter: &cp.Filter{},
		Address: c.Query("address"),
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

	res, err := h.srvs.Service.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get services", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ServiceDelete handles deleting a service by ID.
// @Summary Delete service
// @Description Delete a service by ID
// @Tags service
// @Accept json
// @Produce json
// @Param id query string true "Service ID"
// @Success 200 {object} string "Service deleted"
// @Failure 400 {object} string "Invalid provider ID"
// @Failure 404 {object} string "Service not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/service/{id} [delete]
func (h *Handler) ServiceDelete(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role != "superadmin" && role != "service" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}
	
	id := &cp.ById{Id: c.Query("id")}

	_, err := h.srvs.Service.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete service", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted"})
}
