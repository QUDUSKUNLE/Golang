package closures

func InitSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
