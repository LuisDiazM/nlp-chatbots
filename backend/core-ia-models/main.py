from domain.usecases.chatbot_responses_usecase import ChatbotResponseUsecase
from infraestructure.database.mongoImp import MongoImp
from infraestructure.preprocessing.nlp_preprocessing import PreprocessingNLP
from infraestructure.storage.s3Imp import S3Imp



s3imp = S3Imp()
nlp_pipeline = PreprocessingNLP()
database = MongoImp()
usecase = ChatbotResponseUsecase(s3imp, nlp_pipeline, database)
response = usecase.chatbot_response("6340402b15312bb8a5d9e8ca","tell me a joke")
print(response)