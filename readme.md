Требования
Go 1.21+
Для тестирования API: curl + терминал

Установка
git clone https://github.com/Bogos387/TaskService.git
cd task-service
go mod init task-service
go mod tidy

Запуск
go run main.go

API
(Посредством командной строки (cmd) в Windows)
1. Создание задачи
команда:
curl -X POST -H "Content-Type: application/json" -d "{\"input_data\":\"TaskInputData\"}" http://localhost:8080/tasks/create
ответ:
{"id":"503a4071-3d75-443f-a3dd-c5ac16998c0a"}

2. Получение статуса задачи
команда:
curl http://localhost:8080/tasks/get?id=503a4071-3d75-443f-a3dd-c5ac16998c0a
ответ:
{
"id":"503a4071-3d75-443f-a3dd-c5ac16998c0a",
"status":"completed",
"result":"Processing completed for TaskInputData",
"created_at":"2025-06-23T16:22:52.7226352+03:00",
"started_at":"2025-06-23T16:22:52.7226352+03:00",
"completed_at":"2025-06-23T16:23:12.7233943+03:00",
"input_data":"TaskInputData",
"duration":"20.0007591s"}

3. Удаление задачи
команда:
curl http://localhost:8080/tasks/delete?id=503a4071-3d75-443f-a3dd-c5ac16998c0a


