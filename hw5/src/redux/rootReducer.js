import { combineReducers } from 'redux'
import { connectRouter } from 'connected-react-router'

import directory from './directory/directoryReducer'
import appointment from './appointment/appointmentReducer'

const rootReducer = (history) => combineReducers({
    router: connectRouter(history),
    directory,
    appointment
})

export default rootReducer