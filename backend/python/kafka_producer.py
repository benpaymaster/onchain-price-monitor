from typing import Any, Dict

class KafkaProducer:
	"""
	KafkaProducer sends messages to Kafka topics.
	"""
	def __init__(self, producer: Any = None) -> None:
		"""
		Initialize KafkaProducer with an optional producer instance.
		Args:
			producer (Any): Kafka producer client (can be mocked for testing).
		"""
		self.producer = producer

	def send_price(self, topic: str, value: Dict[str, Any]) -> bool:
		"""
		Send a price message to a Kafka topic.
		Args:
			topic (str): Kafka topic name.
			value (Dict[str, Any]): Message payload.
		Returns:
			bool: True if sent successfully.
		"""
		if self.producer:
			self.producer.send(topic, value)
			return True
		return False
