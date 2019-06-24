Задания:

---Изучение работы со структурами в Go---

Изучить работу со структурами в языке Go
Изучить формат сериализации JSON
Прочитать файл example.json из корня проекта при помощи стандартной библиотеки Go
Дессериализовать структуру данных из файла example.json в структуру данных на языке Go
Распечатать в консоль полученную структуру

---Изучение работы с HTTP протоколом в Go---

Ознакомиться с протоколом HTTP: https://medium.com/@andv/how-to-web-%D0%BA%D1%80%D0%B0%D1%82%D0%BA%D0%BE-%D0%BF%D1%80%D0%BE-http-96d9a3c9df79
Изучить пример создания примитивного HTTP сервера на Go: https://golang.org/doc/articles/wiki/#tmp_3
Написать обработчик запроса /users. Обработчик запроса должен десериализовать структуру данных из файла example.json. Из полученного объекта, обработчик должен взять первого пользователя, сериализовать его обратно в json формат и вернуть клиенту в качестве HTTP ответа со статусом 200.

---Изучение работы с API сторонних сервисов в Go---

Ознакомиться с понятием API и способами его использования
Изучить работу со стандартным пакетом http в Go
Выбрать 1 из 3 предлагаемых сторонних API из дополнения к задаче
Написать функцию, получающую данные из выбранного сервиса и дессериализующую их в заранее опредеденную структуру данных

Для веб-сервера объявить несколько роутов:
корневой роут (/), выводящий приветственные слова "My first API v1"
роут /list, выводящий список всех элементов структуры, полученных из API
роут /list/first, выводящий первый элемент из списка всех элементов структуры, полученных из API
роут /list/latest, выводящий последний элемент из списка всех элементов структуры, полученных из API

Дополнение к задаче (API):
API компании SpaceX, предоставляющее данные о запусках космических летательных аппаратов: https://github.com/r-spacex/SpaceX-API
API криптовалютной биржи Nexchange. Для задачи будем использовать метод получения всех доступных для торговли криптовалют на бирже: https://nexchange2.docs.apiary.io/#reference/0/get-currencies/get-currencies?console=1
API игровых событий Dota 2. Для задачи будем использовать метод получения всех про-команд с их статистикой: https://api.opendota.com/api/teams

---Изучение языка запросов SQL на примере СУБД PostgreSQL---

Установить на рабочую машину СУБД PostgreSQL версии 11.2: https://www.enterprisedb.com/downloads/postgres-postgresql-downloads
Ознакомиться с основными положениями и идеями языка запросов SQL
С помощью утилиты pgAdmin, установленной вместе с пакетом разработки PostgreSQL, создать базу данных mydb, управляемую пользователем postgres
Изучить основной синтаксис языка SQL

При помощи языка запросов SQL создать таблицу user_account со слудующими колонками:
id (является первичным ключом)
firstname
lastname
age
external_id

Со следующими ограничениями целостности таблицы:
все колонки не могут содержать значение NULL
колонка возраста не может быть меньше 0 и больше 150
имя не может быть идентично фамилии
фамилия не может быть идентична имени
внешний идентификатор (external_id) должен быть уникальным. Разрешается воспользоваться типом UUID для этой колонки
внешний идентификатор (external_id) должен генерироваться автоматически при создании новой записи в таблице

Проверить работоспособность ограничений написав скрипт на вставку данных о 3 людях

В собственной директории проекта создать директорию sql, в ней создать 2 файла с расширением .sql:
create_table_user_account.sql (содержит скрипт создания таблицы)
insert_users.sql (содержит скрипт заполнения таблицы)

---Продолжение изучения языка запросов SQL на примере СУБД PostgreSQL: Внешние ключи (Foreign keys)---

Легенда: предполагется, что пользователь будет совершать покупки из имеющегося набора товаров. Товар может иметь 2 типа: подписка (SUB) и единоразовая покупка (COMMON). Товар имеет наименование и стоимость. Покупка является связью имеющихся товаров с аккаунтами пользователей, которые их приобрели.

Изучить работу с внешними ключами (REFERENCES) в PostgreSQL
Изучить работу с датами в PostgreSQL
При помощи языка запросов SQL создать таблицы:

Таблица product (товар) с колонками:
id (является первичным ключом)
name
price
type

Со следующими ограничениями целостности таблицы:
все колонки не могут содержать значение NULL
название товара не может быть пустым
название товара не может быть больше 100 символов
цена товара не может быть меньше 0
цена товара - дробное число
тип товара не может быть пустым

Таблица purchase (покупка) с колонками:
id (является первичным ключом)
user_id
product_id
purchase_date

Со следующими ограничениями целостности таблицы:
все колонки не могут содержать значение NULL
user_account_id - внешний ключ на таблицу user_account
product_id - внешний ключ на таблицу product
purachase_date - имеет тип даты (TIMESTAMP), не может быть меньше текущего времени

---Продолжение изучения языка запросов SQL на примере СУБД PostgreSQL: Объединения (JOINS)---

На основе данных, занесенных в базу данных в задаче №6, написать SQL скрипты, выполняющие выборку данных по следующим критериям:
выбрать всех пользователей, купивших товар с типом подписка (SUB)
выбрать все покупки, совершенные раньше прошлого воскресенья
выбрать всех пользователей, купивших товар дороже 500 условных единиц. Если таких товаров нет, необходимо добавить их в таблицу
выбрать всех пользователей младше 20 лет, купивших товар дороже 500 условных единиц

Добавить каждый скрипт в отдельный файл с расширением .sql в директорию sql

---Создание CRUD сервиса на основе REST---

Изучить методы HTTP: POST, GET, PUT, DELETE
Изучить понятие CRUD
Изучить понятие REST
Изучить способы подключения и использования сторонних библиотек в Go:
Go модули: go mod init (https://github.com/golang/go/wiki/modules)
go get (https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies)
Продумать роутинг разрабатываемого веб-сервиса соответсвующий стилю REST API

Написать REST API для  вэб-сервиса, позволяющий выполнять следующие действия:
Добавлять новый товар в БД (POST)
Просмастривать список всех товаров (GET)
Просматривать информацию о конкретном товаре из БД, используя его id (GET)
Обновлять информацию о товаре в БД, используя его Id (PUT)
Удалять товар из БД, используя его id (DELETE)

Проверить работоспособность сервиса, используя любой графический REST-клиент, например, Postman (https://www.getpostman-beta.com/downloads/)
Исходный код сервиса должен находиться в директории web/services/ вашей директории

---Регистрация и авторизация пользователя на основе REST---

Изучить популярные методы хэширования
Написать скрипт изменения уже существующей таблицы user_account, добавив в нее 2 новые колонки: login и password
колонка login должна быть уникальной

Продумать роутинг для методов регистрации и авторизации
Написать REST API для следущих методов:
Зарегистрировать нового пользователя (POST) - метод должен возвращать JSON-объект, содержащий все колонки таблицы, кроме колонки password
Авторизовать пользователя по логину и паролю (POST) - метод должен возвращать JSON-объект, содержащий все колонки таблицы, кроме колонки password

При регистрации проверять, существует ли пользователь с таким login в базе данных, и если существует, то сообщить пользователю об ошибке регистрации и просьбой выбрать другой логин.