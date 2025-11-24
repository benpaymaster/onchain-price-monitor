import redis
import json
from typing import Any, Dict

class RedisCache:
    """
    RedisCache provides methods to store and retrieve price data in Redis.
    """
    def __init__(self, host: str = 'localhost', port: int = 6379, db: int = 0) -> None:
        self.client = redis.Redis(host=host, port=port, db=db)

    def set_price(self, key: str, price_data: Dict[str, Any]) -> None:
        self.client.set(key, json.dumps(price_data))

    def get_price(self, key: str) -> Dict[str, Any]:
        data = self.client.get(key)
        if data:
            return json.loads(data)
        return {}

    def health_check(self) -> bool:
        try:
            return self.client.ping()
        except redis.ConnectionError:
            return False

# Example usage (commented out for production)
cache = RedisCache()
cache.set_price('latest_price', {'price': 123, 'timestamp': '2025-11-24T00:00:00Z'})
print(cache.get_price('latest_price'))
print('Redis health:', cache.health_check())
