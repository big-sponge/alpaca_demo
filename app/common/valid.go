package common

/**
 * 验证密码格式
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func ValidPassword(input string) (b bool) {
	if len(input) > 32 {
		return false
	}
	return true
}
