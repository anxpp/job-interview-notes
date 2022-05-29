package list

import "errors"

type ArrayList struct {
	Data []interface{}
}

type ArrayListIterator struct {
	al    *ArrayList
	index int
}

func (ali *ArrayListIterator) HasNext() bool {
	return ali.index < ali.al.Len()
}

func (ali *ArrayListIterator) Next() interface{} {
	v, _ := ali.al.Get(ali.index)
	ali.index++
	return v
}

func NewArrayList() *ArrayList {
	return &ArrayList{}
}

func (al *ArrayList) Iterator() Iterator {
	return &ArrayListIterator{
		al: al,
	}
}

func (al *ArrayList) Len() int {
	return len(al.Data)
}

func (al *ArrayList) IsEmpty() bool {
	return al.Len() == 0
}

func (al *ArrayList) Cap() int {
	return cap(al.Data)
}

func (al *ArrayList) Add(x interface{}) {
	al.Data = append(al.Data, x)
}

func (al *ArrayList) FindIndex(x interface{}) int {
	var index = -1
	for i := range al.Data {
		if al.Data[i] == x {
			index = i
			break
		}
	}
	return index
}

func (al *ArrayList) Remove(x interface{}) bool {
	if i := al.FindIndex(x); i < 0 {
		return false
	} else {
		al.Data = append(al.Data[:i], al.Data[i+1:]...)
		return true
	}
}

func (al *ArrayList) Containers(x interface{}) bool {
	return al.FindIndex(x) >= 0
}

func (al *ArrayList) Clear() {
	al.Data = al.Data[:0]
}

func (al *ArrayList) Get(i int) (x interface{}, e error) {
	if i < 0 || i >= al.Len() {
		e = errors.New("index out of range")
		return
	}
	x = al.Data[i]
	return
}

func (al *ArrayList) Head() interface{} {
	if al.Len() == 0 {
		return nil
	}
	return al.Data[0]
}

func (al *ArrayList) Tail() interface{} {
	if al.Len() == 0 {
		return nil
	}
	return al.Data[al.Len()-1]
}

func (al *ArrayList) RPush(x interface{}) {
	al.Add(x)
}

func (al *ArrayList) LPop() interface{} {
	if al.Len() == 0 {
		return nil
	}
	v := al.Data[0]
	al.Data = al.Data[1:]
	return v
}

func (al *ArrayList) Push(x interface{}) {
	al.Add(x)
}

func (al *ArrayList) Pop() interface{} {
	if al.Len() == 0 {
		return nil
	}
	v := al.Data[0]
	al.Data = al.Data[1:]
	return v
}
