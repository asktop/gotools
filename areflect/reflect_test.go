package areflect

import "testing"

func abc() {

}

func TestGetFuncName(t *testing.T) {
	t.Log(GetFuncName(abc))
}
