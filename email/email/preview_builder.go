package email

import (
	"fmt"
	"strings"
)

const previewHeader = "=== Email Preview ==="
const fieldWidth = 12

type EmailPreviewBuilder struct {
	from        string
	to          []string
	cc          []string
	subject     string
	body        string
	attachments []string
}

func NewEmailPreviewBuilder() *EmailPreviewBuilder {
	return &EmailPreviewBuilder{}
}

func (b *EmailPreviewBuilder) SetFrom(address string) EmailBuilder {
	b.from = address
	return b
}

func (b *EmailPreviewBuilder) AddTo(address string) EmailBuilder {
	b.to = append(b.to, address)
	return b
}

func (b *EmailPreviewBuilder) AddCc(address string) EmailBuilder {
	b.cc = append(b.cc, address)
	return b
}

func (b *EmailPreviewBuilder) SetSubject(subject string) EmailBuilder {
	b.subject = subject
	return b
}

func (b *EmailPreviewBuilder) SetBody(body string) EmailBuilder {
	b.body = body
	return b
}

func (b *EmailPreviewBuilder) AddAttachment(filename string) EmailBuilder {
	b.attachments = append(b.attachments, filename)
	return b
}

func (b *EmailPreviewBuilder) GetResult() (string, error) {
	if err := validateEmailState(b.from, b.to, b.subject); err != nil {
		return "", err
	}
	return b.render(), nil
}

func (b *EmailPreviewBuilder) render() string {
	lines := []string{previewHeader}
	lines = append(lines, b.renderRequiredFields()...)
	lines = append(lines, b.renderOptionalFields()...)
	return strings.Join(lines, "\n")
}

func (b *EmailPreviewBuilder) renderRequiredFields() []string {
	return []string{
		formatField("From", b.from),
		formatField("To", strings.Join(b.to, ", ")),
		formatField("Subject", b.subject),
		formatField("Body", b.body),
	}
}

func (b *EmailPreviewBuilder) renderOptionalFields() []string {
	var lines []string
	if len(b.cc) > 0 {
		lines = append(lines, formatField("Cc", strings.Join(b.cc, ", ")))
	}
	if len(b.attachments) > 0 {
		lines = append(lines, formatField("Attachments", strings.Join(b.attachments, ", ")))
	}
	return lines
}

func formatField(label string, value string) string {
	return fmt.Sprintf("%-*s%s", fieldWidth, label+":", value)
}
