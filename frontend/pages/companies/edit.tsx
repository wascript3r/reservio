import {useCompanyInfo} from 'components/company/CompanyInfo'
import {Auth, AuthContext, Role} from 'components/utils/Auth'
import {Err} from 'components/utils/Err'
import Spinner from 'components/utils/Spinner'
import type {NextPage} from 'next'
import {useRouter} from 'next/router'
import {useContext} from 'react'
import CompanyRegForm from '../../components/auth/CompanyRegForm'

const Home: NextPage = () => {
	const auth = useContext(AuthContext) as Auth
	const router = useRouter()

	const {data: company, error, isLoading} = useCompanyInfo(router, auth.getUserID() as string)

	if (!auth.hasAccess(Role.COMPANY)) {
		router.push('/403')
		return <></>
	}

	if (isLoading) {
		return <Spinner/>
	} else if (error) {
		return <Err msg={error.message}/>
	}

	return (
		<CompanyRegForm company={company}/>
	)
}

export default Home
