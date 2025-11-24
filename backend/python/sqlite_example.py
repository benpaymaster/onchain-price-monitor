import sqlite3
from typing import Any, Dict, List

class PriceDatabase:
    """
    PriceDatabase manages price data storage in SQLite.
    """
    def __init__(self, db_path: str = 'prices.db') -> None:
        self.conn = sqlite3.connect(db_path)
        self.create_table()

    def create_table(self) -> None:
        with self.conn:
            self.conn.execute('''
                CREATE TABLE IF NOT EXISTS prices (
                    id INTEGER PRIMARY KEY AUTOINCREMENT,
                    price REAL NOT NULL,
                    timestamp TEXT NOT NULL
                )
            ''')

    def insert_price(self, price: float, timestamp: str) -> None:
        with self.conn:
            self.conn.execute(
                'INSERT INTO prices (price, timestamp) VALUES (?, ?)',
                (price, timestamp)
            )

    def get_latest_price(self) -> Dict[str, Any]:
        cur = self.conn.cursor()
        cur.execute('SELECT price, timestamp FROM prices ORDER BY id DESC LIMIT 1')
        row = cur.fetchone()
        if row:
            return {'price': row[0], 'timestamp': row[1]}
        return {}

    def get_all_prices(self) -> List[Dict[str, Any]]:
        cur = self.conn.cursor()
        cur.execute('SELECT price, timestamp FROM prices ORDER BY id DESC')
        rows = cur.fetchall()
        return [{'price': r[0], 'timestamp': r[1]} for r in rows]

# Example usage (commented out for production)
# db = PriceDatabase()
# db.insert_price(123.45, '2025-11-24T00:00:00Z')
# print(db.get_latest_price())
# print(db.get_all_prices())
