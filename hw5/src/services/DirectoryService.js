import BaseService from './BaseService'

export class DirectoryService extends BaseService {
    findAppointmentStatuses () {
        return super.request({
            url: '/directory/appointment-statuses'
        })
    }
}

export default new DirectoryService()