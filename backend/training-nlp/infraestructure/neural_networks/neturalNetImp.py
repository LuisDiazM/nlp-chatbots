import torch
from torch.utils.data import DataLoader
import torch.nn as nn
import os
import uuid
from infraestructure.neural_networks.nn_models.nn_model import NeuralNet
from infraestructure.neural_networks.data_models.chat_model import ChatDataSet
from domain.models.preprocessing_training_model import PreprocessTrainingModel
from domain.models.model_trained import ModelTrained


class NeuralNetImp:
    batch_size = 8
    hidden_size = 8
    learning_rate = 0.001
    num_epochs = 1000

    def run_nn(self, preprocess_training: PreprocessTrainingModel) -> str:
        data_set = self.__prepare_dataset(
            preprocess_training.x_train, preprocess_training.y_train)
        train_loader = self.__prepare_train_loader(data_set)
        device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
        model = NeuralNet(input_size=preprocess_training.input_size,
                          hidden_size=self.hidden_size, num_classes=preprocess_training.output_size)

        # loss optimizer
        criterion = nn.CrossEntropyLoss()
        optimizer = torch.optim.Adam(model.parameters(), lr=self.learning_rate)

        for epoch in range(self.num_epochs):
            for (words, labels) in train_loader:
                words = words.to(device, dtype=torch.float)
                labels = labels.to(device)
                # forward
                outputs = model.forward(words)
                loss = criterion(outputs, labels)

                # backward
                optimizer.zero_grad()
                loss.backward()
                optimizer.step()

            if epoch % 100 == 0:
                print(
                    f"epoch {epoch+100}/{self.num_epochs}, loss={loss.item()}")

        data_model = self.__generate_data_model(model, preprocess_training)
        model_path = self.__generate_model_path()

        torch.save(data_model.dict(), model_path)
        return model_path
    
    def __generate_model_path(self)->str:
        BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
        model_uid = uuid.uuid4()
        return os.path.join(BASE_DIR, "bin", f"{model_uid}.pth")

    def __generate_data_model(self, model: NeuralNet, preprocess_training: PreprocessTrainingModel) -> ModelTrained:
        data_model = {"model_state": model.state_dict(),
                      "input_size": preprocess_training.input_size,
                      "output_size": preprocess_training.output_size,
                      "hidden_size": self.hidden_size,
                      "all_words": preprocess_training.all_words,
                      "tags": preprocess_training.tags}
        return ModelTrained(**data_model)

    def __prepare_dataset(self, x_train, y_train) -> ChatDataSet:
        return ChatDataSet(x_train=x_train, y_train=y_train)

    def __prepare_train_loader(self, dataset: ChatDataSet) -> DataLoader:
        return DataLoader(dataset=dataset, batch_size=self.batch_size, shuffle=True, num_workers=0)
