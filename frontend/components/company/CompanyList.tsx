import {useQuery} from "react-query";
import axios from "axios";
import Spinner from "../utils/Spinner";
import { Err } from "../utils/Err";
import Link from "next/link";

const CompanyList = () => {
	const {data, error, isLoading} = useQuery<any, Error>("companies", () => {
		return axios.get("/companies").then(res => res.data)
	})

	if (isLoading) {
		return <Spinner/>
	}

	if (error) {
		return <Err msg={error.message}/>
	}

	return (
		<>
			<h2 className="text-center my-5">Registered companies</h2>
			<div className="row row-cols-1 row-cols-md-2 mb-3 text-center">
				{data?.data.companies.map((company: any, index: number) => (
					<div className="col" key={index}>
						<div className="card mb-4 rounded-3 shadow-sm border-primary">
							<div className="card-header py-3 text-white bg-primary border-primary">
								<h4 className="my-0 fw-normal">{company.name}</h4>
							</div>
							<div className="card-body">
								<div className="card-title h5">{company.description}</div>
								<div>
									<span className="card-text text-muted fw-bold">Location: </span>
									<span>{company.address}</span>
								</div>

								<Link href={`/companies/${company.id}`}>
									<button type="button" className="w-100 btn btn-lg btn-outline-primary mt-3">View services
									</button>
								</Link>
							</div>
						</div>
					</div>
				))}
			</div>
		</>
	)
}

export default CompanyList