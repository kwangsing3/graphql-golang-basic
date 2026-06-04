# graphql-golang-basic

基於 **GraphQL + gqlgen** 設計介面、與 **MongoDB Atlas** 互動的 Golang Web Service 範例專案，主要用於快速部署與教學示範。

## 實現的 GraphQL 介面

| 類型 | 方法 | 說明 |
|------|------|------|
| Query | `getStock(code)` | 查詢股票資料及歷史紀錄 |
| Mutation | `createStock(input)` | 新增股票基本資料 |
| Mutation | `insertRecord(input)` | 新增每日交易紀錄 |

## 技術棧

| 層級 | 套件 |
|------|------|
| 語言 | Go 1.19 |
| GraphQL 框架 | [gqlgen v0.17](https://gqlgen.com/) |
| 資料庫 | MongoDB（官方 Go Driver）|
| GraphQL 介面 | gqlgen playground（內建）|

## 快速開始

### 1. 取得原始碼

```bash
git clone https://github.com/kwangsing3/graphql-golang-basic
cd graphql-golang-basic
```

### 2. 設定 MongoDB 連線

編輯 `dbhandler/dbhandler.go`，將 MongoDB Atlas 連線字串替換為你的：

```go
var DB, _ = NewDBHandler("mongodb+srv://<username>:<password>@<cluster>.mongodb.net/test")
```

> MongoDB Atlas 免費叢集可至 [https://www.mongodb.com/atlas](https://www.mongodb.com/atlas) 建立。
> 連線字串格式：`mongodb+srv://user:pass@cluster0.xxxxx.mongodb.net/dbname`

### 3. 安裝依賴並啟動

```bash
go mod tidy
go run server.go
```

伺服器預設監聽 Port **80**，開啟瀏覽器至：

```
http://localhost/
```

即可使用 GraphQL Playground。

## GraphQL Schema

```graphql
type Query {
  stock(code: String!): Stock
}
type Mutation {
  createStock(input: NewStock!): Stock
  insertRecord(input: NewRecord!): DailyRecord
}

type Stock {
  code: String!
  name: String!
  historicalRecord: [DailyRecord]!
}
type DailyRecord {
  date: String!
  tradingVolume: Float!
  tradingPrice: Float!
  openPrice: Float!
  closePrice: Float!
}
```

## 使用範例

**新增股票**
```graphql
mutation {
  createStock(input: { code: "2330", name: "台積電" }) {
    code
    name
  }
}
```

**查詢股票**
```graphql
query {
  stock(code: "2330") {
    name
    historicalRecord {
      date
      closePrice
    }
  }
}
```

## 擴充 Schema

1. 修改 `./graph/schema.graphqls`
2. 重新執行 gqlgen codegen：
   ```bash
   go run github.com/99designs/gqlgen generate
   ```
3. 實作 `./graph/schema.resolvers.go` 中的新方法

## 授權

MIT — 可自由用於商業、個人、教學用途。