package serviceModel

type NetMsg struct {
	Type int
	Open int
	Data interface{}
	Callback chan error
}
