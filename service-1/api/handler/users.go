package handler

import (
	ap "auth/genproto/auth"
	"context"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	_ "github.com/swaggo/swag"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register a new admin
// @Description Register a new admin
// @Tags admin
// @Accept json
// @Produce json
// @Param user body ap.UserCreateReq true "User registration request"
// @Success 201 {object} string "User registered"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/user/add [post]
func (h *Handler) Register(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)
	is_admin := claims.(jwt.MapClaims)["is_admin"].(string)

	if role != "superadmin" && is_admin != "true"{
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	var req ap.UserCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if role != "superadmin"{
		req.Role = role
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	req.Password = string(hashedPassword)

	_, err = h.User.Register(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

// Login godoc
// @Summary Login a user
// @Description Authenticate user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body ap.UserLoginReq true "User login credentials"
// @Success 200 {object} ap.TokenRes "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Invalid username or password"
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	var req ap.UserLoginReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.Login(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Login godoc
// @Summary Login a user
// @Description Authenticate user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body ap.DriverLoginReq true "User login credentials"
// @Success 200 {object} ap.TokenRes "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Invalid number or passport"
// @Router /login/driver [post]
func (h *Handler) LoginDriver(c *gin.Context) {
	var req ap.DriverLoginReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.LoginDriver(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get the profile of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} ap.UserRes
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /profile [get]
func (h *Handler) Profile(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	user, err := h.User.Profile(context.Background(), &ap.ById{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUsers godoc
// @Summary Get users
// @Description Get the profile of the authenticated users
// @Tags admin
// @Accept json
// @Produce json
// @Param role query string false "Role"
// @Param page query integer false "Page"
// @Success 200 {object} ap.GetAllUsersRes
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/user/all [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)
	is_admin := claims.(jwt.MapClaims)["is_admin"].(string)

	if role != "superadmin" && is_admin != "true"{
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	var userRole string

	if role == "superadmin" {
		userRole = c.Query("role")
	} else {
		userRole = role
	}

	req := ap.GetAllUsersReq{
		Role: userRole,
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

	req.Page = int32(page)

	res, err := h.User.GetAllUsers(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get users", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteProfile godoc
// @Summary Delete user profile
// @Description Delete the profil of the authenticated user
// @Tags admin
// @Accept json
// @Produce json
// @Param id query string true "Id"
// @Success 200 {object} string "User profile deleted"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /admin/user/delete [delete]
func (h *Handler) DeleteProfile(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	role := claims.(jwt.MapClaims)["role"].(string)
	is_admin := claims.(jwt.MapClaims)["is_admin"].(string)

	if role != "superadmin" || is_admin != "true"{
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	id := c.Query("id")

	_, err := h.User.DeleteProfile(context.Background(), &ap.ById{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile deleted"})
}

// RefreshToken godoc
// @Summary Get token
// @Description Get the token of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} ap.TokenRes
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /refresh-token [get]
func (h *Handler) RefreshToken(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["user_id"].(string)

	token, err := h.User.RefreshToken(context.Background(), &ap.ById{Id: id})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, token)
}

// RefreshToken godoc
// @Summary Get token
// @Description Get the token of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} ap.TokenRes
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /refresh-token/driver [get]
func (h *Handler) RefreshTokenForDriver(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := claims.(jwt.MapClaims)["id"].(string)
	role := claims.(jwt.MapClaims)["role"].(string)

	if role != "driver"{
		c.JSON(http.StatusForbidden, gin.H{"error": "This page forbidden for you"})
		return
	}

	token, err := h.User.RefreshToken(context.Background(), &ap.ById{Id: id})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, token)
}
