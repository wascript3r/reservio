import type {NextPage} from 'next'
import {useContext} from "react";
import {Auth, AuthContext} from "../../components/utils/Auth";
import {useRouter} from "next/router";
import ClientRegForm from "../../components/auth/ClientRegForm";

const Home: NextPage = () => {
	const auth = useContext(AuthContext) as Auth
	const router = useRouter()

	if (auth.loggedIn()) {
		router.push('/')
		return <></>
	}

	return (
		<ClientRegForm/>
	)
}

export default Home
