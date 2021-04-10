//Given two integers representing the numerator and denominator of a fraction, r
//eturn the fraction in string format. 
//
// If the fractional part is repeating, enclose the repeating part in parenthese
//s. 
//
// If multiple answers are possible, return any of them. 
//
// It is guaranteed that the length of the answer string is less than 104 for al
//l the given inputs. 
//
// 
// Example 1: 
// Input: numerator = 1, denominator = 2
//Output: "0.5"
// Example 2: 
// Input: numerator = 2, denominator = 1
//Output: "2"
// Example 3: 
// Input: numerator = 2, denominator = 3
//Output: "0.(6)"
// Example 4: 
// Input: numerator = 4, denominator = 333
//Output: "0.(012)"
// Example 5: 
// Input: numerator = 1, denominator = 5
//Output: "0.2"
// 
// 
// Constraints: 
//
// 
// -231 <= numerator, denominator <= 231 - 1 
// denominator != 0 
// 
// Related Topics 哈希表 数学 
// 👍 220 👎 0
package main

import (
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func fractionToDecimal(numerator int, denominator int) string {
	// case 0, -5
	if numerator == 0 {
		return "0"
	}

	minus := false

	// 先定义数组 然后再转换
	var res []string

	// 1. 判断正负
	if numerator < 0 {
		minus = !minus
		numerator = -numerator
	}

	if denominator < 0 {
		minus = !minus
		denominator = -denominator
	}

	if minus {
		res = append(res, "-")
	}

	// 整数，小数 integer decimal
	// 但是用余数更贴切
	integer := numerator / denominator
	remainder := numerator % denominator

	if remainder == 0 {
		return strings.Join(res, "") + strconv.Itoa(integer)
	}

	res = append(res, strconv.Itoa(integer), ".")
	// 计算小数是否有循环，通过buf记录每次余数出现时，res的当前的位置
	loopStart := -1
	buf := map[int]int{}

	curPosition := len(res)
	for remainder != 0 {
		if v, ok := buf[remainder]; ok {
			loopStart = v
			break
		} else {
			buf[remainder] = curPosition
			remainder *= 10
			curPosition++
		}


		numerator = remainder / denominator
		remainder = remainder % denominator
		res = append(res, strconv.Itoa(numerator))

	}

	if loopStart != -1 {
		// 塞进去 做括号
		res = append(res[:loopStart], append([]string{"("}, res[loopStart:]...)...)
		res = append(res, ")")
	}
	return strings.Join(res, "")
}
//leetcode submit region end(Prohibit modification and deletion)

/**
func fractionToDecimal(numerator int, denominator int) string {

    if numerator == 0 {
        return "0"
    }

    // 处理符号，顺便把numerator和denominator做abs处理
    minus := false
    if (numerator < 0) {
        minus = !minus
        numerator = -numerator
    }
    if (denominator < 0){
        minus = !minus
        denominator = -denominator
    }

    // 处理整数部分
    result := []string{}
    if minus {
        result = append(result, "-")
    }
    v := numerator / denominator
    remainder := numerator % denominator
    result = append(result, strconv.Itoa(v))
    if remainder == 0 {
        return strings.Join(result, "")
    }

    // 处理小数部分
    result = append(result, ".")
    m := map[int]int{} //remainder到pos
    loopStart := -1 // loop开始的位置
    curPos := len(result)
    for remainder > 0 {
        remainder *= 10
        pos, ok := m[remainder]
        if ok {
            loopStart = pos
            break
        } else {
			// 将余数计入map中
			m[remainder] = curPos
        }

        curPos++
        v = remainder / denominator
        result = append(result, strconv.Itoa(v))
        remainder = remainder % denominator
    }

	// 有循环
    if loopStart != -1 {
        // 塞进去一个(
        result = append(result[:loopStart], append([]string{"("}, result[loopStart:]...)...)
        result = append(result, ")")
    }
    return strings.Join(result, "")
}

作者：speedfirst-2
链接：https://leetcode-cn.com/problems/fraction-to-recurring-decimal/solution/golangshuang-bai-jie-fa-by-speedfirst-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 */