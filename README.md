<h2>GRPC Server</h2>
<h3>Cервис поиска книг и их авторов</h3>
В качестве GRPC-клиента использовал утилиту [Evans](https://github.com/ktr0731/evans),
которой при вызове передаем путь к .proto-файлу с определением сервиса.
Команда для запуска GRPC-клиента: <code>.\evans.exe <директория, в которую склонировали проект>\bookstore\api\proto\bookstore.proto -p 8080</code>

Теперь соберём бинарный файл и запустим базу данных с тестовыми данными в контейнере. Для этого, находясь в корне проекта, выполним <code>make up_db</code> - эта команда создаст из Dockerfile образ БД и запустит из этого образа контейнер. Затем выполним <code>make run_server</code>, в результате будет собран и запущен grpc-сервер, слушающий 8080 порт.

Теперь с помощью Evans можно вызвать удалённый метод: <code>call GetAuthors</code> либо <code>call GetBooks</code>, где в первом случае аргументом будет название искомой книги, во втором - имя автора.
Чтобы протестировать работу методов, попробуйте сделать тестовые запросы: <code>call getBooks</code> с аргументом "Lisa Urry" (без кавычек) и <code>call getAuthors</code> c "Campbell Biology (Campbell Biology Series)"