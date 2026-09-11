package main

import (
	"fmt"
	"log"

	"emailbuilder/email"
)

func main() {
	director := email.NewEmailDirector()

	printWelcomeEmail(director)
	printPasswordResetPreview(director)
	printInvoicePreview(director)
	printValidationError()
}

func printWelcomeEmail(director *email.EmailDirector) {
	objectBuilder := email.NewEmailObjectBuilder()
	director.MakeWelcomeEmail(objectBuilder, "jane.doe@example.com")
	welcomeEmail, err := objectBuilder.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(welcomeEmail)

	previewBuilder := email.NewEmailPreviewBuilder()
	director.MakeWelcomeEmail(previewBuilder, "jane.doe@example.com")
	welcomePreview, err := previewBuilder.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(welcomePreview)
}

func printPasswordResetPreview(director *email.EmailDirector) {
	previewBuilder := email.NewEmailPreviewBuilder()
	director.MakePasswordResetEmail(previewBuilder, "john.smith@example.com", "https://app.example.com/reset/abc123")
	resetPreview, err := previewBuilder.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resetPreview)
}

func printInvoicePreview(director *email.EmailDirector) {
	previewBuilder := email.NewEmailPreviewBuilder()
	director.MakeInvoiceEmail(previewBuilder, "billing@example.com", "invoice_2026_09.pdf")
	invoicePreview, err := previewBuilder.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(invoicePreview)
}

func printValidationError() {
	invalidBuilder := email.NewEmailObjectBuilder()
	invalidBuilder.SetFrom("not-an-address")
	if _, err := invalidBuilder.GetResult(); err != nil {
		fmt.Println("validation error:", err)
	}
}
