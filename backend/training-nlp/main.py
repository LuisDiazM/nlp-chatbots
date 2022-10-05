import asyncio

from domain.usecases.training_usecase import TrainingUsecase

from app.di import Container
from helpers.constants import SUBSCRIPTION_TRAINING_MODEL_COMMAND
from app.controller.handlers import ControllerSubscriptions
from infraestructure.messaging.natsImp import NatsImp
from dependency_injector.wiring import Provide, inject


@inject
async def main(training_usecase: TrainingUsecase = Provide[Container.training_usecase]):

    controllers_instance = ControllerSubscriptions(training_usecase)
    
    # NATS client listen connections
    nats_instance = NatsImp()
    await nats_instance.set_up()
    client = nats_instance.client

    # subscriptors
    await client.subscribe(SUBSCRIPTION_TRAINING_MODEL_COMMAND, cb=controllers_instance.training_model_handler)


if __name__ == '__main__':
    container = Container()
    container.init_resources()
    container.wire(modules=[__name__])

    loop = asyncio.get_event_loop()
    try:
        asyncio.ensure_future(main())
        loop.run_forever()
    finally:
        loop.close()
