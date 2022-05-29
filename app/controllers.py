from dependency_injector.wiring import Provide, inject
from flask_restful import Resource
from flask_restful import reqparse

from app.di import Container
from domain.usecases.chatbot_response_usecase import ChatbotUsecase
from domain.usecases.home_usecase import HomeUseCase


class HomeHandler(Resource):
    @inject
    def __init__(self, home_usecase: HomeUseCase = Provide[Container.home_usecase]) -> None:
        self.home_usecase = home_usecase
        super().__init__()

    def get(self):
        home = self.home_usecase.home()
        return {'spectrum': home}, 200


class ConverseChatbot(Resource):
    def __init__(self) -> None:
        super().__init__()

    def post(self):
        parser = reqparse.RequestParser()
        parser.add_argument("content", type=str, required=True)
        args = parser.parse_args()
        content = args.get("content")

        pairs = [[
            r"my name is (.*)",
            ["Hello, How are you today ?", ]
        ],
            [
            r"hi|hey|hello",
            ["Hello", "Hey there", ]
        ],
            [
            r"what is your name ?",
            ["I am a bot created by Analytics Vidhya. you can call me crazy!", ]
        ],
            [
            r"how are you ?",
            ["I'm doing goodnHow about You ?", ]
        ],
            [
            r"sorry (.*)",
            ["Its alright", "Its OK, never mind", ]
        ],
            [
            r"I am fine",
            ["Great to hear that, How can I help you?", ]
        ],
            [
            r"i'm (.*) doing good",
            ["Nice to hear that", "How can I help you?:)", ]
        ],
            [
            r"(.*) age?",
            ["I'm a computer program dudenSeriously you are asking me this?", ]
        ],
            [
            r"what (.*) want ?",
            ["Make me an offer I can't refuse", ]
        ],
            [
            r"(.*) created ?",
            ["Raghav created me using Python's NLTK library ", "top secret ;)", ]
        ],
            [
            r"(.*) (location|city) ?",
            ['Indore, Madhya Pradesh', ]
        ],
            [
            r"how is weather in (.*)?",
            ["Weather in %1 is awesome like always", "Too hot man here in %1",
                "Too cold man here in %1", "Never even heard about %1"]
        ],
            [
            r"i work in (.*)?",
            ["%1 is an Amazing company, I have heard about it. But they are in huge loss these days.", ]
        ],
            [
            r"(.*)raining in (.*)",
            ["No rain since last week here in %2",
                "Damn its raining too much here in %2"]
        ],
            [
            r"how (.*) health(.*)",
            ["I'm a computer program, so I'm always healthy ", ]
        ],
            [
            r"(.*) (sports|game) ?",
            ["I'm a very big fan of Football", ]
        ],
            [
            r"who (.*) sportsperson ?",
            ["Messy", "Ronaldo", "Roony"]
        ],
            [
            r"who (.*) (moviestar|actor)?",
            ["Brad Pitt"]
        ],
            [
            r"i am looking for online guides and courses to learn data science, can you suggest?",
            ["Crazy_Tech has many great articles with each step explanation along with code, you can explore"]
        ],
            [
            r"quit",
            ["BBye take care. See you soon :) ",
                "It was nice talking to you. See you soon :)"]
        ],
        ]

        usecase = ChatbotUsecase(pairs)
        response_data = usecase.converse_chat(content)
        if response_data is None:
            return "", 200
        return response_data, 200
