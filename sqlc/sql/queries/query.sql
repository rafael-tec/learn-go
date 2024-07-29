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

-- name: DeleteCategoryByID :exec
DELETE FROM categories WHERE id = ?;

-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, price, category_id)
VALUES (?, ?, ?, ?, ?);

-- name: ListCoursesJoinCategory :many
SELECT co.*, co.name as category_name FROM courses co
JOIN categories ca ON co.category_id = ca.id;