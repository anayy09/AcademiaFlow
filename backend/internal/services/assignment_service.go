package services

import (
    "time"

    "github.com/anayy09/academiaflow-backend/internal/database"
    "github.com/anayy09/academiaflow-backend/internal/models"
    "gorm.io/gorm"
)

type AssignmentService struct {
    db *gorm.DB
}

func NewAssignmentService() *AssignmentService {
    return &AssignmentService{
        db: database.GetDB(),
    }
}

type CreateAssignmentRequest struct {
    CourseID       *uint     `json:"course_id"`
    Title          string    `json:"title" binding:"required"`
    Description    string    `json:"description"`
    DueDate        time.Time `json:"due_date" binding:"required"`
    Priority       string    `json:"priority"`
    EstimatedHours int       `json:"estimated_hours"`
}

func (s *AssignmentService) GetUserAssignments(userID uint, status, priority string) ([]models.Assignment, error) {
    var assignments []models.Assignment
    query := s.db.Where("user_id = ?", userID).Preload("Course")

    if status != "" {
        query = query.Where("status = ?", status)
    }
    if priority != "" {
        query = query.Where("priority = ?", priority)
    }

    err := query.Order("due_date ASC").Find(&assignments).Error
    return assignments, err
}

func (s *AssignmentService) CreateAssignment(userID uint, req CreateAssignmentRequest) (*models.Assignment, error) {
    assignment := models.Assignment{
        UserID:         userID,
        CourseID:       req.CourseID,
        Title:          req.Title,
        Description:    req.Description,
        DueDate:        req.DueDate,
        Priority:       req.Priority,
        EstimatedHours: req.EstimatedHours,
        Status:         "pending",
    }

    if assignment.Priority == "" {
        assignment.Priority = "medium"
    }

    err := s.db.Create(&assignment).Error
    if err != nil {
        return nil, err
    }

    // Load the course relationship
    s.db.Preload("Course").First(&assignment, assignment.ID)
    
    return &assignment, err
}

func (s *AssignmentService) GetAssignment(userID, assignmentID uint) (*models.Assignment, error) {
    var assignment models.Assignment
    err := s.db.Where("id = ? AND user_id = ?", assignmentID, userID).
        Preload("Course").
        First(&assignment).Error
    return &assignment, err
}

func (s *AssignmentService) UpdateAssignment(userID, assignmentID uint, updates map[string]interface{}) (*models.Assignment, error) {
    var assignment models.Assignment
    if err := s.db.Where("id = ? AND user_id = ?", assignmentID, userID).First(&assignment).Error; err != nil {
        return nil, err
    }

    err := s.db.Model(&assignment).Updates(updates).Error
    if err != nil {
        return nil, err
    }

    // Reload with relationships
    s.db.Preload("Course").First(&assignment, assignment.ID)
    
    return &assignment, err
}

func (s *AssignmentService) DeleteAssignment(userID, assignmentID uint) error {
    return s.db.Where("id = ? AND user_id = ?", assignmentID, userID).Delete(&models.Assignment{}).Error
}