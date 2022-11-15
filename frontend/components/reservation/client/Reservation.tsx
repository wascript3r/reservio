import {faEdit, faTrash} from '@fortawesome/free-solid-svg-icons'
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'
import {MutateDeleteModal} from 'components/reservation/MutateDeleteModal'
import {MutateFormModal} from 'components/reservation/MutateFormModal'
import moment from 'moment'
import {useState} from 'react'

export const Reservation = ({reservation}: { reservation: any }) => {
	const isPast = moment(reservation.date).isBefore(moment())
	const [showEdit, setShowEdit] = useState(false)
	const [showDelete, setShowDelete] = useState(false)

	const handleCloseEdit = () => setShowEdit(false)
	const handleShowEdit = () => setShowEdit(true)

	const handleCloseDelete = () => setShowDelete(false)
	const handleShowDelete = () => setShowDelete(true)

	return (
		<>
			<tr className={isPast ? 'table-secondary' : ''}>
				<td>{reservation.service.company.name}</td>
				<td>{reservation.service.company.address}</td>
				<td>{reservation.service.title}</td>
				<td>{reservation.date}</td>
				<td>
					{!isPast &&
                        <div className="btn-group">
                            <button className="btn btn-sm btn-primary col-6" onClick={handleShowEdit}>
                                <FontAwesomeIcon icon={faEdit}/>
                            </button>
                            <button className="btn btn-sm btn-danger col-6" onClick={handleShowDelete}>
                                <FontAwesomeIcon icon={faTrash}/>
                            </button>
                        </div>
					}
				</td>
			</tr>

			<MutateFormModal show={showEdit} service={reservation.service} reservation={reservation}
							 onClose={handleCloseEdit}/>
			<MutateDeleteModal show={showDelete} reservation={reservation} onClose={handleCloseDelete}/>
		</>
	)
}
