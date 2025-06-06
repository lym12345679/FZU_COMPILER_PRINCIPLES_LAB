package lexer_test

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"app/lexer"
	"app/utils/log"
)

const (
	Silent = false
)

func LexerAct(str string) (tokens []lexer.Token, errCount int) {
	l := lexer.NewLexer(strings.NewReader(str))
	for {
		token, err := l.NextToken()
		if err != nil && !errors.Is(err, io.EOF) {
			errCount++
			if !Silent {
				fmt.Println(
					log.Sprintf(log.Argument{FrontColor: log.Red, Highlight: true, Format: "Error: %s", Args: []any{err.Error()}}),
				)
			}
		}
		if token.Type == lexer.EOF {
			break
		}
		if err == nil && !Silent {
			fmt.Printf(
				"(%s, %s, %s)\n",
				log.Sprintf(log.Argument{FrontColor: log.Green, Format: "%s", Args: []any{token.Type.ToString()}}),
				log.Sprintf(log.Argument{FrontColor: log.Yellow, Format: "%s", Args: []any{token.Val}}),
				log.Sprintf(log.Argument{FrontColor: log.Blue, Format: "%s", Args: []any{token.SpecificType().ToString()}}),
			)
		}
		tokens = append(tokens, token)
		if errors.Is(err, io.EOF) {
			break
		}
	}
	return
}

type TestCase struct {
	name               string
	str                string
	raisingErrorAnyWay bool
	errorCount         int
	expectedTokens     []lexer.Token
}

