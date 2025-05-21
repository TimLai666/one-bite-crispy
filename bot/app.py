from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import openai
from typing import Optional
import os
from dotenv import load_dotenv

# 載入環境變數
load_dotenv()

# 設定OpenAI API金鑰
openai.api_key = os.getenv("OPENAI_API_KEY")

app = FastAPI(title="聊天機器人 API")

class Message(BaseModel):
    content: str
    system_prompt: Optional[str] = "你是一個有幫助的助手。"

class ChatResponse(BaseModel):
    response: str
    status: str

@app.post("/chat", response_model=ChatResponse)
async def chat_with_llm(message: Message):
    try:
        # 準備與LLM的對話
        messages = [
            {"role": "system", "content": message.system_prompt},
            {"role": "user", "content": message.content}
        ]
        
        # 呼叫OpenAI API
        response = openai.chat.completions.create(
            model="gpt-3.5-turbo",
            messages=messages,
            temperature=0.7,
            max_tokens=1000
        )
        
        # 提取回應文字
        assistant_response = response.choices[0].message.content
        
        return ChatResponse(
            response=assistant_response,
            status="success"
        )
    
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
