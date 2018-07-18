package main

import (
	"fmt"
)

func main() {
	/* 数组长度为 5 */
	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage(balance, 5)

	/* 输出返回的平均值 */
	fmt.Printf("平均值为: %f ", avg)

	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	var ptr *int
	var pptr **int

	fmt.Printf("最大值是 : %d\n", ret)

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

	type Books struct {
		title   string
		author  string
		subject string
		book_id int
	}
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)

	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)

	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	var countryCapitalMap map[string]string
	/* 创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	/* 删除元素 */
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")
	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	var i int = 15
	fmt.Printf("%d 的阶乘是：%d\n", i, Factorial(i))

	var ii int
	for ii = 0; ii < 10; ii++ {
		fmt.Printf("%d\t", Fibonaci(ii))
	}

	fmt.Println("\n")
	fmt.Println("接口测试！")
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(Iphone)
	phone.call()

	//正常情况下
	if result, errorMsg := Divide(100, 10); errorMsg == "NoError" {
		fmt.Println("100/10=", result)
	}
	//当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "NoError" {
		fmt.Println("errorEsg is:", errorMsg)
	}

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

/*阶乘*/
func Factorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Factorial(x-1)
	}
	return

}

/*斐波那契数列*/
func Fibonaci(n int) int {
	if n < 2 {
		return 2
	}
	return Fibonaci(n-2) + Fibonaci(n-1)
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia,I can call you!")
}

type Iphone struct {
}

func (iphone Iphone) call() {
	fmt.Println("I am iphone,I can call you!")
}

//GO错误处理接口
type error interface {
	Error() string
}

type DivideError struct {
	dividee int
	divider int
}

//实现`error`接口
func (de *DivideError) ERROR() string {
	strFormat := `
	Cannot proceed,the divider is zero.
	dividee:%d
	divider:0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

//定义`int`类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.ERROR()
		return
	} else {
		return varDividee / varDivider, "NoError"
	}
}
