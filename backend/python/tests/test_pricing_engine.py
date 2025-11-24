import pytest
from pricing_engine import PricingEngine

def test_basic_price_calculation():
    engine = PricingEngine()
    price = engine.calculate_price(100, 0.01)
    assert price == 101

def test_negative_price():
    engine = PricingEngine()
    price = engine.calculate_price(-100, 0.01)
    assert price == -99
