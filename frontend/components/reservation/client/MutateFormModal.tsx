import {yupResolver} from '@hookform/resolvers/yup'
import axios from 'axios'
import {FormModal} from 'components/reservation/client/FormModal'
import {Err, toastErr} from 'components/utils/Err'
import {FieldValues, useForm} from 'react-hook-form'
import {useMutation, useQuery, useQueryClient} from 'react-query'
import {toast} from 'react-toastify'
import * as yup from 'yup'

export type MutateFormModalProps = {
	show: boolean,
	service: any,
	reservation?: any,
	onClose: () => void,
}

const schema = yup.object().shape({
	date: yup.string().required(),
	comment: yup.string().notRequired().min(5).nullable().transform(value => value === '' ? null : value),
}).required()

function formatNullable(val: any) {
	if (val.comment && typeof val.comment.value !== 'undefined') {
		val.comment = val.comment.value || ''
	}
	return val
}

export const MutateFormModal = (p: MutateFormModalProps) => {
	const queryClient = useQueryClient()
	const companyID = p.service.company ? p.service.company.id : p.service.companyID

	const form = useForm({
		resolver: yupResolver(schema),
		reValidateMode: 'onBlur',
		defaultValues: {
			date: p.reservation?.date,
			comment: p.reservation?.comment,
		},
	})

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
	}, {
		onSuccess: (_data, variables) => form.reset(formatNullable({
			date: p.reservation?.date,
			comment: p.reservation?.comment,
			...variables,
		}))
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
			form={form}
		/>
	)
}
