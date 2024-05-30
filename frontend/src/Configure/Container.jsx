import React, { Component } from 'react';
// import { ipcRenderer } from 'electron';
import Configure from './Configure';

class ConfigureContainer extends Component {
    constructor(props) {
        super(props);
        this.state = {
            showAlert: false,
            config: {
                toggl: { token: null, url: null },
                redmine: { token: null, url: null }
            },
            validation: {
                toggl: null,
                redmine: null,
            }
        };
        this.timeouts = [];
        this.listeners = [];
    }

    componentDidMount() {
        // ipcRenderer.send('ask-config');

        this.listeners.push('receive-config');
        // ipcRenderer.on('receive-config', (event, config) => {
        //     this.setState({ config });
        // });

        this.listeners.push('config-tested');
        // ipcRenderer.on('config-tested', (event, payload) => {
        //     this.setState({ validation: payload });
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

    onChange = (event, type) => {
        const { config } = this.state;
        const { target: { name, value } } = event;

        this.setState({ config: { ...config, [type]: { ...config[type], [name]: value } } });
    };

    onSubmit = () => {
        const { config } = this.state;

        ipcRenderer.send('save-config', config);

        this.setState({ showAlert: true });

        this.timeouts.push(
            setTimeout(() => {
                this.setState({ showAlert: false });
            }, 4000)
        );
    };

    testCredentials = () => {
        const { config } = this.state;

        ipcRenderer.send('test-config', { config });
    };

    render() {
        return (
            <Configure
                {...this.state}
                onChange={this.onChange}
                onSubmit={this.onSubmit}
                testCredentials={this.testCredentials}
            />
        );
    }
}

export default ConfigureContainer;
