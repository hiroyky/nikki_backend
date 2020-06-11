package infrastructures

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

func OpenDatabase(hostName, port, dbName, userName, password string, debugMode bool) (*sql.DB, error) {
	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		userName,
		password,
		hostName,
		port,
		dbName,
	)
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}

	boil.DebugMode = debugMode
	boil.SetDB(db)
	return db, nil
}
