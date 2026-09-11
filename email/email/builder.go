package email

type EmailBuilder interface {
	SetFrom(address string) EmailBuilder
	AddTo(address string) EmailBuilder
	AddCc(address string) EmailBuilder
	SetSubject(subject string) EmailBuilder
	SetBody(body string) EmailBuilder
	AddAttachment(filename string) EmailBuilder
}
