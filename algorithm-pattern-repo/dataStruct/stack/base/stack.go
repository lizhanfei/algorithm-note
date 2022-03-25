package base

type Stack interface {
	Push(val interface{})
	Pop() interface{}
	Length() int
}

func NewStack() Stack {
	return &StackImplV1{
		len: 0,
		val: make([]interface{}, 0),
	}
}

type StackImplV1 struct {
	len int
	val []interface{}
}

func (this *StackImplV1) Push(val interface{}) {
	this.val = append(this.val, val)
	this.len++
}

func (this *StackImplV1) Pop() interface{} {
	if 0 == this.len {
		return nil
	}
	res := this.val[this.len-1]
	this.val = this.val[:this.len-1]
	this.len--
	return res
}
func (this *StackImplV1) Length() int {
	return this.len
}
