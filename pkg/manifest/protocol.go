package manifest

type Protocol int

const (
	Git Protocol = iota
	Url
	DropBox
)

var (
	protocolName = map[Protocol]string{
		Git:     "Git",
		Url:     "Url",
		DropBox: "DropBox",
	}
)

func (p Protocol) String() string {
	return protocolName[p]
}
