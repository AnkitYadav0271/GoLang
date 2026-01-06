package main

import "fmt"

func main() {
	nums:=[]int {1,3,4,5,6,8,5}
	// for i:=0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	//using range key word
    sum:= 0
	for index,num:= range nums {
		fmt.Println(num)
		fmt.Println("Index buddy",index)
		sum += num 
	}

	fmt.Println(sum)

	m:= map[string]string {"firstName":"Happy","lastName":"Yadav"}
	for k,v:= range m {
		fmt.Println(k,v)
	}

	//unicode values will be returned here 
	var whatAString string = "this is only string do not read it "
	for k,v := range whatAString {
		fmt.Println(k,string(v))
	}
}