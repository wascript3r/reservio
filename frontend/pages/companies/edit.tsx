import type {NextPage} from 'next'
import {useContext} from "react";
import {Auth, AuthContext, Role} from "../../components/utils/Auth";
import {useRouter} from "next/router";
import CompanyRegForm from "../../components/auth/CompanyRegForm";
import Spinner from "../../components/utils/Spinner";
import {Err} from "../../components/utils/Err";
import {useCompanyInfo} from "../../components/company/CompanyInfo";

const Home: NextPage = () => {
	const auth = useContext(AuthContext) as Auth
	const router = useRouter()

	if (!auth.hasAccess(Role.COMPANY)) {
		router.push('/403')
		return <></>
	}

	const {data: company, error, isLoading} = useCompanyInfo(router, auth.getUserID() as string)

	if (isLoading) {
		return <Spinner/>
	} else if (error) {
		return <Err msg={error.message}/>
	}

	return (
		<CompanyRegForm company={company}/>
	)
}

export default Home
