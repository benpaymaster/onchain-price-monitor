import pytest
from monitor_stream import MonitorStream

def test_stream_initialization():
    stream = MonitorStream('test_stream')
    assert stream.name == 'test_stream'

def test_stream_run():
    stream = MonitorStream('test_stream')
    result = stream.run()
    assert result is not None
