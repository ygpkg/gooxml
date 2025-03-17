package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ygpkg/gooxml/color"
	"github.com/ygpkg/gooxml/common"
	"github.com/ygpkg/gooxml/document"
	"github.com/ygpkg/gooxml/measurement"

	"github.com/ygpkg/gooxml/schema/soo/wml"
)

func main() {
	// 生成docx文件
	doc := document.New()
	// 添加 table 的行
	addTextRow := func(table document.Table, colWidths []measurement.Distance, cells ...string) {
		row := table.AddRow()
		for i, cellText := range cells {
			cell := row.AddCell()
			cell.Properties().SetWidth(colWidths[i])
			cell.AddParagraph().AddRun().AddText(cellText)
		}
	}
	// 添加居中的大字体文本
	centeredLargerText := func(content string) {
		para := doc.AddParagraph()
		para.Properties().SetStyle("Title")
		para.Properties().SetAlignment(wml.ST_JcCenter)
		run := para.AddRun()
		run.Properties().SetBold(true)
		run.Properties().SetSize(18.0)
		run.AddText(content)
	}
	// 添加居中的小字体文本
	centeredSmallText := func(content string) {
		para := doc.AddParagraph()
		para.Properties().SetStyle("Title")
		para.Properties().SetAlignment(wml.ST_JcCenter)
		run := para.AddRun()
		run.Properties().SetBold(true)
		run.Properties().SetSize(12.0)
		run.AddText(content)
	}
	// 添加左对齐的大字体文本
	leftLargerText := func(content string) {
		para := doc.AddParagraph()
		para.Properties().SetStyle("Heading1")
		para.Properties().SetAlignment(wml.ST_JcLeft)
		run := para.AddRun()
		run.Properties().SetFontFamily("宋体")
		run.Properties().SetBold(true)
		run.Properties().SetSize(14.0)
		run.AddText(content)
	}
	// 添加左对齐的小字体文本
	leftSmallText := func(content string) {
		para := doc.AddParagraph()
		para.Properties().SetStyle("Heading1")
		para.Properties().SetAlignment(wml.ST_JcCenter)
		run := para.AddRun()
		run.Properties().SetSize(14.0)
		run.AddText(content)
	}
	// 添加正文
	mainText := func(content string) {
		para := doc.AddParagraph()
		para.Properties().SetStyle("Content")
		para.Properties().SetFirstLineIndent(measurement.Inch * 0.5)
		run := para.AddRun()
		run.Properties().SetSize(14.0)
		run.Properties().SetFontFamily("宋体")
		content = strings.ReplaceAll(content, "**", "")
		run.AddText(content)
	}
	// 插入图片
	imageText := func(imgURL string) {
		para := doc.AddParagraph()
		run := para.AddRun()
		img, err := common.ImageFromURL(imgURL)
		if err != nil {
			log.Printf("ImageFromURL error: %v", err)
			return
		}
		// 插入图片到段落
		img1ref, err := doc.AddImage(img)
		inl, err := run.AddDrawingInline(img1ref)
		if err != nil {
			log.Printf("AddDrawingInline error: %v", err)
			return
		}
		inl.SetSize(6*measurement.Inch, 6*measurement.Inch)
	}
	centeredLargerText("交底查新报告")
	nowTime := time.Now().Format("2006-01-02")
	centeredSmallText(fmt.Sprintf("生成时间：%s", nowTime))
	leftLargerText("一、查新输入")
	// 添加一个新的表格
	table := doc.AddTable()
	// 设置表格宽度为6英寸宽
	table.Properties().SetWidth(6 * measurement.Inch)
	borders := table.Properties().Borders()
	// 设置表格边框为单线，颜色自动，宽度为0
	borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)

	// 设置列宽
	colWidths := []measurement.Distance{2 * measurement.Inch, 4 * measurement.Inch}
	addTextRow(table, colWidths, fmt.Sprintf("发明名称：%s", "一种基于工业巡检的目标检测技术"))
	addTextRow(table, colWidths, fmt.Sprintf("全部交底书的文字内容\n(备注：交底书的全部内容)"))
	addTextRow(table, colWidths, "交底书全部内容")
	leftLargerText("二、结论及建议")
	// 如果语义相似度高于60%，则提示
	mainText("查新结果不理想（语义相似度高于60%）：根据系统查新分析，系统认为您的技术方案在新创性方面风险较高，建议再行修改优化，以期获得新创性的突破。")
	// mainText("查新结果正向（语义相似度低于60%）：根据系统查新分析，系统认为您的技术方案在新创性方面突破性较强，请点击“查新结果详情”查看具体信息。建议您合理规划权利要求书，以期获得更合理的保护范围；并提供更多具体实施例，以确保权利要求书得得说明书的支持。")
	leftLargerText("三、查新结果（仅供参考）")
	leftSmallText("1.全部交底书文字内容的查新结果")
	leftSmallText("a.统计分析指标")

	// 添加一个新的表格
	table = doc.AddTable()
	// 设置表格宽度为6英寸宽
	table.Properties().SetWidth(6 * measurement.Inch)
	borders = table.Properties().Borders()
	// 设置表格边框为单线，颜色自动，宽度为0
	borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)

	// 设置列宽
	colWidths = []measurement.Distance{2 * measurement.Inch, 4 * measurement.Inch}
	addTextRow(table, colWidths, "TOP1文件的全文语义相似度【中文专利范围】")
	addTextRow(table, colWidths, "XX%")
	addTextRow(table, colWidths, "TOP1文件的技术领域相似度【中文专利范围】")
	addTextRow(table, colWidths, "XX%")
	addTextRow(table, colWidths, "TOP1文件的背景技术相似度【中文专利范围】")
	addTextRow(table, colWidths, "XX%")
	addTextRow(table, colWidths, "TOP1文件的核心技术方案相似度【中文专利范围】")
	addTextRow(table, colWidths, "XX%")
	addTextRow(table, colWidths, "TOP1文件的有益效果相似度【中文专利范围】")
	addTextRow(table, colWidths, "XX%")
	leftSmallText("b.TOP10专利列表【中文专利范围】")

	// 添加一个新的表格
	table = doc.AddTable()
	// 设置表格宽度为6英寸宽
	table.Properties().SetWidth(6 * measurement.Inch)
	borders = table.Properties().Borders()
	// 设置表格边框为单线，颜色自动，宽度为0
	borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)

	// 设置列宽
	colWidths = []measurement.Distance{measurement.Inch * 1, measurement.Inch * 1, measurement.Inch * 1.5,
		measurement.Inch * 2, measurement.Inch * 1.5, measurement.Inch * 1.5,
		measurement.Inch * 1.5}
	addTextRow(table, colWidths, "序号", "相似度", "公开号", "发明名称", "申请人", "公开日", "专利有效性")
	addTextRow(table, colWidths, "1", "61.38%", "CN116681646A", "基于Yolov5的融合空间信息多头预测小目标检测算法", "重庆理工大学", "2023-09-01", "在审")
	addTextRow(table, colWidths, "2", "60.54%", "CN116612065A", "一种基于YOLO v5的输电线路巡检图像缺陷智能识别方法", "贵州电网有限责任公司", "2023-08-18", "在审")
	addTextRow(table, colWidths, "3", "59.87%", "CN110200598B", "一种大型养殖场体征异常禽类检测系统及检测方法", "天津大学", "2020-06-30", "有效")
	addTextRow(table, colWidths, "4", "59.02%", "CN113159278A", "一种分割网络系统", "无锡信捷电气股份有限公司", "2021-07-23", "在审")
	addTextRow(table, colWidths, "5", "58.50%", "CN116310796A", "一种基于深度学习模型的目标图像增广及作物病害识别方法", "上海兰桂骐技术发展股份有限公司", "2023-06-23", "在审")
	addTextRow(table, colWidths, "6", "58.33%", "CN116611478A", "一种基于深度阈值生成对抗网络的工业过程数据增强方法", "西北师范大学", "2023-08-18", "在审")
	addTextRow(table, colWidths, "7", "57.12%", "CN115457322A", "基于深度学习的可回收垃圾识别分类系统", "中国计量大学", "2022-12-09", "失效")
	addTextRow(table, colWidths, "8", "56.79%", "CN117171700B", "一种基于深度学习的钻井溢流预测组合模型及模型适时静默更新与迁移学习方法", "辽宁石油化工大学", "2024-06-28", "有效")
	addTextRow(table, colWidths, "9", "56.78%", "CN116014876A", "配网电缆沟索道式微型巡检机器人", "国网河南省电力公司西峡县供电公司", "2023-04-25", "在审")
	addTextRow(table, colWidths, "10", "56.47%", "CN114529817A", "基于注意力神经网络的无人机光伏故障诊断及定位方法", "东南大学", "2022-05-24", "在审")
	leftLargerText("四、免责声明（仅供参考）")
	leftSmallText("xxxxxxxxxxxxx")
	leftLargerText("五、附件")
	leftSmallText("1. 输入为“全部交底书的文字内容”的中国专利TOP10详情")
	leftSmallText("a. TOP10 专利详情【中文专利范围】")
	leftSmallText("1）TOP1")
	leftSmallText("1. 基础信息")

	// 添加一个新的表格
	table = doc.AddTable()
	// 设置表格宽度为6英寸宽
	table.Properties().SetWidth(6 * measurement.Inch)
	borders = table.Properties().Borders()
	// 设置表格边框为单线，颜色自动，宽度为0
	borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)

	// 设置列宽
	colWidths = []measurement.Distance{7 * measurement.Inch}
	addTextRow(table, colWidths, "公开号")
	addTextRow(table, colWidths, "")
	addTextRow(table, colWidths, "发明名称")
	addTextRow(table, colWidths, "")
	addTextRow(table, colWidths, "公开日")
	addTextRow(table, colWidths, "")
	addTextRow(table, colWidths, "申请日")
	addTextRow(table, colWidths, "")
	addTextRow(table, colWidths, "法律状态")
	addTextRow(table, colWidths, "")
	addTextRow(table, colWidths, "申请人")
	addTextRow(table, colWidths, "")
	addTextRow(table, colWidths, "发明人")
	addTextRow(table, colWidths, "")
	addTextRow(table, colWidths, "当前专利权人")
	addTextRow(table, colWidths, "")
	addTextRow(table, colWidths, "代理机构")
	addTextRow(table, colWidths, "")
	addTextRow(table, colWidths, "IPC 分类号")
	addTextRow(table, colWidths, "")

	{
		leftSmallText("2. 技术特征总结")
		// 添加一个新的表格
		table = doc.AddTable()
		// 设置表格宽度为6英寸宽
		table.Properties().SetWidth(6 * measurement.Inch)
		borders = table.Properties().Borders()
		// 设置表格边框为单线，颜色自动，宽度为0
		borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)

		// 设置列宽
		colWidths = []measurement.Distance{7 * measurement.Inch}
		// 技术问题、技术原理、技术功效都为AI总结里边的内容
		addTextRow(table, colWidths, "技术问题")
		addTextRow(table, colWidths, "")
		addTextRow(table, colWidths, "技术原理")
		addTextRow(table, colWidths, "")
		addTextRow(table, colWidths, "技术功效")
		addTextRow(table, colWidths, "")
		addTextRow(table, colWidths, "核心方案")
		addTextRow(table, colWidths, "")
		addTextRow(table, colWidths, "有益效果")
		addTextRow(table, colWidths, "")
		addTextRow(table, colWidths, "说明书摘要")
		addTextRow(table, colWidths, "")
		addTextRow(table, colWidths, "摘要附图")
		imageText("")
	}
	{
		leftSmallText("3. 具体相似语句对比")
		// 添加一个新的表格
		table = doc.AddTable()
		// 设置表格宽度为6英寸宽
		table.Properties().SetWidth(6 * measurement.Inch)
		borders = table.Properties().Borders()
		// 设置表格边框为单线，颜色自动，宽度为0
		borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)

		// 设置列宽
		colWidths = []measurement.Distance{2 * measurement.Inch, 2 * measurement.Inch,
			2 * measurement.Inch, 1 * measurement.Inch}
		// 技术问题、技术原理、技术功效都为AI总结里边的内容
		addTextRow(table, colWidths, "对比模块", "您的输入", "对比文件中的对应部分", "相似度")
		addTextRow(table, colWidths, "技术领域", "技术领域输入", "技术领域对比文件", "65.46%")
		addTextRow(table, colWidths, "背景技术", "背景技术输入", "背景技术对比文件", "65.46%")
		addTextRow(table, colWidths, "核心技术方案", "核心技术方案输入", "核心技术方案对比文件", "65.46%")
		addTextRow(table, colWidths, "有益效果", "有益效果输入", "有益效果对比文件", "65.46%")
	}
	if err := doc.Validate(); err != nil {
		log.Fatalf("error during validation: %s", err)
	}
	doc.SaveToFile("tables.docx")
}
