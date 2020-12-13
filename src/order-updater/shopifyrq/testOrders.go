package shopifyrq

import (
	"log"
	"strings"
	"time"

	download "../downloader"
	fiken "../fikenrq"
	jsonlines "../jsonlines"
	rbMQ "../rbMQ"
)

func TestOrders() {
	var link string
	var operationResultDataDir string
	var shopifyOrders []ShopifyOrder
	var fileData string
	var fikenSales []fiken.CreateSale

	//start bulkOperation
	ShopifyCallBulkOperation()
	//Check operation status, till the operation is completed
	link = PollOperation()

	//download from the link
	if link != "null" && link != "" {
		//store directory of downloaded file
		operationResultDataDir = download.Download(link)
	} else {
		log.Println("Link was null or empty, no data downloaded.")
	}

	log.Println(operationResultDataDir)
	//TESTING
	// operationResultDataDir = "order-updater/downloader/downloads/Files.jsonl"
	//convert file data from byte to string
	fileData = string(FileReader(operationResultDataDir))

	//Put jsonLines to struct
	error := jsonlines.Decode(strings.NewReader(fileData), &shopifyOrders)
	if error != nil {
		log.Fatal(error)
	}

	companySlug := ""
	url := "https://api.fiken.no/api/v2/companies/" + companySlug
	token := ""
	var bearer = "Bearer " + token
	fikenSales = ConvertFromShopifyToFikenSales(shopifyOrders, url, bearer)
	for i := range fikenSales {
		fiken.CreateFikenSale(url, bearer, fikenSales[i])
		time.Sleep(time.Second * 3)
	}

	//fmt.Println()

}

func TestOrdersWithDocker() {
	var link string
	var operationResultDataDir string
	var shopifyOrders []ShopifyOrder
	var fileData string
	var fikenSales []fiken.CreateSale

	// -- FROM DANIEL Get connection to rbMQ
	// Get connection for this work session
	conn := rbMQ.GetConnection()
	defer conn.Close()

	// -- FROM DANIEL Get channel from connection
	ch := rbMQ.GetChannelFromConnection(conn)
	defer ch.Close()

	// -- FROM DANIEL Start goroutine for log messaging
	log_ch := make(chan string)
	go rbMQ.LogMessage(log_ch, ch)

	sleepDuration := time.Duration(30)
	log.Println("Sleeping for ", sleepDuration, " before sending bulkOperationRequest.")
	time.Sleep(time.Second * sleepDuration)
	//start bulkOperation
	ShopifyCallBulkOperation()
	//Check operation status, till the operation is completed
	link = PollOperation()

	//download from the link
	if link != "null" && link != "" {
		//store directory of downloaded file
		operationResultDataDir = download.Download(link)
	} else {
		log.Println("Link was null or empty, no data downloaded.")
	}

	log.Println(operationResultDataDir)
	//TESTING
	// operationResultDataDir = "order-updater/downloader/downloads/Files.jsonl"
	//convert file data from byte to string
	fileData = string(FileReader(operationResultDataDir))

	//Put jsonLines to struct
	error := jsonlines.Decode(strings.NewReader(fileData), &shopifyOrders)
	if error != nil {
		log.Fatal(error)
	}

	companySlug := "" //TODO: Replace with company slug from MSG

	url := "https://api.fiken.no/api/v2/companies/" + companySlug
	token := "" //TODO: Replace with API key
	var bearer = "Bearer " + token
	fikenSales = ConvertFromShopifyToFikenSales(shopifyOrders, url, bearer)
	for i := range fikenSales {
		fiken.CreateFikenSale(url, bearer, fikenSales[i])
		// FROM DANIEL Send a basic log message with userid 1 (ID should be from DB, message should be response body)
		log_ch <- rbMQ.LogMessageFormatter(1, "a sale was attempted.")

		time.Sleep(time.Second * 3)
	}

	// FROM DANIEL Send a basic log message with userid 1 (ID should be from DB, message should be response body)
	log_ch <- rbMQ.LogMessageFormatter(1, "Sale thing done for noone really.")
}
