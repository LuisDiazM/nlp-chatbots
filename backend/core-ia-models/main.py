import asyncio
import json
from dependency_injector.wiring import Provide, inject
from app.controllers.handlers import ChatbotResponseHandler
from app.core_di import Container
from domain.usecases.chatbot_responses_usecase import ChatbotResponseUsecase
from helpers.constants import QUEUE_CORE_IA, SUBSCRIPTION_CHATBOT_RESPONSE
from infraestructure.messaging.natsImp import NatsImp


@inject
async def main(chatbot_usecase: ChatbotResponseUsecase = Provide[Container.chatbot_usecase]):

    # NATS client listen connections
    nats_instance = NatsImp()
    await nats_instance.set_up()
    client = nats_instance.client

    controller = ChatbotResponseHandler(
        chatbot_response_usecase=chatbot_usecase, nats_client=client)

    #subscribers
    await client.subscribe(SUBSCRIPTION_CHATBOT_RESPONSE ,queue= QUEUE_CORE_IA, cb=controller.chatbot_response_handler)


if __name__ == "__main__":
    container = Container()
    container.init_resources()
    container.wire(modules=[__name__])

    loop = asyncio.get_event_loop()
    try:
        asyncio.ensure_future(main())
        loop.run_forever()
    finally:
        loop.close()
