package res

func Panic (err error) {
	if err != nil {
		panic(err)
	}
}