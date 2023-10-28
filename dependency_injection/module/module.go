package module

//go:generate moq -out internal/mocks/writer_mock.go -pkg=mocks . Writer
type Writer interface {
	WriteEmails(emails ...string) error
	Close() error
}

type Module struct {
	w Writer
}

func NewModule(w Writer) *Module {
	return &Module{w: w}
}

func (m *Module) ProcessEmails(emails ...string) error {
	return m.w.WriteEmails(emails...)
}
