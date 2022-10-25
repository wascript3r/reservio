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

		auth.logout()
		toast.success('You have successfully logged out')
		router.push('/')

	}, [auth, router])

	return <></>
}

export default Home
