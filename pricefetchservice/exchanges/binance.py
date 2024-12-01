import asyncio
import logging

import ccxt.pro as ccxtpro
from kafka import KafkaProducer

from models import PriceEvent

logger = logging.getLogger(__name__)


async def stream(producer: KafkaProducer):
    exchange_name = "binance"
    exchange = ccxtpro.binance()
    symbol = "BTC/USDT"

    while True:
        try:
            ticker = await exchange.watch_ticker(symbol)
            price_event = PriceEvent(
                exchange=exchange_name,
                symbol=ticker["symbol"],
                price=ticker["last"],
                timestamp=ticker["timestamp"],
            )

            producer.send(
                "price-updates", key=exchange_name, value=price_event.__dict__
            )
            logger.info(f"Price Event sent: {price_event}")
        except Exception as e:
            logger.error(f"Error in {exchange_name} stream: {e}")
            await asyncio.sleep(5)
