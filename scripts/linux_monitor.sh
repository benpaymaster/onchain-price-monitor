#!/bin/bash
# Linux resource monitoring script

echo "--- System Resource Usage ---"
echo "CPU Usage:" $(top -bn1 | grep "Cpu(s)" | awk '{print $2 + $4}') "%"
echo "Memory Usage:" $(free -m | awk 'NR==2{printf "%s/%sMB (%.2f%%)", $3,$2,$3*100/$2 }')
echo "Disk Usage:" $(df -h / | awk 'NR==2{print $5}')
echo "Uptime:" $(uptime -p)
