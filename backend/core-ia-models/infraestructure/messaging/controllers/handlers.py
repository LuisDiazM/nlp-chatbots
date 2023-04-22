from typing import Any, Coroutine
import json

from domain.models.test_chatbot_request import TestChatbotRequest
from domain.usecases.chatbot_responses_usecase import ChatbotResponseUsecase


class ChatbotResponseHandler:

    def __init__(self, chatbot_response_usecase: ChatbotResponseUsecase, nats_client:Coroutine[Any, Any, None]) -> None:
        self.chatbot_response_usecase = chatbot_response_usecase
        self.nats_client = nats_client

    async def chatbot_response_handler(self, msg):
        subject = msg.subject
        reply = msg.reply
        data = json.loads(msg.data.decode())
        print("Received a message on '{subject} {reply}': {data}".format(
            subject=subject, reply=reply, data=data))
        request_data = TestChatbotRequest(**data)
        response = self.chatbot_response_usecase.chatbot_response(
            request_data.model_id, request_data.sentence)
        await self.nats_client.publish(reply, f'{response}'.encode())
