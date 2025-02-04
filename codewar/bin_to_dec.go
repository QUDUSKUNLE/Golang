package codewar

func BinToDec(bin string) int {
	dec := 0
	for i, b := range bin {
		if b == '1' {
			dec += 1 << uint(len(bin)-1-i)
		}
	}
	return dec
}
