from typing import List
import nltk
from nltk.stem import PorterStemmer
import numpy as np
from domain.models.dto.training_model import TrainingModel

class PreprocessingNLP:
    input_len = 0
    output_len = 0
    all_words = []
    tags = []

    def __init__(self):
        self.steamer = PorterStemmer()
        self.ignore_words = ["?", ",", "!", "."]

    def tokenization(self, sentence: str) -> List[str]:
        return nltk.word_tokenize(sentence)

    def __steaming(self, word: str) -> str:
        return self.steamer.stem(word=word, to_lowercase=True)

    def bag_of_words(self, tokenize_sentence: List[str], all_words: List[str]) -> np.ndarray:
        """
        :param tokenize_sentence: ["hello", "how", "are", "you"]
        :param all_words: ["hi", "hello", "how", "are", "you", "bye", "thank"]
        :return: [0, 1, 1, 1, 1, 0, 0]
        """
        steam_tokenize_sentence = [self.__steaming(
            word) for word in tokenize_sentence]
        bag = np.zeros(len(all_words), dtype=float)
        for index, word in enumerate(all_words):
            if word in steam_tokenize_sentence:
                bag[index] = 1.0
        return bag


