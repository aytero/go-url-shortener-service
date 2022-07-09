package api

import (
	"testing"
)

//func initTestRouter(testApi Api) {
//	gin.SetMode(gin.TestMode)
//	r := gin.Default()
//	r.GET("/:shortUrl", testApi.GetFullUrlByShort)
//	r.POST("/", testApi.PostUrl)
//}

func TestHandlers(t *testing.T) {
	//var testApi = Api{}
	//initTestRouter(testApi)

	//req := httptest.NewRequest(http.MethodPost, "/", nil)
	//w := httptest.NewRecorder()
	//w := gin.NewResponseWriter()
	//testApi.PostUrl(&gin.Context{Request: req})
	//testApi.PostUrl(ctx)
	//res := w.Result()
	//res := req.Response
	//fmt.Println(res)
	//defer res.Body.Close()
	//_, err := io.ReadAll(res.Body)
	//if err != nil {
	//	t.Errorf("Expected: nil; got: %v\n", err)
	//}
	//if string(data) != "aaa" {
	//	t.Errorf("Expected: %s; got: %v\n", "aaa", string(data))
	//}
}
