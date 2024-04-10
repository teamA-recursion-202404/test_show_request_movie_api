package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 詳細取得のAPIエンドポイント
	// レスポンスを別ファイル res.json に記載しています
	url := "https://api.themoviedb.org/3/movie/359410?language=en-US"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer XXXXXXXXXXX") // XXXX... = APIリードアクセストークン

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Println("Completed")
}