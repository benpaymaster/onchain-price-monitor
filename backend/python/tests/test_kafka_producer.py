import pytest
from kafka_producer import KafkaProducer

class DummyKafka:
    def __init__(self):
        self.sent = []
    def send(self, topic, value):
        self.sent.append((topic, value))
        return True

def test_kafka_send():
    dummy = DummyKafka()
    producer = KafkaProducer(dummy)
    result = producer.send_price('prices', {'price': 123})
    assert result is True
    assert dummy.sent[0][0] == 'prices'
    assert dummy.sent[0][1]['price'] == 123
