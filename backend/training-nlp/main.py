from domain.usecases.training_usecase import TrainingUsecase

usecase = TrainingUsecase()
model = usecase.get_training_data("62f9b0c2da227b1a1066dee4")
print(model)