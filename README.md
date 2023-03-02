# Coinbase websocket client on Golang

## Task

1. Подключиться по WebSocket, к бирже криптовалют CoinBase;
2. Подписаться на цены по трем инструментам: ETH-BTC, BTC-USD, BTC-EUR (данные будут приходить в WebSocket по мере их появления на бирже);
3. Создать таблицу "ticks" в MySql со следующими полями:
   - timestamp (int64) - время цены
   - symbol (string) - название инструментам
   - bid (float64) - цена продажи
   - ask (float64) - цена покупки;
4. Подключиться к MySql и записывать в таблицу "ticks" данные полученные по WebSocket;
5. Записывать данные из трех инструментов (ETH-BTC, BTC-USD, BTC-EUR) в базу данных
   необходимо в 3 потока (каждому инструменту свой поток записи);
6. Репозиторий с проектом выложить на github;
7. Используйте хорошие практики разработки, для возможности дальнейшего расширения
   функционала приложения.

## How to use

Project based on Clean architecture principles.

Requirements:

- Go 1.19+ installed
- Docker installed (to run docker-compose)

```
# Run docker-compose
make compose.start

# Stop docker-compose
make compose.stop

# other command - call help
make help
```

## TODO

- [ ] Logger points
- [ ] Metric data
- [ ] Test implemented


## Resources

- [Coinbase websocket overview](https://docs.cloud.coinbase.com/exchange/docs/websocket-overview)
- [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

