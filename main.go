package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type RequestBody struct {
	Requests
}

type Requests []struct {
	UpdateCells
}

type UpdateCells struct {
	Start
	Rows
	Fields string
}

type Start struct {
	SheetID     int
	RowIndex    int
	ColumnIndex int
}

type Rows []struct {
	Values
}

type Values []struct {
	UserEnteredValue
}

type UserEnteredValue struct {
	StringValue string
}

func main() {
	// .envの値を取得
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error")
	}
	spreadSheetId := os.Getenv("SPREADSHEET_ID")

	// httpクライアント設定
	client := &http.Client{}
	client.Timeout = time.Second * 60

	// リクエストBody作成
	// values := Values{
	// 	{
	// 		UserEnteredValue: UserEnteredValue{
	// 			StringValue: "hoge",
	// 		},
	// 	},
	// }
	// fmt.Println(values)
	requestBody := RequestBody{
		Requests: Requests{
			{
				UpdateCells: UpdateCells{
					Start: Start{
						SheetID:     0,
						RowIndex:    0,
						ColumnIndex: 0,
					},
					Rows: Rows{
						{
							Values: Values{
								{
									UserEnteredValue: UserEnteredValue{
										StringValue: "value01",
									},
								},
								{
									UserEnteredValue: UserEnteredValue{
										StringValue: "value02",
									},
								},
							},
						},
					},
					Fields: "userEnteredValue",
				},
			},
		},
	}

	// リクエスト作成
	b, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("error...")
	}

	requestUrl := fmt.Sprintf("https://sheets.googleapis.com/v4/spreadsheets/%s:batchUpdate", spreadSheetId)
	fmt.Println(requestUrl)
	req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewReader(b))
	if err != nil {
		fmt.Println("error!")
	}

	// リクエストを実行
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error!!")
	}
	fmt.Println(res)
	fmt.Println(1)

}
