package rbMQ

import (
	"fmt"
)

func JobMessageFormatter(userID int, job_type string, shopify_name string, apiShopify string, company_slug string, apiFiken string) string {
	return fmt.Sprintf("%d:%s:%s:%s:%s:%s", userID, job_type, shopify_name, apiShopify, company_slug, apiFiken)
}

func LogMessageFormatter(userID int, description string) string {
	return fmt.Sprintf("%d:%s", userID, description)
}