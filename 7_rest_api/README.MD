# Создай CRUD-ы для твоих постов и комментов. Ответ от АПИ должен реализовать 2 формата представления данных: XML, JSON. Используй стандартный net/http пакет для web сервера.

## 1_rest_api 
launch on localhost:8080

just first version - all handlers in one file handlers.go . 
Version created without db chose logic. Working only with MySQL db.

## 2_rest_api
Have been realised with net/http package from standard library. Demonstration is 
able on [heroku](https://nix-study-api.herokuapp.com/) source.
* Depend on DB sources need to be checkout and if it necessary switch the mode that you need (const Mode) in dbase dir.
