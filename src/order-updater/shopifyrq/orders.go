package shopifyrq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	fiken "../fikenrq"
)

//ShopifyRequest sends a request to the store
func ShopifyRequest(reqMethod, storeURL, query string) []byte {

	var link = "https://" + storeURL + ".myshopify.com/admin/api/graphql.json"
	request, err := http.NewRequest(reqMethod, link, bytes.NewBuffer([]byte(query)))
	request.Header.Set("Content-Type", "application/graphql")
	request.Header.Set("X-Shopify-Access-Token", "")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	// defer response.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	// fmt.Println("Response Status: ", response.Status)
	responseData, _ := ioutil.ReadAll(response.Body)

	return responseData
}

func ShopifyCallBulkOperation() {
	// construct query
	var query string = "mutation {" +
		"bulkOperationRunQuery(" +
		"query:" +
		`"""` +
		"{" +
		"orders" +
		`(query:"test:true")` +
		"{" +
		"edges{" +
		"node{" +
		"id\n" +
		"name\n" +
		"createdAt\n" +
		"fullyPaid\n" +
		"email\n" +
		"currencyCode\n" +
		"transactions{\n" +
		"gateway\n" +
		"createdAt\n" +
		"}" +
		"totalTaxSet{shopMoney{amount}}\n" +
		"originalTotalPriceSet{shopMoney{amount}}\n" +
		"}" +
		"}" +
		"}" +
		"}" +
		`"""` +
		")" +
		"{" +
		"bulkOperation {\n" +
		"id\n" +
		"status\n" +
		"}" +
		"userErrors {" +
		"field\n" +
		"message\n" +
		"}" +
		"}" +
		"}"

	//fmt.Println(query)
	var responseBytes = ShopifyRequest("POST", "", query)

	//get bulkoperation struct
	var bulkOperation BulkOperationResponse

	error := json.Unmarshal(responseBytes, &bulkOperation)
	if error != nil {
		fmt.Println("Failed to put into struct : \n", error)
	}
	//Operation ID to keep track of which operation it is
	var operationID = bulkOperation.Data.BulkOperationRunQuery.BulkOperation.ID

	if operationID == "" {
		log.Fatalln("Something went wrong. No operation running.")
	} else {
		fmt.Println("Operation ID: ", operationID)
		// fmt.Println("Struct is fine")
		var userErrors []interface{}
		userErrors = bulkOperation.Data.BulkOperationRunQuery.UserErrors

		//check for user errors in query
		if len(userErrors) > 0 {

			for i := range userErrors {
				log.Fatalln(userErrors[i])
			}
		}
	}

}

func pollOperationStatus() (string, string) {
	query :=
		"query {\n" +
			"currentBulkOperation {\n" +
			"id\n" +
			"status\n" +
			"errorCode\n" +
			"createdAt\n" +
			"completedAt\n" +
			"objectCount\n" +
			"fileSize\n" +
			"url\n" +
			"partialDataUrl\n" +
			"}" +
			"}"

	var responseBytes = ShopifyRequest("POST", "", query)

	var status BulkOperationStatus

	error := json.Unmarshal(responseBytes, &status)
	if error != nil {
		fmt.Println("Failed to put into struct : \n", error)
	}
	var stat = status.Data.CurrentBulkOperation.Status
	var link = status.Data.CurrentBulkOperation.URL

	// fmt.Println("URL: ", link)

	if status.Data.CurrentBulkOperation.Status != "COMPLETED" {
		fmt.Println("Operation still in progress.")
	} else if status.Data.CurrentBulkOperation.ErrorCode == "null" {
		log.Fatalln("Operation failed with code: ", status.Data.CurrentBulkOperation.ErrorCode)
	} else if status.Data.CurrentBulkOperation.Status == "COMPLETED" {
		log.Println("Operation was completed.")
	}

	return stat, link

}

/*
Checks progress of pollOperation
*/
func PollOperation() string {

	var isDone = false
	var status string
	var link string
	var waitTime int32 = 1
	var maxWaitTime int32 = 3600

	for !isDone {
		status, link = pollOperationStatus()
		if link != "null" && link != "" {
			isDone = true
			break
		}
		if status != "COMPLETED" {
			time.Sleep(time.Second * time.Duration(waitTime))
			fmt.Println("Sleeping... for ", waitTime)
			if waitTime < maxWaitTime {
				waitTime = waitTime + 2 + (waitTime / 2)
			}
		} else if status == "COMPLETED" {
			isDone = true
			break
		}
	}

	if link == "null" || link == "" {

		log.Println("The querry did not return any results.")
	}
	// fmt.Println("link: ", link)
	return link
}

