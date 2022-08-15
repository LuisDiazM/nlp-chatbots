import os
from typing import Dict
from pymongo import MongoClient
from bson import ObjectId
from infraestructure.gateways import DatabaseGateway

class MongoImp(DatabaseGateway):
    client: MongoClient

    def __init__(self):
        self.set_up()

    def set_up(self) -> None:
        self.client = MongoClient(os.getenv("MONGO_URL"))

    def find_by_id(self, id: str, database: str, collection: str) -> Dict:
        collection = self.client.get_database(
            database).get_collection(collection)
        return collection.find_one({"_id": ObjectId(id)})
        
