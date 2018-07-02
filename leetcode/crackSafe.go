package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(crackSafe(4, 6))
	fmt.Println(len("000555545455445444455534553545344535355434543544344435343435533453354334433533343333555245523552545244523452535243523352525542454235425442444234425342434233425242425532453235325432443234325332433233325232423232552245223522542244223422532243223322522242223222255514551355125515451445134512451535143513351235152514251325122515155414541354125415441444134412441534143413341234152414241324122415141415531453135312531543144313431243153314331333123315231423132312231513141313155214521352125215421442134212421532143213321232152214221322122215121412131212155114511351125115411441134112411531143113311231152114211321122115111411131112111155504550355025501550545044503450245014505350435033502350135052504250325022501250515041503150215011505055404540354025401540544044403440244014405340434033402340134052404240324022401240514041403140214011405040405530453035302530153054304430343024301430533043303330233013305230423032302230123051304130313021301130503040303055204520352025201520542044203420242014205320432033202320132052204220322022201220512041203120212011205020402030202055104510351025101510541044103410241014105310431033102310131052104210321022101210511041103110211011105010401030102010105500450035002500150054004400340024001400530043003300230013005200420032002200120051004100310021001100500040003000200010000"))
}

var (
	str string
)

// n password  0 - k-1 digit
func crackSafe(n int, k int) string {
	str = ""
	next := ""
	all := 1
	for i := 0; i < n; i++ {
		next += "0"
		all *= k
	}
	visited := make(map[string]bool)
	visited[next] = true
	str = next
	goNext(next, visited, n, k, all)

	fmt.Println(len(str))
	return str
}

func goNext(next string, visited map[string]bool, n, k, all int) bool {
	//fmt.Println("11: ", str)
	if len(visited) == all {
		return true
	}
	for i := 0; i < k; i++ {
		key := next[1:] + strconv.Itoa(i)
		if _, ok := visited[key]; !ok {
			visited[key] = true
			str += strconv.Itoa(i)
			if goNext(key, visited, n, k, all) {
				return true
			} else {
				delete(visited, key)
				str = str[:len(str)-1]
			}
		}
	}
	return false
}
