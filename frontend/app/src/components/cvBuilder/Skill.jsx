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
    priority: number().positive().integer().required()
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
            console.log(res)
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
            console.log(errors)
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
    async function createSkillDomain(data) {
        try {
            let res = await skillSchema.validate(data, { abortEarly: false });
            console.log(res)
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
            props.setRefresh(!props.refresh)
        } catch (errors) {
            console.log(errors)
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
    </div>
}

export function SkillDependency(props) {
    async function createSkillDependency(data) {
        try {
            let res = await skillDependencySchema.validate(data, { abortEarly: false });
            console.log(res)
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
            props.setRefresh(!props.refresh)
        } catch (errors) {
            console.log(errors)
            return {
                success: false
            }
        }
        return {
            success: true
        }
    }
    return <div>
        <h3>Dependencies</h3>
    </div>
}

export function SkillConflict() {
    return <div>
        <h3>Conflicts</h3>
    </div>
}

function SkillComponent(props) {
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
    return <div className='cv_instance'>
        <ul>
            <li>ID: {props.skill.id}</li>
            <li>Name: {props.skill.name}</li>
            <li>Description: {props.skill.description}</li>
        </ul>
        <button onClick={deleteSkill}>Delete</button>
    </div>
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
    console.log(skills)
    if (skills.length == 0) {
        return <div className="cv_model">
            <h1>Skills</h1>
            <span>No records found!</span>
            <CreateSkillForm refresh={refresh} setRefresh={setRefresh} />
        </div>
    }
    return <div className="cv_model">
        <h1>Skills</h1>
        <div className='cv_instances'>
            {skills.map((skill) => {
                return <SkillComponent skill={skill} key={"skill_" + skill.id} refresh={refresh} setRefresh={setRefresh} />
            })}
        </div>
        <SkillDomain refresh={refresh} setRefresh={setRefresh}></SkillDomain>
        <SkillConflict refresh={refresh} setRefresh={setRefresh}></SkillConflict>
        <SkillDependency refresh={refresh} setRefresh={setRefresh}></SkillDependency>
        <CreateSkillForm refresh={refresh} setRefresh={setRefresh}></CreateSkillForm>
    </div>
}