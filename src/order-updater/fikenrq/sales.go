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

func getSale(link string, bearer string, saleID string) {

	link = link + "/sales" + "/" + saleID

	// Create a new request using http
	req, err := http.NewRequest("GET", link, nil)
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

	var sale GetSale
	// decoding JSON array to
	// the sale struct
	err = json.Unmarshal(bodyBytes, &sale)
	fmt.Println("GETSALE Body: ", sale)
	if err != nil {
		fmt.Println(err)
	}

}

func CreateFikenSale(link string, bearer string, sale CreateSale) int {

	link = link + "/sales"

	jsonReq, err := json.Marshal(sale)
	if err != nil {
		fmt.Println("Failed to parse to JSON: ", err)
	}

	request, err := http.NewRequest("POST", link, bytes.NewBuffer(jsonReq))
	request.Header.Set("Authorization", bearer)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		//fmt.Printf(string(response.StatusCode), ": The HTTP request failed with error %s\n", err)
		log.Println("Error in request to Fiken, when creating a request")
	}
	//##DEBUG
	bodyBytes, err := ioutil.ReadAll(response.Body)
	log.Println("Fiken Response: Response code ", response.StatusCode, ", Response error: ", string(bodyBytes))
	defer response.Body.Close()

	return response.StatusCode
}
