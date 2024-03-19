from dataclasses import dataclass

class Query:

    available_query_args: frozenset = frozenset((
        'page', # int
        'search_field', # values are in available_search_fields
        'enable_snippets', # false by default
        'part_time', # available_part_time
        'only_with_salary', # bool
        'salary', # ?int, 50K, 100K, 150K, 200K, 250K, custom
        'area', # int 
        'L_save_area', # bool
        'experience',
        'schedule', # str
    ))

    available_search_fields: frozenset = frozenset((
        'name', 'company_name', 'description'
    ))

    available_part_time: frozenset = frozenset((
        'employment_part', # Not full day
        'start_after_sixteen', # Evening job
        'only_saturday_and_sunday', # Weekend job
        'employment_project', # Project job
        'from_four_to_six_hours_in_a_day', # 4-6 h/day
    ))

    available_experience: frozenset = frozenset((
        'between1And3',
        'between3And6',
        'moreThan6',
        'noExperience',
    ))

    available_schedule: frozenset = frozenset((
        'remote',
        'fullDay',
        'shift',
    ))

    def __init__(self) -> None:
        pass

    def __repr__(self) -> str:
        pass

@dataclass(slots=True)
class Company:
    name: str
    rating: int | None
    description: str | None
    city: str | None
    domains: list[str] | None
    site_url: str | None


@dataclass(slots=True)
class ItemPreview:
    vacancy_name: str
    vacancy_link: str

    company_name: str
    company_link: str

    vacancy_city: str

    experience_min: int
    experience_max: int

    remote_job: bool

    contacts_preview: bool

