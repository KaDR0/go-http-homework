package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

// Middleware: проверка API ключа и логирование запросов
func apiKeyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path) // Лог запроса

        apiKey := r.Header.Get("X-API-Key")
        if apiKey != "secret123" {
            w.WriteHeader(http.StatusUnauthorized)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
            return
        }

        next.ServeHTTP(w, r)
    })
}

// Обработчик GET /user?id=123
func getUserHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"})
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]int{"user_id": id})
}

// Обработчик POST /user
func createUserHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var data struct {
        Name string `json:"name"`
    }

    if err := json.NewDecoder(r.Body).Decode(&data); err != nil || data.Name == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "invalid name"})
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"created": data.Name})
}

func main() {
    router := mux.NewRouter()

    // Подключаем middleware ко всем маршрутам
    router.Use(apiKeyMiddleware)

    // Регистрируем маршруты
    router.HandleFunc("/user", getUserHandler).Methods("GET")
    router.HandleFunc("/user", createUserHandler).Methods("POST")

    fmt.Println("Server started on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
