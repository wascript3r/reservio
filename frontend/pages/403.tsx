import Link from "next/link";

const NotFound = () => {
	return (
		<div className="header p-3 pb-md-4 mx-auto text-center">
			<h1 className="display-6 fw-normal">403 forbidden</h1>
			<p className="mt-4 fs-5 text-muted">You do not have permissions to view this page. <Link href="/"><a className="link-primary">Return to main page</a></Link></p>
		</div>
	)
}

export default NotFound