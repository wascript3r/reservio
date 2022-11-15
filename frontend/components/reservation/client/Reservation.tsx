import {MutateFormModal} from 'components/reservation/MutateFormModal'
import moment from 'moment'
import {useState} from 'react'

export const Reservation = ({reservation}: { reservation: any }) => {
	const isPast = moment(reservation.date).isBefore(moment())
	const [show, setShow] = useState(false)

	const handleClose = () => setShow(false)
	const handleShow = () => setShow(true)

	return (
		<>
			<tr className={isPast ? 'table-secondary' : ''}>
				<td>{reservation.service.company.name}</td>
				<td>{reservation.service.company.address}</td>
				<td>{reservation.service.title}</td>
				<td>{reservation.date}</td>
				<td>
					{!isPast &&
                        <button className="btn btn-primary" onClick={handleShow}>Edit</button>
					}
				</td>
			</tr>

			<MutateFormModal show={show} service={reservation.service} reservation={reservation} onClose={handleClose}/>
		</>
	)
}
