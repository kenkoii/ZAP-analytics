package app

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//Listing機能をテストする
func TestServerListing(t *testing.T) {
	httptest.NewServer(GetMainEngine())
	testListing(t)
}

//データを引っ張ってくるのをテストする
func TestServerGet(t *testing.T) {
	testGet(t)
}

//POSTした結果をテストする
func TestServerPOST(t *testing.T) {
	testPost(t)
}

//Listingをテストする。goapp serveで実際のサーバは8080Portで既に立っているためそこにGETを送る。
func testListing(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/api/v1/analytics/userproperties?start=08-11-2017&end=08-12-2017&filter=LoginDate")
	if err != nil {
		t.Fatalf("geterror %v", err)
		return
	}
	// 関数を抜ける際に必ずresponseをcloseするようにdeferでcloseを呼ぶ
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("%v", resp.StatusCode)
	}
}

//GETも同様にTodoを取得するためにGETリクエストを適当なURLに行う
//testIDCaseはIDとテストが通るか(通ると想定するケースはTrue)をMapとして持っておく。
//期待した結果ともし違う場合はt.Fatalfメソッドでテストを失敗として処理する。
func testGet(t *testing.T) {
	testIDCase := map[string]bool{"1": true, "2": true, "n": false}
	for key, value := range testIDCase {
		log.Println(key, value)
		resp, _ := http.Get("http://localhost:8080/get/" + key)

		if resp.StatusCode == 200 != value {
			t.Fatalf("%v", resp.StatusCode)
		}
		defer resp.Body.Close()
	}
}

//POSTもGETTodoのGETがPOSTになっただけでやることはほとんど変わらない。
//通ると想定するケースはTrue、通らないと想定するケースはfalseを各String値に紐付けてテストを行う。
func testPost(t *testing.T) {
	strs := map[string]bool{`{"Title" : "Title1" , "Description" : "いろいろやる"}`: true,
		`{"Title" : "Title1" , "Description2" : "いろいろやる"}`:                         false,
		`{"Title2" : "Title1" , "Description" : "いろいろやる"}`:                         false,
		`{"Title" : "gajgapgndfsngdsgkhnpsdknhsdgkhmkgsdhn" , "Description" : ""}`: false,
		`{"Title" : "" , "Description" : "いろいろやる"}`:                                false}

	for key, value := range strs {
		log.Println(key, value)
		url := "http://localhost:8080/post/"

		req, err := http.NewRequest(
			"POST",
			url,
			bytes.NewBuffer([]byte(key)),
		)

		if err != nil {
			t.Fatalf("Error Occured %v", err)
		}

		client := &http.Client{Timeout: time.Duration(15 * time.Second)}
		resp, _ := client.Do(req)
		if resp.StatusCode == 200 != value {
			t.Fatalf("Error Occured: key is %v", key)
		}
		defer resp.Body.Close()
	}
}
