import * as mock from '../MockData'
import BaseController from './BaseController'

class DirectoryController extends BaseController {
    getPath () { 
        return '/directory'
    }

    getHandlers () {
        return [
            {
                path: '/appointment-statuses',
                handler: (vars, params) => { 
                    return mock.getAppointmentStatuses()
                }
            }
        ]
    }
}

export default new DirectoryController ()