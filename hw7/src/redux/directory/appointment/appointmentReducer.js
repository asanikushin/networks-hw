import InitialState from './AppointmentInitialState'

import statusReducer from './status/appointmentStatusReducer'

const initialState = new InitialState()

export default function appointmentReducer (state = initialState, action) {
    let nextState = state

    const status = statusReducer (state.status, action)
    if (status !== state.status) nextState = nextState.setIn(['status'], status)

    return nextState
}
