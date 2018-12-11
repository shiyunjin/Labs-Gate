package serviceModel

type BandwidthMsg struct {
	Type int
	Device string
	Callback chan BandwidthCall
}

type BandwidthCall struct {
	Err error
	Data interface{}
}