import Factory from '../ActionFactory'

import * as actions from '../../redux/directory/appointment/status/list/appointmentStatusListActions'

export default Factory(actions, {
  // здесь мы сами устанавливаем, какой метод
  // подставить в действие action: (params, actions) => actions.someMethod(params)
  action: (params, actions) => actions.load()
})