package domain

import "errors"

var (
	attackErr = errors.New("can't attack without weapon")
)
