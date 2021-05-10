//Reverse bits of a given 32 bits unsigned integer. 
//
// Note: 
//
// 
// Note that in some languages such as Java, there is no unsigned integer type. 
//In this case, both input and output will be given as a signed integer type. They
// should not affect your implementation, as the integer's internal binary represe
//ntation is the same, whether it is signed or unsigned. 
// In Java, the compiler represents the signed integers using 2's complement not
//ation. Therefore, in Example 2 above, the input represents the signed integer -3
// and the output represents the signed integer -1073741825. 
// 
//
// Follow up: 
//
// If this function is called many times, how would you optimize it? 
//
// 
// Example 1: 
//
// 
//Input: n = 00000010100101000001111010011100
//Output:    964176192 (00111001011110000010100101000000)
//Explanation: The input binary string 00000010100101000001111010011100 represen
//ts the unsigned integer 43261596, so return 964176192 which its binary represent
//ation is 00111001011110000010100101000000.
// 
//
// Example 2: 
//
// 
//Input: n = 11111111111111111111111111111101
//Output:   3221225471 (10111111111111111111111111111111)
//Explanation: The input binary string 11111111111111111111111111111101 represen
//ts the unsigned integer 4294967293, so return 3221225471 which its binary repres
//entation is 10111111111111111111111111111111.
// 
//
// 
// Constraints: 
//
// 
// The input must be a binary string of length 32 
// 
// Related Topics 位运算 
// 👍 390 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32 && num > 0; i++ {
		// TODO: 31 而非32
		// num&1取最后一位
		res |= num&1 << (31 - i)
		num >>= 1
	}
	return res
    
}
//leetcode submit region end(Prohibit modification and deletion)

/**
func reverseBits(n uint32) (rev uint32) {
    for i := 0; i < 32 && n > 0; i++ {
        rev |= n & 1 << (31 - i)
        n >>= 1
    }
    return
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/reverse-bits/solution/dian-dao-er-jin-zhi-wei-by-leetcode-solu-yhxz/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */