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

    controllersInstance = ControllerSubscriptions(
        database, preprocess_nlp, neural_net)
    
    natsInstance = NatsImp()
    await natsInstance.set_up()
    client = natsInstance.client

    #subscriptors
    await client.subscribe(SUBSCRIPTION_TRAINING_MODEL, cb=controllersInstance.training_model_handler)


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    try:
        asyncio.ensure_future(main())
        loop.run_forever()
    finally:
        loop.close()
