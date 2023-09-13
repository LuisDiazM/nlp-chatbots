import asyncio
from cmd.config import set_env
from cmd.di import Container

from dependency_injector.wiring import Provide, inject
from domain.helpers.constants import (QUEUE_MANAGE_MODELS, QUEUE_TRAINING_NLP, SUBSCRIPTION_REMOVE_MODELS, SUBSCRIPTION_REMOVE_MODELS_BY_USER,
                                      SUBSCRIPTION_TRAINING_MODEL_COMMAND)
from domain.usecases.storage_models_usecase import StorageModelsUsecase
from domain.usecases.training_usecase import TrainingUsecase
from infraestructure.messaging.controller.handlers import \
    ControllerSubscriptions
from infraestructure.messaging.natsImp import NatsImp
from domain.helpers.loggers import logger
set_env()

@inject
async def main(training_usecase: TrainingUsecase = Provide[Container.training_usecase],
               storage_usecase: StorageModelsUsecase = Provide[Container.storage_usecase]):

    controllers_instance = ControllerSubscriptions(
        training_usecase, storage_usecase)

    # NATS client listen connections
    nats_instance = NatsImp()
    await nats_instance.set_up()
    client = nats_instance.client
    logger.info("training nlp start!!")
    # subscriptors
    await client.subscribe(SUBSCRIPTION_TRAINING_MODEL_COMMAND, queue=QUEUE_TRAINING_NLP, cb=controllers_instance.training_model_handler)
    await client.subscribe(SUBSCRIPTION_REMOVE_MODELS, queue=QUEUE_MANAGE_MODELS, cb=controllers_instance.delete_models_by_trainig_id)
    await client.subscribe(SUBSCRIPTION_REMOVE_MODELS_BY_USER, queue=QUEUE_MANAGE_MODELS, cb=controllers_instance.delete_models_by_user_id)


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
