from dependency_injector.wiring import Provide, inject
import json

from app.core_di import Container
from domain.usecases.chatbot_responses_usecase import ChatbotResponseUsecase


class ChatbotResponseHandler:

    @inject
    def __init__(self, chatbot_response_usecase: ChatbotResponseUsecase = Provide[Container.chatbot_usecase]) -> None:
        self.chatbot_response_usecase = chatbot_response_usecase
        super().__init__()

    def chatbot_response_handler(self):
        # subject = msg.subject
        # reply = msg.reply
        # data = json.loads(msg.data.decode())
        # print("Received a message on '{subject} {reply}': {data}".format(
        #     subject=subject, reply=reply, data=data))
        # request_data = TestChatbotRequest(**data)
        response = self.chatbot_response_usecase.chatbot_response("6340402b15312bb8a5d9e8ca","tell me a joke")
        print(response)
           