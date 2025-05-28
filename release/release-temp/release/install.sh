#!/bin/bash

set -e

echo "=== Westeros Installer ==="

# Проверка прав
if [[ $EUID -ne 0 ]]; then
   echo "Запустите скрипт от root: sudo ./install.sh"
   exit 1
fi

# Запрос переменных окружения
read -p "Введите API_KEY: " API_KEY

# Установка сервиса в /opt
INSTALL_DIR="/opt/westeros"
mkdir -p "$INSTALL_DIR"
cp linux-mobile "$INSTALL_DIR"
chmod 700 "$INSTALL_DIR/linux-mobile"

# Создание .env
echo "API_KEY=$API_KEY" > "$INSTALL_DIR/.env"
chmod 600 "$INSTALL_DIR/.env"

# systemd unit
cat <<EOF > /etc/systemd/system/westeros.service
[Unit]
Description=Westeros Server
After=network.target

[Service]
Type=simple
WorkingDirectory=$INSTALL_DIR
ExecStart=$INSTALL_DIR/linux-mobile
Restart=always
EnvironmentFile=$INSTALL_DIR/.env

[Install]
WantedBy=multi-user.target
EOF

# Разрешить порт 3000
if command -v ufw > /dev/null; then
  ufw allow 3000/tcp
elif command -v firewall-cmd > /dev/null; then
  firewall-cmd --add-port=3000/tcp --permanent
  firewall-cmd --reload
else
  iptables -A INPUT -p tcp --dport 3000 -j ACCEPT
fi

# Запуск сервиса
systemctl daemon-reexec
systemctl daemon-reload
systemctl enable westeros
systemctl start westeros

echo "Установка завершена. Приложение запущено и добавлено в автозагрузку."
