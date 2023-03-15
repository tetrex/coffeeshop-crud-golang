package productsservice

import (
	"context"
	"fmt"
	"time"

	productsschema "github.com/tetrex/coffeeshop-crud-golang/internal/products/productsSchema"
	"github.com/tetrex/coffeeshop-crud-golang/internal/products/productsService/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(dto *dto.CreateProduct) (bool, error) {
	productObj := productsschema.ProductsSchema{
		Name:      dto.Name,
		Price:     dto.Price,
		Qty:       dto.Qty,
		Deleted:   false,
		CreatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
		UpdatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
	}
	collection := productsschema.ProductsCollection()
	result, err := collection.InsertOne(context.TODO(), productObj)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Printf("inserted :: %s", result.InsertedID)

	return true, nil
}

// object ,error,sucess_status
func Find(dto *dto.FindProduct) (bson.M, error, bool) {
	var productObj bson.M
	// convert id to hex id
	objectId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		fmt.Println(err)
		return productObj, err, false
	}

	collection := productsschema.ProductsCollection()
	result := collection.FindOne(context.TODO(), bson.M{"_id": objectId, "deleted": false}) // find non deleted entry
	if err := result.Decode(&productObj); err != nil {
		return productObj, err, false
	}

	return productObj, err, true
}

// sucess_status,error
func Delete(dto *dto.DeleteProduct) (bool, error) {
	// convert id to hex id
	objectId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	collection := productsschema.ProductsCollection()
	collection.FindOneAndUpdate(context.TODO(), bson.M{"_id": objectId}, bson.M{"deleted": true})

	return true, err
}
