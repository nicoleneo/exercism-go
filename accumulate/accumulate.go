package accumulate

const testVersion = 1

func Accumulate(arr []string, apply func(string) string) []string {
	var result []string
	for _, e := range arr {
		result = append(result, apply(e))
	}
	return result
}
