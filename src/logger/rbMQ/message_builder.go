package rbMQ

import (
	"fmt"
)

func JobMessageFormatter(userID int, job_type string, apiShopify string, apiFiken string) string{
	return fmt.Sprintf("User ID:%d||Type:%s||API_Shopify:%s||API_Fiken:%s", userID, job_type, apiShopify, apiFiken)
}