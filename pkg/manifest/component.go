package manifest

import (
	"fmt"
	"os"
)

// https://github.com/WeiDUorg/weidu/blob/devel/src/tp.ml#L98
type WeiduComponent struct {
	tpFile        string
	name          string
	lang          int
	component     int
	componentName string
	subComponent  string
	version       string
	*Meta
}

func (w *WeiduComponent) ToLogString() string {
	if w.subComponent != "" {
		// ~DLCMERGER/DLCMERGER.TP2~ #0 #3 // Merge DLC into game -> Merge all available DLCs: 1.7
		return fmt.Sprintf("~%s%c%s~ #%d #%d // %s -> %s: %s", w.name, os.PathSeparator, w.tpFile, w.lang, w.component, w.componentName, w.subComponent, w.version)
	} else {
		// ~BG1UB/BG1UB.TP2~ #0 #14 // Edie, the Merchant League Applicant: v17
		return fmt.Sprintf("~%s%c%s~ #%d #%d // %s: %s", w.name, os.PathSeparator, w.tpFile, w.lang, w.component, w.componentName, w.version)
	}
}
