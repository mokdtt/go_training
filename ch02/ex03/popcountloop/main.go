package popcountloop

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	count := 0
	for i := uint64(0); i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}
