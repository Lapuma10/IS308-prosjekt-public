package apicaller

import "time"

//FikenContact struct for customer data
//from Fiken
type FikenContact struct {
	ContactEmail string `json:"email"`
	Name         string `json:"name"`
	Supplier     bool   `json:"supplier"`
}

//ShopifyCustomer struct for customer data
//from shopify
type ShopifyCustomer struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// type CustomerList struct {
// 	Data struct {
// 		Customers struct {
// 			Edges []struct {
// 				Node struct {
// 					FirstName string `json:"firstName"`
// 					LastName  string `json:"lastName"`
// 					Email     string `json:"email"`
// 				} `json:"node"`
// 			} `json:"edges"`
// 		} `json:"customers"`
// 	} `json:"data"`
// }

//BulkOperationResponse struct is used for retrieving such a request
type BulkOperationResponse struct {
	Data struct {
		BulkOperationRunQuery struct {
			BulkOperation struct {
				ID     string `json:"id"`
				Status string `json:"status"`
			} `json:"bulkOperation"`
			UserErrors []interface{} `json:"userErrors"`
		} `json:"bulkOperationRunQuery"`
	} `json:"data"`
	Extensions struct {
		Cost struct {
			RequestedQueryCost int `json:"requestedQueryCost"`
			ActualQueryCost    int `json:"actualQueryCost"`
			ThrottleStatus     struct {
				MaximumAvailable   float64 `json:"maximumAvailable"`
				CurrentlyAvailable int     `json:"currentlyAvailable"`
				RestoreRate        float64 `json:"restoreRate"`
			} `json:"throttleStatus"`
		} `json:"cost"`
	} `json:"extensions"`
}

//BulkOperationStatus struct is used for retrieving
//the status and url when bulk operation is done
type BulkOperationStatus struct {
	Data struct {
		CurrentBulkOperation struct {
			ID             string      `json:"id"`
			Status         string      `json:"status"`
			ErrorCode      interface{} `json:"errorCode"`
			CreatedAt      time.Time   `json:"createdAt"`
			CompletedAt    time.Time   `json:"completedAt"`
			ObjectCount    string      `json:"objectCount"`
			FileSize       string      `json:"fileSize"`
			URL            string      `json:"url"`
			PartialDataURL interface{} `json:"partialDataUrl"`
		} `json:"currentBulkOperation"`
	} `json:"data"`
	Extensions struct {
		Cost struct {
			RequestedQueryCost int `json:"requestedQueryCost"`
			ActualQueryCost    int `json:"actualQueryCost"`
			ThrottleStatus     struct {
				MaximumAvailable   float64 `json:"maximumAvailable"`
				CurrentlyAvailable int     `json:"currentlyAvailable"`
				RestoreRate        float64 `json:"restoreRate"`
			} `json:"throttleStatus"`
		} `json:"cost"`
	} `json:"extensions"`
}
