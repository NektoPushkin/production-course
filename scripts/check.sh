set -eu
# Проверка исполняемым файлом, собранным через go test -c
exec ./app.test
# Сервер
exec ./server