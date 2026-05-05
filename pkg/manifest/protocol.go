package manifest

type Protocol int

const (
	Git Protocol = iota
	Url
	Github
)

var (
	protocolName = map[Protocol]string{
		Git: "Git",
		Url: "Url",
	}
)

func (p Protocol) String() string {
	return protocolName[p]
}
