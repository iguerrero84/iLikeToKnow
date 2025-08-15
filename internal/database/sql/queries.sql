-- name: CreateEvent :one
INSERT INTO public.events(
    id,
    title, 
    description,
    start_time,
    end_time,
    created_at
) VALUES (
    @id,
    @title,
    @description,
    @start_time,
    @end_time,
    @created_at
) RETURNING *;

-- name: ListEvents :many
SELECT id, title, description, start_time, end_time, created_at
FROM events
ORDER BY start_time ASC;

-- name: GetEventById :one
SELECT *
FROM
public.events
where id = @id
order by start_time;