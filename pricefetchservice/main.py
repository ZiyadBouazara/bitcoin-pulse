import asyncio
import logging

from dotenv import load_dotenv

from exchanges.exchange import BaseExchange
from exchanges.exchange_config import exchange_config
from infrastructure.kafka_config import create_kafka_topic, get_kafka_producer

load_dotenv()

logging.basicConfig(
    level=logging.ERROR,
    format="[%(asctime)s] %(levelname)s in %(module)s: %(message)s",
    handlers=[logging.StreamHandler()],
)

logger = logging.getLogger(__name__)


async def main():
    create_kafka_topic()
    logger.info("Kafka topic setup complete.")

    producer = get_kafka_producer()
    tasks = []

    for exchange_name, params in exchange_config.items():
        symbol = params["symbol"]
        exchange_stream = BaseExchange(exchange_name, symbol, producer)
        tasks.append(exchange_stream.stream_ticker())
        logger.info(f"Started streaming for exchange: {exchange_name}")

    try:
        await asyncio.gather(*tasks)
    except KeyboardInterrupt:
        logger.info("Shutting down gracefully...")
    finally:
        producer.close()
        logger.info("Kafka producer closed.")


if __name__ == "__main__":
    asyncio.run(main())
