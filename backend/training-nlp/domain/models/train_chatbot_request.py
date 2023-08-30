
from pydantic import BaseModel


class TrainChatbotRequest(BaseModel):
    id: str
    user_id: str = ''