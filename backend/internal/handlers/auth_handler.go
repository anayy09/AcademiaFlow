package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/anayy09/academiaflow-backend/configs"
    "github.com/anayy09/academiaflow-backend/internal/auth"
    "github.com/anayy09/academiaflow-backend/internal/services"
)

type AuthHandler struct {
    userService *services.UserService
    config      *configs.Config
}

func NewAuthHandler(config *configs.Config) *AuthHandler {
    return &AuthHandler{
        userService: services.NewUserService(),
        config:      config,
    }
}

func (h *AuthHandler) Register(c *gin.Context) {
    var req services.RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.userService.Register(req)
    if err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        return
    }

    // Generate JWT token
    token, err := auth.GenerateToken(user.ID, user.Email, user.Username, h.config)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "User registered successfully",
        "token":   token,
        "user":    services.ToUserResponse(user),
    })
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req services.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.userService.Login(req)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    // Generate JWT token
    token, err := auth.GenerateToken(user.ID, user.Email, user.Username, h.config)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "token":   token,
        "user":    services.ToUserResponse(user),
    })
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
    userID := c.GetUint("user_id")
    
    user, err := h.userService.GetUserByID(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "user": services.ToUserResponse(user),
    })
}