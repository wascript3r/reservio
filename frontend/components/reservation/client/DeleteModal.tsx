import BtnSpinner from 'components/utils/BtnSpinner'
import {Button, Modal} from 'react-bootstrap'

export type DeleteModalProps = {
	show: boolean,
	isDeleting: boolean,
	onClose: () => void,
	onSubmit: () => void,
}

export const DeleteModal = (p: DeleteModalProps) => {
	return (
		<Modal show={p.show} onHide={p.onClose}>
			<Modal.Header closeButton>
				<Modal.Title>Cancel confirmation</Modal.Title>
			</Modal.Header>
			<Modal.Body>Are you sure you want to cancel this reservation?</Modal.Body>
			<Modal.Footer>
				<Button variant="secondary" onClick={p.onClose}>
					Close
				</Button>
				<Button variant="danger" onClick={() => p.onSubmit()} disabled={p.isDeleting}>
					{p.isDeleting ? <BtnSpinner/> : 'Cancel'}
				</Button>
			</Modal.Footer>
		</Modal>
	)
}
