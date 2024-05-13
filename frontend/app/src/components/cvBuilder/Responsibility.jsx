import { object, string, number, date } from 'yup';
import { useEffect, useState } from 'react';
import { Formik } from 'formik';

let responsibilitySchema = object().shape({
    skill_id: number().positive().integer().required(),
    priority: number().positive().integer().required(),
    name: string().required(),
    comments: string().required(),
    experience_level: number().positive().integer().required()
});


let responsibilitySynonimSchema = object().shape({
    responsibility_id: number().positive().integer().required(),
    name_id: string().required()
});

let responsibilityConflictSchema = object().shape({
    responsibility_1_id: number().positive().integer().required(),
    responsibility_2_id: number().positive().integer().required(),
    priority: number().positive().integer().required()
});

async function fetchResponsibilities() {
    let response = await fetch('http://0.0.0.0:3001/v1/responsibility/')
    let responsibilities = await response.json()
    if (responsibilities == null) {
        return []
    }
    let validatedObjects = await Promise.all(responsibilities.map(async (object) => {
        await responsibilitySchema.validate(object, { abortEarly: false });
        return object;
    }));
    return await validatedObjects
}

function CreateResponsibilityForm({ refresh, setRefresh }) {
    async function createResponsibility(data) {
        try {
            let res = await responsibilitySchema.validate(data, { abortEarly: false });
            
            await fetch(
                'http://0.0.0.0:3001/v1/responsibility/',
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
            skill_id: 1,
            priority: 1,
            name: 'Python',
            comments: 'Programming language',
            experience_level: 1
        }}
        validationSchema={responsibilitySchema}
        onSubmit={async (values, { setSubmitting, setErrors }) => {

            let data = await createResponsibility(values)
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
                <h3>Create Responsibility</h3>
                <input
                    type="text"
                    name="name"
                    onChange={handleChange}
                    onBlur={handleBlur}
                    value={values.name}
                /> <br />
                <span className='errors'>{errors.name && touched.name && errors.name}</span> <br />
                <input type="text" name="comments" onChange={handleChange} onBlur={handleBlur} value={values.comments} /> <br />
                <span className='errors'>{errors.comments && touched.comments && errors.comments}</span> <br />
                <span>experience_level</span> <br />
                <input type="number" name="experience_level" onChange={handleChange} onBlur={handleBlur} value={values.experience_level} /> <br />
                <span className='errors'>{errors.experience_level && touched.experience_level && errors.experience_level}</span> <br />
                <span>skill_id</span> <br />
                <input type="text" name="skill_id" onChange={handleChange} onBlur={handleBlur} value={values.skill_id} /> <br />
                <span className='errors'>{errors.skill_id && touched.skill_id && errors.skill_id}</span> <br />
                <span>priority</span> <br />
                <input type="number" name="priority" onChange={handleChange} onBlur={handleBlur} value={values.priority} /> <br />
                <span className='errors'>{errors.priority && touched.priority && errors.priority}</span> <br />
                <button type="submit" disabled={isSubmitting}>
                    Create
                </button>
            </form>
        )}
    </Formik>
}   

export function ResponsibilityComponent(props) {
    const [responsibilitySynonims, setResponsibilitySynonims] = useState([])
    const [responsibilityConflicts, setResponsibilityConflicts] = useState([])
    async function fetchResponsibilitySynonims() {
        let response = await fetch(
            'http://0.0.0.0:3001/v1/responsibility/'+ props.responsibility.id +'/synonims',
            {
                method: 'GET',
            }
        )
        let responsibilitysynonims = await response.json()
        if (responsibilitysynonims === null) {
            return []
        }
        let validatedObjects = await Promise.all(responsibilitysynonims.map(async (object) => {
            await responsibilitySynonimSchema.validate(object, { abortEarly: false });
            return object;
        }));
        return validatedObjects;
    }

    
    async function fetchResponsibilityConflicts() {
        let response = await fetch(
            'http://0.0.0.0:3001/v1/responsibility/'+ props.responsibility.id +'/conflicts',
            {
                method: 'GET',
            }
        )
        let responsibilityconflicts = await response.json()
        if (responsibilityconflicts === null) {
            return []
        }
        let validatedObjects = await Promise.all(responsibilityconflicts.map(async (object) => {
            await responsibilityConflictSchema.validate(object, { abortEarly: false });
            return object;
        }));
        return validatedObjects;
    }

    async function setUp(){
        setResponsibilitySynonims(await fetchResponsibilitySynonims())
        setResponsibilityConflicts(await fetchResponsibilityConflicts())
    }
    useEffect(() => {
        setUp()
    }, [])
    async function deleteResponsibility(){
        await fetch(
            'http://0.0.0.0:3001/v1/responsibility/' + props.responsibility.id,
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
        props.setPopup(caller, props.responsibility.id)
    }
    return <tr>
        <td><input type="checkbox" name={'responsibility_' + props.responsibility.name} /></td>
      
        <td>{props.responsibility.name}</td>
        <td>{props.responsibility.comments}</td>
        <td>{props.responsibility.experience_level}</td>
        <td>{props.responsibility.skill_id}</td>
        <td>{props.responsibility.priority}</td>
        <td><button onClick={() => deleteResponsibility()}>Delete</button></td>
    </tr>
}

export function ResponsibilitySynonim(){
    return <div>
        <h3>Synonims</h3>
    </div>
}

export function ResponsibilityConflict(){
    return <div>
        <h3>Conflicts</h3>
    </div>
}

export default function Responsibility(){
    const [responsibilities, setResponsibilities] = useState([])
    const [refresh, setRefresh] = useState(false)
    const [popup, setPopup] = useState(null)
    const [popup_block, setPopupBlock] = useState(null)
    const [selectedResponsibility, setSelectedResponsibility] = useState(0)

    async function setUp() {
        setResponsibilities(await fetchResponsibilities())
    }
    useEffect(() => {
        setUp()
        
    }, [refresh, popup, selectedResponsibility])
    function selectPopup(new_popup, responsibility){
        setSelectedResponsibility(responsibility)
        console.log(responsibility)
        console.log(new_popup)
        switch (new_popup) {
            case "responsibility_synonim":
                setPopupBlock(<ResponsibilitySynonim refresh={refresh} setRefresh={setRefresh} selectedResponsibility={responsibility}></ResponsibilitySynonim>)
                break;
            case "responsibility_conflict":
                setPopupBlock(<ResponsibilityConflict refresh={refresh} setRefresh={setRefresh} selectedResponsibility={responsibility}></ResponsibilityConflict>)
                break;
            default:
                setPopupBlock(<span>Empty</span>)
        }
        setPopup(new_popup)
        
    }
    if (responsibilities.length == 0) {
        return <div className="cv_model">
            <h1>Responsibilities</h1>
            <span>No records found!</span>
            <CreateResponsibilityForm refresh={refresh} setRefresh={setRefresh} />
        </div>
    }
    return <div className="cv_model">
        <div className='cv_instances'>
            <table>
                <caption>
                    Responsibilities
                </caption>
                <thead>
                    <tr>
                        <th></th>
                        <th>Name</th>
                        <th>Comments</th>
                        <th>Experience Level</th>
                        <th>Skill</th>
                        <th>Priority</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {responsibilities.map((responsibility, index) => <ResponsibilityComponent setPopup={selectPopup} key={"responsibility_"+index} responsibility={responsibility} refresh={refresh} setRefresh={setRefresh} />)}
                </tbody>
            </table>
            <CreateResponsibilityForm refresh={refresh} setRefresh={setRefresh}></CreateResponsibilityForm>
        </div>
        <div className={popup !== null? 'popup' : 'hidden'}><button onClick={() => setPopup(null)}>Close</button>{popup_block}</div>
    </div>
}