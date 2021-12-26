import Appointment from './appointment/AppointmentInitialState'

const { Record } = require('immutable')

export default Record({
    appointment: Appointment()
})