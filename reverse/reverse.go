package reverse

//IntReverse 整数反转
//ref: https://app.yinxiang.com/shard/s15/nl/18538491/36adcd5c-aa73-4ac5-bfbb-dbe44d8df3d8
func IntReverse(x int) int {
	reValue := 0
	for {
		//将个位数值与 与前一个数值乘以10 相加
		reValue = reValue*10 + x%10
		//将个位数值剔除
		x = x / 10
		if x == 0 {
			break
		}
	}
	return reValue
}
