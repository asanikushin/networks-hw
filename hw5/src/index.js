import React from 'react'
import ReactDOM from 'react-dom'

import {Provider} from 'react-redux'
import thunkMiddleware from 'redux-thunk'
import {applyMiddleware, compose, createStore} from 'redux'
import {routerMiddleware, ConnectedRouter} from 'connected-react-router'

import {createBrowserHistory} from 'history'

import './index.scss'

import App from './App'

import rootReducer from './redux/rootReducer'

import DirectoryInitialState from './redux/directory/DirectoryInitialState'
import AppointmentInitialState from './redux/appointment/AppointmentInitialState'

function getInitialState () {
    return {
        directory: DirectoryInitialState(),
        appointment: AppointmentInitialState()
    }
}

// создаём кастомную историю
const history = createBrowserHistory()

const store = createStore(
    rootReducer(history),
    getInitialState(),
    compose(applyMiddleware(routerMiddleware(history), thunkMiddleware))
)

ReactDOM.render((
    <Provider store={store}>
        <ConnectedRouter history={history}>
            <App history={history} />
        </ConnectedRouter>
    </Provider>
  ), document.getElementById('root')
);
