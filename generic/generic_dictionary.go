package generic

type Dic interface {
	int | float32
}

func Add[T Dic](m, n T) T {
	return m + n
}
