package devenv

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

var CWD string
var Username string
var Home string
var OS string
var Arch string
var ExeDir string

func init() {
	CWD, _ = os.Getwd()
	user, _ := user.Current()
	ex, _ := os.Executable()
	ExeDir = filepath.Dir(ex)
	Username = user.Username
	Home = user.HomeDir
	Arch = runtime.GOARCH
	OS = runtime.GOOS
}

//$(PWD) -> getwd
//$(OS) -> runtime.GOOS
//$(ARCH) -> runtime.GOARCH
//$(USER) -> user.Current().Username
//$(HOME) -> user.Current().HomeDir
//$(EDIR) -> Exe's dir
func Replace(str string) (ret string) {
	ret = str
	ret = strings.ReplaceAll(ret, `$(PWD)`, CWD)
	ret = strings.ReplaceAll(ret, `$(OS)`, OS)
	ret = strings.ReplaceAll(ret, `$(ARCH)`, Arch)
	ret = strings.ReplaceAll(ret, `$(USER)`, Username)
	ret = strings.ReplaceAll(ret, `$(HOME)`, Home)
	ret = strings.ReplaceAll(ret, `$(EDIR)`, ExeDir)
	return ret
}
