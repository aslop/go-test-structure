package appModule

import (
	"expl_app/app/resources/point"
	// "expl_app/app/resources/session"
	// "expl_app/app/resources/user"
)

type AppModule struct {
	// SessionModule *session.SessionModule
	// UserModule    *user.UserModule
	PointsModule *point.PointModule
}

func NewAppModule(
	// sessionModule *session.SessionModule,
	// userModule *user.UserModule,
	pointModule *point.PointModule,
) *AppModule {
	return &AppModule{
		// SessionModule: sessionModule,
		// UserModule:    userModule,
		PointsModule: pointModule,
	}
}

var AppModuleInstance *AppModule

func NewAppModuleInstance() *AppModule {
	if AppModuleInstance == nil {
		AppModuleInstance = NewAppModule(
			// // Session
			// session.GetSessionModuleInstance(),

			// // User
			// user.GetUserModuleInstance(),

			// Point
			point.GetPointModuleInstance(),
		)
	}
	return AppModuleInstance
}

func GetAppModuleInstance() *AppModule {
	return AppModuleInstance
}

func init() {
	AppModuleInstance = NewAppModuleInstance()
}
