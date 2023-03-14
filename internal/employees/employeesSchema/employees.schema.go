package employeesschema

import (
	"github.com/tetrex/coffeeshop-crud-golang/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ----------------
// collection
var EmployeeCollection = *db.MongoDataBase.Collection("employees")

// ----------------

type EmployeeRoles string

const (
	ADMIN_ROLE   EmployeeRoles = "ADMIN"
	MANAGER_ROLE EmployeeRoles = "MANAGER"
	BARISTA_ROLE EmployeeRoles = "BARISTA"
)

type Employees struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	Role      EmployeeRoles       `bson:"role,omitempty"`
	Name      string              `bson:"name,omitempty"`
	Email     string              `bson:"email,omitempty"`
	Password  string              `bson:"password,omitempty"`
	Deleted   bool                `bson:"bool,omitempty"`
	CreatedAt primitive.Timestamp `bson:"createdAt,omitempty"`
	UpdatedAt primitive.Timestamp `bson:"updatedAt,omitempty"`
}
type EmployeesSchema Employees
