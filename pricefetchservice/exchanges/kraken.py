import asyncio
import json
import logging

import ccxt.pro as ccxtpro
from kafka import KafkaProducer

from exchanges.exchange_list import Exchange
from models.models import PriceEvent

logger = logging.getLogger(__name__)

EXCHANGE_NAME = Exchange.KRAKEN.value
SYMBOL = "BTC/USD"  # BTC/USDT channel is laggy for Kraken ... really weird ... but BTC/USD is still jittery


async def stream(producer: KafkaProducer):
    exchange = ccxtpro.kraken()
    try:
        while True:
            try:
                ticker = await exchange.watch_ticker(SYMBOL)
                price_event = PriceEvent(
                    exchange=EXCHANGE_NAME,
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

                producer.send(
                    "price-updates", key=EXCHANGE_NAME, value=price_event.__dict__
                )
                price_event_json = json.dumps(price_event.__dict__, indent=4)
                logger.info(
                    f"\n🚀🚀🚀 {SYMBOL} Price Event sent 🚀🚀🚀:\n{price_event_json}"
                )
            except Exception as e:
                logger.error(f"Error in {EXCHANGE_NAME} stream: {e}")
                await asyncio.sleep(5)
    except KeyboardInterrupt:
        logger.info(f"Closing {EXCHANGE_NAME} feed")
    finally:
        await exchange.close()
