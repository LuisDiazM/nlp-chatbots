import sys
import app.controllers as controllers
from app.routes import create_app
from app.di import Container

if __name__ == '__main__':
    container = Container()
    container.wire([sys.modules[__name__], controllers])
    app = create_app(container)
    app.run(debug=True)