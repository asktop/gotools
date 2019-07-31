package asort

import (
	"fmt"
	"testing"
)

func TestRStrings(t *testing.T) {
	params := map[string]interface{}{}
	params["c"] = "c"
	params["a"] = "a"
	params["d"] = "d"
	params["b"] = 2.3
	params["e"] = 5.4
	fmt.Println(params)
	fmt.Println(SortParamInterface(params, "&"))
}
