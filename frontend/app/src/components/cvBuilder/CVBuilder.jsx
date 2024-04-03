import Company from "./Company"
import CV from "./CV"
import Domain from "./Domain"
import Employee from "./Employee"
import Responsibility from "./Responsibility"
import Skill from "./Skill"
import Vacancy from "./Vacancy"

import '../../styles/cv_builder.css'

export default function CVBuilder() {
    return <div>
        <h1>CV Builder</h1>
        <Company></Company>
        <CV></CV>
        <Domain></Domain>
        <Employee></Employee>
        <Responsibility></Responsibility>
        <Skill></Skill>
        <Vacancy></Vacancy>



    </div>
}