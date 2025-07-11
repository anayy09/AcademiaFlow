package services

import (
    "errors"

    "github.com/anayy09/academiaflow-backend/internal/auth"
    "github.com/anayy09/academiaflow-backend/internal/database"
    "github.com/anayy09/academiaflow-backend/internal/models"
    "gorm.io/gorm"
)

type UserService struct {
    db *gorm.DB
}

func NewUserService() *UserService {
    return &UserService{
        db: database.GetDB(),
    }
}

type RegisterRequest struct {
    Email     string `json:"email" binding:"required,email"`
    Username  string `json:"username" binding:"required,min=3,max=50"`
    Password  string `json:"password" binding:"required,min=6"`
    FirstName string `json:"first_name" binding:"required"`
    LastName  string `json:"last_name" binding:"required"`
    Program   string `json:"program"`
    Year      int    `json:"year"`
    Advisor   string `json:"advisor"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type UserResponse struct {
    ID        uint   `json:"id"`
    Email     string `json:"email"`
    Username  string `json:"username"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Program   string `json:"program"`
    Year      int    `json:"year"`
    Advisor   string `json:"advisor"`
}

func (s *UserService) Register(req RegisterRequest) (*models.User, error) {
    // Check if user already exists
    var existingUser models.User
    if err := s.db.Where("email = ? OR username = ?", req.Email, req.Username).First(&existingUser).Error; err == nil {
        return nil, errors.New("user with this email or username already exists")
    }

    // Hash password
    hashedPassword, err := auth.HashPassword(req.Password)
    if err != nil {
        return nil, err
    }

    // Create user
    user := models.User{
        Email:     req.Email,
        Username:  req.Username,
        Password:  hashedPassword,
        FirstName: req.FirstName,
        LastName:  req.LastName,
        Program:   req.Program,
        Year:      req.Year,
        Advisor:   req.Advisor,
    }

    if err := s.db.Create(&user).Error; err != nil {
        return nil, err
    }

    return &user, nil
}

func (s *UserService) Login(req LoginRequest) (*models.User, error) {
    var user models.User
    if err := s.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
        return nil, errors.New("invalid credentials")
    }

    if !auth.CheckPasswordHash(req.Password, user.Password) {
        return nil, errors.New("invalid credentials")
    }

    return &user, nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
    var user models.User
    if err := s.db.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (s *UserService) UpdateUser(id uint, updates map[string]interface{}) (*models.User, error) {
    var user models.User
    if err := s.db.First(&user, id).Error; err != nil {
        return nil, err
    }

    if err := s.db.Model(&user).Updates(updates).Error; err != nil {
        return nil, err
    }

    return &user, nil
}

func ToUserResponse(user *models.User) UserResponse {
    return UserResponse{
        ID:        user.ID,
        Email:     user.Email,
        Username:  user.Username,
        FirstName: user.FirstName,
        LastName:  user.LastName,
        Program:   user.Program,
        Year:      user.Year,
        Advisor:   user.Advisor,
    }
}