package jxon

func panicIfErr(err error) {
	if err == nil {
		return
	}
	panic(err)
}
