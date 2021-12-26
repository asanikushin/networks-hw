import { filter, isNumber } from 'underscore'

const USER = 'Иванов Иван Иванович'

const appointments = [
  {
    date: 1556863200000,
    clientName: 'Должанский Николай Сергеевич', 
    statusId: 0,
    statusTitle: 'Завершён',
    holderName: 'Иванов Иван Иванович', 
    compliences: 'Боль в правом ухе', 
    diagnosis: 'Застужено правое ухо'
  },
  {
    date: 1560778200000,
    clientName: 'Петров Пётр Генадьевич', 
    statusId: 2,
    statusTitle: 'Пропущен',
    holderName: 'Иванов Иван Иванович', 
    compliences: 'Боль в горле', 
    diagnosis: 'Ангина'
  },
  {
    date: 1614512974000,
    clientName: 'Буйкевич Галина Петровна', 
    statusId: 3,
    statusTitle: 'Отменён',
    holderName: 'Нестеров Валерий Викторович', 
    compliences: 'Головные боли', 
    diagnosis: 'Мигрень'
  },
  {
    date: 1638197349000,
    clientName: 'Астафьева Ирина Михайловна', 
    statusId: 0,
    statusTitle: 'Завершён',
    holderName: 'Сидоров Генадий Павлович', 
    compliences: 'Тошнота', 
    diagnosis: 'Ротавирус'
  },
]

const appointmentStatuses = [
  { id: 0, title: 'Завершён' },
  { id: 1, title: 'Ожидается' },
  { id: 2, title: 'Пропущен' },
  { id: 3, title: 'Отменён' },
  { id: 4, title: 'Перенесён' },
  { id: 5, title: 'Активен' }
]

export function getAppointments (params) {
  const {
      statusId,
      startDate,
      endDate,
      clientName,
      onlyMe,
    } = params

  return filter(appointments, o => {
      return (startDate ? o.date >= startDate : true) && 
      (endDate ? o.date <= endDate : true) &&
      (clientName ? (clientName.length > 2 ? o.clientName.includes(clientName) : true) : true) && 
      (isNumber(statusId) && statusId >= 0 ? statusId === o.statusId : true) &&
      (onlyMe ? o.holderName === USER : true)
  })
}

export function getAppointmentStatuses () {
  return appointmentStatuses
}