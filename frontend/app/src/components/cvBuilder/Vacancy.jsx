import { object, string, number, date } from 'yup';
import { useEffect, useState } from 'react';
import { Formik } from 'formik';

let vacancySchema = object().shape({
    id: number().positive().integer(),
    name: string().required(),
    company_id: number().positive().integer().required(),
    description: string().required(),
    vacancy_link: string(),
    published_at: date().required(),
    experience: number().positive().integer().required(),
});


let vacancySkillSchema = object().shape({
    vacancy_id: number().positive().integer().required(),
    skill_id: number().positive().integer().required(),
    priority: number().positive().integer().required(),
});

let vacancyDomainSchema = object().shape({
    vacancy_id: number().positive().integer().required(),
    domain_id: number().positive().integer().required(),
    priority: number().positive().integer().required(),
});

async function fetchVacancies() {
    let response = await fetch('http://0.0.0.0:3001/v1/vacancy/');
    let vacancies = await response.json()
    if (vacancies == null) {
        return []
    }
    let validatedObjects = await Promise.all(vacancies.map(async (object) => {
        await vacancySchema.validate(object, { abortEarly: false });
        return object;
    }));
    return await validatedObjects
}

function CreateVacancyForm({ refresh, setRefresh }) {
    async function createVacancy(data) {
        console.log(data)
        try {
            let res = await vacancySchema.validate(data, { abortEarly: false });
            
            await fetch(
                'http://0.0.0.0:3001/v1/vacancy/',
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(res)
                }
            )
            setRefresh(!refresh)
        } catch (errors) {
            return {
                success: false
            }
        }
        return {
            success: true
        }
    }

    return <Formik
        initialValues={{
            name: 'Senior Python Developer',
            company_id: 1,
            description: 'We\'re looking for a senior Python developer to join our team.',
            vacancy_link: 'https://innowise.com/devops-services/',
            published_at: new Date(),
            experience: 5,
        }}
        validationSchema={vacancySchema}
        onSubmit={async (values, { setSubmitting, setErrors }) => {

            let data = await createVacancy(values)
            if (data.success == true) {
                setSubmitting(false);
            } else {
                setErrors({ 'name': 'error!' })
            }

        }}>
        {({
            values,
            errors,
            touched,
            handleChange,
            handleBlur,
            handleSubmit,
            isSubmitting,
        }) => (
            <form onSubmit={handleSubmit} className='frm loginform'>
                <h3>Create Vacancy</h3>
                <input
                    type="text"
                    name="name"
                    onChange={handleChange}
                    onBlur={handleBlur}
                    value={values.name}
                /> <br />
                <span className='errors'>{errors.name && touched.name && errors.name}</span> <br />
                <input type="text" name="description" onChange={handleChange} onBlur={handleBlur} value={values.description} /> <br />
                <span className='errors'>{errors.description && touched.description && errors.description}</span> <br />
                <input type="number" name="experience" onChange={handleChange} onBlur={handleBlur} value={values.experience} /> <br />
                <span className='errors'>{errors.experience && touched.experience && errors.experience}</span> <br />
                <input type="text" name="vacancy_link" onChange={handleChange} onBlur={handleBlur} value={values.vacancy_link} /> <br />
                <span className='errors'>{errors.vacancy_link && touched.vacancy_link && errors.vacancy_link}</span> <br />
                <input type="date" name="published_at" onChange={handleChange} onBlur={handleBlur} value={values.published_at} /> <br />
                <span className='errors'>{errors.published_at && touched.published_at && errors.published_at}</span> <br />
                <input type="number" name="company_id" onChange={handleChange} onBlur={handleBlur} value={values.company_id} /> <br />
                <span className='errors'>{errors.company_id && touched.company_id && errors.company_id}</span> <br />
                <button type="submit" disabled={isSubmitting}>
                    Create
                </button>
            </form>
        )}
    </Formik>
}


