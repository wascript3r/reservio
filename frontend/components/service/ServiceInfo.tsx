import axios from 'axios'
import {MutateFormModal} from 'components/reservation/client/MutateFormModal'
import Link from 'next/link'
import {useContext, useState} from 'react'
import {Button, Modal} from 'react-bootstrap'
import {useMutation, useQueryClient} from 'react-query'
import {toast} from 'react-toastify'
import {Auth, AuthContext, Role} from '../utils/Auth'
import BtnSpinner from '../utils/BtnSpinner'
import {toastErr} from '../utils/Err'

const weekdays = ['monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday', 'sunday']

const ServiceInfo = ({service}: { service: any }) => {
	const auth = useContext(AuthContext) as Auth
	const [show, setShow] = useState(false)

	const handleClose = () => setShow(false)
	const handleShow = () => setShow(true)

	const queryClient = useQueryClient()
	const {mutate: deleteq, isLoading: isDeleting} = useMutation(() => {
		return axios.delete(`/companies/${service.companyID}/services/${service.id}`)
			.then(() => {
				setShow(false)
				toast.success('Service was successfully deleted')
				return queryClient.invalidateQueries(['company', service.companyID, 'services'])
			}).catch(err => toastErr(err))
	})

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
				<MutateFormModal show={show} service={service} onClose={handleClose}/>
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
						<Button variant="danger" onClick={() => deleteq()} disabled={isDeleting}>
							{isDeleting ? <BtnSpinner/> : 'Delete'}
						</Button>
					</Modal.Footer>
				</Modal>
			)}
		</>
	)
}

export default ServiceInfo
