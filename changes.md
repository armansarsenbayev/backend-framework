# Описание реализованных задач (Resty Microservice & Docker)

В рамках выполнения обеих задач был добавлен новый микросервис `notification-service`, настроено взаимодействие через `resty/v2`, а также проект был полностью контейнеризирован с помощью `Docker` и `Docker-Compose`.

## 1. Новый микросервис и интеграция Resty
*   **Добавлен микросервис:** Создана директория `notification-service` с `main.go`. Он работает на порту `8081` и обрабатывает POST запросы на `/notify`.
*   **Установлен Resty v2:** Выполнено `go get github.com/go-resty/resty/v2`.
*   **Создан клиент `clients/notification_client.go`:** Реализует `OnBeforeRequest` и `OnAfterResponse` хуки (middlewares) для логирования каждого запроса и статус-кода ответа.
*   **Интеграция в основной API:** В `handlers/order_handlers.go` при создании заказа асинхронно вызывается клиент Resty, который отправляет запрос в `notification-service`. В код также добавлена возможность изменять `URL` и хост БД через переменные окружения (`NOTIFICATION_SERVICE_URL`, `DB_HOST`).

## 2. Контейнеризация (Docker & Docker-Compose)
*   **Dockerfile (для основного приложения):** В корне проекта создан `Dockerfile`, который собирает монолит (сервис заказов и ресторанов).
*   **Dockerfile (для микросервиса):** В папке `notification-service` создан второй `Dockerfile`, который собирает бинарник `notification-service`.
*   **docker-compose.yml:** Создан файл конфигурации, запускающий сразу 3 контейнера:
    1.  `app`: Основное приложение на порту `8080`.
    2.  `notification`: Микросервис уведомлений на порту `8081`.
    3.  `db`: База данных PostgreSQL 15 на порту `5432` с нужным именем БД (`food_delivery_v2`) и паролем (`1234`).

---

## 🚀 Как запустить и проверить обе задачи

Благодаря Docker-Compose вам больше не нужно запускать сервисы в разных терминалах вручную или поднимать локальную БД. Все разворачивается одной командой!

### Шаг 1: Запуск всего проекта

Откройте терминал в корне проекта и выполните команду:
```bash
docker-compose up --build
```
*Флаг `--build` заставит Docker с нуля скачать модули и собрать оба наших Go приложения.*

Docker автоматически поднимет БД, основной сервер и сервер уведомлений, связав их в одну сеть.

### Шаг 2: Тестирование работы микросервисов и Resty

Откройте второй терминал и сделайте POST запрос для создания заказа (это вызовет срабатывание Resty клиента в основном приложении):

```bash
curl -X POST http://localhost:8080/orders \
-H "Content-Type: application/json" \
-d '{"restaurant_id": 1, "customer_name": "Arman", "total_price": 1500}'
```

*(Если у вас нет curl, вы можете использовать Postman, отправив POST запрос на `http://localhost:8080/orders` с телом в формате JSON).*

### Шаг 3: Проверка логов

Вернитесь в терминал, где запущен `docker-compose up`. Вы должны увидеть логи от обоих приложений, доказывающие, что межсервисное взаимодействие работает:

1.  **Логи Resty от `app`:**
    ```text
    app_1           | [Resty] Requesting: POST http://notification:8081/notify
    app_1           | [Resty] Response Code: 200
    ```
2.  **Логи получения от `notification`:**
    ```text
    notification_1  | Incoming POST /notify
    notification_1  | [Notification Service] Received notification for Order ...
    ```

### Шаг 4: Остановка

Когда закончите тестирование, нажмите `Ctrl+C` в терминале с `docker-compose up` или выполните:
```bash
docker-compose down
```
Это остановит и удалит запущенные контейнеры.
