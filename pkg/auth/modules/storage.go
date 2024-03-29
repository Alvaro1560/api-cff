package modules

import (
	"github.com/jmoiron/sqlx"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/logger"
	"gitlab.ecapture.com.co/gitlab-instance/gitlab-instance-cea63b52/e-capture/indra/api-indra-admin/internal/models"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
)

type ServicesModuleRepository interface {
	Create(m *Module) error
	Update(m *Module) error
	Delete(id string) error
	GetByID(id string) (*Module, error)
	GetAll() ([]*Module, error)
	GetModulesByRoles(roleIDs []string, ids []string, typeArg int) ([]*Module, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesModuleRepository {
	var s ServicesModuleRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return NewModuleSqlServerRepository(db, user, txID)
	case Postgresql:
		return NewModulePsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
