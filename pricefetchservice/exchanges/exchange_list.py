from enum import Enum


class Exchange(Enum):
    # BINANCE = "binance"
    # COINBASE = "coinbase"
    KRAKEN = "kraken"

    @staticmethod
    def list():
        return [exchange.value for exchange in Exchange]
