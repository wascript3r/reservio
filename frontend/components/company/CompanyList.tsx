import {useQuery} from "react-query";
import axios from "axios";
import Spinner from "../utils/Spinner";
import {Err} from "../utils/Err";
import {useContext} from "react";
import {Auth, AuthContext, Role} from "../utils/Auth";
import CompanyCard from "./CompanyCard";

const CompanyList = () => {
	const {data, error, isLoading} = useQuery<any, Error>("companies", () => {
		return axios.get("/companies").then(res => res.data)
	})
	const auth = useContext(AuthContext) as Auth

	if (isLoading) {
		return <Spinner/>
	}

	if (error) {
		return <Err msg={error.message}/>
	}
	const isAdmin = auth.hasAccess(Role.ADMIN)

	return (
		<>
			<h2 className="text-center my-5">Registered companies</h2>
			<div className="row row-cols-1 row-cols-md-2 mb-3 text-center">
				{data?.data.companies.map((company: any, index: number) => (
					<CompanyCard company={company} isAdmin={isAdmin} key={index}/>
				))}
			</div>
		</>
	)
}

export default CompanyList