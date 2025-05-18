import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	statusTodo       = "todo"
	statusInProgress = "in progress"
	statusDone       = "done"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func main() {

	command := os.Args[1]

	if command != "add" && command != "update" && command != "delete" && command != "list" {
		fmt.Println("perintah yang di input tidak terdia")
		return
	}

	description := os.Args[2]
	task := Task{
		Id:          1,
		Description: description,
		Status:      statusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	jsonData, err := json.Marshal(task)
	if err != nil {
		fmt.Println("Error convert data ke json", err)
	}

	file, err := os.Create("task.json")
	if err != nil {
		fmt.Println("Error membuat file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error menulis file json")
		return
	}

	fmt.Println("Berhasil input task ke file json")
	return

}
