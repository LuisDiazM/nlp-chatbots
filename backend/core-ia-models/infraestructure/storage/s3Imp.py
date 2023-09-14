import os
import boto3
from infraestructure.gateways import StorageGateway
from domain.helpers.loggers import logger

class S3Imp(StorageGateway):

    def __init__(self) -> None:
        try:
            self.client = boto3.client('s3', aws_access_key_id=os.getenv("ACCESS_KEY"),
                                    aws_secret_access_key=os.getenv("SECRET_ACCESS_KEY"))
            self.bucket_name = os.getenv("BUCKET_NAME")
        except Exception as e:
            logger.error(f"{str(e)}")

    def download(self, object_name: str) -> str:
        BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
        path_data = os.path.join(BASE_DIR, 'bin', object_name)
        self.client.download_file(self.bucket_name, object_name, path_data)
        return path_data
