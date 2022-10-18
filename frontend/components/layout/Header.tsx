import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faCalendarCheck} from "@fortawesome/free-solid-svg-icons";
import Link from "next/link";

const Header = () => {
	const links = [
		{
			href: "/login",
			text: "Login",
		},
		{
			href: "/registration",
			text: "Registration",
		}
	];

	// Build the links
	const linkElements = links.map((link, index) => {
		const margin = (index === links.length - 1) ? "" : "me-3 "
		const className = `${margin}py-2 text-dark text-decoration-none link-primary`
		return (
			<Link href={link.href} key={index}>
				<a className={className}>{link.text}</a>
			</Link>
		)
	})

	return (
		<header>
			<div className="d-flex flex-column flex-md-row align-items-center pb-3 mb-4 border-bottom nav">
				<a href="/" className="d-flex align-items-center text-dark text-decoration-none">
					<FontAwesomeIcon icon={faCalendarCheck} size="xl"/>
					<span className="ms-2 fs-4">Reservio</span>
				</a>
				<nav className="d-inline-flex mt-2 mt-md-0 ms-md-auto">
					{linkElements}
				</nav>
			</div>
		</header>
	)
}

export default Header