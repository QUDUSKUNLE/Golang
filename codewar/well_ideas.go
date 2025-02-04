package codewar

/*
For every good kata idea there seem to be quite a few bad ones!
In this kata you need to check the provided array (x) for good ideas 'good' and bad ideas 'bad'. If there are one or two good ideas, return 'Publish!', if there are more than 2 return 'I smell a series!'. If there are no good ideas, as is often the case, return 'Fail!'.
*/

const (
	GOOD = "good"
	PUBLISH = "Publish!"
	FAIL    = "Fail!"
	SERIES  = "I smell a series!"
)

func Well(x []string) string {
	publish := 2
	var result string
	for _, val := range x {
		if val == GOOD {
			publish -= 1
			if publish == -1 {
				return SERIES
			}
		}
	}
	switch publish {
	case 1, 0:
		result = PUBLISH
	case 2:
		result = FAIL
	}
	return result
}
