import {useQuery} from "react-query";
import axios from "axios";
import Spinner from "../utils/Spinner";
import {Err} from "../utils/Err";
import {useRouter} from "next/router";
import ServiceInfo from "../service/ServiceInfo";

const CompanyInfo = ({id}: { id: string }) => {
	const router = useRouter()

	const {data: company, error: cerror, isLoading: isCompanyLoading} = useQuery<any, Error>(["company", id], () => {
		return axios.get(`/companies/${id}`)
			.then(res => res.data)
			.catch(err => {
				if (err.response.status === 400 || err.response.status === 404) {
					router.push("/404")
				}
				return Promise.reject(err)
			})
	})
	const {data: services, error: serror, isLoading: isServicesLoading} = useQuery<any, Error>(["services", id], () => {
		return axios.get(`/companies/${id}/services`).then(res => res.data)
	})

	if (isCompanyLoading || isServicesLoading) {
		return <Spinner/>
	}

	if (cerror) {
		return <Err msg={cerror.message}/>
	} else if (serror) {
		return <Err msg={serror.message}/>
	}

	return (
		<>
			<h2 className="text-center my-5">Company information</h2>
			<div className="row row-cols-1 row-cols-md-1 mb-3 text-center">
				<div className="col">
					<div className="card mb-4 rounded-3 shadow-sm border-primary">
						<div className="card-header py-3 text-white bg-primary border-primary">
							<h4 className="my-0 fw-normal">{company?.data.name}</h4>
						</div>
						<div className="card-body">
							<div className="card-title h5">{company?.data.description}</div>
							<div className="row table">
								<table className="col-12 offset-sm-2 col-sm-8 offset-lg-3 col-lg-6">
									<tbody>
									<tr className="row row-cols-2">
										<td className="col card-text text-muted fw-bold text-start">Location</td>
										<td className="col text-end">{company?.data.address}</td>
									</tr>
									<tr className="row row-cols-2">
										<td className="col card-text text-muted fw-bold text-start">Contact email</td>
										<td className="col text-end">{company?.data.email}</td>
									</tr>
									</tbody>
								</table>
							</div>
						</div>
					</div>
				</div>
				<div className="col">
					<h2 className="text-center mt-3 mb-5">Services</h2>
					<div className="row row-cols-1 row-cols-md-2 mb-3 text-center">
						{services?.data.services.length > 0 && services?.data.services.map((service: any, index: number) => (
							<ServiceInfo service={service} key={index}/>
						))}
						{services?.data.services.length === 0 &&
                            <div className="w-100 text-muted text-center">
                                Company doesn't have any services yet
                            </div>
						}
					</div>
				</div>
			</div>
		</>
	)
}

export default CompanyInfo
