package email

type EmailObjectBuilder struct {
	from        string
	to          []string
	cc          []string
	subject     string
	body        string
	attachments []string
}

func NewEmailObjectBuilder() *EmailObjectBuilder {
	return &EmailObjectBuilder{}
}

func (b *EmailObjectBuilder) SetFrom(address string) EmailBuilder {
	b.from = address
	return b
}

func (b *EmailObjectBuilder) AddTo(address string) EmailBuilder {
	b.to = append(b.to, address)
	return b
}

func (b *EmailObjectBuilder) AddCc(address string) EmailBuilder {
	b.cc = append(b.cc, address)
	return b
}

func (b *EmailObjectBuilder) SetSubject(subject string) EmailBuilder {
	b.subject = subject
	return b
}

func (b *EmailObjectBuilder) SetBody(body string) EmailBuilder {
	b.body = body
	return b
}

func (b *EmailObjectBuilder) AddAttachment(filename string) EmailBuilder {
	b.attachments = append(b.attachments, filename)
	return b
}

func (b *EmailObjectBuilder) GetResult() (*Email, error) {
	if err := validateEmailState(b.from, b.to, b.subject); err != nil {
		return nil, err
	}
	return newEmail(b.from, b.to, b.cc, b.subject, b.body, b.attachments), nil
}
