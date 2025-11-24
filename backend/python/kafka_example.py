from kafka import KafkaProducer, KafkaConsumer
import json

# Kafka Producer Example
class MyKafkaProducer:
    def __init__(self, brokers: str):
        self.producer = KafkaProducer(
            bootstrap_servers=brokers,
            value_serializer=lambda v: json.dumps(v).encode('utf-8')
        )

    def send_price(self, topic: str, price_data: dict) -> None:
        self.producer.send(topic, price_data)
        self.producer.flush()

# Kafka Consumer Example
class MyKafkaConsumer:
    def __init__(self, brokers: str, topic: str):
        self.consumer = KafkaConsumer(
            topic,
            bootstrap_servers=brokers,
            value_deserializer=lambda m: json.loads(m.decode('utf-8')),
            auto_offset_reset='earliest',
            enable_auto_commit=True
        )

    def consume(self):
        for message in self.consumer:
            print(f"Received: {message.value}")

# Example usage (commented out for production)
# producer = MyKafkaProducer('localhost:9092')
# producer.send_price('prices', {'price': 123, 'timestamp': '2025-11-24T00:00:00Z'})

# consumer = MyKafkaConsumer('localhost:9092', 'prices')
# consumer.consume()
