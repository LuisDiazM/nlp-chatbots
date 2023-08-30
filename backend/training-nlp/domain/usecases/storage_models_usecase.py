from datetime import datetime
import os

from domain.helpers.constants import DATABASE_TRAINING, COLLECTION_MODELS

from domain.models.dto.models import ModelsDTO

from infraestructure.database.mongoImp import MongoImp
from infraestructure.storage.s3Imp import S3Imp


class StorageModelsUsecase:

    def __init__(self, storage_gateway: S3Imp, database_gateway: MongoImp) -> None:
        self.storage_gateway = storage_gateway
        self.database_gateway = database_gateway

    def upload_ia_model_to_s3(self, path_file: str, user_id:str, training_id: str = None):
        file_upload = self.storage_gateway.upload(path_file)
        bucket_name = self.storage_gateway.bucket_name
        print(bucket_name, file_upload)
        os.remove(path_file)
        if training_id is not None:
            model_data = ModelsDTO(modelName=file_upload, bucketName=bucket_name,
                                   nluIntentId=training_id, created=datetime.now(), userId=user_id)
            self.database_gateway.insert_one(DATABASE_TRAINING, COLLECTION_MODELS, model_data.dict())
