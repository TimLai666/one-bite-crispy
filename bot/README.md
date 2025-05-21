# 聊天機器人 API

這是一個簡單的聊天機器人 API，使用 FastAPI 和 OpenAI 的 GPT 模型來處理訊息。

## 安裝方式

1. 安裝所需套件：
```bash
pip install -r requirements.txt
```

2. 設定環境變數：
   - 複製 `.env.example` 到 `.env`
   - 在 `.env` 中填入你的 OpenAI API 金鑰

3. 執行應用程式：
```bash
uvicorn app:app --reload
```

## API 使用方式

### 發送聊天訊息

**端點：** `POST /chat`

**請求主體：**
```json
{
    "content": "你的訊息內容",
    "system_prompt": "可選的系統提示詞"
}
```

**回應格式：**
```json
{
    "response": "AI 的回應內容",
    "status": "success"
}
```

## API 文件

啟動服務後，可以在以下網址查看完整的 API 文件：
- http://localhost:8000/docs
- http://localhost:8000/redoc
