package hades

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// 将价格格式化为 xx,xx.yy 的形式
func FormatPrice(price float64) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%.2f", price)
}
