package valueobjects

type LinkedList struct {
	len int
	head *node
}

type node struct {
	next *node
	value LinkedListNode
}

type LinkedListNode interface{
	GetIndentifer() string
}

func (l *LinkedList) Add(val LinkedListNode) {
	initNode := &node{
		value : val
	}
	if l.head == nil {
		l.head = node
		l.len = 1
		return
	}
	current := l.head
	if current.next != nil {
		current = current.next
	}
	current.next = initNode
	l.len++
}

func (l *LinkedList) Remove(val LinkedListNode){
	if l.len <= 0 {
		return
	}
	prev *LinkedListNode
	current := l.head
	for current != nil{
		if current.value.GetIndentifer == val.GetIndentifer(){
			if prev == nil {
				l.head = current.next
			} else {
				prev.next = current.next
			}
		} else{
			prev = current
		}
		current = current.next
	}
}

func (l *LinkedList) traverse(f func(value LinkedListNode) error) error {
	current := l.head
	for current != nil {
		err := f(current.value)
		if err != nil {
			return err
		}
		current = current.next
	}
	return nil
}
