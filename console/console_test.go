package console

import(
	"testing"
)

func Test_SetConsoleTitle(t *testing.T){
	r,err := SetConsoleTitle("SetConsoleTitle")
	if r == 0 {
		t.Error(err)
	} else {
		t.Log("success")
	}
}