package webdb

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllCronjobs(db *sql.DB) *sql.Rows {
	results, err := db.Query("SELECT * FROM Cronjobs")

	if err != nil {
		log.Fatal(err)
	}

	return results
}

func GetCronjobsWithAPIKeys(db *sql.DB) *sql.Rows {
	results, err := db.Query("SELECT cronjob_id, last_called_date, interval_days, Cronjobs.user_id as user_id, job_type, Users.shopify_name as shopify_name, Users.api_key_shopify as api_shopify, Users.company_slug as company_slug, Users.api_key_fiken as api_fiken FROM Cronjobs INNER JOIN Users ON Cronjobs.user_id = Users.user_id")

	if err != nil {
		log.Fatal(err)
	}

	return results
}

func GetCronjobDates(db *sql.DB) *sql.Rows {
	results, err := db.Query("SELECT last_called_date FROM Cronjobs")

	if err != nil {
		log.Fatal(err)
	}

	defer results.Close()

	return results
}
