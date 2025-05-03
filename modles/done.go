package modles

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/BurntSushi/toml"
)

type Done struct {
	Task     Task
	DoneDate string
}

func NewDone(id string) (*Done, error) {
	filePath := filepath.Join("list", "done.toml")
	var data struct {
		Tasks map[string]struct {
			ID       string `toml:"id"`
			Doc      string `toml:"doc"`
			DoneDate string `toml:"doneDate"`
		} `toml:"tasks"`
	}
	if _, err := toml.DecodeFile(filePath, &data); err != nil {
		return nil, err
	}
	if taskData, exists := data.Tasks[id]; exists {
		return &Done{
			Task: Task{
				ID:  id,
				Doc: taskData.Doc,
			},
			DoneDate: taskData.DoneDate,
		}, nil
	}
	return nil, fmt.Errorf("task ID %s not found", id)
}

func GetDoneList() ([]Done, error) {
	filePath := filepath.Join("list", "done.toml")
	var data struct {
		Tasks map[string]struct {
			ID       string `toml:"id"`
			Doc      string `toml:"doc"`
			DoneDate string `toml:"doneDate"`
		} `toml:"tasks"`
	}
	if _, err := toml.DecodeFile(filePath, &data); err != nil {
		return nil, err
	}

	var dones []Done
	for id, taskData := range data.Tasks {
		dones = append(dones, Done{
			Task: Task{
				ID:  id,
				Doc: taskData.Doc,
			},
			DoneDate: taskData.DoneDate,
		})
	}
	return dones, nil
}

func DoneList(dones []Done) {
	if len(dones) == 0 {
		fmt.Println("終わらせたタスクはありません！")
		return
	} else {
		fmt.Println("======= DONE LIST =======")
		for _, done := range dones {
			fmt.Printf("ID       : %s\n", done.Task.ID)
			fmt.Printf("Task     : %s\n", done.Task.Doc)
			fmt.Printf("DoneDate : %s\n", done.DoneDate)
			fmt.Println("--------------------------")
		}
	}
}

func AddDoneElement() {
	fmt.Println("完了したタスクのIDを入力してください:")

	var id string
	_, err := fmt.Scan(&id)
	if err != nil {
		panic(err)
	}

	todoPath := filepath.Join("list", "todo.toml")
	var todoData struct {
		Tasks map[string]struct {
			ID       string `toml:"id"`
			Doc      string `toml:"doc"`
			DeadLine string `toml:"deadline"`
		} `toml:"tasks"`
	}
	if _, err = toml.DecodeFile(todoPath, &todoData); err != nil {
		panic(err)
	}

	taskData, ok := todoData.Tasks[id]
	if !ok {
		fmt.Println("指定されたIDのタスクが見つかりません。")
		return
	}

	var doneDate string
	re := regexp.MustCompile(`^\d+/\d+$`)
	for {
		fmt.Println("完了日を入力してください (例: 5/10):")
		_, err = fmt.Scan(&doneDate)
		if err != nil {
			panic(err)
		}
		if re.MatchString(doneDate) {
			break
		} else {
			fmt.Println("形式が正しくありません。例: 5/10 のように入力してください。")
		}
	}

	donePath := filepath.Join("list", "done.toml")
	var doneData struct {
		Tasks map[string]struct {
			ID       string `toml:"id"`
			Doc      string `toml:"doc"`
			DoneDate string `toml:"doneDate"`
		} `toml:"tasks"`
	}
	if _, err = toml.DecodeFile(donePath, &doneData); err != nil {
		doneData.Tasks = make(map[string]struct {
			ID       string `toml:"id"`
			Doc      string `toml:"doc"`
			DoneDate string `toml:"doneDate"`
		})
	}
	if doneData.Tasks == nil {
		doneData.Tasks = make(map[string]struct {
			ID       string `toml:"id"`
			Doc      string `toml:"doc"`
			DoneDate string `toml:"doneDate"`
		})
	}

	doneData.Tasks[id] = struct {
		ID       string `toml:"id"`
		Doc      string `toml:"doc"`
		DoneDate string `toml:"doneDate"`
	}{
		ID:       taskData.ID,
		Doc:      taskData.Doc,
		DoneDate: doneDate,
	}

	delete(todoData.Tasks, id)

	todoFile, err := os.Create(todoPath)
	if err != nil {
		panic(err)
	}
	defer todoFile.Close()
	if err = toml.NewEncoder(todoFile).Encode(todoData); err != nil {
		panic(err)
	}

	doneFile, err := os.Create(donePath)
	if err != nil {
		panic(err)
	}
	defer doneFile.Close()
	if err := toml.NewEncoder(doneFile).Encode(doneData); err != nil {
		panic(err)
	}

	fmt.Printf("タスク(ID: %s) を完了済みに移動しました。\n", id)
}
