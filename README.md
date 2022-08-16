# wb-L0
Wildberries internship task


##Тестовое задание:
###В БД:
1. Развернуть локально postgresql
2. Создать свою бд
3. Настроить своего пользователя.
4. Создать таблицы для хранения полученных данных.
### В сервисе:
5. Подключение и подписка на канал в nats-streaming
6. Полученные данные писать в Postgres
7. Так же полученные данные сохранить in memory в сервисе (Кеш)
8. В случае падения сервиса восстанавливать Кеш из Postgres
9. Поднять http сервер и выдавать данные по id из кеша
10. Сделать просте ши интерфе с отображения полученных данных, для
    их запроса по id
    Доп инфо:
    • Данные статичны, исходя из этого подума те насчет модели хранения
    в Кеше и в pg. Модель в файле model.json
    • В канал могут закинуть что угодно, подума те как избежать проблем
    из-за этого
    • Чтобы проверить работает ли подписка онла н, сдела те себе
    отдельны скрипт, для публикации данных в канал
    • Подума те как не терять данные в случае ошибок или проблем с
    сервисом
    • Nats-streaming разверните локально ( не путать с Nats )
### Бонус задание
11. Покро те сервис автотестами. Будет плюсик вам в карму😊
12. Устро те вашему сервису стресс тест, выясните на что он способен -
    воспользуйтесь утилитами WRK, Vegeta. Попробуйте оптимизировать код

### Модель данных

```json
{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
  "transaction": "b563feb7b2b84b6test",
  "request_id": "",
  "currency": "USD",
  "provider": "wbpay",
  "amount": 1817,
  "payment_dt": 1637907727,
  "bank": "alpha",
  "delivery_cost": 1500,
  "goods_total": 317,
  "custom_fee": 0
  },
  "items": [{
    "chrt_id": 9934930,
    "track_number": "WBILMTESTTRACK",
    "price": 453,
    "rid": "ab4219087a764ae0btest",
    "name": "Mascaras",
    "sale": 30,
    "size": "0",
    "total_price": 317,
    "nm_id": 2389212,
    "brand": "Vivienne Sabo",
    "status": 202
  }],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
```