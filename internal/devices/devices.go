package devices

type Device interface {
	Setup(int)
	ShowColors([]uint32)
	Clear()
	Teardown()
}

var Registry = map[string]Device{}
