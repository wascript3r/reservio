import {Table} from 'components/reservation/company/Table'

export type ServiceReservations = {
	title: string,
	reservations: any[],
}

export const Tables = ({reservations}: { reservations: ServiceReservations[] }) => {
	return (
		<div>
			{reservations.map((sr, index) => (
				<Table key={index} service={sr.title} reservations={sr.reservations}/>
			))}
		</div>
	)
}
