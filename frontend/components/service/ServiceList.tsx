import Spinner from "../utils/Spinner";
import {Err} from "../utils/Err";
import ServiceInfo from "../service/ServiceInfo";
import {useServices} from "../company/CompanyInfo";
import Link from "next/link";

const ServiceList = ({id}: { id: string }) => {
	const {data: services, error, isLoading: isServicesLoading} = useServices(id)

	if (isServicesLoading) {
		return <Spinner/>
	}

	if (error) {
		return <Err msg={error.message}/>
	}

	return (
		<div className="col">
			<h2 className="text-center mt-3 mb-1">My services</h2>
			<div className="text-center mb-5">
				<Link href={`/services/new`}>
					<button type="button"
							className={`btn btn-outline-primary mt-3`}>Create new
					</button>
				</Link>
			</div>
			<div className="row row-cols-1 row-cols-md-2 mb-3 text-center">
				{services?.data.services.length > 0 && services?.data.services.map((service: any, index: number) => (
					<ServiceInfo service={service} key={index}/>
				))}
				{services?.data.services.length === 0 &&
                    <div className="w-100 text-muted text-center">
                        You don&apos;t have any services yet
                    </div>
				}
			</div>
		</div>
	)
}

export default ServiceList
