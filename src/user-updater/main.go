package main

import (
	"log"
	"strconv"
	"strings"
	"time" //temp

	shopify "./apicaller"
	download "./apicaller/downloader"
	jsonlines "./apicaller/jsonlines"
	rbMQ "./rbMQ"
	"github.com/streadway/amqp"
)

func startOperation(userID string, jobType string, shopifyName string, shopifyAPI string, companySlug string, fikenAPI string, ch *amqp.Channel) {
	log.Printf("Sleeping a little before running BulkOperation.")
	time.Sleep(10 * time.Second)
	var link string
	var operationResultDataDir string
	var fileData string
	var customerlist []shopify.ShopifyCustomer
	var fikenList []shopify.FikenContact

	// LOGGING STUFF
	log_ch := make(chan string)
	go rbMQ.LogMessage(log_ch, ch)

	//start bulkOperation
	shopify.ShopifyCallBulkOperation(shopifyName) //refactored
	//Check operation status, till the operation is completed
	link = shopify.PollOperation(shopifyName) //refactored

	//download from the link
	if link != "null" && link != "" {
		//store directory of downloaded file
		operationResultDataDir = download.Download(link)
	} else {
		log.Println("Link was null or empty, no data downloaded.")
	}

	//convert file data from byte to string
	fileData = string(shopify.FileReader(operationResultDataDir))

	//Put jsonLines to struct
	error := jsonlines.Decode(strings.NewReader(fileData), &customerlist)
	if error != nil {
		log.Fatal(error)
	}

	fikenList = shopify.ConvertFromShopifyToFikenContacts(customerlist)
	shopify.ContactExists(fikenList, companySlug, fikenAPI)

	// Basic log
	string_user_id, _ := strconv.Atoi(userID)
	log_ch <- rbMQ.LogMessageFormatter(string_user_id, "Users updated.")
}
func main() {
	// -- Setup for RBMQ communication
	conn := rbMQ.GetConnection()
	defer conn.Close()

	ch := rbMQ.GetChannelFromConnection(conn)
	defer ch.Close()

	msgs := rbMQ.GetChannelDataByName(ch, "job")

	forever := make(chan bool)

	// Listen for messages from rbMQ
	go func() {
		// Runs when messages are received
		for d := range msgs {
			log.Printf("Received a message in user-updater")
			// Splits message body by : into an array
			message := strings.Split(string(d.Body), ":")

			// Get all info from message
			userID := message[0]
			jobType := message[1]
			shopifyName := message[2]
			shopifyAPI := message[3]
			companySlug := message[4]
			fikenAPI := message[5]

			//log.Printf("Testing values: %s, %s, %s, %s, %s, %s", userID, jobType, shopifyName, shopifyAPI, companySlug, fikenAPI)

			// If type = update-user, progress, else, log that a job was received but it wasnt right
			if jobType == "update-user" {
				log.Printf("Job message received, stuff should happen now.")
				go startOperation(userID, jobType, shopifyName, shopifyAPI, companySlug, fikenAPI, ch)
			} else {
				log.Printf("Job message received, but it was the wrong type.")
			}
		}
	}()
	log.Printf(" [*] Waiting for messages.")
	<-forever

}
