package email

type Service interface {
	SendEmail(subj, text string, to []string, contentType string) error
}
