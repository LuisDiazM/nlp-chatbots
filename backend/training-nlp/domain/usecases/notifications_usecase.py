from typing import Dict
from domain.usecases.gateways import Messaging


class EventNotificationUsecase:
    def __init__(self, messaging_gateway: Messaging) -> None:
        self.messaging_gateway = messaging_gateway

    async def publish_notification(self, subject:str, data:Dict)->None:
        await self.messaging_gateway.set_up()
        await self.messaging_gateway.publish_event(subject, data)
        await self.messaging_gateway.shutdown()