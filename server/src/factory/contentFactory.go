package factory

type ContentRepository interface {
	GetContent(id string) struct{}
	SetContent(id string)
}
