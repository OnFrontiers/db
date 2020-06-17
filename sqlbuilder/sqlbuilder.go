package sqlbuilder

import (
	"database/sql"
	"fmt"
	"github.com/upper/db"
)

// Engine represents a SQL database engine.
type Engine interface {
	db.Session

	SQLBuilder
}

func lookupAdapter(adapterName string) (Adapter, error) {
	adapter := db.LookupAdapter(adapterName)
	if sqlAdapter, ok := adapter.(Adapter); ok {
		return sqlAdapter, nil
	}

	return nil, fmt.Errorf("bond: missing SQL adapter %q", adapterName)
}

func BindTx(adapterName string, tx *sql.Tx) (Tx, error) {
	adapter, err := lookupAdapter(adapterName)
	if err != nil {
		return nil, err
	}

	return adapter.NewTx(tx)
}

// Bind creates a binding between an adapter and a *sql.Tx or a *sql.DB.
func BindDB(adapterName string, sess *sql.DB) (Session, error) {
	adapter, err := lookupAdapter(adapterName)
	if err != nil {
		return nil, err
	}

	return adapter.New(sess)
}
