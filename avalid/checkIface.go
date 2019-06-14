package avalid

type checkIface interface {
	Check() (msg string, ok bool)
}
