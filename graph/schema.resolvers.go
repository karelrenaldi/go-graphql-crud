package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"mariadb/graph/generated"
	"mariadb/graph/model"
	"mariadb/internal/handlers"
	"mariadb/internal/models"
)

func (r *mutationResolver) CreateKtp(ctx context.Context, input model.KtpBody) (*models.Ktp, error) {
	return handlers.CreateKtpHandler(ctx, input)
}

func (r *mutationResolver) DeleteKtp(ctx context.Context, id int64) (bool, error) {
	return handlers.DeleteKtpHandler(ctx, id)
}

func (r *mutationResolver) EditKtp(ctx context.Context, id int64, input model.KtpBody) (*models.Ktp, error) {
	return handlers.EditKtpHandler(ctx, id, input)
}

func (r *queryResolver) Ktp(ctx context.Context) ([]*models.Ktp, error) {
	return handlers.GetAllKtpHandler(ctx)
}

func (r *queryResolver) PaginationKtp(ctx context.Context, input model.Pagination) (*model.PaginationResultKtp, error) {
	return handlers.GetPaginationKtpHandler(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
