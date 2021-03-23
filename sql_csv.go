package lib

import (
	"database/sql"
	"fmt"
	"strings"
)

func SqlCsv(db *sql.DB, sql string) string {

	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return "error"
	}
	defer rows.Close()
	colNames, _ := rows.Columns()
	colValues := make([]interface{}, len(colNames))
	for i := range colValues {
		colValues[i] = new(interface{})
	}
	csv := ""
	csv += fmt.Sprintln(strings.Join(colNames, ";"))
	for rows.Next() {
		rows.Scan(colValues...)
		// fmt.Println(*columnValues[col].(*interface{}))
		for c := 0; c < len(colValues); c++ {
			v := fmt.Sprint(*colValues[c].(*interface{}))
			if v == "<nil>" {
				v = ""
			}
			csv += v + ";"
		}
		csv += "\n"
	}
	return csv
}
