SELECT u.*, prod.*, purc.*
FROM product prod,
     purchase purc,
     user_account u
WHERE purc.product_id = prod.id
  AND purc.user_id = u.id
  AND purc.purchase_date < '2019-04-14 00:00:00';

SELECT u.*,
       prod.*,
       purc.*
FROM user_account AS u
         JOIN purchase AS purc ON purc.user_id = u.id AND purc.purchase_date < '2019-04-14 00:00:00'
         JOIN product AS prod ON purc.product_id = prod.id;
