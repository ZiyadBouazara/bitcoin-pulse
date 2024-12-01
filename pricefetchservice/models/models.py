from dataclasses import dataclass

from base.types import Int, Num, Str


@dataclass
class PriceEvent:
    exchange: Str
    symbol: Str
    last: Num
    timestamp: Int
    change: Num
    percentage: Num
    bid: Num
    ask: Num
    open: Num
    high: Num
    low: Num
    quoteVolume: Num
    baseVolume: Num

    def to_json(self) -> bytes:
        import orjson

        return orjson.dumps(self.__dict__)
