package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

type as struct {
	key   string
	value string
}

type ss struct {
	aa []as
	bb as
	cc as
}

func main() {
	keys := make([]string, 3)
	keys[0] = "j"
	keys[1] = "b"
	keys[2] = "C"
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	var buf bytes.Buffer
	maps := make(map[string]interface{})
	maps["aa"] = [2]as{{"aakey", "aavalue"}, {"aakey2", "aavalue2"}}
	maps["bb"] = as{"bbkey", "bbvalue"}
	maps["cc"] = as{"cckey", "ccvalue"}
	buf.WriteString("hello")
	fmt.Println(buf.String())
	fmt.Println(maps["aa"])
	//fmt.Println((maps["aa"]).(string))

	a := ss{
		[]as{as{"aakey", "aavalue"}, as{"aakey2", "aavalue2"}},
		as{"bbkey", "bbvalue"},
		as{"cckey", "ccvalue"},
	}

	fmt.Println("--------------")
	fmt.Println(a)
	bytes, _ := json.Marshal(&a)
	//retMap := make(map[string]interface{})
	var yu ss
	json.Unmarshal(bytes, &yu)
	fmt.Println(yu)
}
