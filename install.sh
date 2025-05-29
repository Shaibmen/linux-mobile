#!/bin/bash

set -e

APP_NAME="linux-mobile"
INSTALL_DIR="/opt/$APP_NAME"
SERVICE_FILE="/etc/systemd/system/$APP_NAME.service"
ENV_FILE="$INSTALL_DIR/.env"
PORT=3000

echo "=== Установка $APP_NAME ==="

# Проверка прав root
if [[ $EUID -ne 0 ]]; then
    echo "Этот скрипт нужно запускать через sudo."
    exit 1
fi

# Проверка наличия бинарника
if [[ ! -f "linux-mobile" ]]; then
    echo "Файл linux-mobile не найден в текущей директории."
    exit 1
fi

# Создание директории
mkdir -p "$INSTALL_DIR"
cp linux-mobile "$INSTALL_DIR/"
chmod +x "$INSTALL_DIR/linux-mobile"

# Заполнение .env
echo -n "Введите API ключ: "
read -r API_KEY

echo "API_KEY=$API_KEY" > "$ENV_FILE"
echo "PORT=$PORT" >> "$ENV_FILE"

chmod 600 "$ENV_FILE"
chown root:root "$ENV_FILE"

# Создание systemd юнита
cat > "$SERVICE_FILE" <<EOF
[Unit]
Description=Linux Mobile Server
After=network.target

[Service]
ExecStart=$INSTALL_DIR/linux-mobile
WorkingDirectory=$INSTALL_DIR
EnvironmentFile=$ENV_FILE
Restart=always
RestartSec=3
User=root

[Install]
WantedBy=multi-user.target
EOF

# Активация сервиса
systemctl daemon-reexec
systemctl daemon-reload
systemctl enable "$APP_NAME"
systemctl restart "$APP_NAME"

# Открытие порта 3000
if command -v ufw &>/dev/null; then
    ufw allow $PORT/tcp
    echo "Открыт порт $PORT через ufw"
else
    echo "UFW не установлен. Убедитесь, что порт $PORT открыт вручную."
fi

echo "$APP_NAME установлен и работает на порту $PORT"
