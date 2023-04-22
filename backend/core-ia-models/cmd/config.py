from dotenv import load_dotenv
from os.path import join, dirname

def set_env():
    dotenv_path = join(dirname(__file__), '.env')
    load_dotenv(dotenv_path)

