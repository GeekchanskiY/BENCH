import Company from "./Company"
import CV from "./CV"
import Domain from "./Domain"
import Employee from "./Employee"
import Responsibility from "./Responsibility"
import Skill from "./Skill"
import Vacancy from "./Vacancy"
import { useState } from "react"

import '../../styles/cv_builder.css'

export default function CVBuilder() {
    const [currpage, setCurrpage] = useState('skill')

    function get_page() {
        switch (currpage) {
            case 'skill':
                return <Skill></Skill>
            case 'domain':
                return <Domain></Domain>
            case 'employee':
                return <Employee></Employee>
            case 'responsibility':
                return <Responsibility></Responsibility>
            case 'cv':
                return <CV></CV>
            case 'company':
                return <Company></Company>
            case 'vacancy':
                return <Vacancy></Vacancy>
            default:
                return <Skill></Skill>
        }
    }

    return <div className="cvbuilder_main">
        
        <aside className="cvbuilder_sidebar">
            <h1>CV Builder</h1>
            <ul>
                <li onClick={() => setCurrpage('skill')}>Skill</li>
                <li onClick={() => setCurrpage('domain')}>Domain</li>
                <li onClick={() => setCurrpage('employee')}>Employee</li>
                <li onClick={() => setCurrpage('responsibility')}>Responsibility</li>
                <li onClick={() => setCurrpage('cv')}>CV</li>
                <li onClick={() => setCurrpage('company')}>Company</li>
                <li onClick={() => setCurrpage('vacancy')}>Vacancy</li>
        
            </ul>
        </aside>
        <section className="cvbuilder_content">
            {get_page()}
        </section>



    </div>
}