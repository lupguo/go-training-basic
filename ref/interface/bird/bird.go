package bird

type Bird struct {
	Color string
	Size  int
}

func (bird *Bird) Eat() {
	panic("implement me")
}

func (bird *Bird) Laugh() {
	panic("implement me")
}
