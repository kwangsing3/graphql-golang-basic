# github.com/kwangsing3/graphql-golang-basic
基於GraphQL的設計介面，與MongoDB互動的Golang Web service範例專案，主要用於快速部屬與教學示範

IDE: visual code



# 範例實現的介面

此範例專案實現幾個主要的 GraphQL 介面特性 :

- 查詢 method (getStock)
- 新增 method (createStock)
- 新增紀錄 method (InsertRecord)

# Where to start with?

```bash
$ git clone https://github.com/kwangsing3/graphql-golang-basic
```

# Develop guide

編輯Graphql結構
``` graphql
# GraphQL schema example
# ./graph/schema.graphqls

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
  date: String! #日期
  tradingVolume: Float! #成交股數
  tradingPrice: Float! #成交金額
  openPrice: Float! #開盤價
  closePrice: Float! #收盤價
}

input NewStock {
  code: String!
  name: String!
}

input NewRecord {
  code: String! #證券代號
  name: String! #證券名稱
  date: String! #日期
  tradingVolume: Float! #成交股數
  tradingPrice: Float! #成交金額
  openPrice: Float! #開盤價
  closePrice: Float! #收盤價
}

```

運行腳本來產生以Graphql結構代碼為主體的程式碼
```bash
$ go run github.com/99designs/gqlgen init
```

最後實作 Reslover
```go
// ./graph/schema.resolvers.go
func (r *queryResolver) Stock(ctx context.Context, code string) (*model.Stock, error) {
	res, err := dbhandler.DB.GetStockByCode(code)
	return res, err
}

.......
```
# Screenshot

```url
http://localhost:4000/
```

- Query 查詢範例
  ![alt text](doc/2023-02-02%2023.12.14.png)

- Mutation 範例
  ![alt text](doc/2023-02-02%2023.13.22.png)

# 作者聲明 Statement

此專案是一個基於 Golang、MongoDB、GraphQL 的範例專案，

```
Copyright (C) <year> <copyright holders>

Permission is hereby granted, free of charge, to any person obtaining a copy of this
software and associated documentation files (the "Software"), to deal in the
Software without restriction, including without limitation the rights to use, copy,
modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
and to permit persons to whom the Software is furnished to do so, subject to the
following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.


THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```

此專案的建立目的主要用於

- 教學、範例
- 快速建置可立即部屬使用的 GraphQL Server

任何獲取此專案源碼的衍伸項目皆可以使用於任何包含但不限於 "商業" 、 "個人" 、 "教學使用"，如此專案的衍伸項目涉及觸犯相關國家的法律或規範，作者及相關貢獻者皆無須承擔相關的法律責任。