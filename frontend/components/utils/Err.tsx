const Err = ({msg}: {msg: string}) => {
	return (
		<div className="alert alert-danger" role="alert">
			{msg}
		</div>
	)
}

export default Err