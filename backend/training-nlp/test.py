import asyncio
import json
from domain.helpers.constants import SUBSCRIPTION_TRAINING_MODEL_COMMAND
import nats

async def main():

    nc = await nats.connect("nats://nlpchatbots:random94566546@localhost:4222")

   
    # Send the request
    await nc.publish(SUBSCRIPTION_TRAINING_MODEL_COMMAND, json.dumps({"id": "64ee7544f0246c7c3fd37f37", "user_id":"luismigueldiazmorales@gmail.com"}).encode())
    
    # try:
    #     msg = await nc.request("query.response.chatbots", json.dumps({"model_id": "64454e441814ae7ca31e216a", "sentence": "Hi rats "}).encode(), timeout=3)
    #     # Use the response
    #     print("Reply:", msg)
    #     pass
    # except asyncio.TimeoutError:
    #     print("Timed out waiting for response")


    # Terminate connection to NATS.
    await nc.drain()

if __name__ == '__main__':
    loop = asyncio.get_event_loop()

    try:
        # asyncio.ensure_future(main())
        loop.run_until_complete(main())
    finally:
        loop.close()