import {QueryTable} from 'components/reservation/client/QueryTable'
import {QueryTables} from 'components/reservation/company/QueryTables'
import {Auth, AuthContext, Role} from 'components/utils/Auth'
import type {NextPage} from 'next'
import {useRouter} from 'next/router'
import {useContext} from 'react'

const Home: NextPage = () => {
	const auth = useContext(AuthContext) as Auth
	const router = useRouter()

	if (!auth.hasAccess(Role.CLIENT) && !auth.hasAccess(Role.COMPANY)) {
		router.push('/403')
		return <></>
	}

	if (auth.hasAccess(Role.CLIENT)) {
		return (
			<QueryTable clientID={auth.getUserID() as string}/>
		)
	} else {
		return (
			<QueryTables companyID={auth.getUserID() as string}/>
		)
	}
}

export default Home
