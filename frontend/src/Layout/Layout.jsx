import React from 'react';
import {Link, Outlet, useLocation} from 'react-router-dom';
import {Breadcrumb, BreadcrumbItem, Col, Container, Row} from 'react-bootstrap';
import {BreadcrumbItemLink} from '../Links/BreadcrumbItemLink.jsx';
import banner from '../assets/images/banner.png';

const Layout = () => {
  const {pathname} = useLocation();

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
          <Outlet />
        </Col>
      </Row>
    </Container>
  )
};

export default Layout
