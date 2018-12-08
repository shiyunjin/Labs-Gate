package serviceModel

type Channel struct {
	NetworkCh chan NetMsg
	StatusCh chan StatusMsg
	Bandwidthch chan BandwidthMsg
}
