# TripPlanner
## アプリケーション概要
TripPlanner は、ユーザーが旅行プランを作成・管理できる Web アプリケーションです。  
バックエンドは Go（Echo フレームワーク）、フロントエンドは Vue を使用して設計しています。

## 環境構築方法
### 必要なツール
- Docker

### Docker を使用したセットアップ
1. リポジトリをクローン
   ```sh
   git clone https://github.com/d-nagano/TripPlanner.git
   cd TripPlanner
   ```

2. コンテナを起動
   ```sh
   // Docker コンテナ起動
   mkae up
   // DB 作成
   make migrate
   ```

3. URLにアクセス
   ```
   http://localhost/signup
   ```
