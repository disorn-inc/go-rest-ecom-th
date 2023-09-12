SELECT p.id, p.title,
       (SELECT array_to_json(array_agg(t))
       FROM (SELECT i.id, i.url, i.filename
       FROM images i
       WHERE i.product_id=p.id) AS "t") as image
FROM products p;

SELECT p.id, p.title,
       (SELECT to_jsonb(t)
       FROM (SELECT i.id, i.url, i.filename
       FROM images i
       WHERE i.product_id=p.id LIMIT 1) AS "t") as image
FROM products p;

SELECT p.id, p.title,
       (SELECT COALESCE(array_to_json(array_agg(t)), '[]'::json)
       FROM (SELECT i.id, i.url, i.filename
       FROM images i
       WHERE i.product_id=p.id) AS "t") as image
FROM products p;

SELECT json_build_object('id', p.id, 'title', p.title) FROM products "p";