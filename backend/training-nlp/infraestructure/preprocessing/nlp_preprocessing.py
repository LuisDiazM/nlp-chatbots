from typing import List
import nltk
from nltk.stem import PorterStemmer
import numpy as np
from domain.models.training_model import TrainingModel

class PreprocessingNLP:
    """
    1. tokenization
    2. steaming
    3. exclude words
    4. bag of words
    """
    input_len = 0
    output_len = 0
    all_words = []
    tags = []

    def __init__(self):
        self.steamer = PorterStemmer()
        self.ignore_words = ["?", ",", "!", "."]

    def prepare_data(self, training_data:TrainingModel) -> tuple:
        xy = []

        for intent in training_data.intents:
            tag = intent.tag
            self.tags.append(tag)
            for pattern in intent.patterns:
                w = self.__tokenization(pattern)
                self.all_words.extend(w)
                xy.append((w, tag))

        self.all_words = self.__exclude_words(self.all_words)
        self.tags = sorted(set(self.tags))

        ## training
        x_train = []
        y_train = []

        for (pattern_sentence, tag) in xy:
            bag = self.__bag_of_words(pattern_sentence, self.all_words)
            x_train.append(bag)
            label = self.tags.index(tag)
            y_train.append(label)

        self.input_len = len(self.all_words)
        self.output_len = len(self.tags)

        return x_train, y_train

    def __tokenization(self, sentence: str) -> List[str]:
        return nltk.word_tokenize(sentence)

    def __steaming(self, word: str) -> str:
        return self.steamer.stem(word=word, to_lowercase=True)

    def __exclude_words(self, all_words: List[str]) -> List[str]:
        steam_words = [self.__steaming(
            word) for word in all_words if word not in self.ignore_words]
        return sorted(set(steam_words))

    def __bag_of_words(self, tokenize_sentence: List[str], all_words: List[str]) -> np.ndarray:
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


