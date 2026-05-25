package manifest

type Protocol int

const (
	Git Protocol = iota
	URL
	Github
)

var protocolName = map[Protocol]string{
	Git: "Git",
	URL: "Url",
}

func (p Protocol) String() string {
	return protocolName[p]
}
