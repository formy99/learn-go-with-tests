package iteration

func Repeater(a string, n int) string {
	var repstr string

	for i := 0; i < n; i++ {
		repstr += a
	}

	return repstr
}
