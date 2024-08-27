package manifest

type Dependency interface {
	GetName() string
	Download(folderPath string) error
}

type DependancyType int

const (
	GIT DependancyType = 0
	URL DependancyType = 1
)
