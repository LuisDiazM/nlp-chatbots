from dependency_injector.wiring import Provide, inject
from app.controllers.handlers import ChatbotResponseHandler
from app.core_di import Container
from domain.usecases.chatbot_responses_usecase import ChatbotResponseUsecase


@inject
def main(chatbot_usecase: ChatbotResponseUsecase = Provide[Container.chatbot_usecase]):
    controller = ChatbotResponseHandler(
        chatbot_response_usecase=chatbot_usecase)
    controller.chatbot_response_handler()


if __name__ == "__main__":
    container = Container()
    container.init_resources()
    container.wire(modules=[__name__])
    main()
