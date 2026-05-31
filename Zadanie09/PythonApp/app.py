import string
from fastapi import FastAPI
from pydantic import BaseModel
import requests
import uvicorn
import random

app = FastAPI()


class UserInput(BaseModel):
    text: str


OPENINGS = [
    "Witaj! W czym mogę ci dzisiaj pomóc?",
    "Jaki piękny wieczór! O czym chcesz pogadać?",
    "Cześć, tutaj twoje AI, jestem do Twojej dyspozycji!",
    "Jak się dzisiaj czujesz?",
    "Kopę lat! Porozmawiajmy."
]

CLOSINGS = [
    "Miło się rozmawiało! Do zobaczenia jutro :)",
    "Do usłyszenia, idę odpocząć ;)",
    "Papa, miłego dnia :)",
    "Dzięki za rozmowę, jak coś to czekam na dalsze pytania!",
    "Uff, właśnie miałem iść odpocząć, żegnaj!"
]

CLOSING_WORDS = ["do widzenia", "papa", "żegnaj", "dobranoc", "kończymy", "żegnaj"]


@app.post("/askai")
def get_ai_response(user_input: UserInput):
    input = user_input.text.lower().strip()

    if input == "[start]":
        return {"response": random.choice(OPENINGS)}

    clear_text = input.translate(str.maketrans('', '', string.punctuation)).strip()

    if clear_text in CLOSING_WORDS:
        return {"response": random.choice(CLOSINGS)}

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
