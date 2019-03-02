package uploads

// In bits for bitshift operations like
// 1 << 20 (1 megabyte)
// 5 << 20 (5 megabytes)
const bitsInMegabyte = 20

func megabytes(num int64) int64 {
	return num << bitsInMegabyte
}
