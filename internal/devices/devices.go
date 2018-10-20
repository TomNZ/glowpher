package devices

type Device interface {
	Setup(int) error
	ShowColors([]uint32) error
	Clear() error
	Teardown()
}

var Registry = map[string]Device{}
