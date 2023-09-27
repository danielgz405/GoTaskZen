package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/danielgz405/GoTaskZen/middleware"
	"github.com/danielgz405/GoTaskZen/models"
	"github.com/danielgz405/GoTaskZen/repository"
	"github.com/danielgz405/GoTaskZen/responses"
	"github.com/danielgz405/GoTaskZen/server"
	"github.com/danielgz405/GoTaskZen/structures"
	"github.com/gorilla/mux"
)

func InsertCategoryHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		user, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = structures.InsertCategoryRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}
		category := models.InsertCategory{}
		loc, error := time.LoadLocation("America/Bogota")
		if error != nil {
			responses.InternalServerError(w, "Invalid request")
			return
		}

		category = models.InsertCategory{
			Name:        req.Name,
			Description: req.Description,
			UserId:      user.Id.Hex(),
			TableId:     req.TableId,
			Color: models.ColorBoard{
				Primary:   req.Primary,
				Secondary: req.Secondary,
			},
			CreatedAt: user.Name + " " + time.Now().In(loc).Format("2006-01-02 15:04:05"),
		}

		createdCategory, err := repository.InsertCategory(r.Context(), &category)
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdCategory)
	}
}

func ListCategorysHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		categorys, err := repository.ListCategorys(r.Context(), params["boardId"])
		if err != nil {
			fmt.Println(err)
			responses.InternalServerError(w, err.Error())
			return
		}
		if categorys == nil {
			categorys = []models.Category{}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(categorys)
	}
}

func UpdateCategoryHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var req = structures.UpdateCategoryRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}
		category := models.UpdateCategory{}

		category = models.UpdateCategory{
			Name:        req.Name,
			Description: req.Description,
			Color: models.ColorBoard{
				Primary:   req.Primary,
				Secondary: req.Secondary,
			},
		}

		updatedCategory, err := repository.UpdateCategory(r.Context(), &category, params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedCategory)
	}
}
func DeleteCategoryHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		err = repository.DeleteCategory(r.Context(), params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		responses.DeleteResponse(w, "Category deleted")
	}
}

//tasks

func InsertTaskHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		user, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = structures.InsertTaskRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}
		task := models.InsertTask{}
		loc, error := time.LoadLocation("America/Bogota")
		if error != nil {
			responses.InternalServerError(w, "Invalid request")
			return
		}

		task = models.InsertTask{
			Name:        req.Name,
			Description: req.Description,
			CreatedAt:   user.Name + " " + time.Now().In(loc).Format("2006-01-02 15:04:05"),
			Priority:    req.Priority,
		}

		params := mux.Vars(r)
		createdTask, err := repository.AddTaskToCategory(r.Context(), task, params["categoryId"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdTask)
	}
}

func UpdateTaskHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var req = structures.UpdateTaskRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}
		task := models.UpdateTask{}

		task = models.UpdateTask{
			Name:        req.Name,
			Description: req.Description,
			Priority:    req.Priority,
			Active:      req.Active,
		}

		updatedTask, err := repository.UpdateTaskToCategory(r.Context(), task, params["id"], params["categoryId"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedTask)
	}
}

func RemoveTaskHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		category, err := repository.RemoveTaskToCategory(r.Context(), params["id"], params["categoryId"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(category)
	}
}
