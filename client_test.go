package pepo

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

var client *Client

func init() {
	godotenv.Load()
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	client = NewClient(key, secret)
}

func TestCreateListSuccess(t *testing.T) {
	defer gock.Off()
	jsonFile, err := os.Open("data/list_create_success.json")
	if err != nil {
		t.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	gock.New(client.GetBaseURL()).
		Post("/list/create").
		Reply(200).
		JSON(byteValue)

	resp, err := client.CreateList("superlist", "list", nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	st.Expect(t, resp.Message, "List created with success")
	st.Expect(t, resp.Data.List.Name, "superlist")
}

func TestCreateListError(t *testing.T) {
	defer gock.Off()
	jsonFile, err := os.Open("data/list_create_error.json")
	if err != nil {
		t.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	gock.New(client.GetBaseURL()).
		Post("/list/create").
		Reply(200).
		JSON(byteValue)

	resp, err := client.CreateList("superlist", "list", nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	st.Expect(t, resp.Error, "VALIDATION_ERROR")
	map1 := map[string]interface{}{
		"list_source": []string{
			"List Source is Mandatory",
		},
	}
	if fmt.Sprint(map1) != fmt.Sprint(resp.ErrorMessage) {
		t.Fatal("Should be equal")
	}
}
