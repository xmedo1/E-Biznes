from fastapi import FastAPI
from pydantic import BaseModel
import requests
import uvicorn

app = FastAPI()


class UserInput(BaseModel):
    text: str


@app.post("/askai")
def get_ai_response(user_input: UserInput):
    query = {
        "model": "llama3",
        "prompt": user_input.text,
        "stream": False
    }

    try:
        answer = requests.post("http://localhost:11434/api/generate", json=query)
        return {"response": answer.json()["response"]}

    except Exception:
        return {"response": "Blad polaczenia z AI."}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=5000)
