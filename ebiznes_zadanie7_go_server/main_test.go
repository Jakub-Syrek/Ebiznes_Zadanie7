package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const statusErrorMessage = "handler returned wrong status code: got %v want %v"

func TestGetProductsOne(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getProducts)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(statusErrorMessage, status, http.StatusOK)
	}

	expected := `[{"id":"1","name":"Product 1","price":10},{"id":"2","name":"Product 2","price":20},{"id":"3","name":"Product 3","price":30}]`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	}
}

func TestHandlePayments_PostTwo(t *testing.T) {
	payment := `{"id":"1","amount":100.00,"cardNumber":"1234567812345678","cardExpiry":"01/23","cardCvv":"123"}`
	req, err := http.NewRequest("POST", "/api/payments", bytes.NewBufferString(payment))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlePayments)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(statusErrorMessage, status, http.StatusOK)
	}

	expected := `{"id":"1","amount":100,"cardNumber":"1234567812345678","cardExpiry":"01/23","cardCvv":"123"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSpace(rr.Body.String()), expected)
	}
}

func TestHandlePayments_GetThree(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/payments", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlePayments)

	handler.ServeHTTP(rr, req)

	

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf(statusErrorMessage, status, http.StatusMethodNotAllowed)
	}
}

func TestHandlePayments_EmptyBodyFour(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/payments", bytes.NewBufferString(""))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlePayments)

	handler.ServeHTTP(rr, req)

	

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf(statusErrorMessage, status, http.StatusBadRequest)
	}
}

func TestHandlePayments_InvalidJSONFive(t *testing.T) {
	payment := `{"id":"1","amount":"invalid","cardNumber":"1234567812345678","cardExpiry":"01/23","cardCvv":"123"}`
	req, err := http.NewRequest("POST", "/api/payments", bytes.NewBufferString(payment))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlePayments)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf(statusErrorMessage, status, http.StatusBadRequest)
	}
}

func TestHandlePayments_InvalidBodySix(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/payments", bytes.NewBufferString("invalid"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlePayments)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf(statusErrorMessage, status, http.StatusBadRequest)
	}
}

func TestHandlePayments_GET_NotAllowedSeven(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/payments", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlePayments)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf(statusErrorMessage, status, http.StatusMethodNotAllowed)
	}
}

func TestGetProducts_POST_NotAllowedEight(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getProducts)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf(statusErrorMessage, status, http.StatusMethodNotAllowed)
	}
}

func TestInvalidPathNine(t *testing.T) {
	req, err := http.NewRequest("GET", "/invalid", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf(statusErrorMessage, status, http.StatusNotFound)
	}
}

func TestHandlePayments_EmptyIDTen(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/payments", strings.NewReader(`{"id":"","amount":100,"cardNumber":"1234567812345678","cardExpiry":"01/23","cardCvv":"123"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlePayments)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(statusErrorMessage, status, http.StatusOK)
	}

	expected := `{"id":"","amount":100,"cardNumber":"1234567812345678","cardExpiry":"01/23","cardCvv":"123"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
