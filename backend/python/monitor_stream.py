from typing import Any, Dict

class MonitorStream:
	"""
	MonitorStream simulates a monitoring data stream.
	"""
	def __init__(self, name: str = "default_stream") -> None:
		"""
		Initialize the stream with a name.
		Args:
			name (str): Name of the stream.
		"""
		self.name = name
		self.is_active = True

	def get_data(self) -> Dict[str, Any]:
		"""
		Simulate getting data from the stream.
		Returns:
			Dict[str, Any]: Example data payload.
		"""
		return {"price": 100, "timestamp": "2025-11-24T00:00:00Z"}

	def run(self) -> Any:
		"""
		Simulate running the stream.
		Returns:
			Any: Example result.
		"""
		return self.get_data()
