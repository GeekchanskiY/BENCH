from dataclasses import dataclass
from config import SEARCH_STRING

from copy import deepcopy, copy


class Query:

    query_args: list = []
    search_url: str = SEARCH_STRING
    page: int = 1
    text: str = None
    search_fields: list[str] = list()
    enable_snippets: bool = False
    part_time: list[str] = list()
    only_with_salary: bool = False
    salary: int | None = None
    area: int | None = None
    L_save_area: int | None = None
    experience: list[str] = list()
    schedule: str | None = None
    

    available_query_args: frozenset = frozenset((
        'text', # str
        'page', # int
        'search_field', # values are in available_search_fields
        'enable_snippets', # false by default
        'part_time', # available_part_time
        'only_with_salary', # bool
        'salary', # ?int, 50K, 100K, 150K, 200K, 250K, custom
        'area', # int 
        'L_save_area', # bool
        'experience', # available_experience
        'schedule', # available_schedule
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

    def __init__(
            self,
            page: int|None = None,
            search_field: str|list|None = None,
            search_query: str|None = None,
            enable_snippets: bool = False,
            part_time: str|list|None = None,
            only_with_salary: bool = False,
            salary: int|None = None,
            area: int|None = None,
            experience: str|None = None,
            schedule: str|list|None = None
        ) -> None:
        
        if area is not None:
            # L_save_args required if search by area enabled
            self.query_args.append(
                ('L_save_args', True)
            )
    
        if page is not None and type(page) == int:
            self.page = page
            self.query_args.append(
                ('page', page)
            )
        
        if experience is not None:
            if type(experience) is not str:
                raise AttributeError('Experience must be str!')
            if experience not in self.available_experience:
                raise AttributeError('Invalid Experience')
            self.query_args.append(
                ('experience', experience)
            )
        
        if search_query is not None:
            if type(search_query) is not str:
                raise AttributeError('Search query must be a str')
            self.query_args.append(
                ('text', search_query)
            )

        if search_field is not None:
            
            if type(search_field) == str:
                if search_field not in self.available_search_fields:
                    raise AttributeError('Invalid search_field')
                self.query_args.append(
                    ('search_field', search_field_param)
                )
            elif type(search_field) == list:
                search_field_param: str
                for search_field_param in search_field:
                    if search_field_param not in self.available_search_fields:
                        raise AttributeError('Invalid search_field in iterable')
                    self.query_args.append(
                        ('search_field', search_field_param)
                    )
                del search_field_param
            else:
                raise AttributeError('search_field must be None, str or list[str]')
        else:
            field: str
            for field in self.available_search_fields:
                self.query_args.append(
                    ('search_field', field)
                )
        print(self.query_args)

    def get_next_page(self) -> 'Query':
        new_query: Query = copy(self)
        new_query.query_args.clear()
        
        new_query.page += 1
        print(new_query.page)
        return new_query

    def __str__(self) -> str:
        for q in self.query_args:
            
            key, value = q
            self.search_url += f'&{key}={value}'
        return self.search_url

    def __repr__(self) -> str:
        return str(self)

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

if __name__ == '__main__':
    q = Query(
        page=4,
        search_query='python'
    )
    print(str(q))
    print(str(q.get_next_page()))