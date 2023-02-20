# wondelful-factorial.http.go

🍣🍣🍣 再帰関数もループ処理も用いずに、階乗を計算する！  

![成果物](./docs/img/fruit.gif)  

考え方は再帰関数と同じです。  
ただし、再帰関数を使わずに、HTTP通信で自分自身に対してリクエストを送っています。  

例えば、`5`の階乗を求めるとします。  
以下のように、`5`から`1`までの数値をリクエストパラメータとして送信します。  

- `http://localhost:8080/5`
- `http://localhost:8080/4`
- `http://localhost:8080/3`
- `http://localhost:8080/2`
- `http://localhost:8080/1`

`/:num`で`num`が`1`の場合にはレスポンスデータに`1`を返すようにしています。  
それ以外の場合には、`num * http(num - 1)`を返しています。  
`http(num - 1)`とは、HTTPリクエストを送信して取得したレスポンスデータです。  
これによって疑似的に再帰関数を実現しています。  

但し、説明の通り再帰関数とは技術的に異なるため、再帰関数もループ処理も用いない斬新な階乗の計算方法です。  

🐶ダフル階乗計算方法ですね！  

## 実行方法

```shell
docker compose up app -d [--build]
```

エンドポイントは`.env`ファイルで適切に設定してください。  
