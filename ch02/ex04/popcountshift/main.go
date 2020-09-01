package popcountshift

func PopCount(x uint64) int {
	count := 0
	for i := uint64(0); i < 64; i++ {
		count += int(x & 1)
		x = x >> 1
	}
	return count
}
