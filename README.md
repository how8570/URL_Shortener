# URL Shortener

## Run code

just simply run the code
  ```bash
  go run ./main.go
  ```

## Known Issues
  - Hash may collision, if you have a lot of short links. If it happens, it will **cause crash because db write lock**.

  - if two client send same url at the same time, **not sure what will happen**.

  
##  Others / DevLog

### 4/4/2022

- 今天來補寫 unit test 跟一些 API 的輸入防呆

- 在寫 test 的時候注意到因為 db 並不會有連線而不能測試到有關 db 操作
  的部分，於是只做了一部分能預先檢查的測試，例如錯誤的參數輸入。
  - 目前已知有 mock db 之類的 lib 可以用，但沒有實作 db 檢測的部分

- 處理 API 輸入防呆/防止重複
  - 去除 url 的 http://、https:// 再存入db

- 實作 Expire time 處理 data base 的部分
  - 時間的部分打算使用 unix timestamp 來存，不過範例給的是 
    "YYYY-MM-DDThh:mm:ssZ" 的樣子。
    所以做了一個 Method 轉換使能夠接受 unix timestamp 或是
    範例的 format

- 修了一些忘記還 db 讀寫全縣的地方

- 寫了一些 Log

### 4/1/2022

- 今天剛好是愚人節，近期終於比較有空了，所以這天晚上像說就來把這份東西
  寫一寫。

- 於是乎今天把兩個 API 部分都先處理到能夠正常使用了，目前用 postman 
  手工測試起來好像沒有太多問題了。應該處理一下像是檢查輸入輸出，寫測試
  等等的部分應該就大功告成了吧 ouo

### 3/??/2022 (Motivation)
- 近期在 Dcard lab 看到了校園招募的資訊，題目是做一個 URL 縮網址服務，
  想到我也還沒有一個縮網址服務，所以就想順便來做一個了。

- 目前先預計使用 Golang + SQLite3 達成，原因為這些工具較為熟悉
  - router 使用 gorilla/mux
