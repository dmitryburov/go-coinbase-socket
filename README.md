# Coinbase websocket client on Golang

![IMG](docs/hero.gif)

## Task

1. Connect to the Coinbase cryptocurrency exchange via WebSocket
2. Subscribe to prices for three instruments: ETH-BTC, BTC-USD, BTC-EUR (data will come in through the WebSocket as they appear on the exchange)
3. Create a "ticks" table in MySql with the following fields:
   - `timestamp` (int64) - price time
   - `symbol` (string) - instrument name
   - `bid` (float64) - selling price
   - `ask` (float64) - buying price
4. Connect to MySql and write data received through WebSocket to the `ticks` table
5. Write data from the three instruments (ETH-BTC, BTC-USD, BTC-EUR) to the database
   in three threads (each instrument has its own write thread)
6. Publish the project repository on Github
7. Use good development practices to enable further application functionality expansion

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

- [x] Logger points
- [ ] Rate limiter
- [ ] Metric data
- [ ] Test implemented


## Resources

- [Coinbase websocket overview](https://docs.cloud.coinbase.com/exchange/docs/websocket-overview)
- [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

