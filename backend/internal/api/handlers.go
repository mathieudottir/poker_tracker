package api

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mathieudottir/poker_tracker/backend/internal/auth"
	"github.com/mathieudottir/poker_tracker/backend/internal/repository"
	"github.com/mathieudottir/poker_tracker/backend/internal/services"
	"github.com/mathieudottir/poker_tracker/backend/internal/watcher"
	"github.com/mathieudottir/poker_tracker/backend/pkg/constants"
)

type API struct {
	authService       *auth.AuthService
	importService     *services.ImportService
	tournamentService *services.TournamentService
	watcherManager    *watcher.WatcherManager
	repo              *repository.Repository
}

func NewAPI(repo *repository.Repository) *API {
	return &API{
		authService:       auth.NewAuthService(repo),
		importService:     services.NewImportService(repo),
		tournamentService: services.NewTournamentService(repo),
		watcherManager:    watcher.NewWatcherManager(),
		repo:              repo,
	}
}

func (api *API) SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Enable CORS
	r.Use(corsMiddleware)

	// Auth endpoints
	r.HandleFunc("/api/auth/register", api.handleRegister).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/auth/login", api.handleLogin).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/auth/dev-login", api.handleDevLogin).Methods("POST", "OPTIONS")

	// User endpoints
	r.HandleFunc("/api/user/{id}", api.handleGetUser).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/user/{id}/settings", api.handleUpdateSettings).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/user/{id}/data", api.handleDeleteUserData).Methods("DELETE", "OPTIONS")

	// Tournament endpoints
	r.HandleFunc("/api/tournaments/{userId}", api.handleGetTournaments).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/tournament/{id}", api.handleGetTournament).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/tournament/{id}/hands", api.handleGetHands).Methods("GET", "OPTIONS")

	// Import endpoints
	r.HandleFunc("/api/import/tournament", api.handleImportTournament).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/import/directory", api.handleImportDirectory).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/import/files", api.handleImportFiles).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/import/logs", api.handleGetImportLogs).Methods("GET", "OPTIONS")

	// Watcher endpoints
	r.HandleFunc("/api/watcher/start", api.handleStartWatcher).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/watcher/stop", api.handleStopWatcher).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/watcher/status/{userId}", api.handleWatcherStatus).Methods("GET", "OPTIONS")

	// Stats endpoints
	r.HandleFunc("/api/stats/{userId}", api.handleGetStats).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/stats/{userId}/multipliers", api.handleGetMultipliers).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/stats/{userId}/bybuyin", api.handleGetByBuyin).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/stats/{userId}/bankroll", api.handleGetBankroll).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/stats/{userId}/ev", api.handleGetEV).Methods("GET", "OPTIONS")

	// Winamax data endpoints
	r.HandleFunc("/api/winamax/multipliers", api.handleGetWinamaxMultipliers).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/winamax/rakeback", api.handleGetRakebackStatuses).Methods("GET", "OPTIONS")

	return r
}

// Auth handlers

func (api *API) handleRegister(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		PlayerName string `json:"player_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := api.authService.Register(r.Context(), req.Username, req.Password, req.PlayerName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respondJSON(w, user)
}

func (api *API) handleLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := api.authService.Login(r.Context(), req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	respondJSON(w, user)
}

func (api *API) handleDevLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := api.authService.DevLogin(r.Context(), req.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	respondJSON(w, user)
}

// User handlers

func (api *API) handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	user, err := api.authService.GetUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondJSON(w, user)
}

func (api *API) handleUpdateSettings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	var req struct {
		PlayerName  string `json:"player_name"`
		WinaStatus  string `json:"wina_status"`
		HHDirectory string `json:"hh_directory"`
		DevMode     bool   `json:"dev_mode"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := api.authService.UpdateUserSettings(r.Context(), userID, req.PlayerName, req.WinaStatus, req.HHDirectory, req.DevMode); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, map[string]string{"status": "success"})
}

func (api *API) handleDeleteUserData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	if err := api.authService.DeleteUserData(r.Context(), userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, map[string]string{"status": "deleted"})
}

// Tournament handlers

