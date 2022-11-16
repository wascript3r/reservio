import {yupResolver} from '@hookform/resolvers/yup'
import {Datepicker, setOptions} from '@mobiscroll/react'
import '@mobiscroll/react/dist/css/mobiscroll.min.css'
import BtnSpinner from 'components/utils/BtnSpinner'
import {extractDirtyFields} from 'components/utils/form'
import moment from 'moment/moment'
import {useCallback, useMemo, useState} from 'react'
import {Button, Modal} from 'react-bootstrap'
import {FieldValues, useForm} from 'react-hook-form'
import * as yup from 'yup'

const numToWeekday = new Map<number, string>([
	[0, 'sunday'],
	[1, 'monday'],
	[2, 'tuesday'],
	[3, 'wednesday'],
	[4, 'thursday'],
	[5, 'friday'],
	[6, 'saturday'],
])

const weekdayLong = new Map<string, string>([
	['SU', 'sunday'],
	['MO', 'monday'],
	['TU', 'tuesday'],
	['WE', 'wednesday'],
	['TH', 'thursday'],
	['FR', 'friday'],
	['SA', 'saturday'],
])

export type FormModalProps = {
	show: boolean,
	service: any,
	reservation?: any,
	invalidTimeSlots: any[]
	isReserving: boolean,
	onClose: () => void,
	onSubmit: (data: FieldValues) => void,
}

const schema = yup.object().shape({
	date: yup.string().required(),
	comment: yup.string().notRequired().min(5).nullable().transform(value => value === '' ? null : value),
}).required()

setOptions({
	theme: 'ios',
	themeVariant: 'light',
})

export const FormModal = (p: FormModalProps) => {
	const {register, handleSubmit, formState: {errors, isDirty, dirtyFields}, setValue, getValues, reset} = useForm({
		resolver: yupResolver(schema),
		reValidateMode: 'onBlur',
		defaultValues: {
			date: p.reservation?.date,
			comment: p.reservation?.comment,
		},
	})

	const handleClose = () => {
		p.onClose()
		reset()
	}

	const [minTime, setMinTime] = useState<string>('')
	const [maxTime, setMaxTime] = useState<string>('')

	const handleDateChange = (event: any) => {
		if (!event.date && !event.value) return

		const date = event.date || moment(event.value).toDate()

		const weekday = numToWeekday.get(date.getDay()) as string
		const schedule = p.service.workSchedule[weekday]
		if (schedule) {
			setMinTime(schedule.from)
			const to = moment(schedule.to, 'HH:mm').subtract(p.service.visitDuration, 'minutes')
			setMaxTime(to.format('HH:mm'))
		}
	}

	const invalidWeekdays = useMemo(() => {
		return Array.from(weekdayLong.keys())
			.map(day => {
				const long = weekdayLong.get(day)
				if (!long) return ''
				return p.service.workSchedule[long] ? '' : day
			})
			.filter(day => day !== '')
			.join(',')
	}, [p.service.workSchedule])

	const invalidSlots = useMemo(() => {
		if (!p.reservation) return p.invalidTimeSlots

		return p.invalidTimeSlots.filter(slot => slot.start !== p.reservation?.date)
	}, [p.invalidTimeSlots, p.reservation])

	const onSubmit = useCallback((data: FieldValues) => {
		const dirtyData = extractDirtyFields(data, dirtyFields)

		if (p.service) {
			if (typeof dirtyData.comment !== 'undefined') {
				dirtyData.comment = {
					value: dirtyData.comment,
				}
			}
		}

		p.onSubmit(dirtyData)
	}, [p, dirtyFields])

	return (
		<Modal show={p.show} onHide={handleClose}>
			<Modal.Header closeButton>
				<Modal.Title>Time reservation</Modal.Title>
			</Modal.Header>
			<form onSubmit={handleSubmit(onSubmit)}>
				<Modal.Body>
					<div className={`form-control ${errors.date ? 'is-invalid' : ''}`}>
						<div className="mbsc-form-group">
							<div className="mbsc-form-group-title">Select visit time</div>
							<Datepicker
								controls={['calendar', 'timegrid']}
								min={moment().format('YYYY-MM-DD')}
								minTime={minTime}
								maxTime={maxTime}
								stepMinute={p.service.visitDuration}
								onOpen={handleDateChange}
								onCellClick={handleDateChange}
								defaultValue={getValues('date')}
								invalid={[
									...invalidSlots,
									{
										recurring: {
											weekDays: invalidWeekdays,
											repeat: 'weekly',
										},
									},
								]}
								onTempChange={(event: any) => {
									handleDateChange(event)
									setValue('date', moment(event.value).format('YYYY-MM-DD HH:mm'), {
										shouldDirty: true,
									})
								}}
								cssClass="booking-datetime"
							/>
						</div>
					</div>
					{errors.date &&
                        <div className="invalid-feedback text-center">{errors.date.message as string}</div>}
					<div className="mt-3">
						<label htmlFor="comment" className="form-label">Comment</label>
						<input {...register('comment')} type="text"
							   className={`form-control ${errors.comment ? 'is-invalid' : ''}`}
							   placeholder="optional"/>
						{errors.comment &&
                            <div
                                className="invalid-feedback text-center">{errors.comment.message as string}</div>}
					</div>
				</Modal.Body>
				<Modal.Footer>
					<Button variant="secondary" onClick={handleClose}>
						Close
					</Button>
					{p.reservation && (
						<button type="submit" className="btn btn-primary"
								disabled={p.isReserving || !isDirty}>
							{p.isReserving ? <BtnSpinner/> : 'Update'}
						</button>
					)}
					{!p.reservation && (
						<button type="submit" className="btn btn-primary" disabled={p.isReserving}>
							{p.isReserving ? <BtnSpinner/> : 'Reserve'}
						</button>
					)}
				</Modal.Footer>
			</form>
		</Modal>
	)
}
