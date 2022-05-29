from typing import List
from nltk.chat.util import Chat


class NLPImp:
    def __init__(self, pairs: List) -> None:
        print("chatbot ...")
        self.chatbot: Chat = Chat(pairs)

    def converse(self, content: str) -> str:
        return self.chatbot.respond(content)
