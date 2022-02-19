package structures

import "fmt"

type HashTable struct {
	values []*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	value string
	next  *bucketNode
}

func hash(value string, size int) int {
	var sum int
	for _, v := range value {
		sum += int(v)
	}
	return sum % size
}

// HashTable を指定した size で作成し初期化して返す
func InitHashTable(size int) *HashTable {
	hashTable := &HashTable{values: make([]*bucket, size)}

	for i := 0; i < size; i++ {
		hashTable.values[i] = &bucket{}
	}
	return hashTable
}

func (h HashTable) GetSize() int {
	return len(h.values)
}

func (h HashTable) GetValues() [][]string {
	result := make([][]string, h.GetSize())
	for i, v := range h.values {
		values := v.GetValues()
		result[i] = make([]string, len(values))
		for k, w := range values {
			result[i][k] = w
		}
	}

	return result
}

func (h HashTable) Insert(value string) error {
	if h.Search(value) {
		return fmt.Errorf("%s is already in the table", value)
	}
	index := hash(value, h.GetSize())

	buc := h.values[index]
	node := &bucketNode{value: value, next: buc.head}
	buc.head = node
	return nil
}

func (h HashTable) Search(value string) bool {
	index := hash(value, h.GetSize())

	buc := *h.values[index]
	node := buc.head
	for node != nil {
		if node.value == value {
			return true
		}

		node = node.next
	}
	return false
}

func (h HashTable) Delete(value string) error {
	if !h.Search(value) {
		return fmt.Errorf("%s is not in the table", value)
	}
	index := hash(value, h.GetSize())
	head := h.values[index].head
	node := head
	for node.next != nil {
		if node.next.value == value {
			node.next = node.next.next
			if head == node {
				h.values[index].head = node.next
			}
			break
		}
		node = node.next
		return nil
	}

	// head が削除対象である場合
	h.values[index].head = nil
	return nil
}

func (b bucket) GetValues() []string {
	var values []string

	node := b.head
	for node != nil {
		values = append(values, node.value)
		node = node.next
	}
	return values
}
