import {FieldValues, useForm} from "react-hook-form";
import {useMutation} from "react-query";
import axios from "axios";
import {useRouter} from "next/router";
import BtnSpinner from "../utils/BtnSpinner";
import {toast} from "react-toastify";
import {toastErr} from "../utils/Err";
import * as yup from "yup";
import {yupResolver} from "@hookform/resolvers/yup";

const schema = yup.object().shape({
	email: yup.string().required(),
	password: yup.string().required().min(8),
	name: yup.string().required().min(3),
	address: yup.string().required().min(5),
	description: yup.string().required().min(5),
}).required();

const CompanyRegForm = () => {
	const router = useRouter()

	const {register, handleSubmit, formState: {errors}} = useForm({
		resolver: yupResolver(schema),
		reValidateMode: 'onBlur'
	})
	const {mutate, isLoading} = useMutation((data: FieldValues) => {
		return axios.post('/companies', data)
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
							<h4 className="my-0 fw-normal">Company registration</h4>
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
								<label htmlFor="name" className="form-label">Company name</label>
								<input {...register('name')} type="text"
									   className={`form-control ${errors.name ? 'is-invalid' : ''}`}/>
								{errors.name &&
                                    <div className="invalid-feedback">{errors.name.message as string}</div>}
							</div>
							<div className="mb-3">
								<label htmlFor="address" className="form-label">Company address</label>
								<input {...register('address')} type="text"
									   className={`form-control ${errors.address ? 'is-invalid' : ''}`}/>
								{errors.address &&
                                    <div className="invalid-feedback">{errors.address.message as string}</div>}
							</div>
							<div className="mb-3">
								<label htmlFor="description" className="form-label">Company description</label>
								<textarea {...register('description')}
										  className={`form-control ${errors.description ? 'is-invalid' : ''}`}/>
								{errors.description &&
                                    <div className="invalid-feedback">{errors.description.message as string}</div>}
							</div>
							<div className="">
								<button type="submit" className="btn btn-primary w-100" disabled={isLoading}>
									{isLoading ? <BtnSpinner/> : 'Register'}
								</button>
							</div>
						</form>
					</div>
				</div>
			</div>
		</>
	);
}

export default CompanyRegForm;