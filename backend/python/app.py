
from flask import Flask, jsonify, request
from pricing_engine import PricingEngine
from monitor_stream import MonitorStream
import logging
from prometheus_flask_exporter import PrometheusMetrics
from flask_jwt_extended import JWTManager, create_access_token, jwt_required

app = Flask(__name__)
app.config['JWT_SECRET_KEY'] = 'super-secret-key'  # Change for production!
metrics = PrometheusMetrics(app)
jwt = JWTManager(app)

# Structured logging setup
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s %(levelname)s %(name)s %(message)s'
)
logger = logging.getLogger("backend")

pricing_engine = PricingEngine()
monitor_stream = MonitorStream()

@app.route('/login', methods=['POST'])
def login():
    username = request.json.get('username', None)
    password = request.json.get('password', None)
    # Simple demo: accept any username/password
    if not username or not password:
        return jsonify({'msg': 'Missing username or password'}), 400
    access_token = create_access_token(identity=username)
    return jsonify(access_token=access_token)

@app.route('/api/price', methods=['GET'])
@jwt_required()
def get_price():
    base_price = float(request.args.get('base_price', 100))
    slippage = float(request.args.get('slippage', 0.01))
    price = pricing_engine.calculate_price(base_price, slippage)
    logger.info(f"Calculated price: base={base_price}, slippage={slippage}, result={price}")
    return jsonify({'base_price': base_price, 'slippage': slippage, 'price': price})

@app.route('/api/monitor', methods=['GET'])
@jwt_required()
def get_monitor_data():
    data = monitor_stream.get_data()
    logger.info(f"Monitor data: {data}")
    return jsonify(data)

@app.route('/health', methods=['GET'])
def health():
    logger.info("Health check requested")
    return jsonify({'status': 'ok'})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