func (api *API) handleGetTournaments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userId"])

	tournaments, err := api.tournamentService.GetTournaments(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, tournaments)
}

func (api *API) handleGetTournament(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	tournament, err := api.tournamentService.GetTournament(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondJSON(w, tournament)
}

func (api *API) handleGetHands(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	hands, err := api.tournamentService.GetHands(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, hands)
}

// Import handlers

func (api *API) handleImportTournament(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID      int    `json:"user_id"`
		HHFile      string `json:"hh_file"`
		SummaryFile string `json:"summary_file"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := api.importService.ImportTournament(r.Context(), req.UserID, req.HHFile, req.SummaryFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, map[string]string{"status": "imported"})
}

func (api *API) handleImportDirectory(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID    int    `json:"user_id"`
		Directory string `json:"directory"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	count, err := api.importService.ImportDirectory(r.Context(), req.UserID, req.Directory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, map[string]interface{}{"status": "imported", "count": count})
}

func (api *API) handleImportFiles(w http.ResponseWriter, r *http.Request) {
	// Set max request body size to 500MB
	r.Body = http.MaxBytesReader(w, r.Body, 500<<20)

	// Parse multipart form (max 500MB)
	if err := r.ParseMultipartForm(500 << 20); err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	userIDStr := r.FormValue("user_id")
	if userIDStr == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		http.Error(w, "No files uploaded", http.StatusBadRequest)
		return
	}

	// Create temp directory
	tempDir, err := os.MkdirTemp("", "poker_import_*")
	if err != nil {
		http.Error(w, "Failed to create temp directory: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tempDir)

	// Save uploaded files to temp directory
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Failed to open uploaded file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		destPath := filepath.Join(tempDir, fileHeader.Filename)
		destFile, err := os.Create(destPath)
		if err != nil {
			http.Error(w, "Failed to create file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer destFile.Close()

		if _, err := io.Copy(destFile, file); err != nil {
			http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Import from temp directory
	count, err := api.importService.ImportDirectory(r.Context(), userID, tempDir)
	if err != nil {
		http.Error(w, "Import failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, map[string]interface{}{"status": "imported", "count": count})
}

func (api *API) handleGetImportLogs(w http.ResponseWriter, r *http.Request) {
	logs := api.importService.GetImportLogs()
	respondJSON(w, logs)
}

// Watcher handlers

func (api *API) handleStartWatcher(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID    int    `json:"user_id"`
		Directory string `json:"directory"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := api.watcherManager.StartWatcher(context.Background(), req.UserID, req.Directory, api.importService); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, map[string]string{"status": "started"})
}

func (api *API) handleStopWatcher(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID int `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.watcherManager.StopWatcher(req.UserID)
	respondJSON(w, map[string]string{"status": "stopped"})
}

func (api *API) handleWatcherStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userId"])

	watcher := api.watcherManager.GetWatcher(userID)
	if watcher == nil {
		respondJSON(w, map[string]interface{}{"running": false})
		return
	}

	respondJSON(w, map[string]interface{}{
		"running":   watcher.IsRunning(),
		"directory": watcher.GetDirectory(),
	})
}

// Stats handlers

func (api *API) handleGetStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userId"])

	user, err := api.authService.GetUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	stats, err := api.tournamentService.GetStats(r.Context(), userID, user.WinaStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, stats)
}

func (api *API) handleGetMultipliers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userId"])

	distribution, err := api.tournamentService.GetMultiplierDistribution(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, distribution)
}

func (api *API) handleGetByBuyin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userId"])

	results, err := api.tournamentService.GetResultsByBuyin(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, results)
}

func (api *API) handleGetBankroll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userId"])

	history, err := api.tournamentService.GetBankrollHistory(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, history)
}

func (api *API) handleGetEV(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userId"])

	history, err := api.tournamentService.GetEVHistory(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, history)
}

// Winamax data handlers

func (api *API) handleGetWinamaxMultipliers(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, constants.WinamaxMultipliers)
}

func (api *API) handleGetRakebackStatuses(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, constants.WinamaxRakebackStatus)
}

// Utility functions

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for all requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, Origin, X-Requested-With")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
