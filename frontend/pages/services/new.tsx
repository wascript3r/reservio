import type {NextPage} from 'next'
import {useContext} from "react";
import {Auth, AuthContext, Role} from "../../components/utils/Auth";
import {useRouter} from "next/router";
import ServiceForm from "../../components/service/ServiceForm";

const Home: NextPage = () => {
	const auth = useContext(AuthContext) as Auth
	const router = useRouter()

	if (!auth.hasAccess(Role.COMPANY)) {
		router.push('/403')
		return <></>
	}

	return (
		<ServiceForm service={null}/>
	)
}

export default Home
