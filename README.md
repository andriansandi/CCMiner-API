[![Buy Me A Coffee](https://img.shields.io/badge/Buy%20Me%20A%20Coffee-donate-yellow)](https://www.buymeacoffee.com/andriansandi)

# Hashrate Server for ccminer

A lightweight web server built with Golang to expose `ccminer` hashrate data via a JSON API.

## Features
- Reads hashrate from `/tmp/ccminer.log`
- Serves hashrate data in JSON format via HTTP
- Lightweight and efficient

## Installation

### 1. Clone this repository
```sh
git clone https://github.com/yourusername/hashrate_server.git
cd hashrate_server
```

### 2. Build the Project
```sh
go build -o hashrate_server
```

### 3. Run the Server
```sh
./hashrate_server
```

### 4. Check the API
```sh
curl http://127.0.0.1:5000/hashrate
```

## Running as a Service (Systemd)
### 1. Clone this repository
```sh
sudo vim /etc/systemd/system/hashrate_server.service
```

### 2. Add the following content
```sh
[Unit]
Description=Hashrate Server for ccminer
After=network.target

[Service]
ExecStart=/usr/local/bin/hashrate_server
Restart=always
User=root
WorkingDirectory=/usr/local/bin

[Install]
WantedBy=multi-user.target
```

### 3. Enable and start the service
```sh
sudo systemctl daemon-reload
sudo systemctl enable hashrate_server
sudo systemctl start hashrate_server
```

### 4. Check the status
```sh
sudo systemctl status hashrate_server
```

## API Response Example
```json
{
  "hashrate": 2.5,
  "unit": "MH/s"
}
```

## Maintainer
Sandi Andrian [andrian.sandi@gmail.com](mailto:andrian.sandi@gmail.com)

<a href="https://www.buymeacoffee.com/andriansandi" target="_blank">
    <img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" height="30">
</a>