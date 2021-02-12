package usecase

import (
	"context"
	"time"

	"github.com/cpartogi/warteg/module/warteg"
	"github.com/cpartogi/warteg/schema/request"
	"github.com/cpartogi/warteg/schema/response"
)

// AuthUsecase will create a usecase with its required repo
type WartegUsecase struct {
	wartegRepo     warteg.Repository
	contextTimeout time.Duration
}

// NewAuthUsecase will create new an contactUsecase object representation of auth.Usecase
func NewWartegUsecase(ar warteg.Repository, timeout time.Duration) warteg.Usecase {
	return &WartegUsecase{
		wartegRepo:     ar,
		contextTimeout: timeout,
	}
}

func (u *WartegUsecase) WartegAdd(ctx context.Context, addw request.Warteg) (wt response.WartegAdd, err error) {
	resp := response.WartegAdd{
		WartegName:        addw.WartegName,
		WartegDesc:        addw.WartegDesc,
		WartegAddr:        addw.WartegAddr,
		WartegContactName: addw.WartegContactName,
		WartegPhone:       addw.WartegPhone,
	}

	req := request.Warteg{
		WartegName:        addw.WartegName,
		WartegDesc:        addw.WartegDesc,
		WartegAddr:        addw.WartegAddr,
		WartegContactName: addw.WartegContactName,
		WartegPhone:       addw.WartegPhone,
	}

	addwarteg, err := u.wartegRepo.WartegAdd(ctx, req)

	if err != nil {
		return resp, err
	}
	return addwarteg, err
}

func (u *WartegUsecase) WartegDelete(ctx context.Context, warteg_id string) (wt response.WartegDelete, err error) {
	resp := response.WartegDelete{
		WartegId: warteg_id,
	}

	delwarteg, err := u.wartegRepo.WartegDelete(ctx, warteg_id)
	if err != nil {
		return resp, err
	}

	return delwarteg, err
}

func (u *WartegUsecase) WartegUpdate(ctx context.Context, warteg_id string, uwt request.WartegUpdate) (wt response.WartegUpdate, err error) {
	resp := response.WartegUpdate{
		WartegId:          warteg_id,
		WartegName:        uwt.WartegName,
		WartegDesc:        uwt.WartegDesc,
		WartegAddr:        uwt.WartegAddr,
		WartegContactName: uwt.WartegContactName,
		WartegPhone:       uwt.WartegPhone,
	}

	req := request.WartegUpdate{
		WartegName:        uwt.WartegName,
		WartegDesc:        uwt.WartegDesc,
		WartegAddr:        uwt.WartegAddr,
		WartegContactName: uwt.WartegContactName,
		WartegPhone:       uwt.WartegPhone,
	}

	upwarteg, err := u.wartegRepo.WartegUpdate(ctx, warteg_id, req)

	if err != nil {
		return resp, err
	}

	return upwarteg, err

}

func (u *WartegUsecase) WartegList(ctx context.Context, warteg_name string) (wl []response.WartegList, err error) {
	resp := []response.WartegList{}

	warteglist, err := u.wartegRepo.WartegList(ctx, warteg_name)

	if err != nil {
		return resp, err
	}

	return warteglist, err

}

func (u *WartegUsecase) WartegDetail(ctx context.Context, warteg_id string) (wd response.WartegDetail, err error) {
	resp := response.WartegDetail{}

	mdetail, err := u.wartegRepo.WartegDetail(ctx, warteg_id)

	if err != nil {
		return resp, err
	}

	return mdetail, err
}
