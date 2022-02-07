package lib

func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEFGHIJKLMN"
	var res string
	for dec > 0 {
		remainder := dec % base
		res = string(charset[remainder]) + res
		dec = dec / base
	}
	
	return res
}
