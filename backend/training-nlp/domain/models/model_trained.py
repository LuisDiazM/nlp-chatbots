from typing import Dict, List
from pydantic import BaseModel


class ModelTrained(BaseModel):
    model_state: Dict
    input_size: int
    output_size: int
    hidden_size: int
    all_words: List[str]
    tags: list
