package unary

func Encode(nums []int) []byte {
	var (
		bs    []byte
		b     int
		index uint
	)

	for _, n := range nums {
		var zeros uint = uint(n - 1)
		if index+zeros >= 8 {
			bs = append(bs, byte(b))
			b = 0
		}

		for ; zeros >= 8; zeros = zeros - 8 {
			bs = append(bs, 0)
		}

		index = (index + zeros) % 8
		b = b | (128 >> index)
		index += 1
	}

	bs = append(bs, byte(b))
	return bs
}

func Decode(bs []byte) []int {
	var nums []int
	zeros := 1
	for _, b := range bs {
		n := int(b)
		mask := 1
		for i := 0; i < 8; i++ {
			if (n & mask) == 0 {
				zeros++
			} else {
				nums = append(nums, zeros)
				zeros = 1
			}
			mask <<= 1
		}
	}

	return nums
}
