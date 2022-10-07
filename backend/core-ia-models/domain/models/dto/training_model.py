from typing import List
from pydantic import BaseModel

class IntentModel(BaseModel):
    patterns: List[str] = []
    responses: List[str] = []
    tag: str = ""

class TrainingModel(BaseModel):
    _id: str = ""
    userId: str = ""
    intents: List[IntentModel] = []

