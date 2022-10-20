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

axios.defaults.baseURL = "https://reservio.hs.vc/api/v1";
const queryClient = new QueryClient()

function MyApp({Component, pageProps}: AppProps) {
	const [auth, setAuth] = useState<Auth | null>(null)

	useEffect(() => {
		setAuth(new Auth())
	}, [])

	return (
		<QueryClientProvider client={queryClient}>
			<AuthContext.Provider value={auth}>
				<Container className={`py-3 ${styles.container}`}>
					<Header/>
					<Component {...pageProps} />
					<Footer/>
				</Container>
			</AuthContext.Provider>
		</QueryClientProvider>
	)
}

export default MyApp
