import Link from "next/link";

const NotFound = () => {
	return (
		<div className="header p-3 pb-md-4 mx-auto text-center">
			<h1 className="display-4 fw-normal">404 page not found</h1>
			<p className="fs-5 text-muted">This page was not found. <Link href="/"><a className="link-primary">Return to main page</a></Link></p>
		</div>
	)
}

export default NotFound