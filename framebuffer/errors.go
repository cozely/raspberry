package framebuffer

type pkgerror struct {
	msg   string
	cause error
}

func wrap(msg string, cause error) pkgerror {
	return pkgerror{msg: msg, cause: cause}
}

func (e pkgerror) Unwrap() error {
	return e.cause
}

func (e pkgerror) Error() string {
	return e.msg + ": " + e.cause.Error()
}
