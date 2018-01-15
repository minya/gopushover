package gopushover

import (
	"github.com/minya/goutils/config"
)

func ReadSettings(spath string) (PushoverSettings, error) {
	var settings PushoverSettings

	errUnmarshal := config.UnmarshalJson(&settings, spath)
	if nil != errUnmarshal {
		return settings, errUnmarshal
	}

	return settings, nil
}

type PushoverSettings struct {
	User  string
	Token string
}
