package handler

import (
	cp "auth/genproto/mobile"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// ServiceCreate handles the creation of a new service.
// @Summary Create service
// @Description Create a new service
// @Tags fine
// @Accept json
// @Produce json
// @Param car body cp.FineCreateReq true "Fine data"
// @Success 200 {object} string "Fine created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/fine/add [post]
func (h *Handler) FineCreate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role != "yhxb" && role != "superadmin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	var req cp.FineCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.StafId = id

	_, err := h.srvs.Fine.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fine created"})
}

// ServiceGetById handles the get a service.
// @Summary Get service
// @Description Get a service
// @Tags fine
// @Accept json
// @Produce json
// @Param id query string false "Fine ID"
// @Success 200 {object} cp.FineGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/fine/{id} [get]
func (h *Handler) FineGetById(c *gin.Context) {
	by_id := cp.ById{Id: c.Query("id")}
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role == "service" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	res, err := h.srvs.Fine.GetById(context.Background(), &by_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get fine", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ServiceGetAll handles getting all service.
// @Summary Get all service
// @Description Get all service
// @Tags fine
// @Accept json
// @Produce json
// @Param page query integer false "Page"
// @Param car_number query string false "CarNumber"
// @Param staf_id query string false "StafId"
// @Param status query string false "Status"
// @Success 200 {object} cp.FineGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/fine/all [get]
func (h *Handler) FineGetAll(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role == "service" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	req := cp.FineGetAllReq{
		Filter: &cp.Filter{},
	}

	car_number := c.Query("car_number")
	status := c.Query("status")
	staf_id := c.Query("staf_id")

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
	req.StafId = staf_id
	req.Status = status

	if role == "driver" {
		car_number := claims.(jwt.MapClaims)["car_number"].(string)
		req.CarNumber = car_number
		req.StafId = ""
	}

	res, err := h.srvs.Fine.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get fines", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ServiceUpdate handles updating an existing service.
// @Summary Update service
// @Description Update an existing service
// @Tags fine
// @Accept json
// @Produce json
// @Param id query string false "Id"
// @Success 200 {object} string "Fine updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Service not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/fine/update/{id} [put]
func (h *Handler) FineUpdate(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)

	if role == "service" {
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	req := cp.FineUpdateReq{
		Id:          c.Query("id"),
	}

	if role == "driver" {
		req.Status = "pending"
	} else {
		req.Status = "paid"
	}

	_, err := h.srvs.Fine.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update fine", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Fine updated"})
}
