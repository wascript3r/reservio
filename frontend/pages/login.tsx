import type {NextPage} from 'next'
import LoginForm from "../components/auth/LoginForm";
import {useContext} from "react";
import {Auth, AuthContext} from "../components/utils/Auth";
import {useRouter} from "next/router";

const Home: NextPage = () => {
	const auth = useContext(AuthContext) as Auth
	const router = useRouter()

	if (auth.loggedIn()) {
		router.push('/')
		return <></>
	}

	return (
		<LoginForm/>
	)
}

export default Home
