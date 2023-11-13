package rest

import "fmt"

func generateEmailBody() string {
	return fmt.Sprintf("Hi,\n\nI hope this email finds you well. We would like to place an order with you for the following items.\n\nPlease let us know if you are able to fulfill this order in this price range. If you have any questions or require further information, please don't hesitate to reach out to us.\n\nThank you for your attention to this matter.\n\nBest regards,\n\nUITS Inventory Team")
}
