# 長野君成長プロジェクト

## Project Rule
- どんなパッケージ・ライブラリを入れてもOK
  - ただし、README に何の何を入れたかと、公式ドキュメントの参照先を記載すること
  - 例えば、ロギングライブラリに logrus を使用するなど

## フレームワーク
- echo
- https://echo.labstack.com/

### Go server with an Nginx proxy and a MariaDB/MySQL database

Project structure:
```
.
├── backend
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── db
│   └── password.txt
├── proxy
│   └── nginx.conf
├── compose.yaml
└── README.md
```
