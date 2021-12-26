import Status from './status/AppointmentStatusInitialState'

const { Record } = require('immutable')

export default Record({
    status: Status()
})