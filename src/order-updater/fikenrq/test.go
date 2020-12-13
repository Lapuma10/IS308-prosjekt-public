package fikenrq

/*
FIKEN API Response codes
All HTTP codes should be expected with their normal semantics. These are some of the common ones:

200 for successful GET
201 for successful POST where you get a Location-header for the created content
400 when invalid content has be sent (for instance a required field is missing, unexpected fields, wrong format, etc)
401 when the user is not authenticated
403 when the user does not have the proper authorization
404 when the requested content is not found
405 When you are trying a method to a resource which doesn't support it (i.e. DELETE on an account).
415 Wrong media type. we accept application/json only.
*/

func Test() {
	/*
		Test function to test functionality
	*/
	companySlug := ""
	url := "https://api.fiken.no/api/v2/companies/" + companySlug
	token := ""
	var bearer = "Bearer " + token

	var sale = CreateSale{
		SaleNumber:         "10002",
		Date:               "2020-11-26",
		Kind:               "external_invoice",
		Settled:            true,
		TotalPaid:          56250,
		OutstandingBalance: 0,
		Lines: []Line{
			Line{
				Description: "Nikey Shoes",
				NetPrice:    45000,
				Vat:         11250,
				Account:     "3000",
				VatType:     "HIGH",
			},
		},
		CustomerID: 1710967200,
		Currency:   "NOK",
		DueDate:    "2020-11-26",
		// Kid:            "5855454756",
		PaymentAccount: "1920:10001",
		PaymentDate:    "2020-11-26",
		// PaymentFee:     0,
	}

	CreateFikenSale(url, bearer, sale)

	// getSale(url, bearer, "1713471449")

}
