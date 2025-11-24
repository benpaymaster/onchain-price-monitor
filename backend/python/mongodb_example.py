from pymongo import MongoClient
from typing import Any, Dict, List

class MongoPriceStore:
    """
    MongoPriceStore manages price data in MongoDB.
    """
    def __init__(self, uri: str = 'mongodb://mongo:27017/', db_name: str = 'price_db') -> None:
        self.client = MongoClient(uri)
        self.db = self.client[db_name]
        self.collection = self.db['prices']

    def insert_price(self, price: float, timestamp: str) -> None:
        self.collection.insert_one({'price': price, 'timestamp': timestamp})

    def get_latest_price(self) -> Dict[str, Any]:
        doc = self.collection.find_one(sort=[('_id', -1)])
        if doc:
            return {'price': doc['price'], 'timestamp': doc['timestamp']}
        return {}

    def get_all_prices(self) -> List[Dict[str, Any]]:
        return list(self.collection.find({}, {'_id': 0}))

# Example usage (commented out for production)
# store = MongoPriceStore()
# store.insert_price(123.45, '2025-11-24T00:00:00Z')
# print(store.get_latest_price())
# print(store.get_all_prices())
