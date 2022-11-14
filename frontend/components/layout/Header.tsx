import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faCalendarCheck} from "@fortawesome/free-solid-svg-icons";
import Link from "next/link";
import {Auth, AuthContext, Role} from "../utils/Auth";
import {useContext} from "react";

const Header = () => {
	const auth = useContext(AuthContext) as Auth

	return (
		<header>
			<div className="d-flex flex-column flex-md-row align-items-center pb-3 mb-4 border-bottom nav">
				<Link href="/">
					<a className="d-flex align-items-center text-dark text-decoration-none">
						<FontAwesomeIcon icon={faCalendarCheck} size="xl"/>
						<span className="ms-2 fs-4">Reservio</span>
					</a>
				</Link>
				<nav className="d-inline-flex mt-2 mt-md-0 ms-md-auto">
					<Link href="/">
						<a className="ms-3 py-2 text-dark text-decoration-none link-primary">Home</a>
					</Link>
					{!auth.loggedIn() && (
						<>
							<Link href="/login">
								<a className="ms-3 py-2 text-dark text-decoration-none link-primary">Login</a>
							</Link>
							<Link href="/registration/client">
								<a className="ms-3 py-2 text-dark text-decoration-none link-primary">Registration</a>
							</Link>
						</>
					)}
					{auth.hasAccess(Role.COMPANY) && (
						<>
							<Link href="/companies/edit">
								<a className="ms-3 py-2 text-dark text-decoration-none link-primary">Edit info</a>
							</Link>
							<Link href="/services">
								<a className="ms-3 py-2 text-dark text-decoration-none link-primary">My services</a>
							</Link>
						</>
					)}
					{auth.loggedIn() && (
						<Link href="/logout">
							<a className="ms-3 py-2 text-dark text-decoration-none link-primary">Logout</a>
						</Link>
					)}
				</nav>
			</div>
		</header>
	)
}

export default Header
