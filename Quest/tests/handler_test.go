package test

import (
	"Quest/internal/handler"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"Quest/internal/types"
	"context"
)

func performRequest(t *testing.T, method, path string, requestBody []byte) *httptest.ResponseRecorder {
	service := &mockService{}
	h := handler.InitHandler(service)
	router := h.Init()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(method, path, bytes.NewReader(requestBody))
	router.ServeHTTP(w, req)

	return w
}

func TestAddUserHandler(t *testing.T) {
	user := types.User{
		Name:    "Dmitriy",
		Balance: 1000,
	}

	userJSON, _ := json.Marshal(user)

	w := performRequest(t, "POST", "/user", userJSON)

	assert.Equal(t, http.StatusOK, w.Code)
	var response types.Response
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "User saved", response.Message)

	responseData, _ := response.Description.(map[string]interface{})
	balance, _ := responseData["balance"].(float64)

	assert.Equal(t, user.Name, responseData["name"])
	assert.Equal(t, user.Balance, int(balance))
}

func TestAddUserHandlerWithoutBody(t *testing.T) {
	w := performRequest(t, "POST", "/user", []byte{})

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response types.Response
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "Can't parse user", response.Message)
}

func TestAddQuestHandler(t *testing.T) {
	quest := types.Quest{
		Name: "Mark",
		Cost: 300,
	}

	questJSON, _ := json.Marshal(quest)

	w := performRequest(t, "POST", "/quest", questJSON)

	assert.Equal(t, http.StatusOK, w.Code)
	var response types.Response
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "Quest saved", response.Message)

	responseData, _ := response.Description.(map[string]interface{})
	cost, _ := responseData["cost"].(float64)

	assert.Equal(t, quest.Name, responseData["name"])
	assert.Equal(t, quest.Cost, int(cost))
}

func TestAddQuestHandlerWithoutBody(t *testing.T) {
	w := performRequest(t, "POST", "/quest", []byte{})

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response types.Response
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "Can't parse quest", response.Message)
}

func TestProcessSignalHandler(t *testing.T) {
	signal := types.Signal{
		UserId:  1,
		QuestId: 2,
	}

	signalJSON, _ := json.Marshal(signal)

	w := performRequest(t, "POST", "/signal", signalJSON)

	assert.Equal(t, http.StatusOK, w.Code)
	var response types.Response
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "Signal processed", response.Message)
}

func TestProcessSignalHandlerWithoutBody(t *testing.T) {
	w := performRequest(t, "POST", "/signal", []byte{})

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response types.Response
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "Can't parse signal", response.Message)
}

type UserHistory struct {
	UserId    int    `db:"user_id" json:"user_id"`
	UserName  string `db:"username" json:"user_name"`
	QuestId   int    `db:"quest_id" json:"quest_id"`
	QuestName string `db:"quest_name" json:"quest_name"`
	Cost      int    `db:"cost" json:"cost"`
}

func performGetUserHistoryRequest(t *testing.T, userID string) *httptest.ResponseRecorder {
	service := &mockService{}
	h := handler.InitHandler(service)
	router := h.Init()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/user/"+userID+"/history", nil)
	router.ServeHTTP(w, req)

	return w
}

func TestGetUserHistoryHandler(t *testing.T) {
	w := performRequest(t, "GET", "/user/1/history", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	var response types.Response
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "Get User history", response.Message)
}

func TestGetUserHistoryHandlerInvalidID(t *testing.T) {
	w := performGetUserHistoryRequest(t, "invalid")

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response types.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Assert the response message
	assert.Equal(t, "Can't parse user id", response.Message)
}

type mockService struct{}

func (s *mockService) AddUser(ctx context.Context, user types.User) error {
	return nil
}

func (s *mockService) AddQuest(ctx context.Context, quest types.Quest) error {
	return nil
}

func (s *mockService) ProcessSignal(ctx context.Context, signal types.Signal) (types.User, error) {
	return types.User{}, nil
}

func (s *mockService) GetUserHistory(ctx context.Context, id int) ([]types.UserHistory, error) {
	return nil, nil
}
