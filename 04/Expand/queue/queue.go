package queue

type Queue []interface{}

func (queue *Queue) Append(value interface{}) {
	*queue = append(*queue, value.(int))
}

func (queue *Queue) Pop() interface{} {
	head := (*queue)[0]
	*queue = (*queue)[1:]
	return head.(int)
}

func (queue *Queue) Isempty() bool {
	return len(*queue) == 0
}
