import React, {useEffect, useState} from 'react';
import {Card, Col, Row} from 'react-bootstrap';
import {ButtonLink} from './Links/ButtonLink.jsx';
import {CanSynchronize} from '../wailsjs/go/main/App.js';

export default function Home() {
  const [canSynchronize, setCanSynchronize] = useState(false);

  useEffect(() => {
    CanSynchronize().then((result) => setCanSynchronize(result));
  }, []);

  return (
    <Row>
      <Col>
        <Card body border="primary">
          <Card.Title>Synchroniser</Card.Title>
          <Card.Text>Synchronisez vos entrées de temps de toggl track avec Redmine.</Card.Text>
          <div className="d-grid">
            <ButtonLink variant="secondary" to="/synchroniser" disabled={!canSynchronize}>Synchroniser</ButtonLink>
          </div>
        </Card>
      </Col>
      <Col>
        <Card body border="secondary">
          <Card.Title>Configurer</Card.Title>
          <Card.Text>Configurer les clés et URLs d’API de toggl track et Redmine.</Card.Text>
          <div className="d-grid">
            <ButtonLink variant="secondary" to="/configurer">Configurer</ButtonLink>
          </div>
        </Card>
      </Col>
    </Row>
  );
}
