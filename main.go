package main

import (
	"fmt"

	"github.com/IR25721/todo_list/modles"
)

func main() {
	for {
		todoList, err := modles.GetTodoList()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		var doneList []modles.Done
		doneList, err = modles.GetDoneList()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		println("Hello! Todo管理へようこそ!\n0:Todoリストを表示\n1:Todoリストに追加\n2:Todoリストから削除\n3:Doneリストを表示\n4:Doneリストに追加\n5:終了")

		var choice int
		_, err = fmt.Scan(&choice)
		if err != nil {
			fmt.Printf("入力エラー: %v\n", err)
			continue
		}

		switch choice {
		case 0:
			modles.PrintList("todo_list", modles.ToTaskItem(todoList))
		case 1:
			modles.AddTodoElement(todoList, doneList)
		case 2:
			modles.DelTodoElement()
		case 3:
			modles.PrintList("done_list", modles.ToTaskItem(doneList))
		case 4:
			modles.AddDoneElement()
		case 5:
			fmt.Println("終了します。")
			return
		default:
			fmt.Println("無効な選択肢です。もう一度選んでください。")
		}
	}
}
