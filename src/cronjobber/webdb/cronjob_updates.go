package webdb

import (
	"database/sql"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func UpdateJobLastCalled(db *sql.DB, job_id int) bool {
	sqlStatement := `
	UPDATE Cronjobs
	SET last_called_date = ?
	WHERE cronjob_id = ?;`
	result, err := db.Exec(sqlStatement, time.Now().Format("2006-01-02"), job_id)


	if err != nil {
		log.Fatal(err)
	}

	count, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}