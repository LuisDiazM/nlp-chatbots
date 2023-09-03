from multiprocessing.connection import Client
import os
from typing import Any, Coroutine
import nats
from domain.helpers.loggers import logger

class NatsImp:

    def __init__(self) -> None:
        self.client: Coroutine[Any, Any, Client]

    async def set_up(self):
        try:
            self.client = await nats.connect(os.getenv("NATS_URL"))
            logger.info("connection NATS successfull")
        except Exception as e:
            logger.error(f"{str(e)}")

    async def shutdown(self):
        await self.client.drain()
