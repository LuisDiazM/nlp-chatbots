
from pydantic import BaseModel


class TrainChatbotRequest(BaseModel):
    id: str
    user_id: str = ''

class DeleteModelsRequest(BaseModel):
    training_id: str

