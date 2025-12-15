package hades

import (
	"bytes"
	"html/template"
	"regexp"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
	blackfriday "github.com/russross/blackfriday/v2"
)

var TemplateFuncMap = template.FuncMap{
	"sub":            Sub,
	"add":            Add,
	"until":          Until,
	"split":          Split,
	"formatPrice":    FormatPrice,
	"SafeHTML":       SafeHTML,
	"SafeJS":         SafeJS,
	"FormatDate":     FormatDate,
	"FormatDateEn":   FormatDateEn,
	"FormatDateCn":   FormatDateCn,
	"GenDescription": GenDescription,
}

// Split splits a string by a separator
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// Sub returns a - b
func Sub(a, b int) int {
	return a - b
}

// Add returns a + b
func Add(a, b int) int {
	return a + b
}

// Until returns a slice of ints from 0 to n-1
func Until(n int) []int {
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = i
	}
	return out
}

// UnescapeHTML
func UnescapeHTML(s string) template.HTML {
	return template.HTML(s)
}

func SafeHTML(s string) template.HTML {
	return template.HTML(s)
}

func SafeJS(s string) template.JS {
	return template.JS(s)
}

// FirstNChars 字符串的前 N 个字符
// fmt.Println(firstN2("世界 Hello", 1)) 	// 世
func FirstNChars(s string, n int) string {
	r := []rune(s)
	if len(r) > n {
		return string(r[:n])
	}
	return s
}

// GenDescription 文章概要
func GenDescription(html string) string {
	return FirstNChars(strip.StripTags(html), 120)
}

// MarkdownToHTML converts markdown to html
func MarkdownToHTML(markdown string) template.HTML {
	descData := []byte(markdown)
	descData = bytes.Replace(descData, []byte("\r"), nil, -1)
	description := string(blackfriday.Run(
		descData,
		blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.HardLineBreak),
	))

	return template.HTML(description)
}

var slugReplaceRegexp = regexp.MustCompile(`[^a-z0-9\-]+`)

// 将字符串转换为 URL 友好的 slug
// 例如：Some Product -> some-product
func Slugify(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}

	value = strings.ToLower(value)
	value = strings.ReplaceAll(value, "_", "-")
	value = strings.ReplaceAll(value, " ", "-")
	value = slugReplaceRegexp.ReplaceAllString(value, "-")
	value = strings.Trim(value, "-")
	if value == "" {
		value = "category"
	}
	return value
}
