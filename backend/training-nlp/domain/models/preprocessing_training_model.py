from typing import Any, List
from pydantic import BaseModel
import numpy as np


class PreprocessTrainingModel(BaseModel):
    X_train: List[Any] = []
    y_train: List[int] = []
    all_words: List[str] = []
    tags: list = []
    input_size: int
    output_size: int
