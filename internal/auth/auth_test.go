package auth

import (
	"net/http/httptest"
	"testing"
)



func TestGetAPIKey(t *testing.T){

	req1 := httptest.NewRequest("GET","/",nil)
	req1.Header.Set("Authorization","abc123")	

	_, err := GetAPIKey(req1.Header)
	if err == nil{
		t.Fatal("Error")
	}

	req2 := httptest.NewRequest("GET","/",nil)
	req2.Header.Set("Authorization","")

	_, err = GetAPIKey(req2.Header)
	if err == nil{
		t.Fatal("Error")
	}

	req3 := httptest.NewRequest("GET","/",nil)
	req3.Header.Set("Authorization","ApiKey abc123")

	got, err := GetAPIKey(req3.Header)
	if err != nil && got == "abc123"{
		t.Fatal("Error")
	}
}