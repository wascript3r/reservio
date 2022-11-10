import type {NextPage} from 'next'
import {useContext} from "react";
import {Auth, AuthContext} from "../../components/utils/Auth";
import {useRouter} from "next/router";
import CompanyRegForm from "../../components/auth/CompanyRegForm";

const Home: NextPage = () => {
	const auth = useContext(AuthContext) as Auth
	const router = useRouter()

	if (auth.loggedIn()) {
		router.push('/')
		return <></>
	}

	return (
		<CompanyRegForm company={null}/>
	)
}

export default Home
