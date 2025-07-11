package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/anayy09/academiaflow-backend/configs"
    "github.com/anayy09/academiaflow-backend/internal/services"
)

type UserHandler struct {
    userService *services.UserService
    config      *configs.Config
}

func NewUserHandler(config *configs.Config) *UserHandler {
    return &UserHandler{
        userService: services.NewUserService(),
        config:      config,
    }
}

type UpdateProfileRequest struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Program   string `json:"program"`
    Year      int    `json:"year"`
    Advisor   string `json:"advisor"`
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
    userID := c.GetUint("user_id")
    
    var req UpdateProfileRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updates := map[string]interface{}{
        "first_name": req.FirstName,
        "last_name":  req.LastName,
        "program":    req.Program,
        "year":       req.Year,
        "advisor":    req.Advisor,
    }

    user, err := h.userService.UpdateUser(userID, updates)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update profile"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Profile updated successfully",
        "user":    services.ToUserResponse(user),
    })
}