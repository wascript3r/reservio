import {FieldValues, useForm} from "react-hook-form";
import {useMutation} from "react-query";
import axios from "axios";
import {useContext} from "react";
import {Auth, AuthContext} from "../utils/Auth";
import {useRouter} from "next/router";
import BtnSpinner from "../utils/BtnSpinner";
import {toast} from "react-toastify";
import {toastErr} from "../utils/Err";
import * as yup from "yup";
import {yupResolver} from "@hookform/resolvers/yup";

const schema = yup.object().shape({
	email: yup.string().required(),
	password: yup.string().required().min(8),
}).required();

const LoginForm = () => {
	const auth = useContext(AuthContext) as Auth
	const router = useRouter()

	const {register, handleSubmit, formState: {errors}} = useForm({
		resolver: yupResolver(schema),
		reValidateMode: 'onBlur'
	})
	const {mutate, isLoading} = useMutation((data: FieldValues) => {
		return axios.post('/users/authenticate', data)
			.then(res => res.data.data)
			.then(data => {
				auth.setToken(data.accessToken)
				auth.setRefreshToken(data.refreshToken)
				toast.success('You have logged in successfully')
				router.push('/')
			}).catch(err => toastErr(err))
	})
	const onSubmit = (data: FieldValues) => mutate(data)

	return (
		<>
			<div className="row mb-3 text-center">
				<div className="offset-1 col-10 offset-md-2 col-md-8 offset-lg-3 col-lg-6">
					<div className="card mb-4 rounded-3 shadow-sm">
						<div className="card-header py-3">
							<h4 className="my-0 fw-normal">Login</h4>
						</div>
						<form className="card-body" onSubmit={handleSubmit(onSubmit)}>
							<div className="mb-3">
								<label htmlFor="email" className="form-label">Email address</label>
								<input {...register('email')} type="email"
									   className={`form-control ${errors.email ? 'is-invalid' : ''}`}
									   placeholder="name@example.com"/>
								{errors.email &&
                                    <div className="invalid-feedback">{errors.email.message as string}</div>}
							</div>
							<div className="mb-3">
								<label htmlFor="password" className="form-label">Password</label>
								<input {...register('password')} type="password"
									   className={`form-control ${errors.password ? 'is-invalid' : ''}`}/>
								{errors.password &&
                                    <div className="invalid-feedback">{errors.password.message as string}</div>}
							</div>
							<div className="">
								<button type="submit" className="btn btn-primary w-100" disabled={isLoading}>
									{isLoading ? <BtnSpinner/> : 'Login'}
								</button>
							</div>
						</form>
					</div>
				</div>
			</div>
		</>
	);
}

export default LoginForm;