package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// add up queries...
		dao := daos.New(db)

		return nil
	}, func(db dbx.Builder) error {
		// add down queries...

		return nil
	})
}
