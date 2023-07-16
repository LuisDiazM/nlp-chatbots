# NLP Chatbots
The project is composed by three microservices and a frontend, the main microservice http-models-server is a http server that comunicate with te other microservices via pub/sub using a NATS as messaging broker:

![Structure](/docs/images/NLP%20chatbots-Page-4.jpg)

## Data model
The microservices training-nlp and core-ia-models share the same database to generate and access the models

![Models](/docs/images/NLP%20chatbots-Modelo%20datos.jpg)