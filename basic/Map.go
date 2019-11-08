package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for key, value := range countryCapitalMap {
		fmt.Println(key, "首都是", value)
	}

	delete(countryCapitalMap, "France")
	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["France"] /*如果确定是真实的,则存在,否则不存在 */
	if ok {
		fmt.Println("France 的首都是", capital)
	} else {
		fmt.Println("France 的首都不存在")
	}
}
