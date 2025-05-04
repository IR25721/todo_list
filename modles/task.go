package modles

import "fmt"

type Task struct {
	ID  string
	Doc string
}

type TaskItem interface {
	PrintElement()
}

func PrintList(title string, items []TaskItem) {
	if len(items) == 0 {
		fmt.Printf("%s はありません!", title)
		return
	}
	fmt.Printf("======= %s =======", title)
	for _, item := range items {
		item.PrintElement()
	}
}

func ToTaskItem[T TaskItem](list []T) []TaskItem {
	items := make([]TaskItem, len(list))
	for i, v := range list {
		items[i] = v
	}
	return items
}
