import asyncio

from helpers.constants import SUBSCRIPTION_TRAINING_MODEL

from infraestructure.neural_networks.neturalNetImp import NeuralNetImp
from infraestructure.preprocessing.nlp_preprocessing import PreprocessingNLP
from infraestructure.database.mongoImp import MongoImp
from app.controller.handlers import ControllerSubscriptions
from infraestructure.messaging.natsImp import NatsImp


async def main():
    database = MongoImp()
    preprocess_nlp = PreprocessingNLP()
    neural_net = NeuralNetImp()

    controllers_instance = ControllerSubscriptions(
        database, preprocess_nlp, neural_net)
    
    nats_instance = NatsImp()
    await nats_instance.set_up()
    client = nats_instance.client

    #subscriptors
    await client.subscribe(SUBSCRIPTION_TRAINING_MODEL, cb=controllers_instance.training_model_handler)


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    try:
        asyncio.ensure_future(main())
        loop.run_forever()
    finally:
        loop.close()
