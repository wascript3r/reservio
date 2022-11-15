import axios from 'axios'
import {useServices} from 'components/company/CompanyInfo'
import {ServiceReservations, Tables} from 'components/reservation/company/Tables'
import {Err} from 'components/utils/Err'
import Spinner from 'components/utils/Spinner'
import {useMemo} from 'react'
import {useQueries} from 'react-query'

export const QueryTables = ({companyID}: { companyID: string }) => {
	const {
		data: services,
		error: serror,
		isLoading: isServicesLoading,
		isFetched: isServicesFetched,
	} = useServices(companyID)

	const results = useQueries(
		services?.data.services.map((service: any) => ({
			queryKey: ['service', service.id, 'reservations'],
			queryFn: () => {
				return axios
					.get(`/companies/${companyID}/services/${service.id}/reservations`)
					.then(res => ({
						serviceID: service.id,
						data: res.data,
					}))
			},
			enabled: isServicesFetched,
		})) || [],
	)
	const isLoading = isServicesLoading || results.some(result => result.isLoading)
	const error = serror || results.find(result => result.error)?.error

	const serviceReservations = useMemo<ServiceReservations[]>(() => {
		if (!results) return []

		return results.map((result: any) => {
			const service = services?.data.services.find((service: any) => service.id === result.data?.serviceID)
			return {
				title: service?.title,
				reservations: result.data?.data.data.reservations,
			}
		})
	}, [results])

	if (isLoading) {
		return <Spinner/>
	} else if (error) {
		return <Err msg={error.message}/>
	}

	return (
		<Tables reservations={serviceReservations}/>
	)
}
