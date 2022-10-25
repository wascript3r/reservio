import {FieldValues, useForm} from "react-hook-form";
import {useMutation} from "react-query";
import axios from "axios";
import {useRouter} from "next/router";
import BtnSpinner from "../utils/BtnSpinner";
import {toast} from "react-toastify";
import {toastErr} from "../utils/Err";
import * as yup from "yup";
import {yupResolver} from "@hookform/resolvers/yup";
import Link from "next/link";

const schema = yup.object().shape({
	email: yup.string().required(),
	password: yup.string().required().min(8),
	firstName: yup.string().required().min(3),
	lastName: yup.string().required().min(3),
	phone: yup.string().required().matches(/^\+?[1-9]\d{1,14}$/, 'must be a valid phone number'),
}).required();

const ClientRegForm = () => {
	const router = useRouter()

	const {register, handleSubmit, formState: {errors}} = useForm({
		resolver: yupResolver(schema),
		reValidateMode: 'onBlur'
	})
	const {mutate, isLoading} = useMutation((data: FieldValues) => {
		return axios.post('/clients', data)
			.then(() => {
				toast.success('You have successfully registered')
				router.push('/login')
			}).catch(err => toastErr(err))
	})
	const onSubmit = (data: FieldValues) => mutate(data)

	return (
		<>
			<div className="row mb-3 text-center">
				<div className="offset-1 col-10 offset-md-2 col-md-8 offset-lg-3 col-lg-6">
					<div className="card mb-4 rounded-3 shadow-sm">
						<div className="card-header py-3">
							<h4 className="my-0 fw-normal">Client registration</h4>
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
							<div className="mb-3">
								<label htmlFor="firstName" className="form-label">First name</label>
								<input {...register('firstName')} type="text"
									   className={`form-control ${errors.firstName ? 'is-invalid' : ''}`}/>
								{errors.firstName &&
                                    <div className="invalid-feedback">{errors.firstName.message as string}</div>}
							</div>
							<div className="mb-3">
								<label htmlFor="lastName" className="form-label">Last name</label>
								<input {...register('lastName')} type="text"
									   className={`form-control ${errors.lastName ? 'is-invalid' : ''}`}/>
								{errors.lastName &&
                                    <div className="invalid-feedback">{errors.lastName.message as string}</div>}
							</div>
							<div className="mb-3">
								<label htmlFor="phone" className="form-label">Phone number</label>
								<input {...register('phone')} type="text"
									   className={`form-control ${errors.phone ? 'is-invalid' : ''}`}/>
								{errors.phone &&
                                    <div className="invalid-feedback">{errors.phone.message as string}</div>}
							</div>
							<div className="mb-2">
								<button type="submit" className="btn btn-primary w-100" disabled={isLoading}>
									{isLoading ? <BtnSpinner/> : 'Register'}
								</button>
							</div>
							<div>
								<Link href="/registration/company">
									<a className="w-100">Register as company instead</a>
								</Link>
							</div>
						</form>
					</div>
				</div>
			</div>
		</>
	);
}

export default ClientRegForm;