func ConvertFromShopifyToFikenSales(shopifyOrders []ShopifyOrder, url, token string) []fiken.CreateSale {

	var fikenList []fiken.CreateSale
	var orders = shopifyOrders
	var contacts []fiken.Contact = fiken.GetContacts(url, token)

	//Loops trough all customers in struct
	for i := range orders {
		var lastTransaction int = (len(orders[i].Transactions)) - 1

		//convert shopify orderTotal from string to Float
		var orderTotalPrice int64 = stringPriceToFikenInt(orders[i].OriginalTotalPriceSet.ShopMoney.Amount)
		var orderTotalTaxes int64 = stringPriceToFikenInt(orders[i].TotalTaxSet.ShopMoney.Amount)
		//Calculate NET price needed by
		var orderTotalBeforeTaxes int64 = orderTotalPrice - orderTotalTaxes

		// log.Println("Net price: ", orderTotalBeforeTaxes, " Taxes: ", orderTotalTaxes, " Total: ", orderTotalPrice)

		var DateCreated string = formatDateToFiken(orders[i].DateCreated)
		var PaymentDate string = formatDateToFiken(orders[i].Transactions[lastTransaction].CreatedAt)

		//create temp FikenSale
		temp := fiken.CreateSale{
			SaleNumber:         orders[i].Name,
			Date:               DateCreated,
			Kind:               "external_invoice",
			Settled:            orders[i].FullyPaid,
			TotalPaid:          orderTotalPrice, //convert total price 562,50kr
			OutstandingBalance: 0,
			Lines: []fiken.Line{
				fiken.Line{
					Description: "Sko og tilbehør",
					NetPrice:    orderTotalBeforeTaxes, //convert net price
					Vat:         orderTotalTaxes,       //convert VAT price
					Account:     "3210",
					VatType:     "HIGH",
				},
			},
			// CustomerID: 1710967200, //should fetch users by their email
			Currency: orders[i].CurrencyCode,
			DueDate:  DateCreated,
			// Kid:            "5855454756",
			// PaymentAccount: "1920:10001",
			PaymentDate: PaymentDate,
			// PaymentFee:     0,
		}

		for x := range contacts {
			// log.Println("CustomerEmail: ", contacts[x].ContactEmail, " == ", orders[i].Email)
			var fikenEmail string = strings.ToLower(contacts[x].ContactEmail)
			var shopifyEmail string = strings.ToLower(orders[i].Email)
			if shopifyEmail == fikenEmail {
				temp.CustomerID = contacts[x].ContactID
				// log.Println("CustomerID: ", temp.CustomerID)
				break
			}
		}

		// log.Println("The Final CustomerID: ", temp.CustomerID)
		//set total Amount paid by customer
		// temp.TotalPaid = orderTotalPrice
		//set NetPrice
		// temp.Lines[0].NetPrice = orderTotalBeforeTaxes
		//set VATPrice
		// temp.Lines[0].Vat = orderTotalTaxes

		// for j := range orders[i].Transactions {
		// 	temp.PaymentDate = orders[i].Transactions[j].CreatedAt
		// }

		var gateway = orders[i].Transactions[lastTransaction].Gateway
		if gateway == "vipps" {
			temp.PaymentAccount = "1960:10002"
		} else if gateway == "stripe" {
			temp.PaymentAccount = "1960:10001"
		} else if gateway == "bogus" {
			temp.PaymentAccount = "1920:10001"
		}
		//add created contact to fikenList
		fikenList = append(fikenList, temp)
	}
	return fikenList
}

//converts prices from shopify string format to fiken, int
func stringPriceToFikenInt(s string) int64 {

	//converting string value to float
	f, _ := strconv.ParseFloat(s, 64)

	// formatting float to fit Fiken value format
	// Multiplying with 100 to move the decimal point 2 places and adding 0.5 to help the rounding in conversion
	f = f*100 + 0.5

	// log.Println("Multiplied ", f)

	//converting float to int
	var ftoInt int64 = int64(f)

	// log.Println("FIxed ", ftoInt)
	return ftoInt
}

func formatDateToFiken(s string) string {
	var date string
	if s != "" {
		date = strings.Split(s, "T")[0]
	} else {

		log.Println("Given date string is empty")
	}

	return date
}