export function VacancySkill(props){
    const [vacancySkills, setVacancySkills] = useState([])
    async function fetchSkills() {
        let response = await fetch('http://0.0.0.0:3001/v1/vacancy/'+ props.vacancy_id +'/skills')
        let skills = await response.json()
        if (skills == null) {
            return []
        }
        let validatedObjects = await Promise.all(skills.map(async (object) => {
            await vacancySkillSchema.validate(object, { abortEarly: false });
            return object;
        }));
        return validatedObjects;
    }

    async function setUp(){
        setVacancySkills(await fetchSkills())
    }
    useEffect(() => {
        setUp()
    }, [])

    async function createVacancySkill(data) {
        try {
            let res = await vacancySkillSchema.validate(data, { abortEarly: false });
            
            await fetch(
                'http://0.0.0.0:3001/v1/vacancy/skill',
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(res)
                }
            )
            setUp()
        } catch (errors) {
            return {
                success: false
            }
        }
        return {
            success: true
        }
    }

    async function deleteVacancySkill(data) {
        await fetch(
            'http://0.0.0.0:3001/v1/vacancy/skill',
            {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            }
        )
        setUp()
    }

    return <div>
        <h1>Skills</h1>
    </div>
}
export function VacancyDomain(props){
    const [vacancyDomain, setVacancyDomain] = useState([])
    async function fetchVacancyDomain(){
        let response = await fetch(
            'http://0.0.0.0:3001/v1/vacancy/' + props.vacancy_id + '/domains',
            {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                },
            }
        )
        let vacancydomain = await response.json()
        if (vacancydomain === null) {
            return []
        }
        let validatedObjects = await Promise.all(vacancydomain.map(async (object) => {
            await vacancyDomainSchema.validate(object, { abortEarly: false });
            return object;
        }));
        return validatedObjects;
    }

    return <div>
        <h1>Domains</h1>
        <div className='cv_instances popup_instance'>
            <table>
                <thead>
                    <tr>
                        <th>Vacancy</th>
                        <th>Skill</th>
                        <th>Priority</th>
                    </tr>
                </thead>
                <tbody>
                    {vacancySkills.map((vacancySkill, index) => (
                        <tr key={"vacancy_skill_table_item_"+index}>
                            <td>{vacancySkill.vacancy_id}</td>
                            <td>{vacancySkill.skill_id}</td>
                            <td>{vacancySkill.priority}</td>
                            <td><button onClick={() => deleteVacancySkill(vacancySkill)}>Delete</button></td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
        <Formik
            initialValues={{
                skill_id: 1,
                vacancy_id: 2,
                priority: 1,
            }}
            validationSchema={vacancySkillSchema}
            onSubmit={async (values, { setSubmitting, setErrors }) => {

                let data = await createVacancySkill(values)
                if (data.success == true) {
                    setSubmitting(false);
                } else {
                    setErrors({ 'comment': 'error!' })
                }

            }}>
            {({
                values,
                errors,
                touched,
                handleChange,
                handleBlur,
                handleSubmit,
                isSubmitting,
            }) => (
                <form onSubmit={handleSubmit} className='frm '>
                    <h3>Create Vacancy Skill</h3>
                    <input
                        type="text"
                        name="skill_id"
                        onChange={handleChange}
                        onBlur={handleBlur}
                        value={values.skill_id}
                    /> <br />
                    <span className='errors'>{errors.skill_id && touched.skill_id && errors.skill_id}</span> <br />
                    <input
                        type="text"
                        name="vacancy_id"
                        onChange={handleChange}
                        onBlur={handleBlur}
                        value={values.vacancy_id}
                    /> <br />
                    <span className='errors'>{errors.vacancy_id && touched.vacancy_id && errors.vacancy_id}</span> <br />
                    <input
                        type="number"
                        name="priority"
                        onChange={handleChange}
                        onBlur={handleBlur}
                        value={values.priority}
                    /> <br />
                    <span className='errors'>{errors.priority && touched.priority && errors.priority}</span> <br />
                    <button type="submit" disabled={isSubmitting}>
                        Create
                    </button>
                </form>
            )}
        </Formik>
    </div>
}

export function VacancyComponent(props){

    const [domains, setDomains] = useState(null)
    const [skills, setSkills] = useState(null)

    async function getVacancyDomains() {
        const response = await fetch(
            'http://0.0.0.0:3001/v1/vacancy/' + props.vacancy.id + '/domains',
            {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                },
            }
        )
        const data = await response.json()
        setDomains(data)
        return data
    }

    async function getVacancySkills() {
        const response = await fetch(
            'http://0.0.0.0:3001/v1/vacancy/' + props.vacancy.id + '/skills',
            {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                },
            }
        )
        const data = await response.json()
        setSkills(data)
        return data
    }

    useEffect(() => {
        getVacancyDomains()
        getVacancySkills()
    }, [])

    async function deleteVacancy(){
        await fetch(
            'http://0.0.0.0:3001/v1/vacancy/' + props.vacancy.id,
            {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json'
                },
            }
        )
        props.setRefresh(!props.refresh)
    }

    function callPopup(caller) {
        props.setPopup(caller, props.vacancy.id)
    }

    return <tr>
    <td><input type="checkbox" name={'vacancy_' + props.vacancy.id} /></td>
    <td>{props.vacancy.id}</td>
    <td>{props.vacancy.name}</td>
    <td>{props.vacancy.description}</td>
    <td className={domains !== null ? "clickable_row" : ""} onClick={() => callPopup("vacancy_domain")}>{domains === null ? "-" : domains.length}</td>
    <td className={skills !== null ? "clickable_row" : ""} onClick={() => callPopup("vacancy_skill")}>{skills === null ? "-" : skills.length}</td>
    
    <td>
      <button onClick={deleteVacancy}>Delete</button>
    </td>
  </tr>

}

