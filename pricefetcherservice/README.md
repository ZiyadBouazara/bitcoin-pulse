# Price Fetcher Service
## Flow Description:

* domain/service/PriceFetchingService calls infrastructure/CoinbaseApiClient to retrieve the latest price of a given stock.

* CoinbaseApiClient makes an HTTP request to Coinbase API and returns the result as a Price object.

* domain/service/PriceFetchingService creates a PriceEvent and calls the relevant PriceProducer.

* PriceProducer (e.g., BitcoinPriceProducer) sends the PriceEvent to the relevant Kafka Topic (in this case bitcoin-price-topic).