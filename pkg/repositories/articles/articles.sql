-- name: GetAll
SELECT id,title,url,point FROM articles

-- name: New
INSERT INTO articles(title,url,point) VALUES(:1,:2,:3)

-- name: DeleteById
DELETE FROM articles where id=:1
