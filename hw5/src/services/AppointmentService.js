import BaseService from './BaseService'

export class AppointmentService extends BaseService {
    find (params) {
        return super.request({
            url: '/appointments',
            params
        })
    }

/*
    findById (appointmentId) {
        return super.request({
            url: `/appointments/${appointmentId}`
        })
    }

    count () {
        return super.request({
            url: '/appointments/count'
        })
    }
*/
}

export default new AppointmentService()