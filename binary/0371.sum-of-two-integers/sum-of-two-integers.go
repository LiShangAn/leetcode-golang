package leetcode

import "fmt"

func getSum(a int, b int) int {
	/*
		required knowledge:
				0 AND 0         0
				0 AND 1         0
				1 AND 0         0
				1 AND 1         1

				OR：
				0 OR 0          0
				0 OR 1          1
				1 OR 0          1
				1 OR 1          1

				XOR：
				0 XOR 0         0
				0 XOR 1         1
				1 XOR 0         1
				1 XOR 1         0

				NOT：
				NOT 0           1
				NOT 1           0

		analysis:
				a:5  >  101
				b:6  >  110
				a+b  > 1011

				a & b = 100
				a | b = 111
				a ^ b = 011

				1000 + 011
					8     3

				>

				1000
				0011

				0000 << 1 = 0
				1011 = 11


				1 1 = 10 = a & b << 1

				101        011
			+ 110  >  + 1000  >
			(5,6)  >    (3,8) > ... (n,0) => n

			兩個數字做^可以實現沒進位的加法!
			所以需要再考慮 1 1 如何進位，可以做 and << 1 來達成
	*/

	fmt.Printf("%d %d\n", a, b)
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return getSum(a&b<<1, a^b)
}
