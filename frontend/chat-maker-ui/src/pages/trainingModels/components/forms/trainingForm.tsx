import React, { useState } from "react";
import { Formik, Field, FieldArray, Form } from "formik";
import Button from "@mui/material/Button";
import Accordion from "@mui/material/Accordion";
import AccordionSummary from "@mui/material/AccordionSummary";
import AccordionDetails from "@mui/material/AccordionDetails";
import Typography from "@mui/material/Typography";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import IconButton from "@mui/material/IconButton";
import DeleteIcon from "@mui/icons-material/Delete";
import AddIcon from "@mui/icons-material/Add";
import CardContent from "@mui/material/CardContent";
import Card from "@mui/material/Card";
import SendIcon from "@mui/icons-material/Send";

import CardActions from "@mui/material/CardActions";
import Tooltip from "@mui/material/Tooltip";
import {
  Intent,
  TrainingModel,
  createTrainingModel,
  updateTrainingModel,
} from "../../../../shared/utilities/trainingModels";
import { useSelector } from "react-redux";
import { CircularProgress } from "@mui/material";
import { ACCESS_TOKEN } from "../../../../shared/utilities/constants";
import styles from './trainingForm.module.scss';

const inputStyles = { marginRight: "10px" };

interface TrainingFormData {
  model: TrainingModel;
  isEdit: boolean;
  closeModal: React.Dispatch<React.SetStateAction<boolean>>;
}

const TrainingChatbotForm: React.FC<TrainingFormData> = ({
  model,
  isEdit,
  closeModal,
}) => {
  const userState = useSelector((store: any) => store.user);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const token = sessionStorage.getItem(ACCESS_TOKEN);

  const handleSubmit = async (values: TrainingModel): Promise<void> => {
    if (isEdit) {
      setIsLoading(true);
      const updateModel: TrainingModel = {
        ...values,
        id: model.id,
      };
      await updateTrainingModel(token ?? "", updateModel);
      setIsLoading(false);
      closeModal(false);
    } else {
      setIsLoading(true);
      const trainingData: TrainingModel = {
        ...values,
        user_id: userState.email,
      };
      await createTrainingModel(token ?? "", trainingData);
      setIsLoading(false);
      closeModal(false);
    }
  };

  return (
    <>
      <div>
        <Formik initialValues={{ ...model }} onSubmit={handleSubmit}>
          {({ values }) => (
            <Form>
              <div className={styles.formContainer}>
                <label htmlFor="title">Nombre del modelo</label>
                <Field id="title" name="title"></Field>
                <label htmlFor="description">Descripción del modelo</label>
                <Field
                  id="description"
                  name="description"
                  as="textarea"
                ></Field>
                <div className={styles.intentsContainer}>
                <FieldArray name="intents">
                  {(intentsFields: {
                    remove: (arg0: number) => void;
                    push: (arg0: Intent) => void;
                  }) => (
                    <>
                      {values?.intents?.map((intentName, index) => (
                        <Accordion key={`${intentName}-${index}`}>
                          <AccordionSummary
                            expandIcon={<ExpandMoreIcon />}
                            aria-controls="panel2a-content"
                            id="panel2a-header"
                          >
                            <Typography>
                              {intentName.tag ?? "intento No 1"}
                            </Typography>
                          </AccordionSummary>
                          <AccordionDetails>
                            <Card sx={{ minWidth: 275 }}>
                              <div key={index}>
                                <CardContent>
                                  <Typography
                                    sx={{ fontSize: 14 }}
                                    color="text.secondary"
                                    gutterBottom
                                  >
                                    <label
                                      htmlFor={`intents[${index}].tag`}
                                      style={inputStyles}
                                    >
                                      Etiqueta del intento
                                    </label>
                                    <Field
                                      id={`intents[${index}].tag`}
                                      name={`intents[${index}].tag`}
                                    />{" "}
                                  </Typography>
                                  <Typography variant="body2">
                                    <FieldArray
                                      name={`intents[${index}].patterns`}
                                    >
                                      {(patternsFields: {
                                        remove: (arg0: number) => void;
                                        push: (arg0: string) => void;
                                      }) => (
                                        <div key={`${values.intents[index]}`}>
                                          <h2>Frases clave</h2>
                                          {values.intents[index].patterns.map(
                                            (_, patternIndex) => (
                                              <div key={patternIndex}>
                                                <Field
                                                  style={{ width: "80%" }}
                                                  id={`intents[${index}].patterns.[${patternIndex}]`}
                                                  name={`intents[${index}].patterns.[${patternIndex}]`}
                                                />

                                                <IconButton
                                                  color="error"
                                                  onClick={() => {
                                                    if (
                                                      values.intents[index]
                                                        .patterns.length > 5
                                                    ) {
                                                      patternsFields.remove(
                                                        patternIndex
                                                      );
                                                    }
                                                  }}
                                                >
                                                  <DeleteIcon />
                                                </IconButton>
                                              </div>
                                            )
                                          )}
                                          <IconButton
                                            onClick={() =>
                                              patternsFields.push("")
                                            }
                                          >
                                            <AddIcon></AddIcon>
                                          </IconButton>
                                        </div>
                                      )}
                                    </FieldArray>
                                  </Typography>
                                </CardContent>

                                <FieldArray
                                  name={`intents[${index}].responses`}
                                >
                                  {(responsesFields: {
                                    remove: (arg0: number) => void;
                                    push: (arg0: string) => void;
                                  }) => (
                                    <div style={{ marginLeft: "15px" }}>
                                      <h2>Posibles respuestas</h2>
                                      {values.intents[index].responses.map(
                                        (_, responseIndex) => (
                                          <div key={responseIndex}>
                                            <Field
                                              style={{ width: "80%" }}
                                              id={`intents[${index}].responses.[${responseIndex}]`}
                                              name={`intents[${index}].responses.[${responseIndex}]`}
                                            />
                                            <IconButton
                                              color="error"
                                              onClick={() => {
                                                if (
                                                  values.intents[index]
                                                    .responses.length > 1
                                                ) {
                                                  responsesFields.remove(
                                                    responseIndex
                                                  );
                                                }
                                              }}
                                            >
                                              <DeleteIcon />
                                            </IconButton>
                                          </div>
                                        )
                                      )}

                                      <IconButton
                                        onClick={() => responsesFields.push("")}
                                      >
                                        <AddIcon></AddIcon>
                                      </IconButton>
                                    </div>
                                  )}
                                </FieldArray>
                              </div>
                              <CardActions>
                                <Button
                                  onClick={() => intentsFields.remove(index)}
                                  size="small"
                                  startIcon={<DeleteIcon />}
                                  color="error"
                                >
                                  Remover intento
                                </Button>
                              </CardActions>
                            </Card>
                          </AccordionDetails>
                        </Accordion>
                      ))}

                      <Tooltip placement="right" title="Agregar intento">
                        <IconButton
                          onClick={() =>
                            intentsFields.push({
                              tag: "",
                              patterns: ["", "", "", "", ""],
                              responses: [""],
                            })
                          }
                        >
                          <AddIcon />
                        </IconButton>
                      </Tooltip>
                    </>
                  )}
                </FieldArray>
                </div>
              </div>
              {isLoading && <CircularProgress />}
              {!isLoading && (
                <Button
                  variant="contained"
                  type="submit"
                  endIcon={<SendIcon />}
                >
                  Entrenar chatbot
                </Button>
              )}
            </Form>
          )}
        </Formik>
      </div>
    </>
  );
};

export default TrainingChatbotForm;
