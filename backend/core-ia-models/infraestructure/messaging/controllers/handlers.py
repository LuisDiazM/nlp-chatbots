from typing import Any, Coroutine
import json
from domain.helpers.constants import FEATURE_RATE_LIMIT, SUBJECT_UPDATE_LICENSE_USAGE

from domain.models.test_chatbot_request import TestChatbotRequest
from domain.usecases.chatbot_responses_usecase import ChatbotResponseUsecase
from domain.helpers.loggers import logger
import time

class ChatbotResponseHandler:

    def __init__(self, chatbot_response_usecase: ChatbotResponseUsecase, nats_client:Coroutine[Any, Any, None]) -> None:
        self.chatbot_response_usecase = chatbot_response_usecase
        self.nats_client = nats_client

    async def chatbot_response_handler(self, msg):
        subject = msg.subject
        reply = msg.reply
        data = json.loads(msg.data.decode())
        logger.info("Received a message on '{subject} {reply}': {data}".format(
            subject=subject, reply=reply, data=data))
        request_data = TestChatbotRequest(**data)
        t1 = time.time()
        response = self.chatbot_response_usecase.chatbot_response(
            request_data.model_id, request_data.sentence)
        t2 = time.time()
        logger.info(f"elapsed time on {data} {t2-t1} seconds")
        await self.nats_client.publish(reply, f'{response}'.encode())
        await self.nats_client.publish(SUBJECT_UPDATE_LICENSE_USAGE, json.dumps({"user_id": request_data.user_id, "feature": FEATURE_RATE_LIMIT}).encode())

