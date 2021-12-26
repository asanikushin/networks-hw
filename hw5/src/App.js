import React, {Component} from 'react';

import {
    Route,
    Switch,
    Redirect,
    withRouter
} from "react-router-dom"

import {connect} from 'react-redux'
import {ConnectedRouter} from 'connected-react-router'

import './App.scss';

import Home from './containers/Home/Home'
import Appointments from './containers/Appointments/Appointments'

class App extends Component {
    render() {
        const {history} = this.props

        return (
            <ConnectedRouter history={history}>
                <div className="App">
                    <Switch>
                        <Route path='/home' component={Home}/>
                        <Route path='/appointments' component={Appointments}/>
                        <Redirect from='/' to='/home'/>
                    </Switch>
                </div>
            </ConnectedRouter>
        );
    }
}

export default withRouter(connect()(App))