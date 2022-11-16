import moment from 'moment'

export const Reservation = ({reservation}: { reservation: any }) => {
	const isPast = moment(reservation.date).isBefore(moment())

	return (
		<tr className={isPast ? 'table-secondary' : ''}>
			<td>{reservation.client.firstName}</td>
			<td>{reservation.client.lastName}</td>
			<td>{reservation.client.phone}</td>
			<td>{reservation.client.email}</td>
			<td>{reservation.date}</td>
			<td>{reservation.comment || '-'}</td>
		</tr>
	)
}
