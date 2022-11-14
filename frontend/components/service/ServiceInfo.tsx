import {Auth, AuthContext, Role} from "../utils/Auth";
import {useContext, useMemo, useState} from "react";
import {Button, Modal} from "react-bootstrap";
import '@mobiscroll/react/dist/css/mobiscroll.min.css'
import {Datepicker, setOptions} from '@mobiscroll/react';
import moment from "moment";
import {useMutation, useQuery} from "react-query";
import axios from "axios";
import {Err, toastErr} from "../utils/Err";
import {FieldValues, useForm} from "react-hook-form";
import {yupResolver} from "@hookform/resolvers/yup";
import {toast} from "react-toastify";
import * as yup from "yup";
import BtnSpinner from "../utils/BtnSpinner";
import Link from "next/link";

setOptions({
	theme: 'ios',
	themeVariant: 'light'
});

const schema = yup.object().shape({
	date: yup.string().required(),
	comment: yup.string().notRequired().min(5).nullable().transform(value => value === '' ? null : value),
}).required();

const weekdays = ['monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday', 'sunday']

const numToWeekday = new Map<number, string>([
	[0, "sunday"],
	[1, "monday"],
	[2, "tuesday"],
	[3, "wednesday"],
	[4, "thursday"],
	[5, "friday"],
	[6, "saturday"],
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

const ServiceInfo = ({service}: { service: any }) => {
	const auth = useContext(AuthContext) as Auth
	const [show, setShow] = useState(false)

	const handleClose = () => {
		setShow(false)
		reset()
	}
	const handleShow = () => setShow(true)

	const [minTime, setMinTime] = useState<string>('')
	const [maxTime, setMaxTime] = useState<string>('')

	const {data: invalid, error: qerror} = useQuery<any, Error>(['services', service.id, 'reservations'], () => {
		return axios.get(`/companies/${service.companyID}/services/${service.id}/reservations`)
			.then(res =>
				res.data.data.reservations.map((r: any) => {
					return {
						start: r.date,
						end: r.date,
					}
				})
			)
	}, {enabled: show, initialData: []})

	const handleDateChange = (event: any) => {
		const weekday = numToWeekday.get(event.date.getDay()) as string
		const schedule = service.workSchedule[weekday]
		if (schedule) {
			setMinTime(schedule.from)
			const to = moment(schedule.to, 'HH:mm').subtract(service.visitDuration, 'minutes')
			setMaxTime(to.format('HH:mm'))
		}
	}

	const invalidWeekdays = useMemo(() => {
		return Array.from(weekdayLong.keys())
			.map(day => {
				const long = weekdayLong.get(day)
				if (!long) return ''
				return service.workSchedule[long] ? '' : day
			})
			.filter(day => day !== '')
			.join(',')
	}, [service.workSchedule])

	const {register, handleSubmit, formState: {errors}, setValue, reset} = useForm({
		resolver: yupResolver(schema),
		reValidateMode: 'onBlur'
	})
	const {mutate, isLoading} = useMutation((data: FieldValues) => {
		return axios.post(`/companies/${service.companyID}/services/${service.id}/reservations`, data)
			.then(() => {
				toast.success('You have successfully made a reservation')
				handleClose()
			}).catch(err => toastErr(err))
	})
	const onSubmit = (data: FieldValues) => mutate(data)

	if (qerror) {
		return <Err msg={qerror.message}/>
	}

	return (
		<>
			<div className="col">
				<div className="card mb-4 rounded-3 shadow-sm">
					<div className="card-header py-3">
						<h4 className="my-0 fw-normal">{service.title}</h4>
					</div>
					<div className="card-body">
						<div className="card-title h5">{service.description}</div>
						<div className="row table">
							<table className="col-12 offset-sm-2 col-sm-8 offset-lg-2 col-lg-8">
								<tbody>
								<tr className="row row-cols-2">
									<td className="col card-text text-muted fw-bold text-start">Specialist
										name
									</td>
									<td className="col text-end">{service.specialistName ?? 'not specified'}</td>
								</tr>
								<tr className="row row-cols-2">
									<td className="col card-text text-muted fw-bold text-start">Specialist
										phone
									</td>
									<td className="col text-end text-primary align-middle">{service.specialistPhone ?? 'not specified'}</td>
								</tr>
								<tr className="row row-cols-2">
									<td className="col card-text text-muted fw-bold text-start">Visit
										duration
									</td>
									<td className="col text-end">{service.visitDuration} <span
										className="fst-italic">minutes</span></td>
								</tr>
								</tbody>
							</table>
						</div>

						<div className="card-title h5 mt-5">Work schedule</div>
						<div className="row table">
							<table className="col-12 offset-sm-2 col-sm-8 offset-lg-2 col-lg-8">
								<tbody>
								{weekdays.map((day: string, index: number) => (
									<tr className="row row-cols-2" key={index}>
										<td className="col card-text text-muted fw-bold text-start">{day.charAt(0).toUpperCase() + day.slice(1)}</td>
										<td className="col text-end">{service.workSchedule[day] ? service.workSchedule[day].from + '-' + service.workSchedule[day].to : 'Closed'}</td>
									</tr>
								))}
								</tbody>
							</table>
						</div>

						{auth.hasAccess(Role.CLIENT) &&
                            <button
                                type="button"
                                className="w-100 btn btn-lg btn-outline-secondary mt-3"
                                onClick={handleShow}>Reserve time
                            </button>
						}

						{auth.hasAccess(Role.COMPANY) && auth.getUserID() === service.companyID &&
                            <div className="row">
                                <div className="col-6">
                                    <Link href={`/services/${service.id}/edit`}>
                                        <button type="button"
                                                className={`w-100 btn btn-outline-primary mt-3`}>Edit
                                        </button>
                                    </Link>
                                </div>
                                <div className="col-6">
                                    <button type="button"
                                            className="w-100 btn btn-outline-danger mt-3" onClick={handleShow}>Delete
                                    </button>
                                </div>
                            </div>
						}
					</div>
				</div>
			</div>

			{auth.hasAccess(Role.CLIENT) && (
				<Modal show={show} onHide={handleClose}>
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
										stepMinute={service.visitDuration}
										onCellClick={handleDateChange}
										invalid={[
											...invalid,
											{
												recurring: {
													weekDays: invalidWeekdays,
													repeat: 'weekly',
												},
											}
										]}
										onTempChange={(event: any) => {
											setValue('date', moment(event.value).format('YYYY-MM-DD HH:mm'))
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
							<button type="submit" className="btn btn-primary" disabled={isLoading}>
								{isLoading ? <BtnSpinner/> : 'Reserve'}
							</button>
						</Modal.Footer>
					</form>
				</Modal>
			)}

			{auth.hasAccess(Role.COMPANY) && (
				<Modal show={show} onHide={handleClose}>
					<Modal.Header closeButton>
						<Modal.Title>Delete confirmation</Modal.Title>
					</Modal.Header>
					<Modal.Body>Are you sure you want to delete this service?</Modal.Body>
					<Modal.Footer>
						<Button variant="secondary" onClick={handleClose}>
							Close
						</Button>
						{/*<Button variant="danger" onClick={() => deleteq()} disabled={isDeleting}>*/}
						{/*	{isDeleting ? <BtnSpinner/> : 'Delete'}*/}
						{/*</Button>*/}
					</Modal.Footer>
				</Modal>
			)}
		</>
	)
}

export default ServiceInfo
