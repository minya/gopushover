package gopushover

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
)

func ReadSettings(spath string) (PushoverSettings, error) {
	var settings PushoverSettings

	if strings.Index(spath, "~/") == 0 {
		user, _ := user.Current()
		spath = path.Join(user.HomeDir, strings.TrimLeft(spath, "~/"))
	} else if strings.Index(spath, "/") != 0 {
		exe, err := os.Executable()
		if err == nil {
			spath = path.Join(filepath.Dir(exe), spath)
		}
	}

	settingsBin, errRead := ioutil.ReadFile(spath)
	if nil != errRead {
		return settings, errRead
	}

	errUnmarshal := json.Unmarshal(settingsBin, &settings)
	if nil != errUnmarshal {
		return settings, errUnmarshal
	}

	return settings, nil
}

type PushoverSettings struct {
	User  string
	Token string
}
