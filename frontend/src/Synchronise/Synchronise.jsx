import React from 'react';
import {Alert, Button, Col, Form, FormGroup, Row, Table, Tooltip} from 'react-bootstrap';
import DatePicker from 'react-datepicker';
// import ReactLoading from 'react-loading';
// import prettyms from 'pretty-ms';
// import moment from 'moment';
import warning from '../assets/images/warning.svg';

import 'react-datepicker/dist/react-datepicker.min.css';
// import '../../assets/css/fix-reactstrap-custominput.css';

const Synchronise = (
  {
    entries, synchronising, loading, tasksSynchronised, dateFrom, dateTo,
    onChangeDate, onChangeTask, onSubmitLoadTasks, onSubmitSynchroniseTasks,
    hasRunningTask
  }) => (
  <>
    {tasksSynchronised && <Alert color="success">Tâches synchronisées !</Alert>}

    <p className="lead">Dates à synchroniser</p>

    <Form className="mb-3 text-center">
      <Row>
        <Col xs={6}>
          <FormGroup>
            <Form.Label htmlFor="dateFrom">Du</Form.Label><br />
            <DatePicker inline calendarStartDay={1} selected={dateFrom}
                        onChange={date => onChangeDate("dateFrom", date)} />
          </FormGroup>
        </Col>
        <Col xs={6}>
          <FormGroup>
            <Form.Label htmlFor="dateTo">Au</Form.Label><br />
            <DatePicker inline calendarStartDay={1} selected={dateTo}
                        onChange={date => onChangeDate("dateTo", date)} />
          </FormGroup>
        </Col>
      </Row>
      <FormGroup>
        <Col className="text-center mt-3">
          <Button color="primary" onClick={onSubmitLoadTasks} disabled={loading}>
            {loading && <ReactLoading className="d-inline-block mr-1" width={20} height={20} />}
            Charger les tâches
          </Button>
        </Col>
      </FormGroup>
    </Form>

    <p className="lead">Tâches</p>

    {hasRunningTask && <Alert color="info">Une tâche est en cours !</Alert>}

    <Form onSubmit={onSubmitSynchroniseTasks} onChange={onChangeTask}>
      <Table striped size="sm">
        <thead>
          <tr>
            <th scope="col" colSpan={2} className="text-center">Tâche</th>
            <th scope="col">Date</th>
            <th scope="col" className="text-center">Durée</th>
            <th scope="col">Commentaire</th>
            <th scope="col">Synchronisée ?</th>
          </tr>
        </thead>
        <tbody>
          {!entries.length &&
            <tr>
              <td colSpan={6} className="text-center">Aucune tâche</td>
            </tr>
          }
          {entries.map((entry, index) => {
            const {
              description, duration, decimalDuration, pastDecimalDuration,
              date, isValid, comment, sync
            } = entry;
            const rowId = `row-${index}`;
            const commentId = `comment-${index}`;
            const syncId = `sync-${index}`;
            const isMuted = !isValid || 0 === decimalDuration;

            return (
              <tr key={index} id={rowId} className={isMuted ? 'text-muted' : ''}>
                <th scope="row">
                  {pastDecimalDuration > 0 &&
                    <>
                      <img src={warning} height={20} id={`warning-${rowId}`} alt="Attention" />
                      <Tooltip placement="right" target={`warning-${rowId}`}>
                        Il y a déjà {pastDecimalDuration} heure(s) enregistrées pour cette tâche.
                      </Tooltip>
                    </>
                  }
                </th>
                <td>{description}</td>
                <td>{moment(date).format('DD/MM/YYYY')}</td>
                <td className="text-center">
                  {prettyms(duration * 1000)}<br />
                  <div className="text-muted small">({decimalDuration}h)</div>
                </td>
                <td>
                  {0 === decimalDuration && <span className="small">Aucun temps !</span>}
                  {!isMuted &&
                    <Form.Control
                      type="text" name="comment" id={commentId} defaultValue={comment}
                      data-id={index} bsSize="sm"
                    />
                  }
                  {!isValid &&
                    <Tooltip placement="top" target={rowId}>
                      Description invalide ! Elle doit commencer par <code>#</code> puis être suivie uniquement de
                      chiffres !
                    </Tooltip>
                  }
                </td>
                <td>
                  {!isMuted &&
                    <Form.Check
                      type="checkbox" name="sync" id={syncId} className="text-center"
                      defaultChecked={sync} data-id={index}
                    />
                  }
                </td>
              </tr>
            );
          })}
        </tbody>
      </Table>

      {!!entries.length &&
        <Row className="mb-3 text-center">
          <Col>
            <Button type="submit" color="primary" disabled={synchronising}>
              {synchronising && <ReactLoading className="d-inline-block mr-1" width={20} height={20} />}
              Synchroniser vers Redmine
            </Button>
          </Col>
        </Row>
      }
    </Form>
  </>
);

export default Synchronise;
