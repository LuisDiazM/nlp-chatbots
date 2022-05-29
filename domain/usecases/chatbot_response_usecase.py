
from typing import List
from infraestructure.nlp.nlp_imp import NLPImp


class ChatbotUsecase:

    def __init__(self, pairs:List) -> None:
        self.nlp = NLPImp(pairs)

    def converse_chat(self, content: str) -> str:
        return self.nlp.converse(content)
