import json
from multiprocessing.connection import Client
import os
from typing import Any, Coroutine, Dict
import nats
from domain.helpers.loggers import logger

class NatsImp:

    def __init__(self) -> None:
        self.client: Coroutine[Any, Any, Client]

    async def set_up(self):
        try:
            url = f"nats://{os.getenv('NATS_URL')}:{os.getenv('NATS_PORT')}"
            logger.info(url)
            self.client = await nats.connect(url)
            logger.info("NATS connected")
        except Exception as e:
            logger.error(str(e))

    async def shutdown(self):
        await self.client.drain()
        await self.client.close()

    async def publish_event(self, subject:str, data:Dict) -> None:
        await self.client.publish(subject, json.dumps(data).encode())
        logger.info(f"publish notification {subject} {data}")