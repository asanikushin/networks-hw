import InitialState from './DirectoryInitialState'

import appointmentReducer from './appointment/appointmentReducer'

const initialState = new InitialState()

export default function dictionaryReducer(state = initialState, action) {
    let nextState = state

    const appointment = appointmentReducer(state.appointment, action)
    if (appointment !== state.appointment) nextState = nextState.setIn(['appointment'], appointment)

    return nextState
}
