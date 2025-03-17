package main

import (
	"log"

	"github.com/ygpkg/gooxml/color"
	"github.com/ygpkg/gooxml/document"
	"github.com/ygpkg/gooxml/measurement"
	"github.com/ygpkg/gooxml/schema/soo/wml"
)

func main() {
	// 创建一个新的 word 文档
	doc := document.New()

	// 定义一个匿名函数来添加表格行
	addRow := func(table document.Table, colWidths []measurement.Distance, cells ...string) {
		row := table.AddRow()
		for i, cellText := range cells {
			cell := row.AddCell()
			cell.Properties().SetWidth(colWidths[i])
			cell.AddParagraph().AddRun().AddText(cellText)
		}
	}

	// 添加一个新的表格
	table := doc.AddTable()
	// 设置表格宽度为6英寸宽
	table.Properties().SetWidth(6 * measurement.Inch)
	borders := table.Properties().Borders()
	// 设置表格边框为单线，颜色自动，宽度为0
	borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)

	// 设置列宽
	colWidths := []measurement.Distance{2 * measurement.Inch, 4 * measurement.Inch}

	// 使用匿名函数添加表格行
	addRow(table, colWidths, "申请人", "清华大学、张三、华为技术有限公司")
	addRow(table, colWidths, "发明人", "李四、王五、赵六")
	addRow(table, colWidths, "联系电话", "18888888888")
	addRow(table, colWidths, "专利名称", "一种去除图像噪声的方法")
	addRow(table, colWidths, "申请类别", "刘曼")
	addRow(table, colWidths, "技术领域（本专利直接所属的技术领域。）", "本专利涉及图像处理技术领域，尤其涉及一种去除图像噪声的方法。")
	addRow(table, colWidths, "背景技术介绍（与本专利最接近的现有技术。）", "高ISO数的图像在生成时通常会伴随着噪声的产生，去除噪声是用以提高图像的清晰度的常用手段。授权公告号为CN***B 的中国发明专利提供了一种去除图像噪声的方法，是使用平均滤波器来去除图像噪声。平均滤波器能通过对局部区域进行平均化处理，降低图像的变化幅度，从而达到提高图像清晰度的目的。")
	addRow(table, colWidths, "背景技术缺陷（对应由本专利所解决的问题和缺点。）", "然而，当图像同时存在平滑区域和细节区域时，平均滤波器会将细节区域中的像素与平滑区域中的像素一起平均，而使得细节区域的图像越来越模糊，导致图像整体的清晰度较差。")
	addRow(table, colWidths, "核心技术方案（本专利的发明点，据此生成一条独权，只写必要技术特征，如有多条独权请添加。）", "一种去除图像噪声的方法，其特征在于，包括以下步骤：1. 接收待处理图像数据；2. 确定滤波窗口，计算所述滤波窗口内各像素点的综合权重；3. 调整双边滤波器重建目标像素的距离权重值和相似权重值；4. 对每个像素点进行滤波处理，得到去除噪声后的图像。")
	addRow(table, colWidths, "有益效果（实施核心技术方案1能够产生的与现有技术的缺陷相对应的有益效果。）", "本专利通过使用双向滤波器，并且调整双向滤波器重建目标像素的距离权重值及相似权重值，能够有效避免细节区域的图像变得模糊，使得调整后的双向滤波器去除噪声的效果提高，使图像变得更加清晰。")
	addRow(table, colWidths, "具体实施方案（核心技术方案1的相关细节方案及作用效果。）", "例如")

	// 验证文档
	if err := doc.Validate(); err != nil {
		log.Fatalf("error during validation: %s", err)
	}
	// 保存文件
	doc.SaveToFile("disclosure-tables.docx")
}
