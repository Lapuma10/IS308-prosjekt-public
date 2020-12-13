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

func ContactExists(ShopifyList []FikenContact, companySlug string, apiKey string) {

	url := "https://api.fiken.no/api/v2/companies/" + companySlug
	token := apiKey // refactored
	var bearer = "Bearer " + token

	FikenList := getContacts(url, bearer)

	for i := range ShopifyList {
		exists := false
		//if email string is empty, ignore.
		if ShopifyList[i].ContactEmail != "" {
			for j := range FikenList {
				log.Println("Fiken email: ", FikenList[j].ContactEmail)
				if FikenList[j].ContactEmail == ShopifyList[i].ContactEmail {
					exists = true
					log.Println("Fiken email: ", FikenList[j].ContactEmail)
					log.Println("Shopify email: ", ShopifyList[i].ContactEmail)
				}
			}
			if exists == false {
				createContact(url, bearer, ShopifyList[i].Name, ShopifyList[i].ContactEmail, false, true)
				log.Println("This customer was not found in fiken: ", ShopifyList[i].Name, " with Email: ", ShopifyList[i].ContactEmail)
			}
		}
	}
}

func getContacts(url string, bearer string) []FikenContact {
	url = url + "/contacts"
	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	// add authorization header to the req
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	//send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Response error.\n[ERROR]", err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	//log.Println(string(bodyBytes))
	var contact []FikenContact

	// decoding JSON array to
	// the contact array
	err = json.Unmarshal(bodyBytes, &contact)
	if err != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}

	return contact
}
func createContact(link string, bearer string, contactName string, contactEmail string, isSupplier bool, isCustomer bool) {

	link = link + "/contacts"
	type InputValues struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Supplier bool   `json:"supplier"`
		Customer bool   `json:"customer"`
		// CustomerID int64  `json:"customerNumber"`
	}

	fmt.Println("Sending customer data to Fiken")
	inputValues := InputValues{
		Name:     contactName,
		Email:    contactEmail,
		Supplier: isSupplier,
		Customer: isCustomer,
		// CustomerID: customerID,
	}

	jsonReq, err := json.Marshal(inputValues)

	request, err := http.NewRequest("POST", link, bytes.NewBuffer(jsonReq))
	request.Header.Set("Authorization", bearer)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		fmt.Printf("The HTTP request was succesfull.")
	}
	defer response.Body.Close()
}
