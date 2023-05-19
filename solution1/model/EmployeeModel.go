package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type EmployeeModel struct {
	EmpID     string             `json:"_id,omitempty" bson:"_id"`
	FullName  string             `json:"full_name,omitempty" bson:"full_name" validate:"required"`
	Address   string             `json:"address,omitempty" bson:"address" validate:"required"`
	Email     string             `json:"email,omitempty" bson:"email" validate:"required,email" `
	Phone     string             `json:"phone,omitempty" bson:"phone" validate:"required"`
	Position  string             `json:"position,omitempty" bson:"position" validate:"required"`
	Salary    float64            `json:"salary,omitempty" bson:"salary" validate:"required,numeric"`
	Status    string             `json:"status,omitempty" bson:"status"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at" `
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at"`
}
