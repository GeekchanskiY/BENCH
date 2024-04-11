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
                <h3>Create Skill</h3>
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


export function VacancySkill(){
    return <div>
        <h1>Skills</h1>
    </div>
}
export function VacancyDomain(){
    return <div>
        <h1>Domains</h1>
    </div>
}

export function VacancyComponent(props){
    return <div className='cv_instance'> 
        {props.vacancy.name}
    </div>

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
            case "skill_conflict":
                
                setPopupBlock()
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
                        <th>Dependencies</th>
                        <th>Conflicts</th>
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