// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: travelChecklist.sql

package db

import (
	"context"
)

const createTravelChecklist = `-- name: CreateTravelChecklist :one
INSERT INTO travel_checklists(
    trip_owner,
    item_task
) VALUES (
    $1, $2
) RETURNING id, trip_owner, item_task, checked
`

type CreateTravelChecklistParams struct {
	TripOwner int64  `json:"trip_owner"`
	ItemTask  string `json:"item_task"`
}

func (q *Queries) CreateTravelChecklist(ctx context.Context, arg CreateTravelChecklistParams) (TravelChecklist, error) {
	row := q.queryRow(ctx, q.createTravelChecklistStmt, createTravelChecklist, arg.TripOwner, arg.ItemTask)
	var i TravelChecklist
	err := row.Scan(
		&i.ID,
		&i.TripOwner,
		&i.ItemTask,
		&i.Checked,
	)
	return i, err
}

const deleteTravelChecklist = `-- name: DeleteTravelChecklist :exec
DELETE FROM travel_checklists WHERE trip_owner = &1
`

func (q *Queries) DeleteTravelChecklist(ctx context.Context) error {
	_, err := q.exec(ctx, q.deleteTravelChecklistStmt, deleteTravelChecklist)
	return err
}

const getTravelChecklist = `-- name: GetTravelChecklist :one
SELECT FROM travel_checklists
WHERE trip_owner = $1
LIMIT 1
FOR NO KEY UPDATE
`

type GetTravelChecklistRow struct {
}

func (q *Queries) GetTravelChecklist(ctx context.Context, tripOwner int64) (GetTravelChecklistRow, error) {
	row := q.queryRow(ctx, q.getTravelChecklistStmt, getTravelChecklist, tripOwner)
	var i GetTravelChecklistRow
	err := row.Scan()
	return i, err
}

const listTravelChecklist = `-- name: ListTravelChecklist :many
SELECT FROM travel_checklists
WHERE trip_owner = $1
LIMIT $2
OFFSET $3
`

type ListTravelChecklistParams struct {
	TripOwner int64 `json:"trip_owner"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

type ListTravelChecklistRow struct {
}

func (q *Queries) ListTravelChecklist(ctx context.Context, arg ListTravelChecklistParams) ([]ListTravelChecklistRow, error) {
	rows, err := q.query(ctx, q.listTravelChecklistStmt, listTravelChecklist, arg.TripOwner, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListTravelChecklistRow{}
	for rows.Next() {
		var i ListTravelChecklistRow
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
