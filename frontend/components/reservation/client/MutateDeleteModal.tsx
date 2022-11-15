import axios from 'axios'
import {DeleteModal} from 'components/reservation/client/DeleteModal'
import {toastErr} from 'components/utils/Err'
import {useMutation, useQueryClient} from 'react-query'
import {toast} from 'react-toastify'

export type MutateDeleteModalProps = {
	show: boolean,
	reservation: any,
	onClose: () => void,
}

export const MutateDeleteModal = (p: MutateDeleteModalProps) => {
	const queryClient = useQueryClient()
	const {mutate: deleteq, isLoading: isDeleting} = useMutation(() => {
		return axios.delete(`/companies/${p.reservation.service.company.id}/services/${p.reservation.service.id}/reservations/${p.reservation.id}`)
			.then(() => queryClient.invalidateQueries('client_reservations'))
			.then(() => {
				toast.success('Reservation was successfully cancelled')
				p.onClose()
			}).catch(err => toastErr(err))
	})

	return (
		<DeleteModal show={p.show} isDeleting={isDeleting} onClose={p.onClose} onSubmit={deleteq}/>
	)
}
