package main

// This service publishes messages to the rabbitMQ queue.
import (
	"log"

	rbMQ "./rbMQ"

	"time"

	db "./webdb"
	//"strconv"
)

type Cronjob struct {
	ID         int    `json:"cronjob_id"`
	Type       string `json:"job_type"`
	LastCalled string `json:"last_called_date"`
	Interval   int    `json:"interval_days"`
	UserID     int    `json:"user_id"`
}

type CronjobWithAPI struct {
	ID         int    `json:"cronjob_id"`
	LastCalled string `json:"last_called_date"`
	Interval   int    `json:"interval_days"`
	UserID     int    `json:"user_id"`
	Type       string `json:"job_type"`
	ShopifyName string `json:"shopify_name"`
	API_Shopify string `json:"api_shopify"`
	CompanySlug string `json:"company_slug"`
	API_Fiken string `json:"api_fiken"`
}

func main() {
	// It waits a little before attempting to connect to anything
	time.Sleep(7 * time.Second)
	// -- Setup for RBMQ communication
	// Get connection for this work session
	conn := rbMQ.GetConnection()
	defer conn.Close()

	//Get channel from connection
	ch := rbMQ.GetChannelFromConnection(conn)
	defer ch.Close()

	//Get DB connection
	connection := db.Connect()
	defer connection.Close()

	//Start goroutine for job messager
	job_ch := make(chan string)
	go rbMQ.PublishJob(job_ch, ch)

	for {
		results := db.GetCronjobsWithAPIKeys(connection)
		defer results.Close()

		for results.Next() {
			var job CronjobWithAPI
			
			err := results.Scan(&job.ID, &job.LastCalled, &job.Interval, &job.UserID, &job.Type, &job.ShopifyName, &job.API_Shopify, &job.CompanySlug, &job.API_Fiken)

			if err != nil {
				panic(err.Error())
			}

			// Publishes job messages in job queue
			if rbMQ.IsXDaysAway(job.LastCalled, job.Interval) {
				// Format message with values in job struct
				messageFetched := rbMQ.JobMessageFormatter(job.UserID, job.Type, job.ShopifyName, job.API_Shopify, job.CompanySlug, job.API_Fiken)
				log.Printf("Sending job-message for user with id %d", job.UserID)
				// Send message to job channel
				job_ch <- messageFetched
				
				// Update the job recently messaged abouts last_called_date
				success := db.UpdateJobLastCalled(connection, job.ID)
				// Check whether the row was updated successfully
				log.Printf("last_called_date updated: %t", success)
			} else {
				log.Printf("Job found, but it was not time to do it.")
			}
		}
		log.Printf("All jobs sent. Waiting 2 hours.")
		time.Sleep(2 * time.Hour)
	}
}
