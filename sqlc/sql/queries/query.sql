-- name: ListCategories :many
SELECT * FROM categories;

-- name: ListCourses :many
SELECT * FROM courses;

-- name: GetCategoryByID :one
SELECT * FROM categories
WHERE id = ?;

-- name: CreateCategory :execresult
INSERT INTO categories (id, name, description)
VALUES (?, ?, ?);

-- name: UpdateCategory :exec
UPDATE categories SET name = ?, description = ?
WHERE id = ?;