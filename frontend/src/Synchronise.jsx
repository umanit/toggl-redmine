import React, {useState} from 'react';
import {Alert, Button, Col, Form, FormGroup, OverlayTrigger, Row, Spinner, Table, Tooltip} from 'react-bootstrap';
import DatePicker, {registerLocale} from 'react-datepicker';
import {fr} from 'date-fns/locale/fr';
import {format, parseJSON} from 'date-fns';
import warning from './assets/images/warning.svg';
import {LoadTasks} from '../wailsjs/go/main/App.js';

import 'react-datepicker/dist/react-datepicker.min.css';

registerLocale('fr', fr);

const today = new Date();
const formattedDate = (jsonDate) => format(parseJSON(jsonDate), 'dd/MM/yyyy');
const forHumans = (seconds) => {
  const units = [
    {label: 'j', value: 86400},
    {label: 'h', value: 3600},
    {label: 'm', value: 60},
    {label: 's', value: 1}
  ];

  let result = [];

  for (const unit of units) {
    const quotient = Math.floor(seconds / unit.value);
    if (quotient > 0) {
      result.push(`${quotient}${unit.label}`);
      seconds -= quotient * unit.value;
    }
  }

  return result.join(' ');
}

export default function Synchronize() {
  const [tasksSynchronised, setTaskSynchronized] = useState(false);
  const [dateFrom, setDateFrom] = useState(today);
  const [dateTo, setDateTo] = useState(today);
  const [taskLoading, setTaskLoading] = useState(false);
  const [hasRunningTask, setHasRunningTask] = useState(false);
  const [entries, setEntries] = useState([]);
  const [synchronising, setSynchronising] = useState(false);

  function loadTasks() {
    setTaskLoading(true);

    LoadTasks(dateFrom, dateTo).then(({Entries, HasRunningTask}) => {
      setEntries([...Object.values(Entries)]);
      setHasRunningTask(HasRunningTask);
      setTaskLoading(false);
    });
  }

  return (
    <>
      {tasksSynchronised && <Alert color="success">Tâches synchronisées !</Alert>}

      <p className="lead">Dates à synchroniser</p>

      <Form className="mb-3 text-center">
        <Row>
          <Col xs={6}>
            <FormGroup>
              <Form.Label htmlFor="dateFrom">Du</Form.Label><br />
              <DatePicker inline selected={dateFrom} locale="fr" onChange={date => setDateFrom(date)} />
            </FormGroup>
          </Col>
          <Col xs={6}>
            <FormGroup>
              <Form.Label htmlFor="dateTo">Au</Form.Label><br />
              <DatePicker inline selected={dateTo} locale="fr" onChange={date => setDateTo(date)} />
            </FormGroup>
          </Col>
        </Row>
        <FormGroup>
          <Col className="text-center mt-3">
            <Button color="primary" onClick={loadTasks} disabled={taskLoading}>
              {taskLoading && <Spinner as="span" animation="grow" size="sm" role="status" aria-hidden="true" />}
              Charger les tâches
            </Button>
          </Col>
        </FormGroup>
      </Form>

      <p className="lead">Tâches</p>

      {hasRunningTask && <Alert color="info" onClose={() => setHasRunningTask(false)} dismissible>
        Attention, une tâche est en cours !
      </Alert>}

      <Form onSubmit={() => console.warn('@todo')}>
        <Table striped size="sm">
          <thead>
            <tr>
              <th scope="col" colSpan={2} className="text-center">Tâche</th>
              <th scope="col">Date</th>
              <th scope="col" className="text-center">Durée</th>
              <th scope="col">Commentaire</th>
              <th scope="col"><abbr title="Synchroniser ?">Sync. ?</abbr></th>
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
                Description, Duration, DecimalDuration, PastDecimalDuration,
                Date, IsValid, Comment, Sync
              } = entry;
              const rowId = `row-${index}`;
              const commentId = `comment-${index}`;
              const syncId = `sync-${index}`;
              const isMuted = !IsValid || 0 === DecimalDuration;

              return (
                <tr key={index} id={rowId} className={isMuted ? 'text-muted' : ''}>
                  <th scope="row">
                    {PastDecimalDuration > 0 &&
                      <OverlayTrigger overlay={
                        <Tooltip id={rowId}>
                          Il y a déjà {PastDecimalDuration} heure(s) enregistrées pour cette tâche.
                        </Tooltip>
                      } placement="right">
                        <img src={warning} height={20} id={`warning-${rowId}`} alt="Attention" />
                      </OverlayTrigger>
                    }
                  </th>
                  <td>{Description}</td>
                  <td>{formattedDate(Date)}</td>
                  <td className="text-center">
                    {forHumans(Duration)}<br />
                    <div className="text-muted small">({DecimalDuration}h)</div>
                  </td>
                  <td>
                    {0 === DecimalDuration && <span className="small">Aucun temps !</span>}
                    {!isMuted &&
                      <Form.Control type="text" name="comment" id={commentId} defaultValue={Comment}
                                    data-id={index} size="sm" />
                    }
                    {!IsValid &&
                      <Tooltip placement="top" target={rowId}>
                        Description invalide ! Elle doit commencer par <code>#</code> puis être suivie uniquement de
                        chiffres !
                      </Tooltip>
                    }
                  </td>
                  <td>
                    {!isMuted &&
                      <Form.Check type="checkbox" name="sync" id={syncId} className="text-center"
                                  defaultChecked={Sync} data-id={index} />
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
                {synchronising && <Spinner as="span" animation="grow" size="sm" role="status" aria-hidden="true" />}
                Synchroniser vers Redmine
              </Button>
            </Col>
          </Row>
        }
      </Form>
    </>
  );
}
