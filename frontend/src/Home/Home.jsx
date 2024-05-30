import React from 'react';
import {Card, CardText, CardTitle, Col, Row} from 'react-bootstrap';
import {ButtonLink} from '../Links/ButtonLink.jsx';

const Home = ({canSynchronise}) => (
  <Row>
    <Col>
      <Card body border="primary">
        <CardTitle>Synchroniser</CardTitle>
        <CardText>Synchronisez vos entrées de temps de toggl track avec Redmine.</CardText>
        <div className="d-grid">
          <ButtonLink variant="secondary" to="/synchroniser" disabled={canSynchronise}>Synchroniser</ButtonLink>
        </div>
      </Card>
    </Col>
    <Col>
      <Card body border="secondary">
        <CardTitle>Configurer</CardTitle>
        <CardText>Configurer les clés et URLs d’API de toggl track et Redmine.</CardText>
        <div className="d-grid">
          <ButtonLink variant="secondary" to="/configurer">Configurer</ButtonLink>
        </div>
      </Card>
    </Col>
  </Row>
);

export default Home;
