import * as yup from "yup";
import {useMutation, useQueryClient} from "react-query";
import {FieldValues, useFieldArray, useForm} from "react-hook-form";
import {yupResolver} from "@hookform/resolvers/yup";
import axios from "axios";
import {toast} from "react-toastify";
import {toastErr} from "../utils/Err";
import {extractDirtyFields} from "../utils/form";
import {useRouter} from "next/router";
import BtnSpinner from "../utils/BtnSpinner";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPlus, faTrash} from "@fortawesome/free-solid-svg-icons";
import {useContext, useEffect} from "react";
import {Auth, AuthContext} from "../utils/Auth";

const weekdays = ['monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday', 'sunday'];

const schema = yup.object().shape({
	title: yup.string().required().min(3).max(100),
	description: yup.string().required().min(5),
	specialistName: yup.string().notRequired().min(5).max(100).nullable().transform(value => value === '' ? null : value),
	specialistPhone: yup.string().notRequired().matches(/^\+?[1-9]\d{1,14}$/, 'must be a valid phone number').nullable().transform(value => value === '' ? null : value),
	visitDuration: yup.number().required().min(1).transform(value => isNaN(value) ? undefined : value),
	workSchedule: yup.array().of(yup.object().shape({
		// Weekday must not already be present in the array
		weekday: yup.string().required('weekday is a required field').oneOf(weekdays, 'weekday is a required field'),
		from: yup.string().required('from is a required field').matches(/^([01]?[0-9]|2[0-3]):[0-5][0-9]$/, 'must be a valid time'),
		to: yup.string().required('to is a required field').matches(/^([01]?[0-9]|2[0-3]):[0-5][0-9]$/, 'must be a valid time'),
	})).required().min(1).test('unique weekdays', 'work schedule must have unique weekdays', (value) => {
		if (!value) return true;
		const weekdays = value.map(v => v.weekday);
		return weekdays.length === new Set(weekdays).size;
	}),
}).required();

function workscheduleToArr(workSchedule: Map<any, any>) {
	if (!workSchedule) return [];
	return Object
		.entries(workSchedule)
		.map(([weekday, {from, to}]) => ({weekday, from, to}))
		.sort((a, b) => weekdays.indexOf(a.weekday) - weekdays.indexOf(b.weekday));
}

function formatNullable(val: any) {
	if (val.specialistName && typeof val.specialistName.value !== 'undefined') {
		val.specialistName = val.specialistName.value || ''
	}
	if (val.specialistPhone && typeof val.specialistPhone.value !== 'undefined') {
		val.specialistPhone = val.specialistPhone.value || ''
	}
	return val
}

