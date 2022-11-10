import {Auth, AuthContext, Role} from "../utils/Auth";
import {useContext, useMemo, useState} from "react";
import {Button, Modal} from "react-bootstrap";
import '@mobiscroll/react/dist/css/mobiscroll.min.css'
import {Datepicker, setOptions} from '@mobiscroll/react';
import moment from "moment";

setOptions({
	theme: 'ios',
	themeVariant: 'light'
});

const weekdays = new Map<string, number>([
	["Monday", 1],
	["Tuesday", 2],
	["Wednesday", 3],
	["Thursday", 4],
	["Friday", 5],
	["Saturday", 6],
	["Sunday", 7],
])

const numToWeekday = new Map<number, string>([
	[0, "sunday"],
	[1, "monday"],
	[2, "tuesday"],
	[3, "wednesday"],
	[4, "thursday"],
	[5, "friday"],
	[6, "saturday"],
])

const weekdayShort = new Map<string, string>([
	['sunday', 'SU'],
	['monday', 'MO'],
	['tuesday', 'TU'],
	['wednesday', 'WE'],
	['thursday', 'TH'],
	['friday', 'FR'],
	['saturday', 'SA'],
])

const sortWorkSchedule = (workSchedule: Map<string, object>) => {
	const arr = Object.entries(workSchedule).map(([day, time]) => {
		return {
			...time,
			day: day[0].toUpperCase() + day.substring(1),
		}
	})
	return arr.sort((a, b) => (weekdays.get(a.day) || 0) - (weekdays.get(b.day) || 0))
}

const ServiceInfo = ({service}: { service: any }) => {
	const auth = useContext(AuthContext) as Auth
	const [show, setShow] = useState(false)

	const handleClose = () => setShow(false)
	const handleShow = () => setShow(true)

	const [minTime, setMinTime] = useState<string>('')
	const [maxTime, setMaxTime] = useState<string>('')
	const [invalid, setInvalid] = useState<any[]>([])

	const handleDateChange = (event: any) => {
		const weekday = numToWeekday.get(event.date.getDay()) as string
		const schedule = service.workSchedule[weekday]
		if (schedule) {
			setMinTime(schedule.from)
			setMaxTime(schedule.to)
		}
	}

	const validWeekdays = useMemo(() => {
		const validDays = Object.keys(service.workSchedule).map(day => {
			return weekdayShort.get(day)
		})
		return validDays.join(',')
	}, [service.workSchedule])

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
									<td className="col text-end">{service.specialistName}</td>
								</tr>
								<tr className="row row-cols-2">
									<td className="col card-text text-muted fw-bold text-start">Specialist
										phone
									</td>
									<td className="col text-end text-primary align-middle">{service.specialistPhone}</td>
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
								{sortWorkSchedule(service.workSchedule).map((day: any, index: number) => (
									<tr className="row row-cols-2" key={index}>
										<td className="col card-text text-muted fw-bold text-start">{day.day}</td>
										<td className="col text-end">{day.from}-{day.to}</td>
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
					</div>
				</div>
			</div>

			<Modal show={show} onHide={handleClose}>
				<Modal.Header closeButton>
					<Modal.Title>Time reservation</Modal.Title>
				</Modal.Header>
				<Modal.Body>
					<div className="mbsc-form-group">
						<div className="mbsc-form-group-title">Select visit time</div>
						<Datepicker
							controls={['calendar', 'timegrid']}
							min={moment().format('YYYY-MM-DD')}
							minTime={minTime}
							maxTime={maxTime}
							stepMinute={service.visitDuration}
							onCellClick={handleDateChange}
							valid={[
								{
									recurring: {
										weekDays: validWeekdays,
										repeat: 'weekly',
									}
								}
							]}
							invalid={invalid}
							cssClass="booking-datetime"
						/>
					</div>
				</Modal.Body>
				<Modal.Footer>
					<Button variant="secondary" onClick={handleClose}>
						Close
					</Button>
					{/*<Button variant="danger" onClick={() => deleteq()} disabled={isDeleting}>*/}
					{/*	{isDeleting ? <BtnSpinner/> : 'Delete'}*/}
					{/*</Button>*/}
				</Modal.Footer>
			</Modal>
		</>
	)
}

export default ServiceInfo
