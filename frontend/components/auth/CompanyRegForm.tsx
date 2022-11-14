import {FieldValues, useForm} from "react-hook-form";
import {useMutation, useQueryClient} from "react-query";
import axios from "axios";
import {useRouter} from "next/router";
import BtnSpinner from "../utils/BtnSpinner";
import {toast} from "react-toastify";
import {toastErr} from "../utils/Err";
import * as yup from "yup";
import {yupResolver} from "@hookform/resolvers/yup";
import {extractDirtyFields} from "../utils/form";

const baseSchema = yup.object().shape({
	name: yup.string().required().min(2),
	address: yup.string().required().min(5),
	description: yup.string().required().min(5),
}).required();

const registerSchema = baseSchema.shape({
	email: yup.string().email().required(),
	password: yup.string().required().min(8),
}).required();

const CompanyRegForm = ({company}: { company: any }) => {
	const router = useRouter()
	const queryClient = useQueryClient()

	const {register, handleSubmit, formState: {errors, isDirty, dirtyFields}, reset} = useForm({
		resolver: yupResolver(company ? baseSchema : registerSchema),
		reValidateMode: 'onBlur',
		defaultValues: {
			email: company?.email,
			password: company?.password,
			name: company?.name,
			address: company?.address,
			description: company?.description,
		}
	})
	const {mutate, isLoading} = useMutation((data: FieldValues) => {
		const url = company ? `/companies/${company.id}` : '/companies'
		const method = company ? 'patch' : 'post'
		return axios({url, method, data})
			.then(() => {
				toast.success(`You have successfully ${company ? 'updated company info' : 'registered'}`)
				if (company) {
					queryClient.invalidateQueries(['company', company.id])
				} else {
					router.push('/login')
				}
			}).catch(err => toastErr(err))
	}, {
		onSuccess: (_data, variables) => reset({
			email: company?.email,
			password: company?.password,
			name: company?.name,
			address: company?.address,
			description: company?.description,
			...variables,
		})
	})
	const onSubmit = (data: FieldValues) => {
		mutate(extractDirtyFields(data, dirtyFields))
	}

	return (
		<>
			<div className="row mb-3 text-center">
				<div className="offset-1 col-10 offset-md-2 col-md-8 offset-lg-3 col-lg-6">
					<div className="card mb-4 rounded-3 shadow-sm">
						<div className="card-header py-3">
							<h4 className="my-0 fw-normal">{company ? 'Update company info' : 'Company registration'}</h4>
						</div>
						<form className="card-body" onSubmit={handleSubmit(onSubmit)}>
							{!company && (
								<>
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
								</>
							)}
							<div className="mb-3">
								<label htmlFor="name" className="form-label">Company name</label>
								<input {...register('name')} type="text"
									   className={`form-control ${errors.name ? 'is-invalid' : ''}`}
								/>
								{errors.name &&
                                    <div className="invalid-feedback">{errors.name.message as string}</div>}
							</div>
							<div className="mb-3">
								<label htmlFor="address" className="form-label">Company address</label>
								<input {...register('address')} type="text"
									   className={`form-control ${errors.address ? 'is-invalid' : ''}`}
								/>
								{errors.address &&
                                    <div className="invalid-feedback">{errors.address.message as string}</div>}
							</div>
							<div className="mb-3">
								<label htmlFor="description" className="form-label">Company description</label>
								<textarea {...register('description')}
										  className={`form-control ${errors.description ? 'is-invalid' : ''}`}
								/>
								{errors.description &&
                                    <div className="invalid-feedback">{errors.description.message as string}</div>}
							</div>
							<div className="">
								{company && (
									<button type="submit" className="btn btn-primary w-100"
											disabled={isLoading || !isDirty}>
										{isLoading ? <BtnSpinner/> : 'Update'}
									</button>
								)}
								{!company && (
									<button type="submit" className="btn btn-primary w-100" disabled={isLoading}>
										{isLoading ? <BtnSpinner/> : 'Register'}
									</button>
								)}
							</div>
						</form>
					</div>
				</div>
			</div>
		</>
	);
}

export default CompanyRegForm;
