SELECT u.*, prod.*
FROM product prod,
     purchase purc,
     user_account u
WHERE purc.product_id = prod.id
  AND purc.user_id = u.id
  AND prod.price > 500
  AND u.age < 20;


SELECT u.*,
       prod.*
FROM user_account AS u
         JOIN purchase AS purc ON purc.user_id = u.id
         JOIN product AS prod ON purc.product_id = prod.id AND prod.price > 500 AND u.age < 20;