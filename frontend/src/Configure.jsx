import React, {useEffect, useState} from 'react';
import {Alert, Button, ButtonGroup, Col, Container, Form, Row, Spinner} from 'react-bootstrap';
import {LoadConfig, SaveConfig, TestCredentials} from '../wailsjs/go/main/App.js';

export default function Configure() {
  const [showAlert, setShowAlert] = useState(false);
  const [config, setConfig] = useState({});
  const [formValidated, setFormValidated] = useState(false);
  const [canTestCredentials, setCanTestCredentials] = useState(false);
  const [testingCredentials, setTestingCredentials] = useState(false);
  const [togglTrackOk, setToggleTrackOk] = useState(null);
  const [redmineOk, setRedmineOk] = useState(null);

  const loadConfig = () => {
    LoadConfig().then(({Config, IsValid}) => {
      setConfig(Config);
      setCanTestCredentials(IsValid);
    });
  }
  const handleSubmit = (event) => {
    event.preventDefault();
    event.stopPropagation();

    const form = event.currentTarget;

    if (form.checkValidity()) {
      const formData = new FormData(form);
      const config = {
        toggl: {
          key: formData.get('toggl[key]'),
          url: formData.get('toggl[url]')
        },
        redmine: {
          key: formData.get('redmine[key]'),
          url: formData.get('redmine[url]')
        }
      };

      SaveConfig(config).then(() => {
        setShowAlert(true);

        loadConfig();
      });
    }

    setFormValidated(true);
  };
  const testCredentials = () => {
    setTestingCredentials(true);

    TestCredentials().then((result) => {
      setToggleTrackOk(result.TogglTrackOk ?? false);
      setRedmineOk(result.RedmineOk ?? false);

      setTestingCredentials(false);
    });
  };

  useEffect(() => loadConfig(), []);

  useEffect(() => {
    let timer;
    if (showAlert) {
      timer = setTimeout(() => {
        setShowAlert(false);
      }, 4000);
    }
    return () => clearTimeout(timer);
  }, [showAlert]);

  return (
    <>
      {showAlert && <Alert variant="success">Configuration enregistrée !</Alert>}

      <p className="lead">
        Veuillez renseigner les clés et les URLs d’API de toggl track et Redmine afin de pour pouvoir utiliser la
        synchronisation.
      </p>

      <Form noValidate validated={formValidated} onSubmit={handleSubmit}>
        <Form.Group as="fieldset">
          <legend>Configuration de toggl track</legend>
          <Container>
            <Row>
              {true === togglTrackOk &&
                <Col xs={12}>
                  <p className="text-success">Connexion à l’API réussie !</p>
                </Col>
              }
              {false === togglTrackOk &&
                <Col xs={12}>
                  <p className="text-danger">Échec de connexion à l’API, veuillez vérifier vos identifiants.</p>
                </Col>
              }
              <Col xs="6">
                <Form.Group>
                  <Form.Label htmlFor="toggl-api-key">Clé d’API</Form.Label>
                  <Form.Control required type="text" id="toggl-api-key" name="toggl[key]" defaultValue={config.toggl?.key} />
                  <Form.Control.Feedback type="invalid">
                    Merci de renseigner la clé d’API à toggl track.
                  </Form.Control.Feedback>
                  <Form.Text>
                    Vous pouvez trouver votre clé d’API dans vos paramètres de compte sur le site de toggle track.
                  </Form.Text>
                </Form.Group>
              </Col>
              <Col xs="6">
                <Form.Group>
                  <Form.Label htmlFor="toggl-api-url">URL de l’API</Form.Label>
                  <Form.Control required type="url" pattern="https://.*" id="toggl-api-url" name="toggl[url]" defaultValue={config.toggl?.url} />
                  <Form.Control.Feedback type="invalid">
                    Merci de renseigner l’URL de l’API toggl track. Elle doit commencer par <code>https://</code>.
                  </Form.Control.Feedback>
                </Form.Group>
              </Col>
            </Row>
          </Container>
        </Form.Group>

        <Form.Group as="fieldset" className="mt-3">
          <legend>Configuration de Redmine</legend>
          <Container>
            <Row>
              {true === redmineOk &&
                <Col xs={12}>
                  <p className="text-success">Connexion à l’API réussie !</p>
                </Col>
              }
              {false === redmineOk &&
                <Col xs={12}>
                  <p className="text-danger">Échec de connexion à l’API, veuillez vérifier vos identifiants.</p>
                </Col>
              }
              <Col xs="6">
                <Form.Label htmlFor="redmine-api-key">Clé d’API</Form.Label>
                <Form.Control required type="text" id="redmine-api-key" name="redmine[key]" defaultValue={config.redmine?.key} />
                <Form.Control.Feedback type="invalid">
                  Merci de renseigner la clé d’API à Redmine.
                </Form.Control.Feedback>
                <Form.Text>
                  Vous trouverez votre clé d’API sur la page de votre compte (<code>/my/account</code>) sous
                  l’intitulé <code>Clé d'accès API</code>.
                </Form.Text>
              </Col>
              <Col xs="6">
                <Form.Group>
                  <Form.Label htmlFor="redmine-api-url">URL de l’API</Form.Label>
                  <Form.Control required type="url" pattern="https://.*" id="redmine-api-url" name="redmine[url]" defaultValue={config.redmine?.url} />
                  <Form.Control.Feedback type="invalid">
                    Merci de renseigner l’URL de l’API Redmine. Elle doit commencer par <code>https://</code>.
                  </Form.Control.Feedback>
                </Form.Group>
              </Col>
            </Row>
          </Container>
        </Form.Group>
        <ButtonGroup className="mt-3">
          <Button disabled={!canTestCredentials || testingCredentials} onClick={() => testCredentials()}>
            {testingCredentials &&
              <Spinner as="span" animation="grow" size="sm" role="status" aria-hidden="true" />
            }
            Tester les identifiants
          </Button>
          <Button variant="secondary" type="submit">Enregistrer</Button>
        </ButtonGroup>
      </Form>
    </>
  );
}
