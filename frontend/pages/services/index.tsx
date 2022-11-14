import type {NextPage} from 'next'
import {useContext} from "react";
import {Auth, AuthContext, Role} from "../../components/utils/Auth";
import {useRouter} from "next/router";
import ServiceList from "../../components/service/ServiceList";

const Home: NextPage = () => {
	const auth = useContext(AuthContext) as Auth
	const router = useRouter()

	if (!auth.hasAccess(Role.COMPANY)) {
		router.push('/403')
		return <></>
	}

	return (
		<ServiceList id={auth.getUserID() as string}/>
	)
}

export default Home
