import type {NextPage} from 'next'
import {useContext} from "react";
import {AuthContext} from "../components/utils/Auth";
import {useRouter} from "next/router";

const Home: NextPage = () => {
	const auth = useContext(AuthContext)
	const router = useRouter()

	if (auth) {
		auth.logout()
		router.push('/')
	}

	return <></>
}

export default Home
