import logging
import os

import orjson
from dotenv import load_dotenv
from kafka import KafkaProducer
from kafka.admin import KafkaAdminClient, NewTopic

load_dotenv()

logger = logging.getLogger(__name__)


KAFKA_BROKER = os.getenv("KAFKA_BROKER")
KAFKA_TOPIC = os.getenv("KAFKA_TOPIC")
KAFKA_PARTITIONS = int(os.getenv("KAFKA_PARTITIONS", 1))
KAFKA_REPLICATION_FACTOR = int(os.getenv("KAFKA_REPLICATION_FACTOR", 1))


def create_kafka_topic():
    admin_client = KafkaAdminClient(
        bootstrap_servers=KAFKA_BROKER, client_id="bitcoin-pulse-admin"
    )
    topic_list = []
    topic_list.append(
        NewTopic(
            name=KAFKA_TOPIC,
            num_partitions=KAFKA_PARTITIONS,
            replication_factor=KAFKA_REPLICATION_FACTOR,
        )
    )
    try:
        admin_client.create_topics(new_topics=topic_list, validate_only=False)
        logger.info(f"Kafka topic '{KAFKA_TOPIC}' created successfully.")
    except Exception as e:
        if "TopicAlreadyExistsError" in str(e):
            logger.warning(f"Kafka topic '{KAFKA_TOPIC}' already exists.")
        else:
            logger.error(f"Failed to create Kafka topic '{KAFKA_TOPIC}': {e}")
    finally:
        admin_client.close()


def get_kafka_producer() -> KafkaProducer:
    producer = KafkaProducer(
        bootstrap_servers=KAFKA_BROKER,
        key_serializer=lambda k: k.encode("utf-8"),
        value_serializer=lambda v: orjson.dumps(v),
        linger_ms=10,
        batch_size=32 * 1024,  # 32 KB
    )
    return producer
