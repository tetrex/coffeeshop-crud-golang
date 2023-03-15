package productsschema

import (
	"github.com/tetrex/coffeeshop-crud-golang/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ----------------
// collection

func ProductsCollection() *mongo.Collection {
	var collection *mongo.Collection = db.MongoDataBase.Collection("products")
	return collection

}

// var collection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
// ----------------

type Product struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	Name      string              `bson:"name,omitempty"`
	Price     float64             `bson:"price,omitempty"`
	Qty       float32             `bson:"qty,omitempty"`
	Deleted   bool                `bson:"deleted,omitempty"`
	CreatedAt primitive.Timestamp `bson:"createdAt,omitempty"`
	UpdatedAt primitive.Timestamp `bson:"updatedAt,omitempty"`
}
type ProductsSchema Product
