import asyncio
import json
import logging

import ccxt.pro as ccxtpro
from kafka import KafkaProducer

from models.models import PriceEvent

logger = logging.getLogger(__name__)


class BaseExchange:
    def __init__(
        self, exchange_name: str, symbol: str, producer: KafkaProducer
    ) -> None:
        self.exchange_name = exchange_name
        self.symbol = symbol
        self.producer = producer
        self.exchange = getattr(ccxtpro, exchange_name)()

    async def stream_ticker(self) -> None:
        if not self.exchange.has["watchTicker"]:
            logger.warning(f"{self.exchange_name} does not support watchTicker")
            return
        try:
            while True:
                await self._fetch_and_send_ticker()
        finally:
            logger.info(f"Closing {self.exchange_name} feed")
            await self.exchange.close()

    async def _fetch_and_send_ticker(self) -> None:
        try:
            ticker = await self.exchange.watch_ticker(self.symbol)
            price_event = self.process_ticker(ticker)
            await self.send_to_kafka(price_event)
        except Exception as e:
            logger.error(f"Error in {self.exchange_name} stream: {e}")
            await asyncio.sleep(5)  # Retry after delay

    def process_ticker(self, ticker) -> PriceEvent:
        price_event = PriceEvent(
            exchange=self.exchange_name,
            symbol=ticker["symbol"],
            last=ticker["last"],
            timestamp=ticker["timestamp"],
            change=ticker["change"],
            percentage=ticker["percentage"],
            bid=ticker["bid"],
            ask=ticker["ask"],
            open=ticker["open"],
            high=ticker["high"],
            low=ticker["low"],
            quoteVolume=ticker["quoteVolume"],
            baseVolume=ticker["baseVolume"],
        )
        return price_event

    async def send_to_kafka(self, price_event: PriceEvent) -> None:
        self.producer.send(
            "price-updates", key=self.exchange_name, value=price_event.__dict__
        )
        price_event_json = json.dumps(price_event.__dict__, indent=4)
        logger.info(
            f"\n🚀🚀🚀 {self.symbol} Price Event sent 🚀🚀🚀:\n{price_event_json}"
        )
