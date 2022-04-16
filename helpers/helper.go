package helpers

func HnadleError(err error) {
	if err != nil {
		panic(err)
	}
}
