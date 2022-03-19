import "testing"
const (
	Monday = iota + 1
	Tuesday
	Wednesday

)

fun TestConstantTry(t *testing.T){
	t.Log(Monday,Tuesday)
}
