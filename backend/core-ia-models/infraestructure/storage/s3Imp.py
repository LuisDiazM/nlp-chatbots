import os
import boto3
from botocore.exceptions import NoCredentialsError
from infraestructure.gateways import StorageGateway


class S3Imp(StorageGateway):

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
            print(f"The file {path_file} was not found")
            return ""
        except NoCredentialsError:
            print("Credentials not available")
            return ""

    def download(self, object_name: str) -> str:
        BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
        path_data = os.path.join(BASE_DIR, 'bin', object_name)
        self.client.download_file(self.bucket_name, object_name, path_data)
        return path_data
