import axios from "axios";
import React from "react";

enum Role {
	CLIENT = 'client',
	COMPANY = 'company',
	ADMIN = 'admin',
}

class Auth {
	private token: string | null
	private refreshToken: string | null
	private userID: string | null
	private role: Role | null

	constructor() {
		this.token = null
		this.refreshToken = null
		this.userID = null
		this.role = null

		const token = localStorage.getItem('accessToken')
		const refreshToken = localStorage.getItem('refreshToken')

		if (token && refreshToken) {
			this.setToken(token)
			this.setRefreshToken(refreshToken)
		}
	}

	private parseToken() {
		if (!this.token) {
			return null
		}
		const arr = this.token.split('.')
		if (arr.length !== 3) {
			return
		}
		const meta = Buffer.from(arr[1], 'base64').toString()
		const {userID, role} = JSON.parse(meta)
		this.userID = userID
		this.role = role
	}

	public getToken(): string | null {
		return this.token
	}

	public getRefreshToken(): string | null {
		return this.refreshToken
	}

	public setToken(token: string): void {
		this.token = token
		this.parseToken()
		localStorage.setItem('accessToken', token)
		axios.defaults.headers.common.Authorization = `Bearer ${token}`
	}

	public logout() {
		this.token = null
		this.refreshToken = null
		this.userID = null
		this.role = null
		localStorage.removeItem('accessToken')
		localStorage.removeItem('refreshToken')
		delete axios.defaults.headers.common.Authorization
	}

	public setRefreshToken(refreshToken: string): void {
		this.refreshToken = refreshToken
		localStorage.setItem('refreshToken', refreshToken)
	}

	public hasAccess(role: Role) {
		return this.role === role
	}

	public isAuth(): boolean {
		return !!this.token
	}

	public getUserID(): string | null {
		return this.userID
	}
}

const AuthContext = React.createContext<Auth | null>(null)

export {Role, Auth, AuthContext}