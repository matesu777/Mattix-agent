package network

type Network struct {
	Name string `json:"name"`
	IPv4 string `json:"ipv4"`
	MAC  string `json:"mac"`

	RxBytes uint64 `json:"rx_bytes"`
	TxBytes uint64 `json:"tx_bytes"`

	RxSpeed uint64 `json:"rx_speed"`
	TxSpeed uint64 `json:"tx_speed"`

	prevRx uint64
	prevTx uint64
}
