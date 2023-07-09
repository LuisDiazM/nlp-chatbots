import React from "react";
import { Formik, Field, FieldArray, Form } from "formik";
import Button from "@mui/material/Button";
import Accordion from "@mui/material/Accordion";
import AccordionSummary from "@mui/material/AccordionSummary";
import AccordionDetails from "@mui/material/AccordionDetails";
import Typography from "@mui/material/Typography";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";

import {
  IntentModel,
  TrainingChatbotModel,
} from "../models/trainingChatbotModel/trainingChat.model";

const initialValues: TrainingChatbotModel = {
  user_id: "",
  intents: [{ patterns: [""], responses: [""], tag: "" }],
};

const TrainingChatbotForm: React.FC = () => {
  const handleSubmit = (values: TrainingChatbotModel) => {
    console.log(values);
  };

  return (
    <Formik initialValues={initialValues} onSubmit={handleSubmit}>
      {({ values }) => (
        <Form>
          <label htmlFor="user_id">Email</label>
          <Field name="user_id" type="email"></Field>
          <FieldArray name="intents">
            {(intentsFields: {
              remove: (arg0: number) => void;
              push: (arg0: IntentModel) => void;
            }) => (
              <Accordion>
                <div>
                  <AccordionSummary
                    expandIcon={<ExpandMoreIcon />}
                    aria-controls="panel1a-content"
                    id="panel1a-header"
                  >
                    <Typography>Accordion 1</Typography>
                    <AccordionDetails>
                      {values.intents.map((_, index) => (
                        <div key={index}>
                          <label htmlFor={`intents[${index}].tag`}>
                            Nombre intento
                          </label>
                          <Field
                            id={`intents[${index}].tag`}
                            name={`intents[${index}].tag`}
                          />

                          <FieldArray name={`intents[${index}].patterns`}>
                            {(patternsFields: {
                              remove: (arg0: number) => void;
                              push: (arg0: string) => void;
                            }) => (
                              <div>
                                {values.intents[index].patterns.map(
                                  (_, patternIndex) => (
                                    <div key={patternIndex}>
                                      <Field
                                        id={`intents[${index}].patterns.[${patternIndex}]`}
                                        name={`intents[${index}].patterns.[${patternIndex}]`}
                                      />
                                      <button
                                        type="button"
                                        onClick={() =>
                                          patternsFields.remove(patternIndex)
                                        }
                                      >
                                        Remover entrada
                                      </button>
                                    </div>
                                  )
                                )}
                                <button
                                  type="button"
                                  onClick={() => patternsFields.push("")}
                                >
                                  Agregar entrada
                                </button>
                              </div>
                            )}
                          </FieldArray>

                          <FieldArray name={`intents[${index}].responses`}>
                            {(responsesFields: {
                              remove: (arg0: number) => void;
                              push: (arg0: string) => void;
                            }) => (
                              <div>
                                {values.intents[index].responses.map(
                                  (_, responseIndex) => (
                                    <div key={responseIndex}>
                                      <Field
                                        id={`intents[${index}].responses.[${responseIndex}]`}
                                        name={`intents[${index}].responses.[${responseIndex}]`}
                                      />
                                      <button
                                        type="button"
                                        onClick={() =>
                                          responsesFields.remove(responseIndex)
                                        }
                                      >
                                        Remover respuesta
                                      </button>
                                    </div>
                                  )
                                )}
                                <button
                                  type="button"
                                  onClick={() => responsesFields.push("")}
                                >
                                  Agregar respuesta
                                </button>
                              </div>
                            )}
                          </FieldArray>

                          <button
                            type="button"
                            onClick={() => intentsFields.remove(index)}
                          >
                            Remover intento
                          </button>
                        </div>
                      ))}
                      <button
                        type="button"
                        onClick={() =>
                          intentsFields.push({
                            tag: "",
                            patterns: [""],
                            responses: [""],
                          })
                        }
                      >
                        Agregar intento
                      </button>
                    </AccordionDetails>
                  </AccordionSummary>
                </div>
              </Accordion>
            )}
          </FieldArray>

          <Button variant="contained" type="submit">
            Entrenar chatbot
          </Button>
        </Form>
      )}
    </Formik>
  );
};

export default TrainingChatbotForm;
