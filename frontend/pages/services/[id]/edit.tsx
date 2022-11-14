import type {NextPage} from 'next'
import {useContext} from "react";
import {Auth, AuthContext, Role} from "../../../components/utils/Auth";
import {useRouter} from "next/router";
import Spinner from "../../../components/utils/Spinner";
import {Err} from "../../../components/utils/Err";
import {useServiceInfo} from "../../../components/company/CompanyInfo";
import ServiceForm from "../../../components/service/ServiceForm";

const Home: NextPage = () => {
	const auth = useContext(AuthContext) as Auth
	const router = useRouter()
	const {id} = router.query

	if (!id) {
		return <></>
	} else if (!auth.hasAccess(Role.COMPANY)) {
		router.push('/403')
		return <></>
	}

	return (
		<ServiceFetch id={id as string}/>
	)
}

const ServiceFetch = ({id}: { id: string }) => {
	const router = useRouter()
	const auth = useContext(AuthContext) as Auth
	const {data: service, error, isLoading} = useServiceInfo(router, auth.getUserID() as string, id as string)

	if (isLoading) {
		return <Spinner/>
	} else if (error) {
		return <Err msg={error.message}/>
	}

	return (
		<ServiceForm service={service}/>
	)
}

export default Home
