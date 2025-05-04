package modles

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/BurntSushi/toml"
)

type Todo struct {
	Task     Task
	DeadLine string
}

func NewTodo(id string) (*Todo, error) {
	filePath := filepath.Join("list", "todo.toml")
	var data struct {
		Todo map[string]struct {
			ID       string `toml:"id"`
			Doc      string `toml:"doc"`
			DeadLine string `toml:"deadline"`
		} `toml:"tasks"`
	}
	if _, err := toml.DecodeFile(filePath, &data); err != nil {
		return nil, err
	}
	if taskData, exists := data.Todo[id]; exists {
		return &Todo{
			Task: Task{
				ID:  id,
				Doc: taskData.Doc,
			},
			DeadLine: taskData.DeadLine,
		}, nil
	}
	return nil, fmt.Errorf("task ID %s not found", id)
}

func GetTodoList() ([]Todo, error) {
	filePath := filepath.Join("list", "todo.toml")
	var data struct {
		Tasks map[string]struct {
			ID       string `toml:"id"`
			Doc      string `toml:"doc"`
			DeadLine string `toml:"deadline"`
		} `toml:"tasks"`
	}
	if _, err := toml.DecodeFile(filePath, &data); err != nil {
		return nil, err
	}

	var todos []Todo
	for id, taskData := range data.Tasks {
		todos = append(todos, Todo{
			Task: Task{
				ID:  id,
				Doc: taskData.Doc,
			},
			DeadLine: taskData.DeadLine,
		})
	}
	return todos, nil
}

func (t Todo) PrintElement() {
	fmt.Printf("ID       : %s\n", t.Task.ID)
	fmt.Printf("Task     : %s\n", t.Task.Doc)
	fmt.Printf("Deadline : %s\n", t.DeadLine)
	fmt.Println("--------------------------")
}

func AddTodoElement(todos []Todo, dones []Done) {
	fmt.Println("タスクを追加してください.\nタスク名:")
	var task string
	var deadline string

	filePath := filepath.Join("list", "todo.toml")
	var data struct {
		Tasks map[string]struct {
			ID       string `toml:"id"`
			Doc      string `toml:"doc"`
			DeadLine string `toml:"deadline"`
		} `toml:"tasks"`
	}
	if _, err := toml.DecodeFile(filePath, &data); err != nil {
		panic(err)
	}

	if data.Tasks == nil {
		data.Tasks = make(map[string]struct {
			ID       string `toml:"id"`
			Doc      string `toml:"doc"`
			DeadLine string `toml:"deadline"`
		})
	}

	maxID := 0
	for k := range data.Tasks {
		if intID, err := strconv.Atoi(k); err == nil {
			if intID > maxID {
				maxID = intID
			}
		}
	}
	newID := strconv.Itoa(maxID + 1)

	_, err := fmt.Scan(&task)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`^\d+/\d+$`)
	for {
		fmt.Println("期限を追加してください (x/y の形式で):")
		_, err = fmt.Scan(&deadline)
		if err != nil {
			panic(err)
		}
		if re.MatchString(deadline) {
			break
		} else {
			fmt.Println("形式が正しくありません。例: 5/10 のように入力してください。")
		}
	}

	data.Tasks[newID] = struct {
		ID       string `toml:"id"`
		Doc      string `toml:"doc"`
		DeadLine string `toml:"deadline"`
	}{
		ID:       newID,
		Doc:      task,
		DeadLine: deadline,
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := toml.NewEncoder(file).Encode(data); err != nil {
		panic(err)
	}

	fmt.Println("タスクが追加されました！")
}

func DelTodoElement() {
	fmt.Println("削除するタスクのIDを入力してください:")

	var id string
	_, err := fmt.Scan(&id)
	if err != nil {
		panic(err)
	}

	filePath := filepath.Join("list", "todo.toml")
	var data struct {
		Tasks map[string]struct {
			ID       string `toml:"id"`
			Doc      string `toml:"doc"`
			DeadLine string `toml:"deadline"`
		} `toml:"tasks"`
	}

	if _, err = toml.DecodeFile(filePath, &data); err != nil {
		panic(err)
	}

	if _, exists := data.Tasks[id]; !exists {
		fmt.Printf("ID %s のタスクは存在しません。\n", id)
		return
	}

	delete(data.Tasks, id)

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := toml.NewEncoder(file).Encode(data); err != nil {
		panic(err)
	}

	fmt.Println("タスクが削除されました！")
}
