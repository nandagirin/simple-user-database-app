package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
)

var (
	app             *fiber.App
	jwtSecret       string
	jwtTokenValid   string
	jwtTokenInvalid string
)

func init() {
	jwtSecret = "09f26e402586e2faa8da4c98a35f1b20d6b033c6097befa8be3486a829587fe2f90a832bd3ff9d42710a4da095a2ce285b009f0c3730cd9b8e1af3eb84df6611"
	os.Setenv("JWT_SECRET", jwtSecret)

	jwtTokenValid = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiYWRtaW4iLCJpYXQiOjE2OTY2NDg4MDJ9.Cz9IZUFnE0BvWf5OEWv5R4YGDO3eSO1RKXHO0QzFRL4"
	jwtTokenInvalid = ".eyJ1c2VyIjoiYWRtaW4iLCJpYXQiOjE2OTY2NDg4MDJ9.Cz9IZUFnE0BvWf5OEWv5R4YGDO3eSO1RKXHO0QzFRL4"

	app = fiber.New(fiber.Config{})
	app.Get("/users", UserList)
	app.Post("/users", UserCreate)
	app.Get("/healthz", Health)
}

func TestUserListOk(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/users", nil)
	req.Header.Set("Authorization", "Bearer "+jwtTokenValid)
	resp, _ := app.Test(req)
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("Expected status code is %d, but got %d", fiber.StatusOK, resp.StatusCode)
	}
}

func TestUserListUnauthorized(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/users", nil)
	req.Header.Set("Authorization", "Bearer "+jwtTokenInvalid)
	resp, _ := app.Test(req)
	if resp.StatusCode != fiber.StatusForbidden {
		t.Fatalf("Expected status code is %d, but got %d", fiber.StatusForbidden, resp.StatusCode)
	}
}

func TestUserCreate(t *testing.T) {
	body := struct {
		User string `json:"user"`
	}{
		User: "test",
	}

	out, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	req := httptest.NewRequest("POST", "http://localhost:8080/users", bytes.NewBuffer(out))
	req.Header.Set("Authorization", "Bearer "+jwtTokenValid)
	resp, _ := app.Test(req)
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("Expected status code is %d, but got %d", fiber.StatusOK, resp.StatusCode)
	}
}

func TestHealth(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/healthz", nil)
	resp, _ := app.Test(req)
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("Expected status code is %d, but got %d", fiber.StatusOK, resp.StatusCode)
	}
}
