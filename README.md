# 快樂登山家 | Golang 後端重構專案

## 專案簡介 (Project Overview)

本專案是 **「快樂登山家 | Hiking Weather Guide」** 的後端服務重構版本，旨在將原有的 Python/FastAPI 技術棧遷移至 Golang，以追求極致的效能、更低的資源佔用和更快的部署速度，特別是為了解決在 Cloud Run 等 Serverless 環境下的冷啟動延遲問題。

這個重構專案不僅是一次技術升級，也是一個為了展示現代後端開發實踐（包含 API 版本控制、分層測試、清晰的專案結構）而設計的專案。

- **原始 Python 專案 Repo**: [hiking-app-python](https://github.com/WuSKai403/hiking-app-python)

---

## API 版本策略

所有由 Golang 實現的新 API Endpoints 將統一使用 `/api/v2` 前綴，以區分於原有的 Python API (`/api/v1`)，確保服務的平滑遷移與向下相容性。

---

## 核心技術棧 (Tech Stack)

| 分類 | 推薦 Golang 套件 | 對應 Python 工具 | 關鍵優勢 |
| :--- | :--- | :--- | :--- |
| **Web 框架** | `Gin` | `FastAPI` | 高效能、輕量級、豐富的中介軟體支援。 |
| **資料庫驅動** | `mongo-go-driver` | `pymongo` | MongoDB 官方驅動，提供完整的 BSON 支援。 |
| **設定管理** | `Viper` | `Pydantic Settings` | 支援多種格式，可熱重載，為業界標準。 |
| **資料驗證** | `go-playground/validator` | `Pydantic` | 透過 Struct Tag 實現靈活的資料驗證。 |
| **日誌** | `uber-go/zap` | `logging` | 高效能的結構化日誌，利於雲端環境整合。 |
| **測試與 Mock** | `testify`, `httptest` | `pytest`, `unittest.mock` | 提供語義化斷言、HTTP 測試伺服器與 Mock 功能。 |

---

## 效能提升與優勢分析 (Golang vs. Python)

| 項目 | Python / FastAPI | Golang | 關鍵優勢 (面試重點) |
| :--- | :--- | :--- | :--- |
| **冷啟動 (Cold Start)** | 較慢 (約 5-10 秒) | **極快 (< 1 秒)** | **完美解決 Cloud Run 痛點**，大幅提升使用者初次請求體驗。 |
| **記憶體佔用** | 較高 | **較低** | 降低 Serverless 營運成本，提升資源利用率。 |
| **併發處理** | `async/await` | **原生 Goroutines** | 模型更簡單直觀，能更高效地利用 CPU 資源處理大量併發請求。 |
| **部署與維運** | 需打包完整環境 | **單一二進位檔** | 極度簡化 Dockerfile，縮小容器體積，加速部署，提升安全性。 |

---

## 測試策略

本專案將採用分層測試策略，確保程式碼的穩定性與可靠性：

1.  **單元測試 (Unit Testing)**:
    *   **目標**: 測試獨立的業務邏輯（如服務層、資料轉換等）。
    *   **工具**: 使用 Go 內建的 `testing` 套件，搭配 `stretchr/testify/assert` 進行斷言。

2.  **整合測試 (Integration Testing)**:
    *   **目標**: 測試 API Endpoints 的完整流程，包含資料庫互動。
    *   **工具**: 使用 `net/http/httptest` 模擬 HTTP 請求，並連接至獨立的測試資料庫。

3.  **Mocking 依賴**:
    *   **目標**: 在單元測試中隔離外部依賴（如 Gemini API、資料庫）。
    *   **方法**: 透過定義 `interface` 和使用 `stretchr/testify/mock` 來建立 Mock 物件。

---

## 本地開發 (Local Development)

### 1. 環境需求
- Go 1.21+

### 2. 初始化專案
```bash
# 下載依賴
go mod tidy
```

### 3. 啟動服務
```bash
# 執行主程式
go run cmd/server/main.go
