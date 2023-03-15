package employeesservice

import (
	"context"
	"fmt"
	"time"

	employeesschema "github.com/tetrex/coffeeshop-crud-golang/internal/employees/employeesSchema"
	"github.com/tetrex/coffeeshop-crud-golang/internal/employees/employeesService/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(dto *dto.Employees) (bool, error) {
	employeeObj := employeesschema.EmployeesSchema{
		Role:      []string{string(dto.Role)},
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  dto.Password,
		Deleted:   false,
		CreatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
		UpdatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
	}
	collection := employeesschema.EmployeeCollection()
	result, err := collection.InsertOne(context.TODO(), employeeObj)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Printf("inserted :: %s", result.InsertedID)

	return true, nil
}

// object ,error,sucess_status
func Find(dto *dto.FindEmployee) (bson.M, error, bool) {
	var userObject bson.M
	// convert id to hex id
	objectId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		fmt.Println(err)
		return userObject, err, false
	}

	collection := employeesschema.EmployeeCollection()
	result := collection.FindOne(context.TODO(), bson.M{"_id": objectId, "deleted": false}) // find non deleted entry
	if err := result.Decode(&userObject); err != nil {
		return userObject, err, false
	}

	return userObject, err, true
}
