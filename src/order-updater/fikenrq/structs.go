package fikenrq

//Contact struct for contact
type Contact struct {
	ContactEmail string `json:"email"`
	Name         string
	Supplier     bool
	ContactID    int64 `json:"contactId,omitempty"`
}

//CreateSale struct for Sales
type CreateSale struct {
	SaleNumber          string `json:"saleNumber"`
	Date                string `json:"date"`
	Kind                string `json:"kind"`
	Settled             bool   `json:"settled"`
	TotalPaid           int64  `json:"totalPaid,omitempty"`
	TotalPaidInCurrency int64  `json:"totalPaidInCurrency,omitempty"`
	OutstandingBalance  int64  `json:"outstandingBalance,omitempty"`
	Lines               []Line `json:"lines"`
	CustomerID          int64  `json:"customerId"`
	Currency            string `json:"currency"`
	DueDate             string `json:"dueDate"`
	Kid                 string `json:"kid,omitempty"`
	PaymentAccount      string `json:"paymentAccount,omitempty"`
	PaymentDate         string `json:"paymentDate,omitempty"`
	PaymentFee          int    `json:"paymentFee,omitempty"`
	ProjectID           int    `json:"projectId,omitempty"`
}

/*
Line struct
*/
type Line struct {
	Description        string `json:"description"`
	NetPrice           int64  `json:"netPrice"`
	Vat                int64  `json:"vat"`
	Account            string `json:"account"`
	VatType            string `json:"vatType"`
	NetPriceInCurrency int64  `json:"netPriceInCurrency,omitempty"`
	VatInCurrency      int64  `json:"vatInCurrency,omitempty"`
	ProjectID          int64  `json:"projectId,omitempty"`
}

/*
GetSale struct for getting a specific sale
*/
type GetSale struct {
	SaleID              int    `json:"saleId"`
	TransactionID       int    `json:"transactionId"`
	SaleNumber          string `json:"saleNumber"`
	Date                string `json:"date"`
	Kind                string `json:"kind"`
	NetAmount           int    `json:"netAmount"`
	VatAmount           int    `json:"vatAmount"`
	Settled             bool   `json:"settled"`
	WriteOff            bool   `json:"writeOff"`
	TotalPaid           int    `json:"totalPaid"`
	TotalPaidInCurrency int    `json:"totalPaidInCurrency"`
	OutstandingBalance  int    `json:"outstandingBalance"`
	Lines               []struct {
		Description        string `json:"description"`
		NetPrice           int    `json:"netPrice"`
		Vat                int    `json:"vat"`
		Account            string `json:"account"`
		VatType            string `json:"vatType"`
		NetPriceInCurrency int    `json:"netPriceInCurrency"`
		VatInCurrency      int    `json:"vatInCurrency"`
		ProjectID          int    `json:"projectId"`
	} `json:"lines"`

	/**
	Customer struct
	*/
	Customer struct {
		ContactID           int    `json:"contactId"`
		Name                string `json:"name"`
		Email               string `json:"email"`
		OrganizationNumber  int64  `json:"organizationNumber"`
		CustomerNumber      int    `json:"customerNumber"`
		CustomerAccountCode string `json:"customerAccountCode"`
		SupplierNumber      int    `json:"supplierNumber"`
		SupplierAccountCode string `json:"supplierAccountCode"`
		Customer            bool   `json:"customer"`
		Supplier            bool   `json:"supplier"`
		ContactPerson       []struct {
			ContactPersonID int    `json:"contactPersonId"`
			Name            string `json:"name"`
			Email           string `json:"email"`
			PhoneNumber     int    `json:"phoneNumber"`
			Address         struct {
				StreetAddress string `json:"streetAddress"`
				City          string `json:"city"`
				PostCode      int    `json:"postCode"`
				Country       string `json:"country"`
			} `json:"address"`
		} `json:"contactPerson"`
		Language string `json:"language"`
		Inactive bool   `json:"inactive"`
		Address  struct {
			StreetAddress string `json:"streetAddress"`
			City          string `json:"city"`
			PostCode      int    `json:"postCode"`
			Country       string `json:"country"`
		} `json:"address"`
		Groups []string `json:"groups"`
	} `json:"customer"`
	Currency     string `json:"currency"`
	DueDate      string `json:"dueDate"`
	Kid          int64  `json:"kid"`
	SalePayments []struct {
		PaymentID   int    `json:"paymentId"`
		Date        string `json:"date"`
		Account     string `json:"account"`
		Amount      int    `json:"amount"`
		AmountInNok int    `json:"amountInNok"`
		Currency    string `json:"currency"`
		Fee         int    `json:"fee"`
	} `json:"salePayments"`

	/**
	Sale Attachments struct
	*/
	SaleAttachments []struct {
		Identifier                                int         `json:"identifier"`
		DownloadURL                               interface{} `json:"downloadUrl"`
		DownloadURLWithFikenNormalUserCredentials interface{} `json:"downloadUrlWithFikenNormalUserCredentials"`
		Comment                                   string      `json:"comment"`
		Type                                      string      `json:"type"`
	} `json:"saleAttachments"`

	PaymentDate struct {
	} `json:"paymentDate"`
}
