package list

type Collection interface {
	IsEmpty() bool
	Len() int
	Cap() int
	Add(x interface{})
	Remove(x interface{}) bool
	Containers(x interface{}) bool
}

type List interface {
	Collection
	Clear()
	Get(index int) (interface{}, error)
	Iterator() Iterator
}

type Queue interface {
	Head() interface{}
	Tail() interface{}
	RPush(x interface{})
	LPop(x interface{})
}

type DeQueue interface {
	Queue
	LPush(x interface{})
	RPop(x interface{})
}

type Stack interface {
	Push(x interface{})
	Pop(x interface{})
	Tail() interface{}
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
