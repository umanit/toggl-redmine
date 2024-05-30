import React, { Component } from 'react';
// import { ipcRenderer } from 'electron';
// import moment from 'moment';
import Synchronise from './Synchronise';

class SynchroniseContainer extends Component {
    constructor(props) {
        super(props);
        this.state = {
            entries: [],
            dateFrom: null, //moment(),
            dateTo: null, //moment(),
            tasksSynchronised: false,
            loading: false,
            synchronising: false,
            hasRunningTask: false
        };
        this.timeouts = [];
        this.listeners = [];
    }

    componentDidMount() {
        this.listeners.push('receive-toggl-entries');
        // ipcRenderer.on('receive-toggl-entries', (event, entries) => {
        //     this.setState({ entries, loading: false });
        // });

        this.listeners.push('receive-toggl-has-running-task');
        // ipcRenderer.on('receive-toggl-has-running-task', (event, hasRunningTask) => {
        //     this.setState({ hasRunningTask });
        // });

        this.listeners.push('tasks-synchronised');
        // ipcRenderer.on('tasks-synchronised', () => {
        //     window.scrollTo(0, 0);
        //
        //     this.setState({
        //         synchronising: false,
        //         tasksSynchronised: true,
        //         entries: []
        //     });
        //
        //     this.timeouts.push(
        //         setTimeout(() => {
        //             this.setState({ tasksSynchronised: false });
        //         }, 4000)
        //     );
        // });
    }

    componentWillUnmount() {
        for (const listener of this.listeners) {
            // ipcRenderer.removeAllListeners(listener);
        }

        for (const timeout of this.timeouts) {
            clearTimeout(timeout);
        }
    }

    onChangeDate = (name, value) => {
        this.setState({ [name]: value });
    };

    onSubmitLoadTasks = () => {
        const { dateFrom, dateTo } = this.state;

        dateFrom.startOf('day');
        dateTo.endOf('day');

        // ipcRenderer.send('load-tasks', {
        //     from: dateFrom.toISOString(),
        //     to: dateTo.toISOString()
        // });
        this.setState({ entries: [], loading: true });
    };

    onChangeTask = event => {
        const target = event.target;
        let entries = [...this.state.entries];
        entries[target.dataset.id][target.name] = 'checkbox' === target.type ? target.checked : target.value;

        this.setState({ entries });
    };

    onSubmitSynchroniseTasks = event => {
        event.preventDefault();

        const { entries } = this.state;

        // ipcRenderer.send('synchronise-tasks', entries);
        this.setState({ synchronising: true });
    };

    render() {
        return (
            <Synchronise
                {...this.state} onChangeDate={this.onChangeDate} onSubmitLoadTasks={this.onSubmitLoadTasks}
                onChangeTask={this.onChangeTask} onSubmitSynchroniseTasks={this.onSubmitSynchroniseTasks}
            />
        );
    }
}

export default SynchroniseContainer;
