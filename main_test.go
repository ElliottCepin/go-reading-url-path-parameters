package main

import (
	"net/http/httptest"
	"testing"
	"strings"
)


func TestGreet(t *testing.T) {
	names := [6]string{"Iyan", "Elliott", "Jordan", "Jocelyn", "Foucalt", "Big"}
	for i := 0; i<len(names); i++ {
		greeting := "Hello, " + names[i] + "!"
		req := httptest.NewRequest("GET", "/greet/" + names[i], nil)
		rec := httptest.NewRecorder()

		serveGreet(rec, req)

		if (rec.Body.String() != greeting) {
			t.Errorf("Got %q, expected %q", rec.Body.String(), greeting)
		}
	}
}

func TestShout(t *testing.T) {
	names := [6]string{"Iyan", "Elliott", "Jordan", "Jocelyn", "Foucalt", "Big"}
	for i := 0; i<len(names); i++ {
		greeting := "Hello, " + names[i] + "!"
		req := httptest.NewRequest("GET", "/shout/" + names[i], nil)
		rec := httptest.NewRecorder()

		serveShout(rec, req)

		if (rec.Body.String() != strings.ToUpper(greeting)) {
			t.Errorf("Got %q, expected %q", rec.Body.String(), strings.ToUpper(greeting))
		}
	}
}
