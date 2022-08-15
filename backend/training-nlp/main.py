from domain.usecases.training_usecase import TrainingUsecase

usecase = TrainingUsecase()
model_path = usecase.generate_model("62fab4f3b11e482d091a6b05")
print(model_path)