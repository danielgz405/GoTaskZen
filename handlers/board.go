package handlers

import (
	"encoding/json"
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

func InsertBoardHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		user, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = structures.InsertBoardRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}
		board := models.InsertBoard{}
		loc, error := time.LoadLocation("America/Bogota")
		if error != nil {
			responses.InternalServerError(w, "Invalid request")
			return
		}

		board = models.InsertBoard{
			Name:        req.Name,
			Description: req.Description,
			UserId:      user.Id.Hex(),
			Saved:       false,
			Color: models.ColorBoard{
				Primary:   req.Primary,
				Secondary: req.Secondary,
			},
			Image:               req.Image,
			Background:          req.Background,
			CreatedAt:           user.Name + " " + time.Now().In(loc).Format("2006-01-02 15:04:05"),
			DesertRef:           req.DesertRef,
			DesertRefBackground: req.DesertRefBackground,
		}

		createdBoard, err := repository.InsertBoard(r.Context(), &board)
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdBoard)
	}
}
func ListBoardsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		boards, err := repository.ListBoards(r.Context())
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		if boards == nil {
			boards = []models.Board{}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(boards)
	}
}

func UpdateBoardHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var req = structures.UpdateBoardRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}
		board := models.UpdateBoard{}

		board = models.UpdateBoard{
			Name:        req.Name,
			Description: req.Description,
			Saved:       req.Saved,
			Color: models.ColorBoard{
				Primary:   req.Primary,
				Secondary: req.Secondary,
			},
			Image:               req.Image,
			Background:          req.Background,
			DesertRef:           req.DesertRef,
			DesertRefBackground: req.DesertRefBackground,
		}

		updatedBoard, err := repository.UpdateBoard(r.Context(), &board, params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedBoard)
	}
}
func DeleteBoardHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		err = repository.DeleteBoard(r.Context(), params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		responses.DeleteResponse(w, "Board deleted")
	}
}

func GetBoardByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Token validation
		_, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		board, err := repository.GetBoardById(r.Context(), params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(board)
	}
}
