SELECT u.*
FROM product prod,
     purchase purc,
     user_account u
WHERE purc.product_id = prod.id
  AND purc.user_id = u.id
  AND prod.type = 'SUB';


SELECT u.*
FROM user_account AS u
         JOIN purchase AS purc ON purc.user_id = u.id
         JOIN product AS prod ON purc.product_id = prod.id AND prod.type = 'SUB';