export default function Vacancy(){
    const [vacancies, setVacancies] = useState([])
    const [refresh, setRefresh] = useState(false)
    const [popup, setPopup] = useState(null)
    const [popup_block, setPopupBlock] = useState(null)
    const [selectedVacancy, setSelectedVacancy] = useState(0)

    async function setUp() {
        setVacancies(await fetchVacancies())
    }

    useEffect(() => {
        setUp()
        
    }, [refresh, popup, selectedVacancy])

    if (vacancies.length == 0) {
        return <div className="cv_model">
            <h1>Vacancies</h1>
            <span>No records found!</span>
            <CreateVacancyForm refresh={refresh} setRefresh={setRefresh} />
        </div>
    }

    function selectPopup(new_popup, vacancy){
        setSelectedVacancy(vacancy)
        console.log(vacancy)
        console.log(new_popup)
        switch (new_popup) {
            case "vacancy_domain":
                
                setPopupBlock(<VacancyDomain vacancy_id={vacancy} refresh={refresh} setRefresh={setRefresh} />)
                break;
            case "vacancy_skill":
                
                setPopupBlock(<VacancySkill vacancy_id={vacancy} refresh={refresh} setRefresh={setRefresh} />)
                break;
            default:
                setPopupBlock(<span>Empty</span>)
        }
        setPopup(new_popup)
        
    }

    return <div className="cv_model">
        <h1>Vacancies</h1>
        <div className='cv_instances'>
            <table>
                <caption>
                    Skills
                </caption>
                <thead>
                    <tr>
                        <th></th>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Description</th>
                        <th>Domains</th>
                        <th>Skills</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {vacancies.map((vacancy, index) => <VacancyComponent setPopup={selectPopup} key={"vacancy_"+index} vacancy={vacancy} refresh={refresh} setRefresh={setRefresh} />)}
                </tbody>
            </table>
            <CreateVacancyForm refresh={refresh} setRefresh={setRefresh}></CreateVacancyForm>
        </div>
        <div className={popup !== null? 'popup' : 'hidden'}><button onClick={() => setPopup(null)}>Close</button>{popup_block}</div>
    </div>
}