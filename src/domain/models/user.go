package models

import "time"

type User struct {
  ID uint `json:"id" gorm:"primary_key,not null"`  
  FirstName string `json:"first_name" gorm:"not null"`
  LastName string `json:"last_name" gorm:"not null"`
  Email string `json:"email" gorm:"not null;unique"`
  Password string `json:"password" gorm:"not null"`
  CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
  UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
