package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"note_server/models"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
)

func setupSuite(tb testing.TB) (func (tb testing.TB)) {
	tb.Log("setting up gorm")
	setupGorm(sqlite.Open("testing.db"))

	return func (tb testing.TB)  {
		os.Remove("./testing.db")
	}
}

func TestPostNote(t *testing.T) {
	setupGorm(sqlite.Open("test.db"))
	t.Log("Setting up server")
	router := SetupServer()


	testNote := models.Note {
		Title: "Test note",
		Descrition: "Test note",
	}

	t.Log("Marshaling note to json")

	noteJson, _ := json.Marshal(testNote)
	httpRecorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/notes/", strings.NewReader(string(noteJson)))
	router.ServeHTTP(httpRecorder, req)

	assert.Equal(t, 200, httpRecorder.Code)

	responseNote := models.Note{}
	if err := json.Unmarshal(httpRecorder.Body.Bytes(), &responseNote); err != nil {
		t.Error("Could not parse json response: ", err)
	}

	t.Log(responseNote)

	assert.Equal(t, testNote.Title, responseNote.Title)
	assert.Equal(t, testNote.Descrition, responseNote.Descrition)

	// assert.Equal(t, nil, responseNote.DeletedAt)
	assert.Equal(t, responseNote.UpdatedAt, responseNote.CreatedAt)
	assert.NotEqual(t, nil, responseNote.CreatedAt)
}