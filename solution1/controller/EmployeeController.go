package controller

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"

	"goenv/configs"
	"goenv/model"
	"goenv/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

var EmployeeCollection = configs.GetCollection(configs.DB, "employee")

func typetime(now time.Time) primitive.DateTime {
	return primitive.DateTime(now.UTC().UnixNano() / int64(time.Millisecond))
}
func GenerateID() string {
	id := time.Now().Format("20060102150405")
	return id
}

func CreateEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var employee model.EmployeeModel

		defer cancel()
		if err := c.ShouldBindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}
		validate := validator.New()

		err := validate.Struct(employee)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		employee.CreatedAt = typetime(time.Now())
		employee.UpdatedAt = typetime(time.Now())
		employee.Status = "active"
		employee.EmpID = "EMP-" + GenerateID()

		checkID := model.EmployeeModel{}
		if err := EmployeeCollection.FindOne(ctx, model.EmployeeModel{EmpID: employee.EmpID}).Decode(&checkID); err == nil {
			employee.EmpID = "EMP-" + GenerateID()
		}

		if _, err := EmployeeCollection.InsertOne(ctx, employee); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
			return
		} else {
			c.JSON(http.StatusCreated, responses.SuccessResponse{Status: http.StatusCreated, Message: "Success Create Employee", Data: map[string]interface{}{"employee": employee}})
		}

	}
}

func GetEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		employee := model.EmployeeModel{}
		id := c.Param("id")
		fmt.Println(id)

		if err := EmployeeCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&employee); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Message: "Success Get Employee ID => " + id, Data: map[string]interface{}{"employee": employee}})
		}

	}
}

func GetEmployees() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var employees []model.EmployeeModel

		cursor, err := EmployeeCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}
		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				return
			}
		}(cursor, ctx)
		for cursor.Next(ctx) {
			var employee model.EmployeeModel
			err := cursor.Decode(&employee)
			if err != nil {
				return
			}
			employees = append(employees, employee)
		}
		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Message: "Success Get Employees", Data: map[string]interface{}{"employees": employees}})
		}
	}
}

func UpdateEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var employee model.EmployeeModel

		id := c.Param("id")

		defer cancel()
		if err := c.ShouldBindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}
		if validationErr := validator.New().Struct(employee); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: validationErr.Error()})
			return
		}

		empExit, _ := EmployeeCollection.CountDocuments(ctx, bson.M{"_id": id})
		if empExit == 0 {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Employee Not Found"})
			return
		}

		update := bson.M{
			"$set": bson.M{
				"full_name":  employee.FullName,
				"address":    employee.Address,
				"email":      employee.Email,
				"phone":      employee.Phone,
				"position":   employee.Position,
				"salary":     employee.Salary,
				"status":     employee.Status,
				"updated_at": typetime(time.Now()),
			},
		}

		result, err := EmployeeCollection.UpdateOne(ctx, bson.M{"_id": id}, update)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		}
		var updatedEmployee model.EmployeeModel
		if result.ModifiedCount == 1 {
			if err := EmployeeCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&updatedEmployee); err != nil {
				c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
				return
			} else {
				c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Message: "Success Update Employee ID => " + id, Data: map[string]interface{}{"employee": updatedEmployee}})
			}
		}

	}
}

func DeleteEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var employee model.EmployeeModel

		defer cancel()

		id := c.Param("id")
		fmt.Println(id)

		if err := EmployeeCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&employee); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		if id == employee.EmpID {
			result, err := EmployeeCollection.DeleteOne(ctx, bson.M{"_id": id})
			if err != nil {
				c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
				return
			} else {
				c.JSON(http.StatusCreated, responses.SuccessResponse{Status: http.StatusCreated, Message: "Success Delete Employee", Data: map[string]interface{}{"result": result}})
			}
		}
	}
}
