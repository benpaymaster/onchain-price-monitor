class PricingEngine:
	"""
	PricingEngine calculates prices with slippage adjustments.
	"""
	def calculate_price(self, base_price: float, slippage: float) -> float:
		"""
		Calculate the adjusted price given a base price and slippage.
		Args:
			base_price (float): The base price of the asset.
			slippage (float): The slippage as a decimal (e.g., 0.01 for 1%).
		Returns:
			float: The adjusted price.
		"""
		return base_price * (1 + slippage)
