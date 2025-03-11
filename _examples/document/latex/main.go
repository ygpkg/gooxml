package main

import (
	"fmt"
	"log"

	"github.com/clearmann/gooxml/document"
)

func main() {
	// 创建新文档
	doc := document.New()

	// 添加段落
	para := doc.AddParagraph()

	// 创建公式运行块
	run := para.AddRun()
	run.Properties().SetMath(true)

	// LaTeX公式示例
	latexFormula := `x^2 + y^2 = r^2`

	// 转换LaTeX为OMML
	ommlFormula, err := document.LatexToOMML(latexFormula)
	if err != nil {
		log.Fatalf("转换LaTeX公式失败: %v", err)
	}

	// 添加转换后的OMML公式
	run.AddText(ommlFormula)

	// 保存文档
	err = doc.SaveToFile("with_latex_formula.docx")
	if err != nil {
		log.Fatalf("保存文档失败: %v", err)
	}

	fmt.Println("文档已成功生成！")
}
