import axios from 'axios'
import {FormModal} from 'components/reservation/client/FormModal'
import {Err, toastErr} from 'components/utils/Err'
import {FieldValues} from 'react-hook-form'
import {useMutation, useQuery, useQueryClient} from 'react-query'
import {toast} from 'react-toastify'

export type MutateFormModalProps = {
	show: boolean,
	service: any,
	reservation?: any,
	onClose: () => void,
}

export const MutateFormModal = (p: MutateFormModalProps) => {
	const queryClient = useQueryClient()
	const companyID = p.service.company ? p.service.company.id : p.service.companyID

	const {
		data: invalidTimeSlots,
		error: qerror,
	} = useQuery<any, Error>(['services', p.service.id, 'reservations'], () => {
		return axios.get(`/companies/${companyID}/services/${p.service.id}/reservations`)
			.then(res =>
				res.data.data.reservations.map((r: any) => {
					return {
						start: r.date,
						end: r.date,
					}
				}),
			)
	}, {enabled: p.show, initialData: []})

	const {mutate, isLoading: isReserving} = useMutation((data: FieldValues) => {
		const url = p.reservation
			? `/companies/${companyID}/services/${p.service.id}/reservations/${p.reservation.id}`
			: `/companies/${companyID}/services/${p.service.id}/reservations`
		const method = p.reservation ? 'PATCH' : 'POST'

		return axios({url, method, data})
			.then(() => queryClient.invalidateQueries('client_reservations'))
			.then(() => {
				toast.success(`You have successfully ${p.reservation ? 'updated' : 'made'} a reservation`)
				p.onClose()
			}).catch(err => toastErr(err))
	})

	const onSubmit = (data: FieldValues) => mutate(data)

	if (qerror) {
		return <Err msg={qerror.message}/>
	}

	return (
		<FormModal
			show={p.show}
			service={p.service}
			reservation={p.reservation}
			invalidTimeSlots={invalidTimeSlots}
			isReserving={isReserving}
			onClose={p.onClose}
			onSubmit={onSubmit}
		/>
	)
}
