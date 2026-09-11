package email

const systemFromAddress = "no-reply@company.com"
const securityFromAddress = "security@company.com"

type EmailDirector struct{}

func NewEmailDirector() *EmailDirector {
	return &EmailDirector{}
}

func (d *EmailDirector) MakeWelcomeEmail(builder EmailBuilder, recipient string) {
	builder.
		SetFrom(systemFromAddress).
		AddTo(recipient).
		SetSubject("Welcome aboard!").
		SetBody("Thanks for joining us. We are excited to have you on board.")
}

func (d *EmailDirector) MakePasswordResetEmail(builder EmailBuilder, recipient string, resetLink string) {
	builder.
		SetFrom(securityFromAddress).
		AddTo(recipient).
		SetSubject("Password reset request").
		SetBody("Click the link below to reset your password: " + resetLink)
}

func (d *EmailDirector) MakeInvoiceEmail(builder EmailBuilder, recipient string, invoiceFile string) {
	builder.
		SetFrom(systemFromAddress).
		AddTo(recipient).
		SetSubject("Your invoice is ready").
		SetBody("Please find your invoice attached.").
		AddAttachment(invoiceFile)
}
