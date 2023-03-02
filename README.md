# Coinbase websocket client on Golang

This is test task.

## Задача

1. Подключиться по WebSocket, к бирже криптовалют CoinBase (documentation: https://docs.cloud.coinbase.com/exchange/docs/websocket-overview)
2. Подписаться на цены по трем инструментам: ETHBTC, BTCUSD, BTCEUR (данные будут приходить в WebSocket по мере их появления на бирже)
3. Создать таблицу "ticks" в MySql со следующими полями:
   timestamp (int64) - время получения цены (брать текущее время Unixtime в милисекундах)
   symbol (string) - название инструментам
   bid (float64) - цена продажи
   ask (float64) - цена покупки
4. Подключиться к MySql и записывать в таблицу "ticks" данные полученные по WebSocket
5. Записывать данные из трех инструментов (ETHBTC, BTCUSD, BTCEUR) в базу данных
   необходимо в 3 потока (каждому инструменту свой поток записи)
6. Репозиторий с проектом выложить на github
7. Используйте хорошие практики разработки, для возможности дальнейшего расширения
   функционала приложения
