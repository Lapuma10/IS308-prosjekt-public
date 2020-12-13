package rbMQ

import (
	"fmt"
)

func JobMessageFormatter(userID int, job_type string, apiShopify string, apiFiken string) string {
	return fmt.Sprintf("User ID:%d||Type:%s||API_Shopify:%s||API_Fiken:%s", userID, job_type, apiShopify, apiFiken)
}

func LogMessageFormatter(userID int, description string) string {
	return fmt.Sprintf("User ID:%d||Description:%s", userID, description)
}