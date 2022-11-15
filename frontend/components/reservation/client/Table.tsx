import {Reservation} from './Reservation'

export const Table = ({reservations}: { reservations: any[] }) => {
	return (
		<table className="table">
			<thead>
			<tr>
				<th>Company</th>
				<th>Address</th>
				<th>Service</th>
				<th>Date</th>
				<th>Actions</th>
			</tr>
			</thead>
			<tbody>
			{reservations.map((reservation, index) => (
				<Reservation reservation={reservation} key={index}/>
			))}
			</tbody>
		</table>
	)
}
