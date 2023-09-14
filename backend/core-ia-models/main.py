import asyncio
from dependency_injector.wiring import Provide, inject
from cmd.config import set_env
from cmd.core_di import Container
from domain.usecases.chatbot_responses_usecase import ChatbotResponseUsecase
from domain.helpers.constants import QUEUE_CORE_IA, SUBSCRIPTION_CHATBOT_RESPONSE
from infraestructure.messaging.controllers.handlers import ChatbotResponseHandler
from infraestructure.messaging.natsImp import NatsImp
from domain.helpers.loggers import logger
set_env()

@inject
async def main(chatbot_usecase: ChatbotResponseUsecase = Provide[Container.chatbot_usecase]):

    # NATS client listen connections
    nats_instance = NatsImp()
    await nats_instance.set_up()
    client = nats_instance.client
    controller = ChatbotResponseHandler(
        chatbot_response_usecase=chatbot_usecase, nats_client=client)
    logger.info("core ia models start!!!")
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