const ServiceForm = ({service}: { service: any }) => {
	const router = useRouter()
	const queryClient = useQueryClient()
	const auth = useContext(AuthContext) as Auth

	const {register, handleSubmit, formState: {errors, isDirty, dirtyFields}, reset, control} = useForm({
		resolver: yupResolver(schema),
		reValidateMode: 'onBlur',
		defaultValues: {
			title: service?.title,
			description: service?.description,
			specialistName: service?.specialistName,
			specialistPhone: service?.specialistPhone,
			visitDuration: service?.visitDuration.toString(),
			workSchedule: workscheduleToArr(service?.workSchedule),
		}
	})
	const {fields, append, remove} = useFieldArray({control, name: "workSchedule"});

	const {mutate, isLoading} = useMutation((data: FieldValues) => {
		const url = service ? `/companies/${service.companyID}/services/${service.id}` : `/companies/${auth.getUserID()}/services`
		const method = service ? 'patch' : 'post'
		return axios({url, method, data})
			.then(() => {
				toast.success(`You have successfully ${service ? 'updated service info' : 'created new service'}`)
				if (service) {
					queryClient.invalidateQueries(['service', service.id])
				}
				router.push('/services')
			}).catch(err => {
				toastErr(err)
				return Promise.reject(err)
			})
	}, {
		onSuccess: (_data, variables) => reset(formatNullable({
			title: service?.title,
			description: service?.description,
			specialistName: service?.specialistName,
			specialistPhone: service?.specialistPhone,
			visitDuration: service?.visitDuration.toString(),
			workSchedule: workscheduleToArr(service?.workSchedule),
			...variables,
		}))
	})
	const onSubmit = (data: FieldValues) => {
		const dirtyData = extractDirtyFields(data, dirtyFields)

		if (service) {
			if (typeof dirtyData.specialistName !== 'undefined') {
				dirtyData.specialistName = {
					value: dirtyData.specialistName,
				}
			}

			if (typeof dirtyData.specialistPhone !== 'undefined') {
				dirtyData.specialistPhone = {
					value: dirtyData.specialistPhone,
				}
			}
		}

		if (dirtyData.workSchedule) {
			dirtyData.workSchedule = dirtyData.workSchedule?.reduce((acc: any, cur: any) => {
				acc[cur.weekday] = {from: cur.from, to: cur.to}
				return acc
			}, {})
		}

		mutate(dirtyData)
	}

	useEffect(() => {
		if (!service && fields.length === 0) {
			append({weekday: '', from: '', to: ''}, {shouldFocus: false})
		}
	}, [])

	return (
		<>
			<div className="row mb-3 text-center">
				<div className="offset-1 col-10 offset-md-2 col-md-8 offset-lg-3 col-lg-6">
					<div className="card mb-4 rounded-3 shadow-sm">
						<div className="card-header py-3">
							<h4 className="my-0 fw-normal">{service ? 'Update service info' : 'Create new service'}</h4>
						</div>
						<form className="card-body" onSubmit={handleSubmit(onSubmit)}>
							<div className="mb-3">
								<label htmlFor="title" className="form-label">Title</label>
								<input {...register('title')} type="text"
									   className={`form-control ${errors.title ? 'is-invalid' : ''}`}
								/>
								{errors.title &&
                                    <div className="invalid-feedback">{errors.title.message as string}</div>}
							</div>
							<div className="mb-3">
								<label htmlFor="description" className="form-label">Description</label>
								<input {...register('description')} type="text"
									   className={`form-control ${errors.description ? 'is-invalid' : ''}`}/>
								{errors.description &&
                                    <div className="invalid-feedback">{errors.description.message as string}</div>}
							</div>
							<div className="mb-3">
								<label htmlFor="specialistName" className="form-label">Specialist name</label>
								<input {...register('specialistName')} type="text"
									   className={`form-control ${errors.specialistName ? 'is-invalid' : ''}`}
									   placeholder="optional"
								/>
								{errors.specialistName &&
                                    <div className="invalid-feedback">{errors.specialistName.message as string}</div>}
							</div>
							<div className="mb-3">
								<label htmlFor="specialistPhone" className="form-label">Specialist phone</label>
								<input {...register('specialistPhone')} type="text"
									   className={`form-control ${errors.specialistPhone ? 'is-invalid' : ''}`}
									   placeholder="optional"
								/>
								{errors.specialistPhone &&
                                    <div className="invalid-feedback">{errors.specialistPhone.message as string}</div>}
							</div>
							<div className="mb-3">
								<label htmlFor="visitDuration" className="form-label">Visit duration</label>
								<input {...register('visitDuration')} type="number"
									   className={`form-control ${errors.visitDuration ? 'is-invalid' : ''}`}
									   placeholder="30"
									   disabled={!!service}
								/>
								{errors.visitDuration &&
                                    <div className="invalid-feedback">{errors.visitDuration.message as string}</div>}
							</div>
							<div className="mb-3 border-top pt-3 mt-4 px-3">
								<div className="mb-3">
									<label htmlFor="workSchedule" className="form-label">Work schedule</label>
									<button type="button" className="btn btn-sm btn-primary ms-2"
											onClick={() => append({weekday: '', from: '', to: ''})}><FontAwesomeIcon icon={faPlus}/>
									</button>
								</div>
								{fields.map((item, index) => (
									<div key={item.id} className="row mb-2">
										<div className={`col-12 mb-3${index > 0 ? ' border-top py-4 mt-3' : ''}`}>
											<select {...register(`workSchedule.${index}.weekday` as const)}
													className={`form-select ${errors.workSchedule?.[index]?.weekday ? 'is-invalid' : ''}`}
													defaultValue="">
												<option disabled={true} value="">Weekday</option>
												{weekdays.map((val, index) => (
													<option key={index} value={val}>{val.charAt(0).toUpperCase() + val.slice(1)}</option>
												))}
											</select>
											{errors.workSchedule?.[index]?.weekday &&
												<div className="invalid-feedback">{errors.workSchedule?.[index]?.weekday.message as string}</div>
											}
										</div>
										<div className="row">
											<div className="col-5">
												<label htmlFor="from" className="form-label">From</label>
											</div>
											<div className="col-5">
												<label htmlFor="to" className="form-label">To</label>
											</div>
										</div>
										<div className="col-5">
											<input {...register(`workSchedule.${index}.from`)} type="time"
												   className={`form-control ${errors.workSchedule?.[index]?.from ? 'is-invalid' : ''}`}
											/>
											{errors.workSchedule?.[index]?.from &&
                                                <div
                                                    className="invalid-feedback">{errors.workSchedule?.[index]?.from.message as string}</div>}
										</div>
										<div className="col-5">
											<input {...register(`workSchedule.${index}.to`)} type="time"
												   className={`form-control ${errors.workSchedule?.[index]?.to ? 'is-invalid' : ''}`}
											/>
											{errors.workSchedule?.[index]?.to &&
                                                <div
                                                    className="invalid-feedback">{errors.workSchedule?.[index]?.to.message as string}</div>}
										</div>
										<div className="col-2">
											<button type="button" className="btn btn-danger"
													onClick={() => remove(index)} disabled={fields.length === 1}>
												<FontAwesomeIcon icon={faTrash}/>
											</button>
										</div>
									</div>
								))}
								{errors.workSchedule &&
                                    <div className="invalid-feedback d-block">{errors.workSchedule.message as string}</div>
								}
							</div>
							<div className="">
								{service && (
									<button type="submit" className="btn btn-primary w-100"
											disabled={isLoading || !isDirty}>
										{isLoading ? <BtnSpinner/> : 'Update'}
									</button>
								)}
								{!service && (
									<button type="submit" className="btn btn-primary w-100" disabled={isLoading}>
										{isLoading ? <BtnSpinner/> : 'Create'}
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

export default ServiceForm;
