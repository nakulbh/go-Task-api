package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nakulbh/goTask/models"
)

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request payload")
		return
	}

	task.ID = insertTask(task)

	res := response{
		ID:      insertID,
		Message: "Task created succesfully",
	}
	json.NewEncoder(w).Encode(task)
}
