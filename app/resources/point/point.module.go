package point

import (
	"expl_app/app/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type PointModule struct {
	Repository *pointRepository
}

var (
	PointModuleInstance *PointModule
)

func NewPointModule(db *mongo.Database) *PointModule {
	PointModuleInstance = &PointModule{
		Repository: GetPointRepositoryInstance(),
	}

	return PointModuleInstance
}

// ==================================================
// Get instance
// ==================================================
func GetPointModuleInstance() *PointModule {
	if PointModuleInstance == nil {
		PointModuleInstance = NewPointModule(db.Database)
	}

	return PointModuleInstance
}
