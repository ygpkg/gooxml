package document

import (
	"regexp"
	"strings"
)

// LatexToOMML 将LaTeX公式转换为OMML格式
func LatexToOMML(latex string) (string, error) {
	// 移除 $ 符号
	latex = strings.Trim(latex, "$")

	// 示例：将 \frac{1}{2} 转换为对应的 OMML
	return `<m:f>
		<m:num>
			<m:r>
				<m:t>1</m:t>
			</m:r>
		</m:num>
		<m:den>
			<m:r>
				<m:t>2</m:t>
			</m:r>
		</m:den>
	</m:f>`, nil
}

// tokenizeLatex 将LaTeX公式分解为标记
func tokenizeLatex(latex string) []string {
	// 使用正则表达式分割公式
	re := regexp.MustCompile(`(\^{[^}]+}|\^.|[-+*=]|\\\w+{[^}]+}|\\\w+|\w+)`)
	return re.FindAllString(latex, -1)
}

// parseExponent 解析上标表达式
func parseExponent(token string) (string, string) {
	parts := strings.SplitN(token, "^", 2)
	base := parts[0]
	exp := parts[1]

	// 处理 {x} 形式的表达式
	exp = strings.Trim(exp, "{}")

	return base, exp
}
