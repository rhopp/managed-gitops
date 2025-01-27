package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/redhat-appstudio/managed-gitops/backend/routes"
	"github.com/redhat-appstudio/managed-gitops/backend/util"
)

func TestServer(t *testing.T) {
	serverURL := "http://localhost:8090"

	server := routes.RouteInit()
	go func() {
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Println("Error on ListenAndServe:", err)
		}
	}()
	defer server.Close()

	if err := util.WaitForServerUp(serverURL); err != nil {
		t.Errorf("%v", err)
	}

	// GET should successfully pass with 200
	resp_app, err_app := http.Get(serverURL + "/api/v1/application/")
	if err_app != nil {
		t.Errorf("unexpected error in GET /api/v1/application/: %v", err_app)
	}
	if resp_app.StatusCode == http.StatusMethodNotAllowed {
		t.Errorf("unexpected response: %v, expected: %v", resp_app.StatusCode, http.StatusOK)
	}

	// GET should give a 405
	resp_opr, err_opr := http.Get(serverURL + "/api/v1/operation/")
	if err_opr != nil {
		t.Errorf("unexpected error in GET /api/v1/operation/: %v", err_opr)
	}
	if resp_opr.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("unexpected response: %v, expected: %v", resp_opr.StatusCode, http.StatusOK)
	}

	// GET should successfully pass with 200
	resp, err := http.Get(serverURL + "/api/v1/managedenvironment/")
	if err != nil {
		t.Errorf("unexpected error in GET /api/v1/managedenvironment/: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected response: %v, expected: %v", resp.StatusCode, http.StatusOK)
	}
}
