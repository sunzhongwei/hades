package hades

import "regexp"

// 判断是否包含俄语字符
func HasRussianChars(text string) bool {
	// Use Unicode property for Cyrillic scripts which covers the needed ranges
	// 西里尔字母（俄文：Кириллица，英文：Cyrillic）又称斯拉夫字母，
	// 是以希腊字母为基础创制的拼音文字体系
	// 当前该字母仍为俄罗斯、乌克兰等东欧国家及部分中亚地区的官方文字系统
	pattern := `\p{Cyrillic}`
	matched, _ := regexp.MatchString(pattern, text)
	return matched
}

// 文本中包含 URL 的数量。包含 http/https 链接
func CountURLs(text string) int {
	pattern := `(http|https)://[^\s]+`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(text, -1)
	return len(matches)
}
