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
	if proto, ok := protocolName[p]; ok {
		return proto
	}
	return ""
}
