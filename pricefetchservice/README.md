# Price Fetch Service

This service streams real-time cryptocurrency price data from over 100 exchanges (e.g., Binance, Coinbase) using WebSockets and sends the data to a Kafka topic for processing by other microservices.

## Features

-   Streams live price data from multiple exchanges.
-   Publishes data to a single Kafka topic (`bitcoin-price-topic`).
-   Partitioning ensures that messages for each exchange are logically grouped.

## Architecture

This service focuses solely on **producers**:

-   Fetches price data from WebSocket APIs of various exchanges using CCXT.
-   Publishes the data to a Kafka topic using the exchange name as the partition key.

**Note**: Consumers for this data reside in other microservices, such as `stockservice-go` or `alert-service`.

---

## Getting Started

### Prerequisites

1. **Kafka**: Ensure Kafka is installed and running locally or on a server.
2. **Python 3.9+**: Install Python and Poetry for dependency management.
3. **Dependencies**: Install required Python packages.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/ZiyadBouazara/bitcoin-pulse.git
    cd bitcoin-pulse
    ```

2. Install dependencies using Poetry:

    ```bash
    poetry install
    ```

3. Create a `.env` file:

    ```bash
    cp .env.example .env
    ```

4. Update `.env` with your environment-specific configurations.

---

### Usage

#### Starting the Kafka Producer Service

1. Ensure Kafka is running with the `bitcoin-price-topic` topic created:

    ```bash
    bin/kafka-topics.sh --create --topic price-updates \
        --bootstrap-server localhost:9092 \
        --partitions 100 \
        --replication-factor 1
    ```

2. Run the producer service:

    ```bash
    poetry run python producers/producer_service.py
    ```

3. The service will:
    - Connect to WebSocket APIs.
    - Stream real-time data to Kafka.

---

### Configuration

Environment variables are defined in the `.env` file for easy configuration.

| Variable          | Description                 | Example                                          |
| ----------------- | --------------------------- | ------------------------------------------------ |
| `KAFKA_BROKER`    | Kafka broker address        | `localhost:9092`                                 |
| `KAFKA_TOPIC`     | Kafka topic name            | `bitcoin-price-topic`                            |
| `BINANCE_WS_URL`  | Binance WebSocket endpoint  | `wss://stream.binance.com:9443/ws/btcusdt@trade` |
| `COINBASE_WS_URL` | Coinbase WebSocket endpoint | `wss://ws-feed.exchange.coinbase.com`            |

---

## Extending the Service

1. **Add a New Exchange**:

    - Create a new file in `exchanges/` (e.g., `kraken.py`).
    - Implement a WebSocket client for the exchange using CCXT.
    - Register the new exchange in `producer_service.py`.

2. **Scaling**:
    - Adjust Kafka partitions to handle additional exchanges.

---

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for feedback.

---

## License

This project is licensed under the MIT License. See `LICENSE` for details.

---

## **4. Running the Service**

With everything set up:

1. **Start Kafka**:

    - Ensure Zookeeper and Kafka are running.
    - a docker-compose.yaml is provided at the root of bitcoin-pulse for this.

2. **Run Producer**:

    ```bash
    poetry run python producers/producer_service.py
    ```

3. **Monitor Kafka**:
    - Use Kafka CLI tools or third-party tools like Kafka Manager to verify that messages are flowing into the `price-updates` topic.

---

This setup provides a reusable and developer-friendly structure for your service. Let me know if you’d like further assistance with any specific part!
