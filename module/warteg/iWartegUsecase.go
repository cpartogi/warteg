package warteg

import (
	"context"

	"github.com/cpartogi/warteg/schema/request"
	"github.com/cpartogi/warteg/schema/response"
)

// Usecase is
type Usecase interface {
	WartegAdd(ctx context.Context, addw request.Warteg) (wt response.WartegAdd, err error)
	WartegDelete(ctx context.Context, warteg_id string) (wt response.WartegDelete, err error)
	WartegUpdate(ctx context.Context, warteg_id string, uwt request.WartegUpdate) (wt response.WartegUpdate, err error)
	WartegList(ctx context.Context, warteg_name string) (wl []response.WartegList, err error)
	WartegDetail(ctx context.Context, warteg_id string) (wd response.WartegDetail, err error)
}
