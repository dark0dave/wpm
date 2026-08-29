package manifest

import (
	"fmt"
	"os"
)

// https://github.com/WeiDUorg/weidu/blob/devel/src/tp.ml#L98
type WeiduComponent struct {
	name          string
	tpFile        string
	lang          uint
	component     uint
	componentName string
	subComponent  string
	version       string
	*Meta
}

func (w *WeiduComponent) ToLogString() (out string) {
	if w.name != "" {
		out = fmt.Sprintf("~%s%c%s~ #%d #%d", w.name, os.PathSeparator, w.tpFile, w.lang, w.component)
	} else {
		out = fmt.Sprintf("~%s~ #%d #%d", w.tpFile, w.lang, w.component)
	}
	if w.componentName != "" {
		out = fmt.Sprintf("%s // %s", out, w.componentName)
	}
	if w.subComponent != "" {
		out = fmt.Sprintf("%s -> %s", out, w.subComponent)
	}
	if w.version != "" {
		out = fmt.Sprintf("%s: %s", out, w.version)
	}
	return out
}
