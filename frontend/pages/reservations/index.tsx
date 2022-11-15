import {QueryTable} from 'components/reservation/client/QueryTable'
import {Auth, AuthContext, Role} from 'components/utils/Auth'
import type {NextPage} from 'next'
import {useRouter} from 'next/router'
import {useContext} from 'react'

const Home: NextPage = () => {
	const auth = useContext(AuthContext) as Auth
	const router = useRouter()

	if (!auth.hasAccess(Role.CLIENT)) {
		router.push('/403')
		return <></>
	}

	return (
		<QueryTable clientID={auth.getUserID() as string}/>
	)
}

export default Home
