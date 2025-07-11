package services

import (
    "github.com/anayy09/academiaflow-backend/internal/database"
    "github.com/anayy09/academiaflow-backend/internal/models"
    "gorm.io/gorm"
)

type CourseService struct {
    db *gorm.DB
}

func NewCourseService() *CourseService {
    return &CourseService{
        db: database.GetDB(),
    }
}

type CreateCourseRequest struct {
    CourseName string `json:"course_name" binding:"required"`
    CourseCode string `json:"course_code" binding:"required"`
    Instructor string `json:"instructor"`
    Credits    int    `json:"credits"`
    Semester   string `json:"semester" binding:"required"`
    Status     string `json:"status"`
}

func (s *CourseService) GetUserCourses(userID uint) ([]models.Course, error) {
    var courses []models.Course
    err := s.db.Where("user_id = ?", userID).Find(&courses).Error
    return courses, err
}

func (s *CourseService) CreateCourse(userID uint, req CreateCourseRequest) (*models.Course, error) {
    course := models.Course{
        UserID:     userID,
        CourseName: req.CourseName,
        CourseCode: req.CourseCode,
        Instructor: req.Instructor,
        Credits:    req.Credits,
        Semester:   req.Semester,
        Status:     req.Status,
    }

    if course.Status == "" {
        course.Status = "enrolled"
    }

    err := s.db.Create(&course).Error
    return &course, err
}

func (s *CourseService) GetCourse(userID, courseID uint) (*models.Course, error) {
    var course models.Course
    err := s.db.Where("id = ? AND user_id = ?", courseID, userID).First(&course).Error
    return &course, err
}

func (s *CourseService) UpdateCourse(userID, courseID uint, updates map[string]interface{}) (*models.Course, error) {
    var course models.Course
    if err := s.db.Where("id = ? AND user_id = ?", courseID, userID).First(&course).Error; err != nil {
        return nil, err
    }

    err := s.db.Model(&course).Updates(updates).Error
    return &course, err
}

func (s *CourseService) DeleteCourse(userID, courseID uint) error {
    return s.db.Where("id = ? AND user_id = ?", courseID, userID).Delete(&models.Course{}).Error
}