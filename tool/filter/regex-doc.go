package filter

import (
	"regexp"
)

// /http://www.cnblogs.com/golove/p/3269099.html

//**********************************************************//
// 复合：
//         xy             匹配 xy（x 后面跟随 y）
//         x|y            匹配 x 或 y (优先匹配 x)

//**********************************************************//
// 重复：
//         x*             匹配零个或多个 x，优先匹配更多(贪婪)
//         x+             匹配一个或多个 x，优先匹配更多(贪婪)
//         x?             匹配零个或一个 x，优先匹配一个(贪婪)
//         x{n,m}         匹配 n 到 m 个 x，优先匹配更多(贪婪)
//         x{n,}          匹配 n 个或多个 x，优先匹配更多(贪婪)
//         x{n}           只匹配 n 个 x
//         x*?            匹配零个或多个 x，优先匹配更少(非贪婪)
//         x+?            匹配一个或多个 x，优先匹配更少(非贪婪)
//         x??            匹配零个或一个 x，优先匹配零个(非贪婪)
//         x{n,m}?        匹配 n 到 m 个 x，优先匹配更少(非贪婪)
//         x{n,}?         匹配 n 个或多个 x，优先匹配更少(非贪婪)
//         x{n}?          只匹配 n 个 x
//**********************************************************//

// 可以将“命名字符类”作为“字符类”的元素：

//         [\d]           匹配数字 (相当于 \d)
//         [^\d]          匹配非数字 (相当于 \D)
//         [\D]           匹配非数字 (相当于 \D)
//         [^\D]          匹配数字 (相当于 \d)
//         [[:name:]]     命名的“ASCII 类”包含在“字符类”中 (相当于 [:name:])
//         [^[:name:]]    命名的“ASCII 类”不包含在“字符类”中 (相当于 [:^name:])
//         [\p{Name}]     命名的“Unicode 类”包含在“字符类”中 (相当于 \p{Name})
//         [^\p{Name}]    命名的“Unicode 类”不包含在“字符类”中 (相当于 \P{Name})
//**********************************************************//

// 转义序列：

//         \a             匹配响铃符    （相当于 \x07）
//                        注意：正则表达式中不能使用 \b 匹配退格符，因为 \b 被用来匹配单词边界，
//                        可以使用 \x08 表示退格符。
//         \f             匹配换页符    （相当于 \x0C）
//         \t             匹配横向制表符（相当于 \x09）
//         \n             匹配换行符    （相当于 \x0A）
//         \r             匹配回车符    （相当于 \x0D）
//         \v             匹配纵向制表符（相当于 \x0B）
//         \123           匹配 8  進制编码所代表的字符（必须是 3 位数字）
//         \x7F           匹配 16 進制编码所代表的字符（必须是 3 位数字）
//         \x{10FFFF}     匹配 16 進制编码所代表的字符（最大值 10FFFF  ）
//         \Q...\E        匹配 \Q 和 \E 之间的文本，忽略文本中的正则语法

//         \\             匹配字符 \
//         \^             匹配字符 ^
//         \$             匹配字符 $
//         \.             匹配字符 .
//         \*             匹配字符 *
//         \+             匹配字符 +
//         \?             匹配字符 ?
//         \{             匹配字符 {
//         \}             匹配字符 }
//         \(             匹配字符 (
//         \)             匹配字符 )
//         \[             匹配字符 [
//         \]             匹配字符 ]
//         \|             匹配字符 |
//**********************************************************//

// “Perl 类”取值如下：

//     \d             数字 (相当于 [0-9])
//     \D             非数字 (相当于 [^0-9])
//     \s             空白 (相当于 [\t\n\f\r ])
//     \S             非空白 (相当于[^\t\n\f\r ])
//     \w             单词字符 (相当于 [0-9A-Za-z_])
//     \W             非单词字符 (相当于 [^0-9A-Za-z_])
//**********************************************************//

// “ASCII 类”取值如下

//     [:alnum:]      字母数字 (相当于 [0-9A-Za-z])
//     [:alpha:]      字母 (相当于 [A-Za-z])
//     [:ascii:]      ASCII 字符集 (相当于 [\x00-\x7F])
//     [:blank:]      空白占位符 (相当于 [\t ])
//     [:cntrl:]      控制字符 (相当于 [\x00-\x1F\x7F])
//     [:digit:]      数字 (相当于 [0-9])
//     [:graph:]      图形字符 (相当于 [!-~])
//     [:lower:]      小写字母 (相当于 [a-z])
//     [:print:]      可打印字符 (相当于 [ -~] 相当于 [ [:graph:]])
//     [:punct:]      标点符号 (相当于 [!-/:-@[-反引号{-~])
//     [:space:]      空白字符(相当于 [\t\n\v\f\r ])
//     [:upper:]      大写字母(相当于 [A-Z])
//     [:word:]       单词字符(相当于 [0-9A-Za-z_])
//     [:xdigit:]     16 進制字符集(相当于 [0-9A-Fa-f])

//**********************************************************//

//**********************************************************//
// “Unicode 类”取值如下---普通类：

