import Link from "next/link";
import BtnSpinner from "../utils/BtnSpinner";
import {useMutation, useQueryClient} from "react-query";
import axios from "axios";
import {toast} from "react-toastify";
import {toastErr} from "../utils/Err";
import {useState} from "react";
import {Button, Modal} from "react-bootstrap";

const CompanyCard = ({company, isAdmin}: { company: any, isAdmin: boolean }) => {
	const color = company.approved ? 'primary' : 'secondary'
	const queryClient = useQueryClient()
	const [show, setShow] = useState(false)

	const handleClose = () => setShow(false)
	const handleShow = () => setShow(true)

	const {mutate: approve, isLoading: isApproving} = useMutation((a: boolean) => {
		return axios.patch(`/companies/${company.id}`, {approved: a})
			.then(() => {
				toast.success(`Company was successfully ${a ? 'approved' : 'disapproved'}`)
				return queryClient.invalidateQueries('companies')
			}).catch(err => toastErr(err))
	})
	const {mutate: deleteq, isLoading: isDeleting} = useMutation(() => {
		return axios.delete(`/companies/${company.id}`)
			.then(() => {
				setShow(false)
				toast.success('Company was successfully deleted')
				return queryClient.invalidateQueries('companies')
			}).catch(err => toastErr(err))
	})

	return (
		<>
			<div className="col">
				<div className={`card mb-4 rounded-3 shadow-sm border-${color}`}>
					<div className={`card-header py-3 text-white bg-${color} border-${color}`}>
						<h4 className="my-0 fw-normal">{company.name}</h4>
					</div>
					<div className="card-body">
						<div className="card-title h5">{company.description}</div>
						<div>
							<span className="card-text text-muted fw-bold">Location: </span>
							<span>{company.address}</span>
						</div>

						<Link href={`/companies/${company.id}`}>
							<button type="button"
									className={`w-100 btn btn-lg btn-outline-${color} mt-3`}>View
								services
							</button>
						</Link>
						{isAdmin &&
                            <div className="row">
                                <div className="col-6">
									{!company.approved &&
                                        <button type="button"
                                                className="w-100 btn btn-outline-secondary mt-3" disabled={isApproving}
                                                onClick={() => approve(true)}>{isApproving ?
											<BtnSpinner/> : 'Approve'}
                                        </button>
									}
									{company.approved &&
                                        <button type="button"
                                                className="w-100 btn btn-outline-primary mt-3" disabled={isApproving}
                                                onClick={() => approve(false)}>{isApproving ?
											<BtnSpinner/> : 'Disapprove'}
                                        </button>
									}
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

			<Modal show={show} onHide={handleClose}>
				<Modal.Header closeButton>
					<Modal.Title>Delete confirmation</Modal.Title>
				</Modal.Header>
				<Modal.Body>Are you sure you want to delete this company?</Modal.Body>
				<Modal.Footer>
					<Button variant="secondary" onClick={handleClose}>
						Close
					</Button>
					<Button variant="danger" onClick={() => deleteq()} disabled={isDeleting}>
						{isDeleting ? <BtnSpinner/> : 'Delete'}
					</Button>
				</Modal.Footer>
			</Modal>
		</>
	)
}

export default CompanyCard