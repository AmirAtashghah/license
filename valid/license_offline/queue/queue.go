package queue

const maxQueueSize = 100

type Queue struct {
	items  []string
	exists map[string]bool
}

func NewQueue() *Queue {
	return &Queue{
		items:  make([]string, 0, maxQueueSize),
		exists: make(map[string]bool),
	}
}

func (q *Queue) Enqueue(item string) {
	if len(q.items) >= maxQueueSize {
		// Remove the oldest item
		oldest := q.items[0]
		q.items = q.items[1:]
		delete(q.exists, oldest)
	}
	q.items = append(q.items, item)
	q.exists[item] = true
}

func (q *Queue) Dequeue() (string, bool) {
	if len(q.items) == 0 {
		return "", false
	}
	item := q.items[0]
	q.items = q.items[1:]
	delete(q.exists, item)
	return item, true
}

func (q *Queue) Exists(item string) bool {
	_, found := q.exists[item]
	return found
}