//     C                 -其他-          (other)
//     Cc                控制字符        (control)
//     Cf                格式            (format)
//     Co                私人使用区      (private use)
//     Cs                代理区          (surrogate)
//     L                 -字母-          (letter)
//     Ll                小写字母        (lowercase letter)
//     Lm                修饰字母        (modifier letter)
//     Lo                其它字母        (other letter)
//     Lt                首字母大写字母  (titlecase letter)
//     Lu                大写字母        (uppercase letter)
//     M                 -标记-          (mark)
//     Mc                间距标记        (spacing mark)
//     Me                关闭标记        (enclosing mark)
//     Mn                非间距标记      (non-spacing mark)
//     N                 -数字-          (number)
//     Nd                十進制数字      (decimal number)
//     Nl                字母数字        (letter number)
//     No                其它数字        (other number)
//     P                 -标点-          (punctuation)
//     Pc                连接符标点      (connector punctuation)
//     Pd                破折号标点符号  (dash punctuation)
//     Pe                关闭的标点符号  (close punctuation)
//     Pf                最后的标点符号  (final punctuation)
//     Pi                最初的标点符号  (initial punctuation)
//     Po                其他标点符号    (other punctuation)
//     Ps                开放的标点符号  (open punctuation)
//     S                 -符号-          (symbol)
//     Sc                货币符号        (currency symbol)
//     Sk                修饰符号        (modifier symbol)
//     Sm                数学符号        (math symbol)
//     So                其他符号        (other symbol)
//     Z                 -分隔符-        (separator)
//     Zl                行分隔符        (line separator)
//     Zp                段落分隔符      (paragraph separator)
//     Zs                空白分隔符      (space separator)
//**********************************************************//

// “Unicode 类”取值如下---脚本类：

//     Arabic                  阿拉伯文
//     Armenian                亚美尼亚文
//     Balinese                巴厘岛文
//     Bengali                 孟加拉文
//     Bopomofo                汉语拼音字母
//     Braille                 盲文
//     Buginese                布吉文
//     Buhid                   布希德文
//     Canadian_Aboriginal     加拿大土著文
//     Carian                  卡里亚文
//     Cham                    占族文
//     Cherokee                切诺基文
//     Common                  普通的，字符不是特定于一个脚本
//     Coptic                  科普特文
//     Cuneiform               楔形文字
//     Cypriot                 塞浦路斯文
//     Cyrillic                斯拉夫文
//     Deseret                 犹他州文
//     Devanagari              梵文
//     Ethiopic                衣索比亚文
//     Georgian                格鲁吉亚文
//     Glagolitic              格拉哥里文
//     Gothic                  哥特文
//     Greek                   希腊
//     Gujarati                古吉拉特文
//     Gurmukhi                果鲁穆奇文
//     Han                     汉文
//     Hangul                  韩文
//     Hanunoo                 哈鲁喏文
//     Hebrew                  希伯来文
//     Hiragana                平假名（日语）
//     Inherited               继承前一个字符的脚本
//     Kannada                 坎那达文
//     Katakana                片假名（日语）
//     Kayah_Li                克耶字母
//     Kharoshthi              卡罗须提文
//     Khmer                   高棉文
//     Lao                     老挝文
//     Latin                   拉丁文
//     Lepcha                  雷布查文
//     Limbu                   林布文
//     Linear_B                B类线形文字（古希腊）
//     Lycian                  利西亚文
//     Lydian                  吕底亚文
//     Malayalam               马拉雅拉姆文
//     Mongolian               蒙古文
//     Myanmar                 缅甸文
//     New_Tai_Lue             新傣仂文
//     Nko                     Nko文
//     Ogham                   欧甘文
//     Ol_Chiki                桑塔利文
//     Old_Italic              古意大利文
//     Old_Persian             古波斯文
//     Oriya                   奥里亚文
//     Osmanya                 奥斯曼亚文
//     Phags_Pa                八思巴文
//     Phoenician              腓尼基文
//     Rejang                  拉让文
//     Runic                   古代北欧文字
//     Saurashtra              索拉什特拉文（印度县城）
//     Shavian                 萧伯纳文
//     Sinhala                 僧伽罗文
//     Sundanese               巽他文
//     Syloti_Nagri            锡尔赫特文
//     Syriac                  叙利亚文
//     Tagalog                 塔加拉文
//     Tagbanwa                塔格巴努亚文
//     Tai_Le                  德宏傣文
//     Tamil                   泰米尔文
//     Telugu                  泰卢固文
//     Thaana                  塔安那文
//     Thai                    泰文
//     Tibetan                 藏文
//     Tifinagh                提非纳文
//     Ugaritic                乌加里特文
//     Vai                     瓦伊文
//     Yi                      彝文
//**********************************************************//

//  to suit "." must use \.
var (
	//(http://)(.)*?(baidu)(.)*?(\.com)
	//(http://)[\d|\D]*?(baidu)[\d|\D]*?(.com)  wrong
	ReTest = regexp.MustCompile(`(http://|https://)(.)*?((?i)test)(.)*?(VIMI)`)

	//*********************************************************//
	ReHtml  = regexp.MustCompile(`(https|http)://(\w|[[:punct:]])*?(\.html)`)
	ReTitle = regexp.MustCompile(`<title>[\d|\D]*?(</title>)`)
)

type VimiRegexp struct {
	BeginWith   []string
	MustContain []string
	EndWith     []string
}

func (s *VimiRegexp) GetRegexp() *regexp.Regexp {
	str := ""
	str += s.SliceToString(s.BeginWith)
	str += "(.)*?"
	str += s.SliceToString(s.MustContain)
	str += "(.)*?"
	str += s.SliceToString(s.EndWith)

	// if need debug
	// fmt.Println(str)

	return regexp.MustCompile(str)
}

func (s *VimiRegexp) SliceToString(args []string) string {
	ans := "((?i)" // ignore capital or lower case
	for j, str := range args {
		for i := 0; i < len(str); i++ {
			switch str[i] {
			case '\\', '^', '$', '.', '*', '+', '?', '{', '}', '(', ')', '[', ']', '|':
				ans += `\`
				ans += string(str[i])
			default:
				ans += string(str[i])
			}
		}
		if j != len(args)-1 {
			ans += "|"
		}
	}
	ans += ")"
	return ans
}
