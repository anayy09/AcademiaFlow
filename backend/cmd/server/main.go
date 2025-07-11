package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/anayy09/academiaflow-backend/configs"
    "github.com/anayy09/academiaflow-backend/internal/database"
    "github.com/anayy09/academiaflow-backend/internal/handlers"
    "github.com/anayy09/academiaflow-backend/internal/middleware"
)

func main() {
    // Load configuration
    config := configs.LoadConfig()

    // Connect to database
    database.Connect(config)
    database.Migrate()

    // Set Gin mode
    if config.Server.Host != "localhost" {
        gin.SetMode(gin.ReleaseMode)
    }

    // Initialize Gin router
    router := gin.Default()

    // Add middleware
    router.Use(middleware.CORSMiddleware())
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    // Initialize handlers
    authHandler := handlers.NewAuthHandler(config)
    userHandler := handlers.NewUserHandler(config)
    courseHandler := handlers.NewCourseHandler(config)
    assignmentHandler := handlers.NewAssignmentHandler(config)

    // Health check endpoint
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status":  "ok",
            "message": "AcademiaFlow API is running",
        })
    })

    // API v1 routes
    v1 := router.Group("/api/v1")
    {
        // Authentication routes (public)
        auth := v1.Group("/auth")
        {
            auth.POST("/register", authHandler.Register)
            auth.POST("/login", authHandler.Login)
        }

        // Protected routes
        protected := v1.Group("/")
        protected.Use(middleware.AuthMiddleware(config))
        {
            // User routes
            users := protected.Group("/users")
            {
                users.GET("/profile", authHandler.GetProfile)
                users.PUT("/profile", userHandler.UpdateProfile)
            }

            // Course routes
            courses := protected.Group("/courses")
            {
                courses.GET("/", courseHandler.GetCourses)
                courses.POST("/", courseHandler.CreateCourse)
                courses.GET("/:id", courseHandler.GetCourse)
                courses.PUT("/:id", courseHandler.UpdateCourse)
                courses.DELETE("/:id", courseHandler.DeleteCourse)
            }

            // Assignment routes
            assignments := protected.Group("/assignments")
            {
                assignments.GET("/", assignmentHandler.GetAssignments)
                assignments.POST("/", assignmentHandler.CreateAssignment)
                assignments.GET("/:id", assignmentHandler.GetAssignment)
                assignments.PUT("/:id", assignmentHandler.UpdateAssignment)
                assignments.DELETE("/:id", assignmentHandler.DeleteAssignment)
                assignments.PATCH("/:id/status", assignmentHandler.UpdateStatus)
            }
        }
    }

    // Start server
    log.Printf("Server starting on %s:%s", config.Server.Host, config.Server.Port)
    log.Fatal(router.Run(":" + config.Server.Port))
}