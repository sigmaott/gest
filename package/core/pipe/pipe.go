package pipe

type IPipe interface {
	Bind(data any) error
}
