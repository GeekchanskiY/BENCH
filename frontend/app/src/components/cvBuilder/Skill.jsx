import { object, string, number, date } from 'yup';
import { useEffect, useState } from 'react';
import { Formik } from 'formik';

let skillSchema = object().shape({
    id: number().positive().integer(),
    name: string().required(),
    description: string().required()
});


let skillDomainSchema = object().shape({
    skill_id: number().positive().integer().required(),
    domain_id: number().positive().integer().required(),
    priority: number().positive().integer().required()
});

let skillConflictSchema = object().shape({
    skill_1: number().positive().integer().required(),
    skill_2: number().positive().integer().required(),
    priority: number().positive().integer().required(),
    comment: string().required()
});

let skillDependencySchema = object().shape({
    parent_skill: number().positive().integer().required(),
    child_skill: number().positive().integer().required(),
});


async function fetchSkills() {
    let response = await fetch('http://0.0.0.0:3001/v1/skill/')
    let skills = await response.json()
    if (skills == null) {
        return []
    }
    let validatedObjects = await Promise.all(skills.map(async (object) => {
        await skillSchema.validate(object, { abortEarly: false });
        return object;
    }));
    return await validatedObjects
}

function CreateSkillForm({ refresh, setRefresh }) {
    async function createSkill(data) {

        try {
            let res = await skillSchema.validate(data, { abortEarly: false });
            
            let response = await fetch(
                'http://0.0.0.0:3001/v1/skill/',
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
            name: 'Python',
            description: 'Programming language'
        }}
        validationSchema={skillSchema}
        onSubmit={async (values, { setSubmitting, setErrors }) => {

            let data = await createSkill(values)
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
                <button type="submit" disabled={isSubmitting}>
                    Create
                </button>
            </form>
        )}
    </Formik>
}

export function SkillDomain(props) {
    const [skillDomain, setSkillDomain] = useState([])

    async function fetchSkillDomain(){
        let response = await fetch(
            'http://0.0.0.0:3001/v1/skill/domain',
            {
                method: 'GET',
            }
        )
        let skilldomain = await response.json()
        if (skilldomain.length == 0) {
            setSkillDomain([])
        }
        let validatedObjects = await Promise.all(skilldomain.map(async (object) => {
            await skillDomainSchema.validate(object, { abortEarly: false });
            return object;
        }));
        return validatedObjects;
    }

    async function setUp(){
        setSkillDomain(await fetchSkillDomain())
    }
    useEffect(() => {
        setUp()
    }, [])
    async function createSkillDomain(data) {
        try {
            
            let res = await skillDomainSchema.validate(data, { abortEarly: false });
           
            let response = await fetch(
                'http://0.0.0.0:3001/v1/skill/domain',
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(res)
                }
            )
            props.setRefresh(!props.refresh)
        } catch (errors) {
            return {
                success: false
            }
        }
        return {
            success: true
        }
    }
    return <div>
        <h3>Domains</h3>
        <div className='cv_instances'>
            {skillDomain.map((skill, index) => {
                return <div key={index} className='cv_instance'>
                    <ul>
                        <li>Skill_id: {skill.skill_id}</li>
                        <li>Domain_id: {skill.domain_id}</li>
                        <li>Priority: {skill.priority}</li>
                    </ul>
                    
                </div>
            })}
        </div>
        <Formik
            initialValues={{
                skill_id: 1,
                domain_id: 2,
                priority: 1,
            }}
            validationSchema={skillDomainSchema}
            onSubmit={async (values, { setSubmitting, setErrors }) => {

                let data = await createSkillDomain(values)
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
                    <h3>Create Skill Domain</h3>
                    <input
                        type="text"
                        name="skill_id"
                        onChange={handleChange}
                        onBlur={handleBlur}
                        value={values.skill_id}
                    /> <br />
                    <span className='errors'>{errors.parent_skill && touched.parent_skill && errors.parent_skill}</span> <br />
                    <input
                        type="text"
                        name="domain_id"
                        onChange={handleChange}
                        onBlur={handleBlur}
                        value={values.domain_id}
                    /> <br />
                    <span className='errors'>{errors.child_skill && touched.child_skill && errors.child_skill}</span> <br />
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

export function SkillDependency(props) {
    const [skillDependency, setSkillDependency] = useState([])
    async function fetchSkilldeps() {
        let response = await fetch(
            'http://0.0.0.0:3001/v1/skill/dependency',
            {
                method: 'GET',
            }
        )
        let skilldeps = await response.json()
        if (skilldeps.length == 0) {
            setSkillDependency([])
        }
        let validatedObjects = await Promise.all(skilldeps.map(async (object) => {
            await skillDependencySchema.validate(object, { abortEarly: false });
            return object;
        }));
        return validatedObjects;

    }
    async function setUp() {
        setSkillDependency(await fetchSkilldeps())
        
    }
    useEffect(() => {
        setUp()
    }, [])
    async function createSkillDependency(data) {
        try {
            // console.log(data)
            let res = await skillDependencySchema.validate(data, { abortEarly: false });
           
            let response = await fetch(
                'http://0.0.0.0:3001/v1/skill/dependency',
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
            console.error(errors)
            return {
                success: false
            }
        }
        return {
            success: true
        }
    }
    async function deleteSkillDependency(skill) {
        console.log(skillDependency)
        await fetch("http://0.0.0.0:3001/v1/skill/dependency", {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(skill)
        })
    }
    return <div>
        <h3>Dependencies</h3>
        <div className='cv_instances'>
            {skillDependency.map((skill, num) => {
                return <div key={"skill_dependency_" + num} className='cv_instance'><ul>
                    <li>Parent: {skill.parent_skill}</li>
                    <li>Child: {skill.child_skill}</li>
                    <li><button onClick={() => deleteSkillDependency(skill)}>Delete</button></li>
                </ul></div>
            })}
        </div>
        <Formik
            initialValues={{
                parent_skill: 1,
                child_skill: 2
            }}
            validationSchema={skillDependencySchema}
            onSubmit={async (values, { setSubmitting, setErrors }) => {

                let data = await createSkillDependency(values)
                setUp()
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
                    <h3>Create Skill Dependency</h3>
                    <input
                        type="text"
                        name="parent_skill"
                        onChange={handleChange}
                        onBlur={handleBlur}
                        value={values.parent_skill}
                    /> <br />
                    <span className='errors'>{errors.parent_skill && touched.parent_skill && errors.parent_skill}</span> <br />
                    <input
                        type="text"
                        name="child_skill"
                        onChange={handleChange}
                        onBlur={handleBlur}
                        value={values.child_skill}
                    /> <br />
                    <span className='errors'>{errors.child_skill && touched.child_skill && errors.child_skill}</span> <br />
                    <button type="submit" disabled={isSubmitting}>
                        Create
                    </button>
                </form>
            )}
        </Formik>
    </div>
}

export function SkillConflict() {
    const [skillConflict, setSkillConflict] = useState([])

    async function fetchSkillConflict() {
        let response = await fetch(
            'http://0.0.0.0:3001/v1/skill/conflict',
            {
                method: 'GET',
            }
        )
        let skillconflict = await response.json()
        if (skillconflict.length == 0) {
            setSkillConflict([])
        }
        let validatedObjects = await Promise.all(skillconflict.map(async (object) => {
            await skillConflictSchema.validate(object, { abortEarly: false });
            return object;
        }));
        return validatedObjects;
    }

    async function setUp(){
        setSkillConflict(await fetchSkillConflict())
    }
    useEffect(() => {
        setUp()
    }, [])
    async function createSkillConflict(data) {
        try {
            let res = await skillConflictSchema.validate(data, { abortEarly: false });
           
            let response = await fetch(
                'http://0.0.0.0:3001/v1/skill/conflict',
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(res)
                }
            )
        } catch (errors) {
            return {
                success: false
            }
        }
        return {
            success: true
        }
    }
    return <div>
        <h3>Conflicts</h3>
        <div className='cv_instances'>
            {skillConflict.map((skill, index) => {
                return <div key={index} className='cv_instance'>
                    <ul>
                        <li>Skill_1_id: {skill.skill_1}</li>
                        <li>Skill_2_id: {skill.skill_2}</li>
                        <li>Comment: {skill.comment}</li>
                        <li>Priority: {skill.priority}</li>
                    </ul>
                    
                </div>
            })}
        </div>
        <Formik
            initialValues={{
                skill_1: 1,
                skill_2: 2,
                comment: 'You cant use them both!',
                priority: 1
            }}
            validationSchema={skillConflictSchema}
            onSubmit={async (values, { setSubmitting, setErrors }) => {

                let data = await createSkillConflict(values)
                setUp()
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
                    <h3>Create Skill Domain</h3>
                    <input
                        type="text"
                        name="skill_1"
                        onChange={handleChange}
                        onBlur={handleBlur}
                        value={values.skill_1}
                    /> <br />
                    <span className='errors'>{errors.skill_1 && touched.skill_1 && errors.skill_1}</span> <br />
                    <input
                        type="text"
                        name="skill_2"
                        onChange={handleChange}
                        onBlur={handleBlur}
                        value={values.skill_2}
                    /> <br />
                    <span className='errors'>{errors.skill_2 && touched.skill_2 && errors.skill_2}</span> <br />
                    <input
                        type="number"
                        name="priority"
                        onChange={handleChange}
                        onBlur={handleBlur}
                        value={values.priority}
                    /> <br />
                    <span className='errors'>{errors.priority && touched.priority && errors.priority}</span> <br />
                    <input
                        type="text"
                        name="comment"
                        onChange={handleChange}
                        onBlur={handleBlur}
                        value={values.comment}
                    /> <br />
                    <span className='errors'>{errors.comment && touched.comment && errors.comment}</span> <br />
                    <button type="submit" disabled={isSubmitting}>
                        Create
                    </button>
                </form>
            )}
        </Formik>
    </div>
}

function SkillComponent(props) {
    const [skillDependency, setSkillDependency] = useState([])
    const [skillConflict, setSkillConflict] = useState([])
    const [skillDomain, setSkillDomain] = useState([])

    async function deleteSkill() {
        await fetch(
            'http://0.0.0.0:3001/v1/skill/' + props.skill.id,
            {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json'
                },
            }

        )
        props.setRefresh(!props.refresh)
    }

    async function getSkillDependencies() {
        const response = await fetch(
            'http://0.0.0.0:3001/v1/skill/' + props.skill.id + '/dependencies',
            {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                },
            }
        )
        const data = await response.json()
        if (data != null){
            console.log("skill_id:" + props.skill.id)
            console.log(data)
        }
        setSkillDependency(data)
        return data
    }

    useEffect(() => {
        getSkillDependencies()
    }, [])

    return  <tr>
        <td>{props.skill.id}</td>
        <td>{props.skill.name}</td>
        <td>{props.skill.description}</td>
        <td></td>
        <td>{skillDependency === null ? "No dependencies" : skillDependency.length}</td>
        <td></td>
        <td>
            <button onClick={deleteSkill}>Delete</button>
        </td>
    </tr>
}

export default function Skill() {
    const [skills, setSkills] = useState([])
    const [refresh, setRefresh] = useState(false)
    async function setUp() {
        setSkills(await fetchSkills())
    }
    useEffect(() => {
        setUp()
    }, [refresh])
    if (skills.length == 0) {
        return <div className="cv_model">
            <h1>Skills</h1>
            <span>No records found!</span>
            <CreateSkillForm refresh={refresh} setRefresh={setRefresh} />
        </div>
    }
    return <div className="cv_model">
        <div className='cv_instances'>
            <table>
                <caption>
                    Skills
                </caption>
                <thead>
                    <tr>
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
                    {skills.map((skill, index) => <SkillComponent key={"skill_"+index} skill={skill} refresh={refresh} setRefresh={setRefresh} />)}
                </tbody>
            </table>
            <CreateSkillForm refresh={refresh} setRefresh={setRefresh}></CreateSkillForm>
        </div>
        

        <SkillDomain refresh={refresh} setRefresh={setRefresh}></SkillDomain>
        <SkillConflict refresh={refresh} setRefresh={setRefresh}></SkillConflict>
        <SkillDependency refresh={refresh} setRefresh={setRefresh}></SkillDependency>

    </div>
}