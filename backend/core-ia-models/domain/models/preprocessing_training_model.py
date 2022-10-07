from typing import Any, List
from pydantic import BaseModel


class PreprocessTrainingModel(BaseModel):
    x_train: List[Any] = []
    y_train: List[int] = []
    all_words: List[str] = []
    tags: list = []
    input_size: int
    output_size: int
