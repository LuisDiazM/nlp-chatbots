import asyncio
import json
import nats
from nats.errors import ConnectionClosedError, TimeoutError, NoServersError

async def main():

    nc = await nats.connect("nats://localhost:4222")

   
    await nc.publish("training.chatbots", json.dumps({"id": "62fab4f3b11e482d091a6b05"}).encode())


    # Terminate connection to NATS.
    await nc.drain()

if __name__ == '__main__':
    loop = asyncio.get_event_loop()

    try:
        # asyncio.ensure_future(main())
        loop.run_until_complete(main())
    finally:
        loop.close()