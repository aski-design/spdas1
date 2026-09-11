package email

import "fmt"

type Email struct {
	from        string
	to          []string
	cc          []string
	subject     string
	body        string
	attachments []string
}

func newEmail(from string, to []string, cc []string, subject string, body string, attachments []string) *Email {
	return &Email{
		from:        from,
		to:          copyStrings(to),
		cc:          copyStrings(cc),
		subject:     subject,
		body:        body,
		attachments: copyStrings(attachments),
	}
}

func (e *Email) From() string {
	return e.from
}

func (e *Email) To() []string {
	return copyStrings(e.to)
}

func (e *Email) Cc() []string {
	return copyStrings(e.cc)
}

func (e *Email) Subject() string {
	return e.subject
}

func (e *Email) Body() string {
	return e.body
}

func (e *Email) Attachments() []string {
	return copyStrings(e.attachments)
}

func (e *Email) String() string {
	return fmt.Sprintf("Email{from=%s, to=%v, subject=%s}", e.from, e.to, e.subject)
}

func copyStrings(source []string) []string {
	if len(source) == 0 {
		return nil
	}
	destination := make([]string, len(source))
	copy(destination, source)
	return destination
}
