package xerror

func CheckError(err *error, f func() error) {
	if *err != nil {
		return
	}
	*err = f()
}
