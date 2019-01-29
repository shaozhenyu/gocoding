package main

import (
	"fmt"
	"strconv"
	"strings"
	"sort"
)

func main() {
	nums := []int{824,938,1399,5607,6973,5703,9609,4398,8247}
	nums = []int{121,12}
	nums = []int{9051,5526,2264,5041,1630,5906,6787,8382,4662,4532,6804,4710,4542,2116,7219,8701,8308,957,8775,4822,396,8995,8597,2304,8902,830,8591,5828,9642,7100,3976,5565,5490,1613,5731,8052,8985,2623,6325,3723,5224,8274,4787,6310,3393,78,3288,7584,7440,5752,351,4555,7265,9959,3866,9854,2709,5817,7272,43,1014,7527,3946,4289,1272,5213,710,1603,2436,8823,5228,2581,771,3700,2109,5638,3402,3910,871,5441,6861,9556,1089,4088,2788,9632,6822,6145,5137,236,683,2869,9525,8161,8374,2439,6028,7813,6406,7519}
	fmt.Println(largestNumber(nums))
}

type Node struct {
	v string
	s string
}

func largestNumber(nums []int) string {
	ret := ""
	add := make([]string, 10)
	for i := 0; i < 10; i++ {
		add[i] = strings.Repeat(strconv.Itoa(i), 10)
	}
	node := make([]Node, len(nums))
	for i := 0; i < len(nums); i++ {
		v := strconv.Itoa(nums[i])
		s := v + add[int(v[0] - '0')]
		node[i] = Node{v, s[:10]}
	}
	sort.Slice(node, func(i, j int) bool {
		if node[i].s == node[j].s {
			return node[i].v + node[j].v > node[j].v + node[i].v
		}
		return node[i].s > node[j].s
	})
	for i := 0; i < len(node); i++ {
		ret += node[i].v
	}
	if len(ret) > 0 && ret[0] == '0' {
		return "0"
	}
	return ret
}

