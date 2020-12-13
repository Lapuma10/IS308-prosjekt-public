package shopifyrq

import "time"

/*
Orders struct for shopify orders
ID = Globally unique identifier.
Name = unique identifier for the order, displayed on the order(Faktura nummer)
DisplayFulfillmentStatus = Shows sumarry of fulfillment status
Email = Email address provided by the customer.
FullyPaid = Whether the order has been paid in full.

*/
// type ShopifyOrder struct {
// 	Data struct {
// 		Orders struct {
// 			Edges []struct {
// 				Node struct {
// 					ID          string `json:"id"`
// 					Name        string `json:"name"`
// 					DateCreated string `json:"createdAt"`
// 					FullyPaid   bool   `json:"fullyPaid"`
// 					Email       string `json:"email"`
// 					// LineItems    []LineItem    `json:"lineItems"`
// 					CurrencyCode          string                `json:"currencyCode"`
// 					Transactions          []Transaction         `json:"transactions"`
// 					TotalTaxSet           TotalTaxSet           `json:"totalTaxSet"`
// 					OriginalTotalPriceSet OriginalTotalPriceSet `json:"originalTotalPriceSet"`
// 				} `json:"node"`
// 			} `json:"edges"`
// 		} `json:"orders"`
// 	} `json:"data"`
// }

type ShopifyOrder struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DateCreated string `json:"createdAt"`
	FullyPaid   bool   `json:"fullyPaid"`
	Email       string `json:"email"`
	// LineItems    []LineItem    `json:"lineItems"`
	CurrencyCode          string                `json:"currencyCode"`
	Transactions          []Transaction         `json:"transactions"`
	TotalTaxSet           TotalTaxSet           `json:"totalTaxSet"`
	OriginalTotalPriceSet OriginalTotalPriceSet `json:"originalTotalPriceSet"`
}

//Transaction holds transaction information for each order
type Transaction struct {
	Gateway   string `json:"gateway"`
	CreatedAt string `json:"createdAt"`
}

type OriginalTotalPriceSet struct {
	ShopMoney ShopMoney `json:"shopMoney"`
}

type TotalTaxSet struct {
	ShopMoney ShopMoney `json:"shopMoney"`
}

//LineItem holds each line item in an order
// type LineItem struct {
// 	Name                 string               `json:"name"`
// 	DiscountAllocation   []DiscountAllocation `json:"discountAllocations"`
// 	OriginalUnitPriceSet ShopMoney            `json:"originalUnitPriceSet"`
// 	Quantity             int                  `json:"quantity"`
// 	TaxLines             []TaxLine            `json:"taxLines,omitempty"`
// }

// //DiscountAllocation holds the discount amount for each line item
// type DiscountAllocation struct {
// 	AllocatedAmountSet AllocatedAmountSet `json:"allocatedAmountSet"`
// }

// //AllocatedAmountSet holds the discount amounts in ShopMoney
// type AllocatedAmountSet struct {
// 	ShopMoney ShopMoney `json:"shopMoney"`
// }

//ShopMoney holds amount and currencyCode like "NOK", "USD" etc.
type ShopMoney struct {
	Amount       string `json:"amount"` // SHOULD BE FLOAT, But returns in string
	CurrencyCode string `json:"currencyCode,omitempty"`
}

// //TaxLine holds amount of tax in a PriceSet struct
// type TaxLine struct {
// 	PriceSet       PriceSet `json:"priceSet"`
// 	Rate           float32  `json:"rate"`
// 	RatePercentage float32  `json:"ratePercentage"`
// }

// //PriceSet holds Shopmoney
// type PriceSet struct {
// 	ShopMoney ShopMoney `json:"shopMoney"`
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
