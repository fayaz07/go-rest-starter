package helpers

import "context"

type ContextPurpose string

const (
	DatabaseConnection ContextPurpose = "db_connect"
	DatabaseReadWrite  ContextPurpose = "db_read_write"
)

type AppContext struct {
	Purpose    ContextPurpose
	Ctx        context.Context
	cancelFunc context.CancelFunc
}

func (c *AppContext) Cancel() {
	c.cancelFunc()
}
