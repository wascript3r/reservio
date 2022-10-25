import type {NextPage} from 'next'
import LoginForm from "../components/auth/LoginForm";
import {useContext} from "react";
import {AuthContext} from "../components/utils/Auth";
import {useRouter} from "next/router";

const Home: NextPage = () => {
	const auth = useContext(AuthContext)
	const router = useRouter()

	if (!auth) {
		return <></>
	} else if (auth.loggedIn()) {
		router.push('/')
		return <></>
	}

	return (
		<LoginForm/>
	)
}

export default Home
