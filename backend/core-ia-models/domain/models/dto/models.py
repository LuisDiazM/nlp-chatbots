
from datetime import datetime
from pydantic import BaseModel


class ModelsDTO(BaseModel):
    modelName: str
    bucketName: str
    nluIntentId: str
    created: datetime
    userId: str
