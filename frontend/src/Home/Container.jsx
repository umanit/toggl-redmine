import React, {Component} from 'react';
// import {ipcRenderer} from 'electron';
import Home from './Home';

class HomeContainer extends Component {
    constructor(props) {
        super(props);
        this.state = {
            canSynchronise: false
        };
        this.listeners = [];
    }

    componentDidMount() {
        // ipcRenderer.send('ask-can-synchronise');

        this.listeners.push('receive-can-synchronise');
        // ipcRenderer.on('receive-can-synchronise', (event, canSynchronise) => {
        //     this.setState(canSynchronise);
        // });
    }

    componentWillUnmount() {
        for (const listener of this.listeners) {
            // ipcRenderer.removeAllListeners(listener);
        }
    }

    render() {
        return (
            <Home {...this.state} />
        );
    }
}

export default HomeContainer;
