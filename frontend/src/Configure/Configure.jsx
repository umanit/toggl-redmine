import React from 'react';
import {Alert, Button, ButtonGroup, Col, Container, Form, FormGroup, FormText, Row} from 'react-bootstrap';

const Configure = ({config, onChange, onSubmit, testCredentials, showAlert, validation}) => {
  const redmineValid = true === validation.redmine ? {valid: true} : (false === validation.redmine ? {invalid: true} : {});
  const togglValid = true === validation.toggl ? {valid: true} : (false === validation.toggl ? {invalid: true} : {});

  return (
    <>
      {showAlert && <Alert color="success">Configuration saved!</Alert>}

      <p className="lead">
        Veuillez renseigner les clés et les URLs d’API de toggl track et Redmine afin de pour pouvoir utiliser la
        synchronisation.
      </p>

      <Form>
        <FormGroup as="fieldset">
          <legend>Configuration de toggl track</legend>
          <Container>
            <Row>
              <Col xs="6">
                <FormGroup>
                  <Form.Label htmlFor="toggl-api-token">Clé d’API</Form.Label>
                  <Form.Control
                    type="text"
                    id="toggl-api-token"
                    name="token"
                    defaultValue={config.toggl.token}
                    onChange={e => onChange(e, 'toggl')}
                    {...togglValid}
                  />
                  {true === validation.toggl &&
                    <Form.Control.Feedback valid>Connexion à l’API réussie !</Form.Control.Feedback>
                  }
                  {false === validation.toggl &&
                    <Form.Control.Feedback>
                      Impossible de contacter l’API. La clé d’API est-elle correcte ?
                    </Form.Control.Feedback>
                  }
                  <FormText color="muted">
                    Vous pouvez trouver votre clé d’API dans vos paramètres de compte sur le site de toggle track.
                  </FormText>
                </FormGroup>
              </Col>
              <Col xs="6">
                <FormGroup>
                  <Form.Label htmlFor="toggl-api-url">URL de l’API</Form.Label>
                  <Form.Control
                    type="text"
                    id="toggl-api-url"
                    name="url"
                    defaultValue={config.toggl.url}
                    onChange={e => onChange(e, 'toggl')}
                  />
                </FormGroup>
              </Col>
            </Row>
          </Container>
        </FormGroup>
        <FormGroup as="fieldset" className="mt-3">
          <legend>Configuration de Redmine</legend>
          <Container>
            <Row>
              <Col xs="6">
                <Form.Label htmlFor="redmine-api-token">Clé d’API</Form.Label>
                <Form.Control
                  type="text"
                  id="redmine-api-token"
                  name="token"
                  defaultValue={config.redmine.token}
                  onChange={e => onChange(e, 'redmine')}
                  {...redmineValid}
                />
                {true === validation.toggl &&
                  <Form.Control.Feedback valid>Connexion à l’API réussie !</Form.Control.Feedback>
                }
                {false === validation.toggl &&
                  <Form.Control.Feedback>
                    Impossible de contacter l’API. La clé d’API est-elle correcte ?
                  </Form.Control.Feedback>
                }
                <FormText color="muted">
                  Vous trouverez votre clé d’API sur la page de votre compte (<code>/my/account</code>) sous
                  l’intitulé <code>Clé d'accès API</code>.
                </FormText>
              </Col>
              <Col xs="6">
                <FormGroup>
                  <Form.Label htmlFor="redmine-api-url">URL de l’API</Form.Label>
                  <Form.Control
                    type="text"
                    id="redmine-api-url"
                    name="url"
                    defaultValue={config.redmine.url}
                    onChange={e => onChange(e, 'redmine')}
                  />
                </FormGroup>
              </Col>
            </Row>
          </Container>
        </FormGroup>
        <ButtonGroup className="mt-3">
          <Button color="primary" onClick={() => testCredentials('redmine')}>Tester les identifiants</Button>
          <Button variant="secondary" onClick={onSubmit}>Enregistrer</Button>
        </ButtonGroup>
      </Form>
    </>
  );
};

export default Configure;
