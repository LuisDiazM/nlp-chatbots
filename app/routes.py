from app.di import Container
from flask_restful import Api
from flask import Flask
from app.controllers import *


def create_app(container: Container) -> Flask:
    app = Flask(__name__)
    api = Api(app, prefix="/api/v1")
    app.container = container

    # routes
    api.add_resource(HomeHandler, "/home")
    api.add_resource(ConverseChatbot, "/converse/web")
    return app
