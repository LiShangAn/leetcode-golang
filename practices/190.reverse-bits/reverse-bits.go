package leetcode

func reverseBits(num uint32) uint32 {

	var result uint32 = 0

	for i := 0; i < 32; i++ { // i 0~31
		result += (num >> i) & 1 << (31 - i)
	}

	return result
}
