import {faCalendarCheck} from '@fortawesome/free-solid-svg-icons'
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'
import Link from 'next/link'
import {useRouter} from 'next/router'
import {useCallback, useContext} from 'react'
import {Nav, Navbar} from 'react-bootstrap'
import {Auth, AuthContext, Role} from '../utils/Auth'

const Header = () => {
	const auth = useContext(AuthContext) as Auth
	const router = useRouter()

	const activeClass = useCallback((path: string) => {
		return router.pathname === path ? '' : ''
	}, [router.pathname])

	return (
		<Navbar collapseOnSelect expand="lg" variant="light" className="border-bottom mb-5">
			<Navbar.Brand>
				<Link href="/">
					<a className="d-flex align-items-center text-dark text-decoration-none">
						<FontAwesomeIcon icon={faCalendarCheck} size="xl"/>
						<span className="ms-2 fs-4">Reservio</span>
					</a>
				</Link>
			</Navbar.Brand>
			<Navbar.Toggle aria-controls="responsive-navbar-nav"/>
			<Navbar.Collapse id="responsive-navbar-nav">
				<Nav className="ms-auto text-center">
					<Link href="/">
						<a className={`ms-3 py-2 text-dark text-decoration-none link-primary ${activeClass('/')}`}>Home</a>
					</Link>
					{!auth.loggedIn() && (
						<>
							<Link href="/login">
								<a className={`ms-3 py-2 text-dark text-decoration-none link-primary ${activeClass('/login')}`}>Login</a>
							</Link>
							<Link href="/registration/client">
								<a className={`ms-3 py-2 text-dark text-decoration-none link-primary ${activeClass('/registration/client')}${activeClass('/registration/company')}`}>Registration</a>
							</Link>
						</>
					)}
					{auth.hasAccess(Role.COMPANY) && (
						<>
							<Link href="/companies/edit">
								<a className={`ms-3 py-2 text-dark text-decoration-none link-primary ${activeClass('/companies/edit')}`}>Edit info</a>
							</Link>
							<Link href="/services">
								<a className={`ms-3 py-2 text-dark text-decoration-none link-primary ${activeClass('/services')}`}>My services</a>
							</Link>
							<Link href="/reservations">
								<a className={`ms-3 py-2 text-dark text-decoration-none link-primary ${activeClass('/reservations')}`}>Reservations</a>
							</Link>
						</>
					)}
					{auth.hasAccess(Role.CLIENT) && (
						<>
							<Link href="/reservations">
								<a className={`ms-3 py-2 text-dark text-decoration-none link-primary ${activeClass('/reservations')}`}>My reservations</a>
							</Link>
						</>
					)}
					{auth.loggedIn() && (
						<Link href="/logout">
							<a className={`ms-3 py-2 text-dark text-decoration-none link-primary`}>Logout</a>
						</Link>
					)}
				</Nav>
			</Navbar.Collapse>
		</Navbar>
	)
}

export default Header
