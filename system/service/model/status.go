package serviceModel

type StatusMsg struct {
	Type int
	Data interface{}
	Callback chan interface{}
}
