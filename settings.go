package gopushover

import (
	"github.com/minya/goutils/config"
	"strings"
)

func ReadSettings(spath string) (PushoverSettings, error) {
	var settings PushoverSettings

	errUnmarshal := config.UnmarshalJson(&settings, strings.TrimLeft(spath, "~/"))
	if nil != errUnmarshal {
		return settings, errUnmarshal
	}

	return settings, nil
}

type PushoverSettings struct {
	User  string
	Token string
}
