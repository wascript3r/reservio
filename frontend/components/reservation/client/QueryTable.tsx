import axios from 'axios'
import {Table} from 'components/reservation/client/Table'
import {Err} from 'components/utils/Err'
import Spinner from 'components/utils/Spinner'
import {useQuery} from 'react-query'

export const QueryTable = ({clientID}: { clientID: string }) => {
	const {data, error, isLoading} = useQuery<any, Error>('client_reservations', () => {
		return axios
			.get('/clients/' + clientID + '/reservations')
			.then(res => res.data)
	})

	if (isLoading) {
		return <Spinner/>
	} else if (error) {
		return <Err msg={error.message}/>
	}

	return (
		<Table reservations={data?.data.reservations}/>
	)
}
