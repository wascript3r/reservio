import {Button, Modal} from 'react-bootstrap'

export type ViewModalProps = {
	show: boolean,
	reservation: any,
	onClose: () => void,
}

export const ViewModal = (p: ViewModalProps) => {
	return (
		<Modal show={p.show} onHide={p.onClose}>
			<Modal.Header closeButton>
				<Modal.Title>Reservation details</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<div className="row table">
					<table className="col-12 offset-sm-2 col-sm-8 offset-lg-2 col-lg-8">
						<tbody>
						<tr className="row row-cols-2">
							<td className="col card-text text-muted fw-bold text-start">Company name</td>
							<td className="col text-end">{p.reservation.service.company.name}</td>
						</tr>
						<tr className="row row-cols-2">
							<td className="col card-text text-muted fw-bold text-start">Company address</td>
							<td className="col text-end">{p.reservation.service.company.address}</td>
						</tr>
						<tr className="row row-cols-2">
							<td className="col card-text text-muted fw-bold text-start">Service</td>
							<td className="col text-end">{p.reservation.service.title}</td>
						</tr>
						<tr className="row row-cols-2">
							<td className="col card-text text-muted fw-bold text-start">Specialist
								name
							</td>
							<td className="col text-end">{p.reservation.service.specialistName ?? 'not specified'}</td>
						</tr>
						<tr className="row row-cols-2">
							<td className="col card-text text-muted fw-bold text-start">Specialist
								phone
							</td>
							<td className="col text-end text-primary align-middle">{p.reservation.service.specialistPhone ?? 'not specified'}</td>
						</tr>
						<tr className="row row-cols-2">
							<td className="col card-text text-muted fw-bold text-start">Date</td>
							<td className="col text-end">{p.reservation.date}</td>
						</tr>
						<tr className="row row-cols-2">
							<td className="col card-text text-muted fw-bold text-start">Visit
								duration
							</td>
							<td className="col text-end">{p.reservation.service.visitDuration} <span
								className="fst-italic">minutes</span></td>
						</tr>
						</tbody>
					</table>
				</div>
			</Modal.Body>
			<Modal.Footer>
				<Button variant="secondary" onClick={p.onClose}>
					Close
				</Button>
			</Modal.Footer>
		</Modal>
	)
}
