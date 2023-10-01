import os
from typing import Dict
from pymongo import MongoClient
from bson import ObjectId
from infraestructure.gateways import DatabaseGateway
from domain.helpers.loggers import logger

class MongoImp(DatabaseGateway):
    client: MongoClient

    def __init__(self):
        self.set_up()

    def set_up(self) -> None:
        try:
            url = f"mongodb://{os.getenv('MONGO_USER')}:{os.getenv('MONGO_PASSWORD')}@{os.getenv('MONGO_URL')}:{os.getenv('MONGO_PORT')}"
            logger.info(url)
            self.client = MongoClient(url)
            logger.info("Mongo connected")
        except Exception as e:
            logger.error(str(e))

    def find_by_id(self, id: str, database: str, collection: str) -> Dict:
        collection = self.client.get_database(
            database).get_collection(collection)
        return collection.find_one({"_id": ObjectId(id)})
        
    def shutdown(self):
        self.client.close()
    
        
