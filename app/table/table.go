package table

import (
	"os"

	"github.com/mitchellh/go-wordwrap"
	"github.com/olekukonko/tablewriter"
)

func Table(res [][]string) {
	// 创建一个新的表格写入器
	table := tablewriter.NewWriter(os.Stdout)

	for _, row := range res {
		similarity := row[0]
		answer := wrapText(row[1], 160) // 将第二个字段固定长度为 160 并换行
		table.Append([]string{similarity, answer})
	}

	// 渲染表格
	table.Render()
}

func wrapText(s string, width int) string {
	wrappedText := wordwrap.WrapString(s, uint(width))
	return wrappedText
}
