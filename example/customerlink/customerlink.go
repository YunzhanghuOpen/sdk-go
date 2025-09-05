package customerlink

import (
	"fmt"

	"github.com/YunzhanghuOpen/sdk-go/example/base"
)

// Example 样例
func Example() {
	link, err := base.GetCustomerLink("rsa", "https://www.example.com", "testmemberid")
	if err != nil {
		// 发生异常
		fmt.Println(err)
		return
	}
	// 操作成功
	fmt.Println(link)
}
