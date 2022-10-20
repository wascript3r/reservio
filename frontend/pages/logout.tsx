import type {NextPage} from 'next'
import {useContext, useEffect} from "react";
import {AuthContext} from "../components/utils/Auth";
import {useRouter} from "next/router";
import {toast} from "react-toastify";

const Home: NextPage = () => {
	const auth = useContext(AuthContext)
	const router = useRouter()

	useEffect(() => {
		if (!auth) {
			return
		}

		if (auth.getToken() === null) {
			toast.success('You have successfully logged out')
			router.push('/')
			return
		}

		auth.logout()
	}, [auth])

	return <></>
}

export default Home
