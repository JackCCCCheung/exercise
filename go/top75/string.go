package top75

// GcdOfStrings 1071. 字符串的最大公因子
// 对于字符串s 和t，只有在s = t + ... + t（t 自身连接 1 次或多次）时，我们才认定“t 能除尽 s”。
// 给定两个字符串str1和str2。返回 最长字符串x，要求满足x 能除尽 str1 且 x 能除尽 str2 。
// str1 = "ABABAB", str2 = "ABAB" => "AB"
func GcdOfStrings(str1 string, str2 string) string {
	var short, long string
	if len(str1) > len(str2) {
		short = str2
		long = str1
	} else {
		short = str1
		long = str2
	}
	var result = ""
	for i := 0; i < len(short); i++ {
		gcd := findGcd(short, long, i+1)
		if len(gcd) > len(result) {
			result = gcd
		}
	}
	return result
}

func findGcd(short string, long string, index int) string {
	if len(short)%index != 0 || len(long)%index != 0 {
		return ""
	}
	shortCount, longCount := len(short)/index, len(long)/index
	var shortAppend, longAppend string
	indexStr := short[0:index]
	for i := 0; i < shortCount; i++ {
		shortAppend = shortAppend + indexStr
	}
	for i := 0; i < longCount; i++ {
		longAppend = longAppend + indexStr
	}
	if longAppend == long && shortAppend == short {
		return indexStr
	}
	return ""
}
