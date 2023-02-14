package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/kwangsing3/graphql-golang-basic/dbhandler"
	"github.com/kwangsing3/graphql-golang-basic/graph/model"
)

// CreateStock is the resolver for the createStock field.
func (r *mutationResolver) CreateStock(ctx context.Context, input model.NewStock) (*model.Stock, error) {
	res, err := dbhandler.DB.InsertStock(input)
	return res, err
}

// InsertRecord is the resolver for the insertRecord field.
func (r *mutationResolver) InsertRecord(ctx context.Context, input model.NewRecord) (*model.DailyRecord, error) {
	data := model.DailyRecord{
		Date:          input.Date,
		TradingVolume: input.TradingVolume,
		TradingPrice:  input.TradingPrice,
		OpenPrice:     input.OpenPrice,
		HighestPrice:  input.HighestPrice,
		LowestPrice:   input.LowestPrice,
		ClosePrice:    input.ClosePrice,
		PriceDiff:     input.PriceDiff,
		TransAmount:   input.TransAmount,
	}
	_, err := dbhandler.DB.InsertRecord(input.Code, data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// Stock is the resolver for the stock field.
func (r *queryResolver) Stock(ctx context.Context, code string) (*model.Stock, error) {
	res, err := dbhandler.DB.GetStockByCode(code)
	return res, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
