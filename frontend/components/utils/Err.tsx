import {toast} from "react-toastify";

const Err = ({msg}: {msg: string}) => {
	return (
		<div className="alert alert-danger" role="alert">
			{msg}
		</div>
	)
}

const toastErr = (err: any) => {
	if (err.response.data && err.response.data.error && err.response.data.error.message) {
		toast.error(err.response.data.error.message)
	} else {
		toast.error(err.message)
	}
}

export { Err, toastErr }