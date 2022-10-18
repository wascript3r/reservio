import type {NextPage} from 'next'
import CompanyInfo from "components/company/CompanyInfo";
import {useRouter} from "next/router";

const Home: NextPage = () => {
	const router = useRouter()
	const {id} = router.query

	return (
		<CompanyInfo id={id as string}/>
	)
}

export default Home
