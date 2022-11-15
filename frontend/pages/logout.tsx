import {Auth, AuthContext} from 'components/utils/Auth'
import type {NextPage} from 'next'
import {useRouter} from 'next/router'
import {useContext, useEffect} from 'react'
import {toast} from 'react-toastify'

const Home: NextPage = () => {
	const auth = useContext(AuthContext) as Auth
	const router = useRouter()

	useEffect(() => {
		if (!auth) {
			return
		}

		auth.logout()
		toast.success('You have successfully logged out')
		router.push('/')

	}, [])

	return <></>
}

export default Home
