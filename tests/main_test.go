package tests

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"owen2020/app/apputil/applog"
	"owen2020/routes"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// test print无屏幕输出，需要使用testing.Log testing.Logf
// https://stackoverflow.com/questions/23205419/how-do-you-print-in-a-go-test-using-the-testing-package
func TestGetList(t *testing.T) {
	err := godotenv.Load("/Users/owen/go/src/learn-go/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := routes.SetUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/activity/v1/activity", nil)
	router.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	applog.Logger.Info(string(body))
	t.Log(string(body))
	t.Log(w.Code)

	assert.Equal(t, 200, w.Code)
}

func Test0001(test *testing.T) {
	fmt.Println("TestEncrypt = ", "123")
}
