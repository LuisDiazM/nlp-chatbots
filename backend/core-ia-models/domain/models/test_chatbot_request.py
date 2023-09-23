from pydantic import BaseModel


class TestChatbotRequest(BaseModel):
    model_id: str
    sentence: str
    user_id : str
