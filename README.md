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

2. 環境変数の設定
   ```sh
   cd backend
   cp .env.example .env
   ```

3. コンテナを起動
   ```sh
   cd ..
   // Docker コンテナ起動
   make up
   // データベースのマイグレーション適用
   make migrate
   ```

4. ユーザー登録ページにアクセスする
   ```
   http://localhost/signup
   ```
