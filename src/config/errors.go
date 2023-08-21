package config

import "errors"

var ErrFailedToLoadEnv = errors.New("failed to load env file")
var ErrConvertTimeout = errors.New("failed to convert timeout value to integer")
