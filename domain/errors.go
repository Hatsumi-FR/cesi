package domain

import "errors"

var (
	attackErr   = errors.New("can't attack without weapon")
	overhealErr = errors.New("can't cast heal spell")
)
