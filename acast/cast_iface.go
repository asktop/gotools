package acast

func ToIface(in interface{}) interface{} {
	if err, ok := in.(error); ok {
		return err.Error()
	} else {
		return in
	}
}
