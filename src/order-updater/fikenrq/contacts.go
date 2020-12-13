package fikenrq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetContacts(url string, bearer string) []Contact {

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
		log.Println("Response error.\n[ERRO]", err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	//log.Println(string(bodyBytes))
	var contact []Contact

	// decoding JSON array to
	// the contact array
	err = json.Unmarshal(bodyBytes, &contact)
	if err != nil {
		fmt.Println(err)
	}
	// for i := range contact {
	// 	if !contact[i].Supplier {
	// 		fmt.Println("Contact Email: ", contact[i].ContactEmail, "\nContact Name: ", contact[i].Name)
	// 	}
	// }
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
