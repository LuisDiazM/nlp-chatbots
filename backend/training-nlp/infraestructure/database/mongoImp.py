import os
from typing import Dict
from pymongo import MongoClient
from bson import ObjectId
from domain.helpers.loggers import logger


class MongoImp:
    client: MongoClient

    def __init__(self):
        self.set_up()
        logger.info(f"connection Mongo successfull")

    def set_up(self) -> None:
        url = f"mongodb://{os.getenv('MONGO_USER')}:{os.getenv('MONGO_PASSWORD')}@{os.getenv('MONGO_URL')}:{os.getenv('MONGO_PORT')}"
        logger.info(url)
        self.client = MongoClient(url)

    def find_by_id(self, id: str, database: str, collection: str) -> Dict:
        col = self.client.get_database(
            database).get_collection(collection)
        return col.find_one({"_id": ObjectId(id)})

    def shutdown(self):
        self.client.close()

    def insert_one(self, database: str, collection: str, data: Dict) -> None:
        col = self.client[database][collection]
        col.insert_one(data)

    def delete_by_id(self, database: str, collection: str, id: str) -> bool:
        try:
            col = self.client[database][collection]
            filter_to_delete = {"_id": ObjectId(id)}
            col.delete_one(filter_to_delete)
            return True
        except Exception as e:
            logger.error(f"could not delete document {str(e)}")
            return False

    def find_documents(self, database: str, collection: str, filter: Dict):
        col = self.client.get_database(
            database).get_collection(collection)
        cursor = col.find(filter)
        documents = []
        for doc in cursor:
            documents.append(doc)
        return documents