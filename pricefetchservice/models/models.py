from dataclasses import dataclass


@dataclass
class PriceEvent:
    exchange: str
    symbol: str
    price: float
    timestamp: int

    def to_json(self) -> bytes:
        import orjson

        return orjson.dumps(self.__dict__)
