import os
import boto3
from botocore.exceptions import NoCredentialsError
from domain.helpers.loggers import logger


class S3Imp:

    def __init__(self) -> None:
        self.client = boto3.client('s3', aws_access_key_id=os.getenv("ACCESS_KEY"),
                                   aws_secret_access_key=os.getenv("SECRET_ACCESS_KEY"))
        self.bucket_name = os.getenv("BUCKET_NAME")

    def upload(self, path_file: str) -> str:
        try:
            file_name = os.path.basename(path_file)
            self.client.upload_file(path_file, self.bucket_name, file_name)
            return file_name
        except FileNotFoundError:
            logger.warning(f"The file {path_file} was not found")
            return ""
        except NoCredentialsError:
            logger.error("Credentials not available")
            return ""

    def delete_file(self, path: str) -> bool:
        try:
            self.client.delete_object(Bucket=self.bucket_name, Key=path)
            return True
        except Exception as e:
            logger.error(f"file s3 could not deleted {str(e)}")
            return False
