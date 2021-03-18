package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cpartogi/warteg/constant"
	"github.com/cpartogi/warteg/schema/request"
	"github.com/cpartogi/warteg/schema/response"
)

const addWarteg = `-- name: AddWarteg :one
INSERT INTO tb_warteg (
	warteg_id,
    warteg_name,
    warteg_desc,
    warteg_addr,
	warteg_contact_name,
	warteg_phone
) VALUES (
    uuid(),
    ?,
    ?,
    ?,
	?,
	?
)
`

func (q *Queries) WartegAdd(ctx context.Context, addw request.Warteg) (wt response.WartegAdd, err error) {
	result, err := q.db.ExecContext(ctx, addWarteg,
		addw.WartegName,
		addw.WartegDesc,
		addw.WartegAddr,
		addw.WartegContactName,
		addw.WartegPhone,
	)

	if err != nil {
		return
	}

	rows, _ := result.RowsAffected()

	if rows != 1 {
		return
	}

	i := response.WartegAdd{
		WartegName:        addw.WartegName,
		WartegDesc:        addw.WartegDesc,
		WartegAddr:        addw.WartegAddr,
		WartegContactName: addw.WartegContactName,
		WartegPhone:       addw.WartegPhone,
	}

	return i, err
}

const deleteWarteg = `-- name: DeleteWarteg :one
UPDATE tb_warteg SET is_delete=1, updated_date=CURRENT_TIMESTAMP(3) WHERE warteg_id = ?
`

func (q *Queries) WartegDelete(ctx context.Context, warteg_id string) (wt response.WartegDelete, err error) {
	result, err := q.db.ExecContext(ctx, deleteWarteg, warteg_id)

	if err != nil {
		return
	}

	rows, _ := result.RowsAffected()

	if rows != 1 {
		err = constant.ErrNotFound
	}

	i := response.WartegDelete{
		WartegId: warteg_id,
	}

	return i, err
}

const updateWarteg = `-- name: UpdateWarteg :one
UPDATE tb_warteg SET warteg_name=?, warteg_desc=?, warteg_addr=?, warteg_contact_name=?, warteg_phone=?, updated_date=CURRENT_TIMESTAMP(3) WHERE warteg_id = ?
`

func (q *Queries) WartegUpdate(ctx context.Context, warteg_id string, uwt request.WartegUpdate) (wt response.WartegUpdate, err error) {
	result, err := q.db.ExecContext(ctx, updateWarteg,
		uwt.WartegName,
		uwt.WartegDesc,
		uwt.WartegAddr,
		uwt.WartegContactName,
		uwt.WartegPhone,
		warteg_id,
	)

	if err != nil {
		return
	}

	rows, err := result.RowsAffected()

	if rows != 1 {
		err = constant.ErrNotFound
	}

	i := response.WartegUpdate{
		WartegId:          warteg_id,
		WartegName:        uwt.WartegName,
		WartegDesc:        uwt.WartegDesc,
		WartegAddr:        uwt.WartegAddr,
		WartegContactName: uwt.WartegContactName,
		WartegPhone:       uwt.WartegPhone,
	}

	return i, err
}

func (q *Queries) WartegList(ctx context.Context, warteg_name string) (wl []response.WartegList, err error) {
	qWarteg := `SELECT warteg_id, warteg_name, warteg_addr FROM tb_warteg WHERE is_delete=0 AND warteg_name like '%%%s%%'  ORDER BY warteg_name LIMIT 50`

	listWarteg := fmt.Sprintf(qWarteg, warteg_name)

	rows, err := q.db.QueryContext(ctx, listWarteg)

	if err != nil {
		return
	}

	defer rows.Close()

	var y []response.WartegList
	var i response.WartegList

	c := 0

	for rows.Next() {
		_ = rows.Scan(
			&i.WartegId,
			&i.WartegName,
			&i.WartegAddr,
		)
		y = append(y, i)
		c++
	}

	//return not found
	if c == 0 {
		err = constant.ErrNotFound
	}
	return y, err

}

const getWartegDetail = `-- name: WartegDetail :one
SELECT warteg_id, warteg_name, warteg_desc, warteg_addr, warteg_contact_name, warteg_phone FROM tb_warteg
WHERE warteg_id = ?
`

func (q *Queries) WartegDetail(ctx context.Context, warteg_id string) (wd response.WartegDetail, err error) {
	row := q.db.QueryRowContext(ctx, getWartegDetail, warteg_id)
	var i response.WartegDetail
	err = row.Scan(
		&i.WartegId,
		&i.WartegName,
		&i.WartegDesc,
		&i.WartegAddr,
		&i.WartegContactName,
		&i.WartegPhone,
	)

	if err == sql.ErrNoRows {
		err = constant.ErrNotFound
	}

	return i, err
}
