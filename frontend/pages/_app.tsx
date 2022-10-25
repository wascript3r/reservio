import "styles/globals.css"
import "bootstrap/dist/css/bootstrap.min.css"
import type {AppProps} from 'next/app'
import Header from "../components/layout/Header";
import {Container} from "react-bootstrap";
import styles from "styles/Home.module.css"
import Footer from "../components/layout/Footer";
import axios from "axios";
import {QueryClient, QueryClientProvider} from 'react-query'
import React, {useEffect, useState} from "react";
import {Auth, AuthContext} from "../components/utils/Auth";
import {toast, ToastContainer} from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import {useRouter} from "next/router";

axios.defaults.baseURL = "https://reservio.hs.vc/api/v1";

const queryClient = new QueryClient()

function MyApp({Component, pageProps}: AppProps) {
	const [auth, setAuth] = useState<Auth | null>(null)
	const router = useRouter()

	useEffect(() => {
		const auth = new Auth()
		setAuth(auth)

		console.log('Adding interceptors')
		const id = axios.interceptors.response.use(
			res => res,
			err => {
				if (err.config.url.includes('/tokens')) {
					console.log('Failed to refresh token, logging out')
					auth.logout()
					toast.error('Your session has expired, please log in again')
					router.push('/login')
				} else if (err.response.status === 401 && err.response.data?.error?.name === 'token_invalid_or_expired') {
					const refreshToken = auth.getRefreshToken()
					if (refreshToken) {
						console.log('Refreshing token')
						return axios.post('/tokens', {refreshToken}).then(res => {
							console.log('Token refreshed')
							auth.setToken(res.data.data.accessToken)
							return Promise.reject(err)
						})
					}
				}

				return Promise.reject(err)
			}
		)

		return () => {
			console.log('Removing interceptors')
			axios.interceptors.response.eject(id)
		}
	}, [])

	if (!auth) {
		return <></>
	}

	return (
		<>
			<ToastContainer/>
			<QueryClientProvider client={queryClient}>
				<AuthContext.Provider value={auth}>
					<Container className={`py-3 ${styles.container}`}>
						<Header/>
						<Component {...pageProps} />
						<Footer/>
					</Container>
				</AuthContext.Provider>
			</QueryClientProvider>
		</>
	)
}

export default MyApp