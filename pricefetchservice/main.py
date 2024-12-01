import asyncio
import importlib
import logging
import os

from dotenv import load_dotenv

from exchanges.exchange_list import Exchange
from infrastructure.kafka_config import create_kafka_topic, get_kafka_producer

load_dotenv()

logging.basicConfig(
    level=logging.INFO,
    format="[%(asctime)s] %(levelname)s in %(module)s: %(message)s",
    handlers=[logging.StreamHandler()],
)

logger = logging.getLogger(__name__)


async def main():
    create_kafka_topic()
    logger.info("Kafka topic setup complete.")

    producer = get_kafka_producer()
    tasks = []
    exchanges = Exchange.list()

    for name in exchanges:
        module_name = f"exchanges.{name}"
        try:
            module = importlib.import_module(module_name)
            tasks.append(module.stream(producer))
            logger.info(f"Started streaming for exchange: {name}")
        except ModuleNotFoundError:
            logger.error(f"Exchange module '{module_name}' not found.")
    try:
        await asyncio.gather(*tasks)
    except KeyboardInterrupt:
        logger.info("Shutting down gracefully...")
    finally:
        producer.close()
        logger.info("Kafka producer closed.")


if __name__ == "__main__":
    asyncio.run(main())
