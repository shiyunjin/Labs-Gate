package serviceModel

type BandwidthMsg struct {
	Type int
	Device string
	Callback chan interface{}
}
