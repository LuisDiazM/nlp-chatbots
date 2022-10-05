
from pydantic import BaseModel


class TrainChatbotRequest(BaseModel):
    id: str
    userId: str = ''