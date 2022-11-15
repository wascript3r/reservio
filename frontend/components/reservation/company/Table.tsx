import {Reservation} from './Reservation'

export const Table = ({service, reservations}: { service: string, reservations: any[] }) => {
	return (
		<table className="table mb-5">
			<thead>
			<tr>
				<th colSpan={6} className="h3 text-center">{service}</th>
			</tr>
			<tr>
				<th>First name</th>
				<th>Last name</th>
				<th>Phone</th>
				<th>Email</th>
				<th>Date</th>
				<th>Comment</th>
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