var testCase = []TestCase{
	{
		name: "Hello World",
		str: `
package main

import (
	"fmt"
)

func main() {
	var a int = 1
	fmt.Println("Hello, World!")
}
`,
		expectedTokens: make([]lexer.Token, 0),
	},
	{
		name: "Data Type Judgment",
		str: `
// 整数
0
-1
2147483647
-2147483648
0x1A2B3C4D
0X1a2b3c4d

// 浮点数
0.0
-0.1
3.141592653589793
00.0

// 字符串
""
"Hello, 世界!"
"Escape: \\n \\t \\\""

// 字符
'a'
'\n'
'\''
'中'
`,
		expectedTokens: []lexer.Token{
			{Type: lexer.INTEGER, Val: "0"},
			{Type: lexer.OPERATOR, Val: "-"},
			{Type: lexer.INTEGER, Val: "1"},
			{Type: lexer.INTEGER, Val: "2147483647"},
			{Type: lexer.OPERATOR, Val: "-"},
			{Type: lexer.INTEGER, Val: "2147483648"},
			{Type: lexer.INTEGER, Val: "0x1A2B3C4D"},
			{Type: lexer.INTEGER, Val: "0X1a2b3c4d"},
			{Type: lexer.FLOAT, Val: "0.0"},
			{Type: lexer.OPERATOR, Val: "-"},
			{Type: lexer.FLOAT, Val: "0.1"},
			{Type: lexer.FLOAT, Val: "3.141592653589793"},
			{Type: lexer.FLOAT, Val: "0.0"},
			{Type: lexer.STRING, Val: ""},
			{Type: lexer.STRING, Val: "Hello, 世界!"},
			{Type: lexer.STRING, Val: "Escape: \\n \\t \\\""},
			{Type: lexer.CHAR, Val: "a"},
			{Type: lexer.CHAR, Val: "\n"},
			{Type: lexer.CHAR, Val: "'"},
			{Type: lexer.CHAR, Val: "中"},
		},
	},
	{
		name: "Operator Judgment",
		str:  `/ % = == != < <= > >= && || ++ -- ! & | ^ << >>`,
		expectedTokens: []lexer.Token{
			{Type: lexer.OPERATOR, Val: "/"},
			{Type: lexer.OPERATOR, Val: "%"},
			{Type: lexer.OPERATOR, Val: "="},
			{Type: lexer.OPERATOR, Val: "=="},
			{Type: lexer.OPERATOR, Val: "!="},
			{Type: lexer.OPERATOR, Val: "<"},
			{Type: lexer.OPERATOR, Val: "<="},
			{Type: lexer.OPERATOR, Val: ">"},
			{Type: lexer.OPERATOR, Val: ">="},
			{Type: lexer.OPERATOR, Val: "&&"},
			{Type: lexer.OPERATOR, Val: "||"},
			{Type: lexer.OPERATOR, Val: "++"},
			{Type: lexer.OPERATOR, Val: "--"},
			{Type: lexer.OPERATOR, Val: "!"},
			{Type: lexer.OPERATOR, Val: "&"},
			{Type: lexer.OPERATOR, Val: "|"},
			{Type: lexer.OPERATOR, Val: "^"},
			{Type: lexer.OPERATOR, Val: "<<"},
			{Type: lexer.OPERATOR, Val: ">>"},
		},
	},
	{
		name: "Delimiter Judgment",
		str:  `( ) { } [ ] , ; . :`,
		expectedTokens: []lexer.Token{
			{Type: lexer.DELIMITER, Val: "("},
			{Type: lexer.DELIMITER, Val: ")"},
			{Type: lexer.DELIMITER, Val: "{"},
			{Type: lexer.DELIMITER, Val: "}"},
			{Type: lexer.DELIMITER, Val: "["},
			{Type: lexer.DELIMITER, Val: "]"},
			{Type: lexer.DELIMITER, Val: ","},
			{Type: lexer.DELIMITER, Val: ";"},
			{Type: lexer.DELIMITER, Val: "."},
			{Type: lexer.DELIMITER, Val: ":"},
		},
	},
	{
		name: "Reserved Word Judgment",
		str: `break case chan const continue default defer do else
false for func go goto if import
interface map package range return select
struct switch true type var rune`,
		expectedTokens: []lexer.Token{
			{Type: lexer.RESERVED, Val: "break"},
			{Type: lexer.RESERVED, Val: "case"},
			{Type: lexer.RESERVED, Val: "chan"},
			{Type: lexer.RESERVED, Val: "const"},
			{Type: lexer.RESERVED, Val: "continue"},
			{Type: lexer.RESERVED, Val: "default"},
			{Type: lexer.RESERVED, Val: "defer"},
			{Type: lexer.RESERVED, Val: "do"},
			{Type: lexer.RESERVED, Val: "else"},
			{Type: lexer.RESERVED, Val: "false"},
			{Type: lexer.RESERVED, Val: "for"},
			{Type: lexer.RESERVED, Val: "func"},
			{Type: lexer.RESERVED, Val: "go"},
			{Type: lexer.RESERVED, Val: "goto"},
			{Type: lexer.RESERVED, Val: "if"},
			{Type: lexer.RESERVED, Val: "import"},
			{Type: lexer.RESERVED, Val: "interface"},
			{Type: lexer.RESERVED, Val: "map"},
			{Type: lexer.RESERVED, Val: "package"},
			{Type: lexer.RESERVED, Val: "range"},
			{Type: lexer.RESERVED, Val: "return"},
			{Type: lexer.RESERVED, Val: "select"},
			{Type: lexer.RESERVED, Val: "struct"},
			{Type: lexer.RESERVED, Val: "switch"},
			{Type: lexer.RESERVED, Val: "true"},
			{Type: lexer.RESERVED, Val: "type"},
			{Type: lexer.RESERVED, Val: "var"},
			{Type: lexer.RESERVED, Val: "rune"},
		},
	},
	{
		name: "Identifier Judgment",
		str: `// 合法标识符
a
A
abc
ABC
a1
A1
_abc
_123
变量名
变量123
π`,
		expectedTokens: []lexer.Token{
			{Type: lexer.IDENTIFIER, Val: "a"},
			{Type: lexer.IDENTIFIER, Val: "A"},
			{Type: lexer.IDENTIFIER, Val: "abc"},
			{Type: lexer.IDENTIFIER, Val: "ABC"},
			{Type: lexer.IDENTIFIER, Val: "a1"},
			{Type: lexer.IDENTIFIER, Val: "A1"},
			{Type: lexer.IDENTIFIER, Val: "_abc"},
			{Type: lexer.IDENTIFIER, Val: "_123"},
			{Type: lexer.IDENTIFIER, Val: "变量名"},
			{Type: lexer.IDENTIFIER, Val: "变量123"},
			{Type: lexer.IDENTIFIER, Val: "π"},
		},
	},
	{
		name: "Type Judgment",
		str: `//布尔类型
bool
//有符号整数
int
//浮点数类型
float
//字符串类型
string
//字节类型
byte`,
		expectedTokens: []lexer.Token{
			{Type: lexer.TYPE, Val: "bool"},
			{Type: lexer.TYPE, Val: "int"},
			{Type: lexer.TYPE, Val: "float"},
			{Type: lexer.TYPE, Val: "string"},
			{Type: lexer.TYPE, Val: "byte"},
		},
	},
	{
		name: "Annotation Judgment",
		str: `// 单行注释
/* 多行注释 */
/* 多行注释中包含单行注释
// 这是单行注释
*/

var x = 42 // 变量声明后的注释
// 单行注释中包含代码片段
// var a = 10

/* 多行注释中包含代码片段
func test() {
    var b = 20
}
*/

// 注释中包含特殊字符
// !@#$%^&*()_+-={}[]|:;"'<>,.?/ \n \t \\

// 注释与代码混合
var e = 50 /* 这是一个注释 */ + 10`,
		expectedTokens: []lexer.Token{
			{Type: lexer.RESERVED, Val: "var"},
			{Type: lexer.IDENTIFIER, Val: "x"},
			{Type: lexer.OPERATOR, Val: "="},
			{Type: lexer.INTEGER, Val: "42"},
			{Type: lexer.RESERVED, Val: "var"},
			{Type: lexer.IDENTIFIER, Val: "e"},
			{Type: lexer.OPERATOR, Val: "="},
			{Type: lexer.INTEGER, Val: "50"},
			{Type: lexer.OPERATOR, Val: "+"},
			{Type: lexer.INTEGER, Val: "10"},
		},
	},
	{
		name: "Wrong Identifier Judgment",
		str: `// 非法标识符
// 非法字符
@ #
// 存在非法字符的字符串
abc@ 123#
`,
		expectedTokens: make([]lexer.Token, 0),
		errorCount:     4,
	},
	{
		name: "Wrong Number Judgment",
		str: `// 非法数字
// 非法整数
123abc
0XGHI
0xghi
0x123.456
0x
0X
// 非法浮点数
123.456.789

// 不支持的科学计数法
1e10

// 不支持的八进制数
0777

// 不支持的二进制数
0b1010

// 错误前缀
001
`,
		expectedTokens: make([]lexer.Token, 0),
		errorCount:     11,
	},
	{
		name: "Multiline String Using Double Quotes",
		str: `
"Multi-line:
Line 1
Line 2"
`,
		expectedTokens:     make([]lexer.Token, 0),
		raisingErrorAnyWay: true,
	},
	{
		name: "String Not Closed",
		str: `
// 字符串未闭合
"Hello, World!
`,
		expectedTokens:     make([]lexer.Token, 0),
		raisingErrorAnyWay: true,
	},
	{
		name: "Char Not Closed",
		str: `// 字符未闭合
'a
`,
		expectedTokens:     make([]lexer.Token, 0),
		raisingErrorAnyWay: true,
	},
	{
		name: "Char Too Long(More Than 1 Character)",
		str: `// 字符过长
'abc'
'123'
'abc123'
'中文'
`,
		expectedTokens:     make([]lexer.Token, 0),
		raisingErrorAnyWay: true,
	},
	{
		name: "When Meeting Error, Lexer Should Not Stop",
		str: `// 遇到错误时，词法分析器不应该停止
@var x = 42
`,
		expectedTokens: []lexer.Token{
			{},
			{Type: lexer.IDENTIFIER, Val: "var"},
			{Type: lexer.IDENTIFIER, Val: "x"},
			{Type: lexer.OPERATOR, Val: "="},
			{Type: lexer.INTEGER, Val: "42"},
		},
		raisingErrorAnyWay: true,
	},
	{
		name: "Operator Mixed",
		str: `// 操作符混合时，应当匹配最长的操作符后，停止此次匹配
<=> >=< !=> >-< === ====
`,
		expectedTokens: []lexer.Token{
			{Type: lexer.OPERATOR, Val: "<="},
			{Type: lexer.OPERATOR, Val: ">"},
			{Type: lexer.OPERATOR, Val: ">="},
			{Type: lexer.OPERATOR, Val: "<"},
			{Type: lexer.OPERATOR, Val: "!="},
			{Type: lexer.OPERATOR, Val: ">"},
			{Type: lexer.OPERATOR, Val: ">"},
			{Type: lexer.OPERATOR, Val: "-"},
			{Type: lexer.OPERATOR, Val: "<"},
			{Type: lexer.OPERATOR, Val: "=="},
			{Type: lexer.OPERATOR, Val: "="},
			{Type: lexer.OPERATOR, Val: "=="},
			{Type: lexer.OPERATOR, Val: "=="},
		},
	},
	{
		name:           "Empty Input",
		str:            ``,
		expectedTokens: []lexer.Token{},
	},
	{
		name: "unicode 4 bit",
		str: `// unicode 4 bit
"\u0041"
"\u0042\u0043"
"\u0043\u0044\u0045"
"\u0044\u0045\u0046\u0047"
"\u0048\u0065\u006C\u006C\u006F\u002c\u0020\u0057\u006F\u0072\u006C\u0064\u0021"
`,
		expectedTokens: []lexer.Token{
			{Type: lexer.STRING, Val: "\u0041"},
			{Type: lexer.STRING, Val: "\u0042\u0043"},
			{Type: lexer.STRING, Val: "\u0043\u0044\u0045"},
			{Type: lexer.STRING, Val: "\u0044\u0045\u0046\u0047"},
			{Type: lexer.STRING, Val: "Hello, World!"},
		},
	},
	{
		name: "unicode 8 bit",
		str: `// unicode 8 bit
"\U0001F600"
"\U0001F601\U0001F602"
"\U0001F602\U0001F603\U0001F604"
"\U0001F603\U0001F604\U0001F605\U0001F606"
"\U0001F604\U0001F605\U0001F606\U0001F607\U0001F608"
// 你好，世界！
"\U00004F60\U0000597D\U0000FE50\U00004e16\U0000754c\U0000ff01"
`,
		expectedTokens: []lexer.Token{
			{Type: lexer.STRING, Val: "\U0001F600"},
			{Type: lexer.STRING, Val: "\U0001F601\U0001F602"},
			{Type: lexer.STRING, Val: "\U0001F602\U0001F603\U0001F604"},
			{Type: lexer.STRING, Val: "\U0001F603\U0001F604\U0001F605\U0001F606"},
			{Type: lexer.STRING, Val: "\U0001F604\U0001F605\U0001F606\U0001F607\U0001F608"},
			{Type: lexer.STRING, Val: "\U00004F60\U0000597D\U0000FE50\U00004e16\U0000754c\U0000ff01"},
		},
	},
	{
		name: "unicode 4 bit with error",
		str: `// unicode 4 bit with error
"\u004K"
"\u003"
"\u"`,
		raisingErrorAnyWay: true,
	},
	{
		name: "unicode 8 bit with error",
		str: `// unicode 8 bit with error
"\U0001F60"
"\U0001F60G"
"\U"`,
		raisingErrorAnyWay: true,
	},
	{
		name: "mixed unicode",
		str: `// mixed unicode
"\u0041abcd\U0001F600efghijklmnop\U0001F601qrstuvwxyz\U0001F602"`,
		expectedTokens: []lexer.Token{
			{Type: lexer.STRING, Val: "\u0041abcd\U0001F600efghijklmnop\U0001F601qrstuvwxyz\U0001F602"},
		},
	},
	{
		name: "octal 2 bit",
		str: `// octal 2 bit
"\001"
"\002\003"
"\003\004\005"
"\004\005\006\007"`,
		expectedTokens: []lexer.Token{
			{Type: lexer.STRING, Val: "\001"},
			{Type: lexer.STRING, Val: "\002\003"},
			{Type: lexer.STRING, Val: "\003\004\005"},
			{Type: lexer.STRING, Val: "\004\005\006\007"},
		},
	},
	{
		name: "octal 2 bit with error",
		str: `// octal 2 bit with error
"\01"
"\001\002\003\004\005\006\007\010"
"\0"`,
		raisingErrorAnyWay: true,
	},
	{
		name: "mixed unicode and octal",
		str: `// mixed unicode and octal
"\u0041abcd\001efghijklmnop\U0001F601qrstuvwxyz\003"`,
		expectedTokens: []lexer.Token{
			{Type: lexer.STRING, Val: "\u0041abcd\001efghijklmnop\U0001F601qrstuvwxyz\003"},
		},
	},
	{
		name: "backtick string",
		str: "`" + `
// This is a backtick string
// It can contain any characters, including newlines and quotes
"Hello, World!"
// It will ignore escape sequences like \n, \t, and \"
\n\t\"\a` + "`",
		expectedTokens: []lexer.Token{
			{Type: lexer.STRING, Val: `
// This is a backtick string
// It can contain any characters, including newlines and quotes
"Hello, World!"
// It will ignore escape sequences like \n, \t, and \"
\n\t\"\a`},
		},
	},
}

func TestLexer(t *testing.T) {
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			tokens, errCount := LexerAct(tc.str)
			if tc.raisingErrorAnyWay {
				if errCount == 0 {
					t.Errorf("Expected at least one error, got %d", errCount)
				}
				return
			}
			if errCount != tc.errorCount {
				t.Errorf("Expected %d errors, got %d", tc.errorCount, errCount)
			}
			if len(tc.expectedTokens) == 0 {
				return
			}
			if len(tokens) != len(tc.expectedTokens) {
				t.Errorf("Expected %d tokens, got %d", len(tc.expectedTokens), len(tokens))
				return
			}
			for i, token := range tokens {
				if token.Type != tc.expectedTokens[i].Type || token.Val != tc.expectedTokens[i].Val {
					t.Errorf("Expected token %d to be %v, got %v", i, tc.expectedTokens[i], token)
				}
			}
		})
	}
}
