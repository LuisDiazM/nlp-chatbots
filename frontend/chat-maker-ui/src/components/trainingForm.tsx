import React from "react";
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

import {
  IntentModel,
  TrainingChatbotModel,
} from "../models/trainingChatbotModel/trainingChat.model";
import CardActions from "@mui/material/CardActions";
import Tooltip from "@mui/material/Tooltip";

const inputStyles = { marginRight: "10px" };

const initialValues: TrainingChatbotModel = {
  user_id: "",
  intents: [{ patterns: [""], responses: [""], tag: "" }],
};

const TrainingChatbotForm: React.FC = () => {
  const handleSubmit = (values: TrainingChatbotModel) => {
    console.log(values);
  };

  return (
    <>
      <div style={{ margin: "auto", width: "50%" }}>
        <Formik initialValues={initialValues} onSubmit={handleSubmit}>
          {({ values }) => (
            <Form>
              {/* <label htmlFor="user_id" style={inputStyles}>
                Email
              </label>
              <Field name="user_id" type="email"></Field> */}

              <FieldArray name="intents">
                {(intentsFields: {
                  remove: (arg0: number) => void;
                  push: (arg0: IntentModel) => void;
                }) => (
                  <div>
                    {values.intents.map((intentName, index) => (
                      <Accordion>
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
                                      <div>
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
                                                      .patterns.length > 1
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

                              <FieldArray name={`intents[${index}].responses`}>
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
                                                values.intents[index].responses
                                                  .length > 1
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
                            patterns: [""],
                            responses: [""],
                          })
                        }
                      >
                        <AddIcon />
                      </IconButton>
                    </Tooltip>
                  </div>
                )}
              </FieldArray>

              <Button variant="contained" type="submit" endIcon={<SendIcon />}>
                Entrenar chatbot
              </Button>
            </Form>
          )}
        </Formik>
      </div>
    </>
  );
};

export default TrainingChatbotForm;
