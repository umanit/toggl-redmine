import {useState} from 'react';
import {Link, Route} from 'react-router-dom';
import {Breadcrumb, BreadcrumbItem, Col, Container, Row} from 'react-bootstrap';
import banner from './assets/images/banner.png';
import {Greet} from '../wailsjs/go/main/App';

function App() {
  const [resultText, setResultText] = useState("Please enter your name below 👇");
  const [name, setName] = useState('');
  const updateName = (e) => setName(e.target.value);
  const updateResultText = (result) => setResultText(result);

  function greet() {
    Greet(name).then(updateResultText);
  }

  return (
    <Route render={() => (
      <Container id="App">
        <Row>
          <Col>
            <Link to="/">
              <img src={banner} alt="toggl track - Redmine bridge" className="mt-4 mb-5" />
            </Link>
          </Col>
        </Row>
        <Row>
          <Col>
            <Breadcrumb>
              <BreadcrumbItem>Accueil</BreadcrumbItem>
            </Breadcrumb>
          </Col>
        </Row>
        <Row>
          <Col>

          </Col>
        </Row>
      </Container>
    )} />
  )
}

export default App
