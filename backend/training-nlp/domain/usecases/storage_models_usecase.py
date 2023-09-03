from datetime import datetime
import os
import time

from domain.helpers.constants import DATABASE_TRAINING, COLLECTION_MODELS

from domain.models.dto.models import ModelsDTO
from domain.helpers.loggers import logger
from domain.usecases.gateways import DatabaseGateway, StorageGateway


class StorageModelsUsecase:

    def __init__(self, storage_gateway: StorageGateway, database_gateway: DatabaseGateway) -> None:
        self.storage_gateway = storage_gateway
        self.database_gateway = database_gateway

    def upload_ia_model_to_s3(self, path_file: str, user_id:str, training_id: str = None):
        file_upload = self.storage_gateway.upload(path_file)
        bucket_name = self.storage_gateway.bucket_name
        logger.info(f"{bucket_name} - {file_upload}")
        os.remove(path_file)
        if training_id is not None:
            model_data = ModelsDTO(modelName=file_upload, bucketName=bucket_name,
                                   nluIntentId=training_id, created=datetime.now(), userId=user_id)
            self.database_gateway.insert_one(DATABASE_TRAINING, COLLECTION_MODELS, model_data.dict())

    def delete_models_by_training_id(self, training_id:str):
        filters = {"nluIntentId":training_id}
        documents = self.database_gateway.find_documents(DATABASE_TRAINING, COLLECTION_MODELS, filters)
        start_time = time.time()
        if len(documents)>0:
            for doc in documents:
                model_id = doc.get("_id","")
                model_path = doc.get("modelName","")
                self.storage_gateway.delete_file(model_path)
                self.database_gateway.delete_by_id(DATABASE_TRAINING, COLLECTION_MODELS, model_id)
        end_time = time.time()
        elapsed_time = end_time - start_time
        logger.info(f"models with training_id {training_id} were deleted in {elapsed_time} seconds")