package leetcode

func defangIPaddr(address string) string {

	result := ""
	for i := 0; i < len(address); i++ {

		if address[i] == '.' {

			result = result + "[.]"

		} else {

			result += string(address[i])
		}
	}
	return result
}
