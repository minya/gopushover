package gopushover

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
	"path"
	"strings"
)

func ReadSettings(spath string) (PushoverSettings, error) {
	var settings PushoverSettings

	if strings.Index(spath, "~/") == 0 {
		user, _ := user.Current()
		spath = path.Join(user.HomeDir, strings.TrimLeft(spath, "~/"))
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
