package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/anayy09/academiaflow-backend/configs"
    "github.com/anayy09/academiaflow-backend/internal/services"
)

type AssignmentHandler struct {
    assignmentService *services.AssignmentService
    config           *configs.Config
}

func NewAssignmentHandler(config *configs.Config) *AssignmentHandler {
    return &AssignmentHandler{
        assignmentService: services.NewAssignmentService(),
        config:           config,
    }
}

func (h *AssignmentHandler) GetAssignments(c *gin.Context) {
    userID := c.GetUint("user_id")
    
    // Optional filters
    status := c.Query("status")
    priority := c.Query("priority")
    
    assignments, err := h.assignmentService.GetUserAssignments(userID, status, priority)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch assignments"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"assignments": assignments})
}

func (h *AssignmentHandler) CreateAssignment(c *gin.Context) {
    userID := c.GetUint("user_id")
    
    var req services.CreateAssignmentRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    assignment, err := h.assignmentService.CreateAssignment(userID, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create assignment"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message":    "Assignment created successfully",
        "assignment": assignment,
    })
}

func (h *AssignmentHandler) GetAssignment(c *gin.Context) {
    userID := c.GetUint("user_id")
    assignmentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
        return
    }

    assignment, err := h.assignmentService.GetAssignment(userID, uint(assignmentID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Assignment not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"assignment": assignment})
}

func (h *AssignmentHandler) UpdateAssignment(c *gin.Context) {
    userID := c.GetUint("user_id")
    assignmentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
        return
    }

    var updates map[string]interface{}
    if err := c.ShouldBindJSON(&updates); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    assignment, err := h.assignmentService.UpdateAssignment(userID, uint(assignmentID), updates)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update assignment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":    "Assignment updated successfully",
        "assignment": assignment,
    })
}

func (h *AssignmentHandler) DeleteAssignment(c *gin.Context) {
    userID := c.GetUint("user_id")
    assignmentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
        return
    }

    err = h.assignmentService.DeleteAssignment(userID, uint(assignmentID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete assignment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Assignment deleted successfully"})
}

func (h *AssignmentHandler) UpdateStatus(c *gin.Context) {
    userID := c.GetUint("user_id")
    assignmentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
        return
    }

    var req struct {
        Status string `json:"status" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    assignment, err := h.assignmentService.UpdateAssignment(userID, uint(assignmentID), map[string]interface{}{
        "status": req.Status,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update assignment status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":    "Assignment status updated successfully",
        "assignment": assignment,
    })
}