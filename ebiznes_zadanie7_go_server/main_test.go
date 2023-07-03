package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const statusErrorMessage = "handler returned wrong status code: got %v want %v"
const bodyErrorMessage = "handler returned unexpected body: got %v want %v"
const apiPayments = "/api/payments"

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
		t.Errorf(bodyErrorMessage,
			strings.TrimSpace(rr.Body.String()), expected)
	}
}

func TestHandlePaymentsPostTwo(t *testing.T) {
	payment := `{"id":"1","amount":100.00,"cardNumber":"1234567812345678","cardExpiry":"01/23","cardCvv":"123"}`
	req, err := http.NewRequest("POST", apiPayments, bytes.NewBufferString(payment))
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
		t.Errorf(bodyErrorMessage,
			strings.TrimSpace(rr.Body.String()), expected)
	}
}

func TestHandlePaymentsGetThree(t *testing.T) {
	req, err := http.NewRequest("GET", apiPayments, nil)
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

func TestHandlePaymentsEmptyBodyFour(t *testing.T) {
	req, err := http.NewRequest("POST", apiPayments, bytes.NewBufferString(""))
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

func TestHandlePaymentsInvalidJSONFive(t *testing.T) {
	payment := `{"id":"1","amount":"invalid","cardNumber":"1234567812345678","cardExpiry":"01/23","cardCvv":"123"}`
	req, err := http.NewRequest("POST", apiPayments, bytes.NewBufferString(payment))
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

func TestHandlePaymentsInvalidBodySix(t *testing.T) {
	req, err := http.NewRequest("POST", apiPayments, bytes.NewBufferString("invalid"))
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

func TestHandlePaymentsGETNotAllowedSeven(t *testing.T) {
	req, err := http.NewRequest("GET", apiPayments, nil)
	if err != nil {
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlePayments)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf(statusErrorMessage, status, http.StatusMethodNotAllowed)
	}
}

func TestGetProductsPOSTNotAllowedEight(t *testing.T) {
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

func TestHandlePaymentsEmptyIDTen(t *testing.T) {
	req, err := http.NewRequest("POST", apiPayments, strings.NewReader(`{"id":"","amount":100,"cardNumber":"1234567812345678","cardExpiry":"01/23","cardCvv":"123"}`))
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
		t.Errorf(bodyErrorMessage,
			rr.Body.String(), expected)
	}
}
