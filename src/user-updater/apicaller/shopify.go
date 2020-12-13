package apicaller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	return responseData
}
func ShopifyCallBulkOperation(shopifyName string) {
	// construct query
	var query string = "mutation {" +
		"bulkOperationRunQuery(" +
		"query:" +
		`"""` +
		"{" +
		"customers{" +
		"edges{" +
		"node{" +
		"firstName\n" +
		"lastName\n" +
		"email\n" +
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

	var responseBytes = ShopifyRequest("POST", shopifyName, query) // refactored

	//get bulkoperation struct
	var bulkOperation BulkOperationResponse

	error := json.Unmarshal(responseBytes, &bulkOperation)
	if error != nil {
		fmt.Println("Failed to put into struct : \n", error)
	}
	//Operation ID to keep track of which operation it is
	var operationID = bulkOperation.Data.BulkOperationRunQuery.BulkOperation.ID

	if operationID == "" {
		log.Println("Something went wrong. No operation running.")
		log.Fatalln("Error: ", bulkOperation.Data.BulkOperationRunQuery.UserErrors)
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

func ConvertFromShopifyToFikenContacts(customerlist []ShopifyCustomer) []FikenContact {
	var fikenList []FikenContact
	//Loops trough all customers in struct
	for i := range customerlist {
		//create temp FikenContact
		temp := FikenContact{
			Name:         customerlist[i].FirstName + " " + customerlist[i].LastName,
			ContactEmail: customerlist[i].Email,
			Supplier:     false,
		}
		//add created contact to fikenList
		fikenList = append(fikenList, temp)
	}
	return fikenList
}

func pollOperationStatus(shopifyName string) (string, string) {
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

	var responseBytes = ShopifyRequest("POST", shopifyName, query) //refactored

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
func PollOperation(shopifyName string) string { //refactored

	var isDone = false
	var status string
	var link string
	var waitTime int32 = 1
	var maxWaitTime int32 = 3600

	for !isDone {
		status, link = pollOperationStatus(shopifyName) //refactored
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
		log.Println("The query did not return any results.")
	}
	// fmt.Println("link: ", link)
	return link
}
