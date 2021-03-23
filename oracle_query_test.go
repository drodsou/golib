package lib

import (
	"database/sql"
)

// QueryTest checks if oracle queries work
func OracleQueryTest(DB *sql.DB) string {

	if DB == nil {
		return "ERROR: app.DB is nil"
	}

	rows, err := DB.Query("select sysdate from dual")
	// rows, err := app.DB.Query("select apellidos from datos_personales where id_datos_personales = 360")
	if err != nil {
		return "Error running query. " + err.Error()
	}
	defer rows.Close()
	var result string
	for rows.Next() {
		rows.Scan(&result)
	}

	return result
}
