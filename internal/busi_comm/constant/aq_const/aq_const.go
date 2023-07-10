package aq_const

type aqConst struct {
	Pass    int32
	Warning int32
	Reject  int32
}

var AqConst = &aqConst{
	Pass:    1,
	Warning: 2,
	Reject:  3,
}
