import React, {useEffect, useState} from 'react';
import {Link, Outlet, useLocation} from 'react-router-dom';
import {Alert, Breadcrumb, BreadcrumbItem, Col, Container, Row} from 'react-bootstrap';
import {BreadcrumbItemLink} from './Links/BreadcrumbItemLink.jsx';
import banner from './assets/images/banner.png';
import {EventsOff, EventsOn} from "../wailsjs/runtime/runtime.js";
import {GetLogsPath} from "../wailsjs/go/main/App.js";

const Layout = () => {
  const [errorOccured, setErrorOccured] = useState(false);
  const [logsPath, setLogsPath] = useState('');
  const {pathname} = useLocation();

  useEffect(() => {
    EventsOn("goError", () => setErrorOccured(true));

    return () => EventsOff("goError");
  }, []);

  useEffect(() => {
    GetLogsPath().then(setLogsPath);
  }, []);

  return (
    <Container>
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
            {"/" === pathname
              ? <BreadcrumbItem active>Accueil</BreadcrumbItem>
              : <BreadcrumbItemLink to="/">Accueil</BreadcrumbItemLink>
            }
            {"/configurer" === pathname
              ? <BreadcrumbItem active>Configurer</BreadcrumbItem>
              : ""
            }
            {"/synchroniser" === pathname
              ? <BreadcrumbItem active>Synchroniser</BreadcrumbItem>
              : ""
            }
          </Breadcrumb>
        </Col>
      </Row>
      <Row>
        <Col>
          {errorOccured && <Alert variant="danger">
            Une erreur est survenue ! Merci d’aller voir les logs dans <code>{logsPath}</code>.
          </Alert>}
          <Outlet />
        </Col>
      </Row>
    </Container>
  )
};

export default Layout
