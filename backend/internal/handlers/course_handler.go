package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/anayy09/academiaflow-backend/configs"
    "github.com/anayy09/academiaflow-backend/internal/services"
)

type CourseHandler struct {
    courseService *services.CourseService
    config        *configs.Config
}

func NewCourseHandler(config *configs.Config) *CourseHandler {
    return &CourseHandler{
        courseService: services.NewCourseService(),
        config:        config,
    }
}

func (h *CourseHandler) GetCourses(c *gin.Context) {
    userID := c.GetUint("user_id")
    
    courses, err := h.courseService.GetUserCourses(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch courses"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"courses": courses})
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {
    userID := c.GetUint("user_id")
    
    var req services.CreateCourseRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    course, err := h.courseService.CreateCourse(userID, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create course"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Course created successfully",
        "course":  course,
    })
}

func (h *CourseHandler) GetCourse(c *gin.Context) {
    userID := c.GetUint("user_id")
    courseID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    course, err := h.courseService.GetCourse(userID, uint(courseID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"course": course})
}

func (h *CourseHandler) UpdateCourse(c *gin.Context) {
    userID := c.GetUint("user_id")
    courseID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    var updates map[string]interface{}
    if err := c.ShouldBindJSON(&updates); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    course, err := h.courseService.UpdateCourse(userID, uint(courseID), updates)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update course"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Course updated successfully",
        "course":  course,
    })
}

func (h *CourseHandler) DeleteCourse(c *gin.Context) {
    userID := c.GetUint("user_id")
    courseID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    err = h.courseService.DeleteCourse(userID, uint(courseID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete course"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}