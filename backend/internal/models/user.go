package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    Email     string         `json:"email" gorm:"uniqueIndex;not null"`
    Username  string         `json:"username" gorm:"uniqueIndex;not null"`
    Password  string         `json:"-" gorm:"not null"` // "-" excludes from JSON
    FirstName string         `json:"first_name"`
    LastName  string         `json:"last_name"`
    Program   string         `json:"program"`   // PhD, MS, etc.
    Year      int            `json:"year"`      // Year in program
    Advisor   string         `json:"advisor"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Course struct {
    ID          uint           `json:"id" gorm:"primaryKey"`
    UserID      uint           `json:"user_id" gorm:"not null"`
    User        User           `json:"-" gorm:"foreignKey:UserID"`
    CourseName  string         `json:"course_name" gorm:"not null"`
    CourseCode  string         `json:"course_code" gorm:"not null"`
    Instructor  string         `json:"instructor"`
    Credits     int            `json:"credits"`
    Semester    string         `json:"semester"` // Fall 2024, Spring 2025, etc.
    Grade       string         `json:"grade"`
    Status      string         `json:"status"` // enrolled, completed, dropped
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Assignment struct {
    ID          uint           `json:"id" gorm:"primaryKey"`
    UserID      uint           `json:"user_id" gorm:"not null"`
    CourseID    *uint          `json:"course_id"` // Optional - can be independent
    User        User           `json:"-" gorm:"foreignKey:UserID"`
    Course      *Course        `json:"course,omitempty" gorm:"foreignKey:CourseID"`
    Title       string         `json:"title" gorm:"not null"`
    Description string         `json:"description"`
    DueDate     time.Time      `json:"due_date"`
    Priority    string         `json:"priority"` // high, medium, low
    Status      string         `json:"status"`   // pending, in_progress, completed
    EstimatedHours int         `json:"estimated_hours"`
    ActualHours    int         `json:"actual_hours"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}