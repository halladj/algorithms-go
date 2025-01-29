package stack

type Stack interface {
	StackPush(value interface{}) error
	StackPop() (interface{}, error)
	StackPrint()
}
