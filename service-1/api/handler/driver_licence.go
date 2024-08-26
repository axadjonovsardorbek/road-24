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
// @Tags licence
// @Accept json
// @Produce json
// @Param licence body cp.DriverLicenceCreateReq true "DriverLicence data"
// @Success 200 {object} string "DriverLicence created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/driver/licence/add [post]
func (h *Handler) LicenceCreate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)
	is_admin := claims.(jwt.MapClaims)["is_admin"].(string)

	if role != "yidxp" && role != "superadmin" && is_admin == "false" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	var req cp.DriverLicenceCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err := h.srvs.DriverLicence.Create(context.Background(), &req)

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
// @Tags licence
// @Accept json
// @Produce json
// @Param id query string true "DriverLicence ID"
// @Success 200 {object} cp.DriverLicenceGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/driver/licence/{id} [get]
func (h *Handler) LicenceGetById(c *gin.Context) {
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

	by_id := cp.ById{}

	by_id.Id = c.Query("id")

	res, err := h.srvs.DriverLicence.GetById(context.Background(), &by_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get licence", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ServiceGetAll handles getting all service.
// @Summary Get all service
// @Description Get all service
// @Tags licence
// @Accept json
// @Produce json
// @Param page query integer false "Page"
// @Param category query string false "Category"
// @Param exp_at query string false "ExpAt"
// @Success 200 {object} cp.DriverLicenceGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/driver/licence/all [get]
func (h *Handler) LicenceGetAll(c *gin.Context) {
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

	req := cp.DriverLicenceGetAllReq{
		Filter: &cp.Filter{},
	}

	category := c.Query("category")
	exp_at := c.Query("exp_at")

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
	req.Category = category
	req.ExpAt = exp_at

	res, err := h.srvs.DriverLicence.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get licences", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ServiceDelete handles deleting a service by ID.
// @Summary Delete service
// @Description Delete a service by ID
// @Tags licence
// @Accept json
// @Produce json
// @Param id query string true "DriverLicence ID"
// @Success 200 {object} string "DriverLicence deleted"
// @Failure 400 {object} string "Invalid provider ID"
// @Failure 404 {object} string "DriverLicence not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/driver/licence/{id} [delete]
func (h *Handler) LicenceDelete(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)
	is_admin := claims.(jwt.MapClaims)["is_admin"].(string)

	if role != "yidxp" && role != "superadmin" && is_admin == "false" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}
	
	id := &cp.ById{Id: c.Param("id")}

	_, err := h.srvs.DriverLicence.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete car", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "DriverLicence deleted"})
}
